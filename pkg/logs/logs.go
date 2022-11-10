package logs

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/services"
)

type (
	logsServiceFinal struct {
		sm services.ServiceManager
	}
)

func NewLogsService(ctx context.Context) services.LogsService {
	return &logsServiceFinal{}
}

func (f *logsServiceFinal) Start(ctx context.Context) error {
	return nil
}
func (f *logsServiceFinal) Close(ctx context.Context) error {
	return nil
}
func (f *logsServiceFinal) Healthy(ctx context.Context) error {
	return nil
}
func (f *logsServiceFinal) WithServiceManager(sm services.ServiceManager) services.LogsService {
	f.sm = sm
	return f
}
func (f *logsServiceFinal) ServiceManager() services.ServiceManager {
	return f.sm
}
func (f *logsServiceFinal) Info(s string)  {}
func (f *logsServiceFinal) Warn(s string)  {}
func (f *logsServiceFinal) Error(s string) {}
func (f *logsServiceFinal) Debug(s string) {}
