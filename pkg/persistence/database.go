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

var version string

func NewDatabase(ctx context.Context) services.Database {
	return &mysqlDatabaseFinal{}
}

func (n *mysqlDatabaseFinal) Start(ctx context.Context) error {
	sm := n.ServiceManager()
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(db:3306)/"+dbName)
	slogs := sm.LogsService()
	if err != nil {
		slogs.Error(ctx, err.Error())
		return err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	n.db = db
	if errH := n.Healthy(ctx); errH != nil {
		slogs.Error(ctx, errH.Error())
		return errH
	}
	slogs.Info(ctx, "Database Started!")
	return nil
}

func (n *mysqlDatabaseFinal) Close(ctx context.Context) error {
	return n.db.Close()
}

func (n *mysqlDatabaseFinal) Healthy(ctx context.Context) error {
	db := n.db
	return db.QueryRow("SELECT VERSION()").Scan(&version)
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
	err := n.db.QueryRow("SELECT * FROM planet WHERE id = :0", id).Scan(&p.Id, &p.Name, &p.Climate, &p.Terrain, &p.FilmURLs, &p.URL)
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
	rows, err := n.db.Query("SELECT * FROM planet WHERE name LIKE '%:0%'", name)
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
		if err := rows.Scan(&p.Id, &p.Name, &p.Climate, &p.Terrain, &p.FilmURLs, &p.URL); err != nil {
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
	stmt, err := n.db.PrepareContext(ctx, "INSERT INTO planet(id, name, climate, terrain, filmsurl, url) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		//log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Id, p.Name, p.Climate, p.Terrain, p.FilmURLs, p.URL)
	if err != nil {
		//log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		//log.Printf("Error %s when finding rows affected", err)
		return err
	}
	//log.Printf("%d products created ", len(rows))
	return nil
}

func (n *mysqlDatabaseFinal) UpdatePlanet(ctx context.Context, p *models.Planet) error {
	//TODO: Finish the update
	stmt, err := n.db.PrepareContext(ctx, "UPDATE planet(id, name, climate, terrain, filmsurl, url) SET (?, ?, ?, ?, ?, ?)")
	if err != nil {
		//log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Id, p.Name, p.Climate, p.Terrain, p.FilmURLs, p.URL)
	if err != nil {
		//log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		//log.Printf("Error %s when finding rows affected", err)
		return err
	}
	//log.Printf("%d products created ", len(rows))
	return nil
}

func (n *mysqlDatabaseFinal) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	rows, err := n.db.Query("SELECT * FROM planet")
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
		if err := rows.Scan(&p.Id, &p.Name, &p.Climate, &p.Terrain, &p.FilmURLs, &p.URL); err != nil {
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
		//log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		//log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		//log.Printf("Error %s when finding rows affected", err)
		return err
	}
	//log.Printf("%d products created ", len(rows))
	return nil
}

func (n *mysqlDatabaseFinal) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	stmt, err := n.db.PrepareContext(ctx, "DELETE FROM planet WHERE name = ?")
	if err != nil {
		//log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, exactName)
	if err != nil {
		//log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		//log.Printf("Error %s when finding rows affected", err)
		return err
	}
	//log.Printf("%d products created ", len(rows))
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
