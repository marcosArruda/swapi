package services

import (
	"context"
)

type (
	noOpsPersistenceService struct {
		sm ServiceManager
	}
)

func NewNoOpsPersistenceService(ctx context.Context) PersistenceService {
	return &noOpsPersistenceService{}
}

func (n *noOpsPersistenceService) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsPersistenceService) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsPersistenceService) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsPersistenceService) WithServiceManager(sm ServiceManager) PersistenceService {
	n.sm = sm
	return n
}

func (n *noOpsPersistenceService) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsPersistenceService) WithDatabase(db Database) PersistenceService {
	return n
}
func (n *noOpsPersistenceService) Insert() error {
	return nil
}
