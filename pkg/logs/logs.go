package logs

import (
	"context"
	"log"

	"github.com/marcosArruda/swapi/pkg/services"
	"go.uber.org/zap"
)

type (
	logsServiceFinal struct {
		sm     services.ServiceManager
		logger *zap.Logger
	}
)

func NewLogsService(ctx context.Context) services.LogsService {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	return &logsServiceFinal{logger: logger}
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
	f.logger.Info(s)
	//fmt.Println("(INFO) " + s)
}
func (f *logsServiceFinal) Warn(ctx context.Context, s string) {
	f.logger.Warn(s)
	//fmt.Println("(WARN) " + s)
}
func (f *logsServiceFinal) Error(ctx context.Context, s string) {
	f.logger.Error(s)
	//fmt.Println("(ERROR) " + s)
}
func (f *logsServiceFinal) Debug(ctx context.Context, s string) {
	f.logger.Debug(s)
	//fmt.Println("(DEBUG) " + s)
}
