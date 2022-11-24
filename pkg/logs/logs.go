package logs

import (
	"context"
	"log"

	"github.com/marcosArruda/swapi/pkg/services"
	"go.uber.org/zap"
)

type (
	logsServiceFinal struct {
		sm              services.ServiceManager
		logger          *zap.Logger
		AppNameField    zap.Field
		AppVersionField zap.Field
		AppEnvField     zap.Field
	}
)

var (
	AppNameKey    string = "service"
	AppVersionKey string = "version"
	AppEnvKey     string = "env"
	AppName       string = "swapiapp"
	AppVersion    string = "1.0"
	AppEnv        string = "PROD"
)

func NewLogsService() services.LogsService {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	return &logsServiceFinal{logger: logger}
}

func (f *logsServiceFinal) Start(ctx context.Context) error {
	f.AppNameField = zap.String(AppNameKey, AppName)
	f.AppVersionField = zap.String(AppVersionKey, AppVersion)
	f.AppEnvField = zap.String(AppEnvKey, AppEnv)
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
	f.logger.Info(s, f.AppEnvField, f.AppNameField, f.AppVersionField)
}
func (f *logsServiceFinal) Warn(ctx context.Context, s string) {
	f.logger.Warn(s, f.AppEnvField, f.AppNameField, f.AppVersionField)
}
func (f *logsServiceFinal) Error(ctx context.Context, s string) {
	f.logger.Error(s, f.AppEnvField, f.AppNameField, f.AppVersionField)
}
func (f *logsServiceFinal) Debug(ctx context.Context, s string) {
	f.logger.Debug(s, f.AppEnvField, f.AppNameField, f.AppVersionField)
}
