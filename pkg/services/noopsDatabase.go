package services

import (
	"context"
)

type (
	noOpsDatabase struct {
		sm ServiceManager
	}
)

func NewNoOpsDatabase(ctx context.Context) Database {
	return &noOpsDatabase{}
}

func (n *noOpsDatabase) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsDatabase) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsDatabase) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsDatabase) WithServiceManager(sm ServiceManager) Database {
	n.sm = sm
	return n
}

func (n *noOpsDatabase) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsDatabase) Connect() error {
	return nil
}
