package LogExt

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
)

func init() {
	fmt.Println("LogExt init")
}
func LogErr(arg0 interface{}, args ...interface{}) {
	log.LOGGER("Test").Info(arg0, args...)
}

func LogWarn(arg0 interface{}, args ...interface{}) {
	log.LOGGER("Test").Warn(arg0, args...)
}

func LogInfo(arg0 interface{}, args ...interface{}) {
	log.LOGGER("Test").Info(arg0, args...)
}

func LogDebug(arg0 interface{}, args ...interface{}) {
	log.LOGGER("Test").Debug(arg0, args...)
}
