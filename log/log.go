package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errorLog = NewLog("\033[31m[error]\033[0m\t")
	infoLog  = NewLog("\033[33m[info]\033[0m\t")
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

func NewLog(prefix string) *log.Logger {
	// io.Discard: 默认为 Disable 模式
	return log.New(io.Discard, prefix, log.LstdFlags|log.Lshortfile)
}

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	Disable = iota
	ErrorLevel
	InfoLevel
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	//for _, logger := range loggers {
	//	logger.SetOutput(io.Discard)
	//}
	if level >= ErrorLevel {
		errorLog.SetOutput(os.Stdout)
	}
	if level >= InfoLevel {
		infoLog.SetOutput(os.Stdout)
	}
}
