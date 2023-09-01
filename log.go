package logger

import (
	"fmt"
	"time"
)

var Version string = "1.0"

func Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}
