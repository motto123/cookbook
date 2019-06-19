package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func init() {
	InitLogger("./log/lottery.log", "debug")
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func InitLogger(path string, logLevel string) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    10, // megabytes(MB)
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})
	level := zap.DebugLevel
	if logLevel != "debug" {
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		//zap.DebugLevel,
		level,
	)
	Logger = zap.New(core, zap.AddCaller())
	//Logger.Sync()
	Sugar = Logger.Sugar()

}

func InitLogger1(path string) {
	atom := zap.NewAtomicLevel()
	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	Logger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),

		zapcore.Lock(os.Stdout),
		atom,
	), zap.AddCaller())

	atom.SetLevel(zap.DebugLevel)
	Sugar = Logger.Sugar()
}
