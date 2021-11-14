package log

import (
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.Lshortfile|log.LstdFlags)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m", log.Lshortfile|log.LstdFlags)
	loggers  = []*log.Logger{errorLog, infoLog}
	//锁
	mu sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)
