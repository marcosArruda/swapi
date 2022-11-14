package logs

import (
	"context"
	"fmt"

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
	f.Info(ctx, "Staring LogsService")
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
func (f *logsServiceFinal) Info(ctx context.Context, s string) {
	fmt.Println("(INFO) " + s)
}
func (f *logsServiceFinal) Warn(ctx context.Context, s string) {
	fmt.Println("(WARN) " + s)
}
func (f *logsServiceFinal) Error(ctx context.Context, s string) {
	fmt.Println("(ERROR) " + s)
}
func (f *logsServiceFinal) Debug(ctx context.Context, s string) {
	fmt.Println("(DEBUG) " + s)
}
