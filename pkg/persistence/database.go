package persistence

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		n.sm.LogsService().Error(err.Error())
		return err
	}

	n.db = db
	if errH := n.Healthy(ctx); errH != nil {
		n.sm.LogsService().Error(err.Error())
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
	return nil, nil
}

func (n *mysqlDatabaseFinal) InsertPlanet(ctx context.Context, p *models.Planet) (*models.Planet, error) {
	return nil, nil
}
