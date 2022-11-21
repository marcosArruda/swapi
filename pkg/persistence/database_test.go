package persistence

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

func NewManagerForTestsDatabase() (services.ServiceManager, context.Context) {
	asyncWorkChannel := make(chan func() error)
	stop := make(chan struct{})

	os.Setenv("DB_NAME", "dummyName")
	os.Setenv("DB_USER", "dummyUser")
	os.Setenv("DB_PASSWORD", "dummyPassword")
	os.Setenv("DB_HOSTPORT", "dummyHostPort")

	ctx := context.Background()
	ctx = context.WithValue(ctx, logs.AppEnvKey, "TESTS")
	ctx = context.WithValue(ctx, logs.AppNameKey, logs.AppName)
	ctx = context.WithValue(ctx, logs.AppVersionKey, logs.AppVersion)
	return services.NewManager(asyncWorkChannel, stop), ctx
}

func buildMock(t *testing.T, errorIn int) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	if errorIn == -2 {
		mock.ExpectClose()
	}

	expect := []*sqlmock.ExpectedExec{}
	expect = append(expect, mock.ExpectExec("CREATE TABLE IF NOT EXISTS film").WillReturnResult(sqlmock.NewResult(1, 1)))
	expect = append(expect, mock.ExpectExec("CREATE TABLE IF NOT EXISTS planet").WillReturnResult(sqlmock.NewResult(1, 1)))
	expect = append(expect, mock.ExpectExec("CREATE TABLE IF NOT EXISTS planet_film").WillReturnResult(sqlmock.NewResult(1, 1)))

	mock.ExpectQuery("SELECT VERSION").WillReturnRows(mock.NewRows([]string{"version"}).AddRow("1.0"))

	if errorIn >= 0 {
		expect[errorIn].WillReturnError(errors.New("some error"))
	}
	return db
}

func buildTransactionsMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_mysqlDatabaseFinal_buildConnection(t *testing.T) {
	type args struct {
		ctx context.Context
		db  *sql.DB
	}
	sm, ctx := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name:    "successMocked",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: ctx, db: db},
			wantErr: false,
		},
		{
			name:    "successPROD",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: ctx, db: nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.buildConnection(tt.args.ctx, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlDatabaseFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	sm, ctx := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()

	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: context.WithValue(ctx, "mockDb", buildMock(t, -1))},
			wantErr: false,
		},
		{
			name:    "errorFilm",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: context.WithValue(ctx, "mockDb", buildMock(t, 0))},
			wantErr: true,
		},
		{
			name:    "errorPlanet",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: context.WithValue(ctx, "mockDb", buildMock(t, 1))},
			wantErr: true,
		},
		{
			name:    "errorPlanetFilm",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: context.WithValue(ctx, "mockDb", buildMock(t, 2))},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
			defer tt.args.ctx.Value("mockDb").(*sql.DB).Close()
		})
	}
}

func Test_mysqlDatabaseFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	sm, ctx := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	dbService.Start(context.WithValue(ctx, "mockDb", buildMock(t, -2)))
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name: "success", //just success because Database.Close() just closes the connection
			// and is intermitent if the connection was not yet created.
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	dbService.Start(context.WithValue(ctx, "mockDb", buildMock(t, -1)))
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx},
			wantErr: false,
		},
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
	sm, _ := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	tests := []struct {
		name string
		n    *mysqlDatabaseFinal
		args args
		want services.Database
	}{
		{
			name: "success",
			n:    dbService.(*mysqlDatabaseFinal),
			args: args{sm: sm},
			want: dbService,
		},
		{
			name: "successNil",
			n:    dbService.(*mysqlDatabaseFinal),
			args: args{sm: nil},
			want: dbService,
		},
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
	sm, _ := NewManagerForTestsDatabase()
	tests := []struct {
		name string
		n    *mysqlDatabaseFinal
		want services.ServiceManager
	}{
		{
			name: "success",
			n:    sm.WithDatabase(NewDatabase()).Database().(*mysqlDatabaseFinal),
			want: sm,
		},
		{
			name: "successNil",
			n:    NewDatabase().(*mysqlDatabaseFinal),
			want: nil,
		},
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
	sm, ctx := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	db, mock := buildTransactionsMock(t)
	mock.ExpectBegin()
	ctx = context.WithValue(ctx, "mockDb", db)
	dbService.Start(ctx)
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: ctx},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.n.BeginTransaction(tt.args.ctx)
			if (err == nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.BeginTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
				return
			}
		})
	}
}

func Test_mysqlDatabaseFinal_CommitTransaction(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	sm, _ := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	db, mock := buildTransactionsMock(t)
	mock.ExpectBegin()
	mock.ExpectCommit()
	tx, _ := db.Begin()

	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{tx: tx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.CommitTransaction(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.CommitTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
				return
			}
		})
	}
}

func Test_mysqlDatabaseFinal_RollbackTransaction(t *testing.T) {
	type args struct {
		tx *sql.Tx
	}
	sm, _ := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	db, mock := buildTransactionsMock(t)
	mock.ExpectBegin()
	mock.ExpectRollback()
	tx, _ := db.Begin()
	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{tx: tx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RollbackTransaction(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.RollbackTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
				return
			}
		})
	}
}

func Test_mysqlDatabaseFinal_GetPlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	sm, ctx := NewManagerForTestsDatabase()
	dbService := sm.WithDatabase(NewDatabase()).Database()
	db, mock := buildTransactionsMock(t)

	basicPlanet := &models.Planet{
		Id:      1,
		Name:    "Terra",
		Climate: "tropical",
		Terrain: "terra",
		Films: []*models.Film{
			{
				Id:        1,
				Title:     "Filme da terra",
				EpisodeID: 1,
				Created:   "800 quintilhões de anos atras",
				Director:  "Único",
				URL:       "https://something.com/api/film/1/",
			},
		},
		URL: "https://something.com/api/planet/1/",
	}

	expPlanet := mock.ExpectQuery("FROM planet")
	expPlanet.WillReturnRows(sqlmock.NewRows([]string{"id", "name", "climate", "terrain", "url"}).
		FromCSVString("1,Terra,tropical,terra,https://something.com/api/planet/1/"))

	expFilm := mock.ExpectQuery("FROM film")
	expFilm.WillReturnRows(sqlmock.NewRows([]string{"id", "title", "episode_id", "created", "director", "url"}).
		FromCSVString("1,title,1,800 quintilhões de anos atras,Único,https://something.com/api/film/1/"))

	ctx = context.WithValue(ctx, "mockDb", db)
	dbService.Start(ctx)

	tests := []struct {
		name    string
		n       *mysqlDatabaseFinal
		args    args
		want    *models.Planet
		wantErr bool
	}{
		{
			name:    "success",
			n:       dbService.(*mysqlDatabaseFinal),
			args:    args{ctx: ctx, id: 1},
			want:    basicPlanet,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.GetPlanetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlDatabaseFinal.GetPlanetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !planetSuperficialDeepEqual(got, tt.want) {
				t.Errorf("mysqlDatabaseFinal.GetPlanetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func planetSuperficialDeepEqual(p1 *models.Planet, p2 *models.Planet) bool {
	return p1.Id == p2.Id && p1.Name == p2.Name && p1.Terrain == p2.Terrain && p1.Climate == p2.Climate
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
