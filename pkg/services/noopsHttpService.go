package services

import (
	"context"
)

type (
	noOpsHttpService struct {
		sm ServiceManager
	}
)

func NewHttpService(ctx context.Context) HttpService {
	return &noOpsHttpService{}
}

func (n *noOpsHttpService) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsHttpService) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsHttpService) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsHttpService) WithServiceManager(sm ServiceManager) HttpService {
	n.sm = sm
	return n
}

func (n *noOpsHttpService) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsHttpService) WithDatabase(db Database) HttpService {
	return n
}

func (n *noOpsHttpService) StartListening(ctx context.Context) error {
	return nil
}
