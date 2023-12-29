package docker_exposer

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	l, _ := zap.NewDevelopment(zap.AddCaller())
	log = l.Sugar()
}

func DefaultLogger() *zap.SugaredLogger {
	return log
}
