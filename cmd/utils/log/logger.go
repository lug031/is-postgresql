package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	logLevel := getLogLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout("02-01-2006 15:04:05.000")
	encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields: map[string]interface{}{
			"pid":      os.Getpid(),
			"hostname": getHostname(),
		},
	}

	var err error
	log, err = config.Build()
	if err != nil {
		panic("Falló al inicializar logger: " + err.Error())
	}

	defer log.Sync()
}

// Obtener instancia del logger
func GetLogger() *zap.Logger {
	return log
}

func getLogLevel() zapcore.Level {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		return zap.DebugLevel // pordefecto
	}

	logLevel, err := zapcore.ParseLevel(logLevelStr)
	if err != nil {
		return zap.InfoLevel // pordefecto
	}

	return logLevel
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic("Falló al obtener hostname: " + err.Error())
	}
	return hostname
}
