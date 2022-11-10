package services

import (
	"context"
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
		Info(s string)
		Warn(s string)
		Error(s string)
		Debug(s string)
	}

	Database interface {
		GenericService
		WithServiceManager(sm ServiceManager) Database
		ServiceManager() ServiceManager
		Connect() error
	}

	PersistenceService interface {
		GenericService
		WithServiceManager(sm ServiceManager) PersistenceService
		ServiceManager() ServiceManager
		WithDatabase(db Database) PersistenceService
		Insert() error
	}

	GenericApiService interface {
		GenericService
		WithServiceManager(sm ServiceManager) GenericApiService
		ServiceManager() ServiceManager
		Get()
	}

	SwApiService interface {
		GenericService
		WithServiceManager(sm ServiceManager) SwApiService
		ServiceManager() ServiceManager
		Get()
	}

	HttpService interface {
		GenericService
		WithServiceManager(sm ServiceManager) HttpService
		ServiceManager() ServiceManager
		StartListening(ctx context.Context) error
	}

	ServiceManager interface {
		GenericService
		WithLogsService(ls LogsService) ServiceManager
		LogsService() LogsService
		WithPersistenceService(p PersistenceService) ServiceManager
		PersistenceService() PersistenceService
		WithHttpService(h HttpService) ServiceManager
		HttpService() HttpService

		WithSwApiService(sw SwApiService) ServiceManager
		SwApiService() SwApiService
	}

	serviceManagerFinal struct {
		logsService        LogsService
		persistenceService PersistenceService
		swApiService       SwApiService
		httpService        HttpService
	}
)

func NewManager(ctx context.Context) ServiceManager {
	return &serviceManagerFinal{}
}

func (m *serviceManagerFinal) Start(ctx context.Context) error {
	if err := m.logsService.Start(ctx); err != nil {
		return err
	}

	if err := m.persistenceService.Start(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	if err := m.swApiService.Start(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	if err := m.httpService.Start(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	return nil
}

func (m *serviceManagerFinal) Close(ctx context.Context) error {
	if err := m.logsService.Close(ctx); err != nil {
		return err
	}

	if err := m.persistenceService.Close(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	if err := m.swApiService.Close(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	if err := m.httpService.Close(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	return nil
}

func (m *serviceManagerFinal) Healthy(ctx context.Context) error {
	if err := m.logsService.Healthy(ctx); err != nil {
		return err
	}

	if err := m.persistenceService.Healthy(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	if err := m.swApiService.Healthy(ctx); err != nil {
		m.logsService.Error(err.Error())
		return err
	}

	if err := m.httpService.Healthy(ctx); err != nil {
		m.logsService.Error(err.Error())
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

func (m *serviceManagerFinal) WithSwApiService(sw SwApiService) ServiceManager {
	m.swApiService = sw.WithServiceManager(m)
	return m
}

func (m *serviceManagerFinal) SwApiService() SwApiService {
	return m.swApiService
}
