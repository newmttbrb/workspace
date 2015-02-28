package logger

import (
	"log"
	"os"
)

type Logger struct {
	name string
	file *os.File
}

func LoggerFactory(name string) (pconn *Logger) {
	return &Logger{name, nil}
}

func (l *Logger) Start() (err error) {
	l.file, err = os.Create(l.name)
	if err == nil {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.SetOutput(l.file)
	}
	return
}

func (l *Logger) Stop() {
	if l.file != nil {
		l.file.Close()
	}
}
