package persistence

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

func Test_mysqlDatabaseFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_createTablesIfNotExists(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.createTablesIfNotExists(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.createTablesIfNotExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_WithServiceManager(t *testing.T) {
	type args struct {
		sm services.ServiceManager
	}
	tests := []struct {
		name string
		n    *mysqlDatabaseFinal
		args args
		want services.Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.WithServiceManager(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.WithServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_ServiceManager(t *testing.T) {
	tests := []struct {
		name string
		n    *mysqlDatabaseFinal
		want services.ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_BeginTransaction(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    *sql.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.BeginTransaction(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.BeginTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_CommitTransaction(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.CommitTransaction(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.CommitTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_RollbackTransaction(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RollbackTransaction(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.RollbackTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_GetPlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    *models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.GetPlanetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.GetPlanetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.GetPlanetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_SearchPlanetsByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.SearchPlanetsByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.SearchPlanetsByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.SearchPlanetsByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_fillFilms(t *testing.T) {
	type args struct {
		ctx context.Context
		p   *models.Planet
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.fillFilms(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.fillFilms() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_InsertPlanet(t *testing.T) {
	type args struct {
		ctx                 context.Context
		tx                  *sql.Tx
		readyToInsertPlanet *models.Planet
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.InsertPlanet(tt.args.ctx, tt.args.tx, tt.args.readyToInsertPlanet); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.InsertPlanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_UpdatePlanet(t *testing.T) {
	type args struct {
		ctx context.Context
		p   *models.Planet
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.UpdatePlanet(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.UpdatePlanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_basicInsertFilm(t *testing.T) {
	type args struct {
		ctx context.Context
		tx  *sql.Tx
		f   *models.Film
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.basicInsertFilm(tt.args.ctx, tt.args.tx, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.basicInsertFilm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_insertChildrenFilms(t *testing.T) {
	type args struct {
		ctx context.Context
		tx  *sql.Tx
		p   *models.Planet
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.insertChildrenFilms(tt.args.ctx, tt.args.tx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.insertChildrenFilms() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_assertfilmPlanetRelationshipExists(t *testing.T) {
	type args struct {
		ctx context.Context
		tx  *sql.Tx
		p   *models.Planet
		f   *models.Film
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.assertfilmPlanetRelationshipExists(tt.args.ctx, tt.args.tx, tt.args.p, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.assertfilmPlanetRelationshipExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_filmAloneExists(t *testing.T) {
	type args struct {
		ctx    context.Context
		tx     *sql.Tx
		idFilm int
	}
	tests := []struct {
		name string
		n    *mysqlDatabaseFinal
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.filmAloneExists(tt.args.ctx, tt.args.tx, tt.args.idFilm); got != tt.want {
				t.Errorf("mysqlDatabaseFinal.filmAloneExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_ListAllPlanets(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.ListAllPlanets(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.ListAllPlanets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.ListAllPlanets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_RemovePlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		tx  *sql.Tx
		id  int
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetById(tt.args.ctx, tt.args.tx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.RemovePlanetById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_RemovePlanetByExactName(t *testing.T) {
	type args struct {
		ctx       context.Context
		tx        *sql.Tx
		exactName string
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetByExactName(tt.args.ctx, tt.args.tx, tt.args.exactName); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.RemovePlanetByExactName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_emptyAndNameError(t *testing.T) {
	type args struct {
		name string
		err  error
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.emptyAndNameError(tt.args.name, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.emptyAndNameError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.emptyAndNameError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_emptyAndGenericError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.emptyAndGenericError(tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.emptyAndGenericError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.emptyAndGenericError() = %v, want %v", got, tt.want)
			}
		})
	}
}
