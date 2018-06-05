package lib

import (
	"github.com/golang/glog"
)

type DefaultCacheLogger struct{}

func (l DefaultCacheLogger) Infof(format string, args ...interface{}) {
	glog.Infof(format, args)
}

func (l DefaultCacheLogger) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args)
}
