package main

import (
	"io"
	"log"
)

var Logger *logger

type logger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func InitLogger(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Logger = &logger{
		log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
