package kgo_logger

import (
	"io"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
)

//Default file name for the logger
const (
	DefaultFileName = "app.log"
)

//Loggable type inplemented for all the custom projects
type Loggable interface {
	InfoF(format string, args ...interface{})
	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
	FatalF(format string, args ...interface{})
	PrintF(format string, args ...interface{})
}

//Logger is an adapter type around logrus.
type Logger struct {
	textLogger *logrus.Logger
}

//TODO: implement log rotation

//NewLogger returns a new logger creating a new file in the process
func NewLogger(level string, logDir string, logFile string) (*Logger, error) {
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	defaultfilename := defaultfileName(logFile)
	file, err := createLogFile(logDir, defaultfilename)
	if err != nil {
		return nil, err
	}

	logger := &logrus.Logger{
		Level:     parsedLevel,
		Out:       file,
		Formatter: new(logrus.TextFormatter),
	}
	return &Logger{
		textLogger: logger,
	}, nil
}

//InfoF prints the log as an info with the particlar formating
func (l *Logger) InfoF(format string, args ...interface{}) {
	l.textLogger.Infof(format, args...)
}

//Error prints the log as an error field with no formatting
func (l *Logger) Error(args ...interface{}) {
	l.textLogger.Error(args...)
}

//ErrorF prints the log as an error with formatting
func (l *Logger) ErrorF(format string, args ...interface{}) {
	l.textLogger.Errorf(format, args...)
}

//FatalF prints the log with a particual format and quits the application.
func (l *Logger) FatalF(format string, args ...interface{}) {
	l.textLogger.Fatalf(format, args...)
}

//PrintF prints the log with particlar formatting
func (l *Logger) PrintF(format string, args ...interface{}) {
	l.textLogger.Printf(format, args)
}

//internal helper functions
func createLogFile(logDir string, logFile string) (io.Writer, error) {
	if ok, err := dirExists(logDir); !ok {
		return nil, err
	}

	file, err := os.Create(path.Join(logDir, logFile))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func dirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func defaultfileName(file string) string {
	if file == "" {
		return DefaultFileName
	}
	return file
}
