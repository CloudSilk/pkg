package main

import (
	"github.com/CloudSilk/pkg/utils/log"

	"github.com/sirupsen/logrus"
)

func main() {
	log.SetServiceName("log-example2")
	log.Info(nil, "hello logger!")
	log.Debug(nil, "debug")
	log.SetLevel(logrus.DebugLevel)
	log.Debug(nil, "debug")
	log.UseJSONFormatter()
	log.Info(nil, "hello logger!")
	log.Infof(nil, "test format:%s", "hello logger")
	// log.Panic(nil, "panic")
}
