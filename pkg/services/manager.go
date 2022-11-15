package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/peterhellberg/swapi"
)

var (
	EmptyPlanetSlice = make([]*models.Planet, 0)
)

type (
	GenericService interface {
		Start(ctx context.Context) error
		Close(ctx context.Context) error
		Healthy(ctx context.Context) error
	}

	LogsService interface {
		GenericService
		WithServiceManager(sm ServiceManager) LogsService
		ServiceManager() ServiceManager
		Info(ctx context.Context, s string)
		Warn(ctx context.Context, s string)
		Error(ctx context.Context, s string)
		Debug(ctx context.Context, s string)
	}

	Database interface {
		GenericService
		WithServiceManager(sm ServiceManager) Database
		ServiceManager() ServiceManager
		GetPlanetById(ctx context.Context, id int) (*models.Planet, error)
		SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error)
		ListAllPlanets(ctx context.Context) ([]*models.Planet, error)
		InsertPlanet(ctx context.Context, p *models.Planet) error
		UpdatePlanet(ctx context.Context, p *models.Planet) error
		RemovePlanetById(ctx context.Context, id int) error
		RemovePlanetByExactName(ctx context.Context, exactName string) error
	}

	PersistenceService interface {
		GenericService
		WithServiceManager(sm ServiceManager) PersistenceService
		ServiceManager() ServiceManager
		UpsertPlanet(ctx context.Context, p *models.Planet) error
		GetPlanetById(ctx context.Context, id int) (*models.Planet, error)
		SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error)
		ListAllPlanets(ctx context.Context) ([]*models.Planet, error)
		RemovePlanetById(ctx context.Context, id int) error
		RemovePlanetByExactName(ctx context.Context, exactName string) error
	}

	SwApiService interface {
		GenericService
		WithServiceManager(sm ServiceManager) SwApiService
		ServiceManager() ServiceManager
		GetPlanetById(ctx context.Context, id int) (*swapi.Planet, error)
		PutOnline()
		PutOffline()
		IsOnline() bool
	}

	PlanetFinderService interface {
		GenericService
		WithServiceManager(sm ServiceManager) PlanetFinderService
		ServiceManager() ServiceManager
		GetPlanetById(ctx context.Context, id int) (*models.Planet, error)
		SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error)
		ListAllPlanets(ctx context.Context) ([]*models.Planet, error)
		RemovePlanetById(ctx context.Context, id int) error
		RemovePlanetByExactName(ctx context.Context, exactName string) error
	}

	HttpService interface {
		GenericService
		WithServiceManager(sm ServiceManager) HttpService
		ServiceManager() ServiceManager
		GetPlanetById(c *gin.Context)
		SearchPlanetsByName(c *gin.Context)
		ListAllPlanets(c *gin.Context)
		RemovePlanetById(c *gin.Context)
		RemovePlanetByExactName(c *gin.Context)
	}

	ServiceManager interface {
		GenericService
		WithLogsService(ls LogsService) ServiceManager
		LogsService() LogsService
		WithDatabase(db Database) ServiceManager
		Database() Database
		WithPersistenceService(p PersistenceService) ServiceManager
		PersistenceService() PersistenceService
		WithSwApiService(sw SwApiService) ServiceManager
		SwApiService() SwApiService
		WithPlanetFinderService(p PlanetFinderService) ServiceManager
		PlanetFinderService() PlanetFinderService
		WithHttpService(h HttpService) ServiceManager
		HttpService() HttpService
	}

	serviceManagerFinal struct {
		logsService         LogsService
		database            Database
		persistenceService  PersistenceService
		swApiService        SwApiService
		planetFinderService PlanetFinderService
		httpService         HttpService
	}
)

func NewManager(ctx context.Context) ServiceManager {
	return &serviceManagerFinal{
		logsService:         NewNoOpsLogsService(ctx),
		database:            NewNoOpsDatabase(ctx),
		persistenceService:  NewNoOpsPersistenceService(ctx),
		swApiService:        NewNoOpsSwService(ctx),
		planetFinderService: NewNoOpsPlanetFinderService(ctx),
		httpService:         NewNoOpsHttpService(ctx),
	}
}

func (m *serviceManagerFinal) Start(ctx context.Context) error {
	if err := m.logsService.Start(ctx); err != nil {
		return err
	}

	if err := m.database.Start(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.persistenceService.Start(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.swApiService.Start(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.planetFinderService.Start(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.httpService.Start(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (m *serviceManagerFinal) Close(ctx context.Context) error {
	if err := m.logsService.Close(ctx); err != nil {
		return err
	}

	if err := m.persistenceService.Close(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.swApiService.Close(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.httpService.Close(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (m *serviceManagerFinal) Healthy(ctx context.Context) error {
	if err := m.logsService.Healthy(ctx); err != nil {
		return err
	}

	if err := m.persistenceService.Healthy(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.swApiService.Healthy(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	if err := m.httpService.Healthy(ctx); err != nil {
		m.logsService.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (m *serviceManagerFinal) WithLogsService(ls LogsService) ServiceManager {
	m.logsService = ls.WithServiceManager(m)
	return m
}

func (m *serviceManagerFinal) LogsService() LogsService {
	return m.logsService
}

func (m *serviceManagerFinal) WithHttpService(h HttpService) ServiceManager {
	m.httpService = h.WithServiceManager(m)
	return m
}

func (m *serviceManagerFinal) HttpService() HttpService {
	return m.httpService
}

func (m *serviceManagerFinal) WithPersistenceService(p PersistenceService) ServiceManager {
	m.persistenceService = p.WithServiceManager(m)
	return m
}

func (m *serviceManagerFinal) PersistenceService() PersistenceService {
	return m.persistenceService
}

func (m *serviceManagerFinal) WithDatabase(db Database) ServiceManager {
	m.LogsService().Info(context.Background(), "Doing WithDatabase")
	m.database = db.WithServiceManager(m)
	return m
}

func (m *serviceManagerFinal) Database() Database {
	return m.database
}

func (m *serviceManagerFinal) WithSwApiService(sw SwApiService) ServiceManager {
	m.swApiService = sw.WithServiceManager(m)
	return m
}

func (m *serviceManagerFinal) SwApiService() SwApiService {
	return m.swApiService
}

func (m *serviceManagerFinal) WithPlanetFinderService(p PlanetFinderService) ServiceManager {
	m.planetFinderService = p.WithServiceManager(m)
	return m
}
func (m *serviceManagerFinal) PlanetFinderService() PlanetFinderService {
	return m.planetFinderService
}
