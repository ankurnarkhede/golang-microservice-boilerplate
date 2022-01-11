package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// LoggingConfig stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default.
type LoggingConfig struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Get JSON encoder if JSON logs are enabled
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// InitialiseLogger : Initializes the zap logger.
func InitialiseLogger(config LoggingConfig) *zap.Logger {
	var cores []zapcore.Core

	// Configure standard output logging
	if config.EnableConsole {
		var consoleLevel zapcore.Level

		_ = consoleLevel.UnmarshalText([]byte(config.ConsoleLevel))
		consoleWriter := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat), consoleWriter, consoleLevel)
		cores = append(cores, core)
	}

	// Configure file logging
	if config.EnableFile {
		var fileLevel zapcore.Level

		_ = fileLevel.UnmarshalText([]byte(config.FileLevel))
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileLocation,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat), fileWriter, fileLevel)
		cores = append(cores, core)
	}

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the cores together.
	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	Logger = zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	defer Logger.Sync()
	Logger.Info("Initialized the logger", zap.String("loggerType", "zap"))

	return Logger
}

// GetLogger : Gets the logger instance.
func GetLogger() *zap.Logger {
	return Logger
}
