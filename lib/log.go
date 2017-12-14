package lib

import (
	"fmt"
	"log"
)

// LogInfo - info message
func LogInfo(msg ...interface{}) {
	log.Println("[INFO]", fmt.Sprint(msg...))
}

// LogError - error message
func LogError(msg ...interface{}) {
	log.Println("[ERROR]", fmt.Sprint(msg...))
}
