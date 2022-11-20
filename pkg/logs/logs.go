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

var (
	AppNameKey      string = "AppName"
	AppVersionKey   string = "AppVersion"
	AppEnvKey       string = "AppEnv"
	AppName         string = "swapiapp"
	AppVersion      string = "1.0"
	AppEnv          string = "PROD"
	AppNameField    *zap.Field
	AppVersionField *zap.Field
	AppEnvField     *zap.Field
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
	name := zap.String(AppNameKey, ctx.Value(AppNameKey).(string))
	AppNameField = &name
	version := zap.String(AppVersionKey, ctx.Value(AppVersionKey).(string))
	AppVersionField = &version
	env := zap.String(AppEnvKey, ctx.Value(AppEnvKey).(string))
	AppEnvField = &env
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
	f.logger.Info(s, *AppEnvField, *AppNameField, *AppVersionField)
}
func (f *logsServiceFinal) Warn(ctx context.Context, s string) {
	f.logger.Warn(s, *AppEnvField, *AppNameField, *AppVersionField)
}
func (f *logsServiceFinal) Error(ctx context.Context, s string) {
	f.logger.Error(s, *AppEnvField, *AppNameField, *AppVersionField)
}
func (f *logsServiceFinal) Debug(ctx context.Context, s string) {
	f.logger.Debug(s, *AppEnvField, *AppNameField, *AppVersionField)
}
