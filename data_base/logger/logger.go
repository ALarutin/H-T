package logger

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Trace   *log.Logger
	Debug   *log.Logger
	Warning *log.Logger
	Fatal   *log.Logger
)

var (
	fileOut = os.Stdout
	fileErr = os.Stderr
)

func init() {
	Info = log.New(fileOut,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Error = log.New(fileErr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Trace = log.New(fileOut,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Debug = log.New(fileOut,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Warning = log.New(fileOut,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Fatal = log.New(fileErr,
		"FATAL: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

/*В любой пакет нужно испортировать ""github.com/go-park-mail-ru/2019_1_SleeplessNights/log""
log.<log>.Println("commit")*/
