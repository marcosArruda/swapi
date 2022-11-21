package services

import (
	"context"
	"fmt"
)

type (
	noOpsLogsService struct {
		sm ServiceManager
	}
)

func NewNoOpsLogsService() LogsService {
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
func (f *noOpsLogsService) Info(ctx context.Context, s string) {
	fmt.Println("(TESTS-INFO) " + s)
}
func (f *noOpsLogsService) Warn(ctx context.Context, s string) {
	fmt.Println("(TESTS-WARN) " + s)
}
func (f *noOpsLogsService) Error(ctx context.Context, s string) {
	fmt.Println("(TESTS-ERROR) " + s)
}
func (f *noOpsLogsService) Debug(ctx context.Context, s string) {
	fmt.Println("(TESTS-DEBUG) " + s)
}
