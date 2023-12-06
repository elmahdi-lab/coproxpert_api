package logging

import (
	"ithumans.com/coproxpert/config"
	"os"
	"sync"
)

var (
	globalLogger AppLogger
	once         sync.Once
)

func GetLogger() AppLogger {
	once.Do(func() {
		if os.Getenv("ENV") == config.Development {
			globalLogger = &LocalLogger{}
		} else {
			globalLogger = &GoogleCloudLogger{}
		}
	})
	return globalLogger
}
