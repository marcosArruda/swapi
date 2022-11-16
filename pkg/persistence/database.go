package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

type (
	mysqlDatabaseFinal struct {
		sm services.ServiceManager
		db *sql.DB
	}
)

var (
	version               string
	planetCreateTable     = "CREATE TABLE IF NOT EXISTS planet (id INT PRIMARY KEY, name VARCHAR(255) NOT NULL, climate VARCHAR(50) NOT NULL,terrain VARCHAR(40), url VARCHAR(255),INDEX (name))"
	filmCreateTable       = "CREATE TABLE IF NOT EXISTS film (id INT PRIMARY KEY, title VARCHAR(255) NOT NULL, episodeid INT NOT NULL, director VARCHAR(50) NOT NULL, created VARCHAR(40) NOT NULL, url VARCHAR(255) NOT NULL,INDEX (title))"
	manyToManyCreateTable = "CREATE TABLE IF NOT EXISTS planet_film (filmid INT NOT NULL, planetid INT NOT NULL, INDEX (filmid, planetid))"
)

func NewDatabase(ctx context.Context) services.Database {
	return &mysqlDatabaseFinal{}
}

func (n *mysqlDatabaseFinal) Start(ctx context.Context) error {
	sm := n.ServiceManager()
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHostPort := os.Getenv("DB_HOSTPORT")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHostPort, dbName))

	if err != nil {
		sm.LogsService().Error(ctx, err.Error())
		return err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	n.db = db
	if err = n.Healthy(ctx); err != nil {
		sm.LogsService().Error(ctx, err.Error())
		return err
	}
	sm.LogsService().Info(ctx, "Database Started!")
	//sm.LogsService().Info(ctx, "Creating Tables ...")
	//if err = n.createTablesIfNotExists(ctx); err != nil {
	//	sm.LogsService().Error(ctx, "Error creating tables: "+err.Error())
	//	return err
	//}

	//n.InsertPlanet(ctx, &models.Planet{
	//	Id:      9999,
	//	Name:    "Terra",
	//	Climate: "Good",
	//	Terrain: "solid",
	//	URL:     "http://something.com/planet/9999",
	//})
	//p, err := n.GetPlanetById(ctx, 9999)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("Planet{id: %d, name: %s, climate: %s, terrain: %s, url: %s}", p.Id, p.Name, p.Climate, p.Terrain, p.URL)
	//fmt.Println("")
	//if err = n.RemovePlanetById(ctx, 9999); err != nil {
	//	return err
	//}
	sm.LogsService().Info(ctx, "Basic Tables Created!")
	return nil
}

func (n *mysqlDatabaseFinal) createTablesIfNotExists(ctx context.Context) error {
	_, err := n.db.ExecContext(ctx, filmCreateTable)
	if err != nil {
		return err
	}
	_, err = n.db.ExecContext(ctx, planetCreateTable)
	if err != nil {
		return err
	}
	_, err = n.db.ExecContext(ctx, manyToManyCreateTable)
	if err != nil {
		return err
	}
	return nil
}

func (n *mysqlDatabaseFinal) Close(ctx context.Context) error {
	return n.db.Close()
}

func (n *mysqlDatabaseFinal) Healthy(ctx context.Context) error {
	return n.db.QueryRow("SELECT VERSION()").Scan(&version)
}

func (n *mysqlDatabaseFinal) WithServiceManager(sm services.ServiceManager) services.Database {
	n.sm = sm
	return n
}

func (n *mysqlDatabaseFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *mysqlDatabaseFinal) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	p := &models.Planet{}
	err := n.db.QueryRow("SELECT id, name, climate, terrain, url FROM planet WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Climate, &p.Terrain, &p.URL)
	if err == sql.ErrNoRows {
		return nil, messages.NoPlanetFound
	}
	if err != nil {
		msg := fmt.Sprintf("Something went wrong searching locally by Planet with ID %d: %s", id, err.Error())
		return nil, &messages.PlanetError{Msg: msg, PlanetId: id}
	}
	return p, nil
}

func (n *mysqlDatabaseFinal) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	rows, err := n.db.Query("SELECT id, name, climate, terrain, url FROM planet WHERE name LIKE '%?%'", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return services.EmptyPlanetSlice, messages.NoPlanetFound
		}
		return n.emptyAndNameError(name, err)
	}
	defer rows.Close()
	planets := []*models.Planet{}
	for rows.Next() {
		var p models.Planet
		if err := rows.Scan(&p.Id, &p.Name, &p.Climate, &p.Terrain, &p.URL); err != nil {
			return n.emptyAndNameError(name, err)
		}
		planets = append(planets, &p)
	}
	if err := rows.Err(); err != nil {
		return n.emptyAndNameError(name, err)
	}
	return planets, nil
}

func (n *mysqlDatabaseFinal) InsertPlanet(ctx context.Context, p *models.Planet) error {
	ctx = context.Background()
	stmt, err := n.db.PrepareContext(context.Background(), "INSERT INTO planet(id, name, climate, terrain, url) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when preparing SQL statement: %s", err.Error()))
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Id, p.Name, p.Climate, p.Terrain, p.URL)
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when inserting row into planet table: %s ", err.Error()))
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when when finding rows affected: %s ", err.Error()))
		return err
	}
	n.sm.LogsService().Info(ctx, fmt.Sprintf("Planet Inserted! Planet ID: %d ", p.Id))
	return nil
}

func (n *mysqlDatabaseFinal) UpdatePlanet(ctx context.Context, p *models.Planet) error {
	//TODO: Finish the update
	stmt, err := n.db.PrepareContext(ctx, "UPDATE planet SET name = ?, climate = ?, terrain = ?, url = ? WHERE id = ?")
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when preparing SQL statement: %s", err.Error()))
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Name, p.Climate, p.Terrain, p.URL, p.Id)
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when updating planet table: %s ", err.Error()))
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when when finding rows affected: %s ", err.Error()))
		return err
	}

	n.sm.LogsService().Info(ctx, fmt.Sprintf("Planet Inserted! Planet ID: %d ", p.Id))
	return nil
}

func (n *mysqlDatabaseFinal) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	rows, err := n.db.Query("SELECT id, name, climate, terrain, url FROM planet")
	if err != nil {
		if err == sql.ErrNoRows {
			return services.EmptyPlanetSlice, messages.NoPlanetFound
		}
		return n.emptyAndGenericError(err)
	}
	defer rows.Close()
	planets := []*models.Planet{}
	for rows.Next() {
		var p models.Planet
		if err := rows.Scan(&p.Id, &p.Name, &p.Climate, &p.Terrain, &p.URL); err != nil {
			return n.emptyAndGenericError(err)
		}
		planets = append(planets, &p)
	}
	if err := rows.Err(); err != nil {
		return n.emptyAndGenericError(err)
	}
	return planets, nil
}

func (n *mysqlDatabaseFinal) RemovePlanetById(ctx context.Context, id int) error {
	stmt, err := n.db.PrepareContext(ctx, "DELETE FROM planet WHERE id = ?")
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when preparing SQL statement: %s", err.Error()))
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when deleting planet: %s", err.Error()))
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when when finding rows affected: %s", err.Error()))
		return err
	}

	n.sm.LogsService().Info(ctx, fmt.Sprintf("Planet Removed! Planet ID: %d ", id))
	return nil
}

func (n *mysqlDatabaseFinal) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	stmt, err := n.db.PrepareContext(ctx, "DELETE FROM planet WHERE name = ?")
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when preparing SQL statement: %s", err.Error()))
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, exactName)
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when deleting planet: %s", err.Error()))
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		n.sm.LogsService().Error(ctx, fmt.Sprintf("Error when when finding rows affected: %s", err.Error()))
		return err
	}
	n.sm.LogsService().Info(ctx, fmt.Sprintf("Planet Removed! Planet Name: %s ", exactName))
	return nil
}

func (n *mysqlDatabaseFinal) emptyAndNameError(name string, err error) ([]*models.Planet, error) {
	baseMsg := "Something went wrong searching locally by Planets with name "
	msg := fmt.Sprintf("%s'%s': %s", baseMsg, name, err.Error())
	return services.EmptyPlanetSlice, &messages.PlanetError{Msg: msg, PlanetName: name}
}

func (n *mysqlDatabaseFinal) emptyAndGenericError(err error) ([]*models.Planet, error) {
	baseMsg := "Something went wrong searching locally by All Planets: "
	msg := fmt.Sprintf("%s%s", baseMsg, err.Error())
	return services.EmptyPlanetSlice, &messages.PlanetError{Msg: msg}
}
