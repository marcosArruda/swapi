package services

import (
	"context"
)

type (
	noOpsGenericApiService struct {
		sm ServiceManager
	}
)

func NewNoOpsGenericApiService(ctx context.Context) GenericApiService {
	return &noOpsGenericApiService{}
}

func (n *noOpsGenericApiService) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsGenericApiService) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsGenericApiService) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsGenericApiService) WithServiceManager(sm ServiceManager) GenericApiService {
	n.sm = sm
	return n
}

func (n *noOpsGenericApiService) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsGenericApiService) WithDatabase(db Database) GenericApiService {
	return n
}
func (n *noOpsGenericApiService) Get() {}
