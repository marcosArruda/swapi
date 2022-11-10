package services

import (
	"context"
)

type (
	noOpsLogsService struct {
		sm ServiceManager
	}
)

func NewNoOpsLogsService(ctx context.Context) LogsService {
	return &noOpsLogsService{}
}

func (f *noOpsLogsService) Start(ctx context.Context) error {
	return nil
}
func (f *noOpsLogsService) Close(ctx context.Context) error {
	return nil
}
func (f *noOpsLogsService) Healthy(ctx context.Context) error {
	return nil
}
func (f *noOpsLogsService) WithServiceManager(sm ServiceManager) LogsService {
	f.sm = sm
	return f
}
func (f *noOpsLogsService) ServiceManager() ServiceManager {
	return f.sm
}
func (f *noOpsLogsService) Info(s string)  {}
func (f *noOpsLogsService) Warn(s string)  {}
func (f *noOpsLogsService) Error(s string) {}
func (f *noOpsLogsService) Debug(s string) {}
