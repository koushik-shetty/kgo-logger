package kgologger_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	logger "github.com/koushik-shetty/kgologger"
)

func readFileContents(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}

func removeFile(t *testing.T, file string) {
	err := os.Remove(file)
	assert.NoError(t, err, "failed to remove file")
}

func TestNewLoggerCreatesAFile(t *testing.T) {
	logDir := os.TempDir()
	logFile := logger.DefaultFileName
	log, err := logger.NewLogger("info", logDir, logFile)

	assert.NotNil(t, log, "expected logger to be created")
	assert.NoError(t, err, "Expected new logger not to fail")

	filename := path.Join(logDir, logFile)
	_, err = os.Stat(filename)

	assert.True(t, !os.IsNotExist(err), "Expected file to be created")
	removeFile(t, filename)
}

func TestLoggerWritesEntries(t *testing.T) {
	logDir := os.TempDir()
	logFile := logger.DefaultFileName
	log, err := logger.NewLogger("info", logDir, logFile)

	assert.NoError(t, err, "Expected new logger not to fail")

	filename := path.Join(logDir, logFile)

	log.ErrorF("TestLog: %s", "testdata")

	fileContents, err := readFileContents(filename)
	assert.NoError(t, err, "Expected to read created file")

	assert.Contains(t, fileContents, "TestLog: testdata", "Expected file to contain log entry")
	removeFile(t, filename)
}

func TestLoggerWritesEntriesOfAppropriateLevel(t *testing.T) {
	logDir := os.TempDir()
	logFile := logger.DefaultFileName
	log, err := logger.NewLogger("info", logDir, logFile)

	assert.NoError(t, err, "Expected new logger not to fail")

	filename := path.Join(logDir, logFile)

	log.InfoF("TestLog: %s", "InfoLog")
	log.ErrorF("TestLog: %s", "ErrorLog")

	fileContents, err := readFileContents(filename)
	assert.NoError(t, err, "Expected to read created file")

	scanner := bufio.NewScanner(bytes.NewBuffer([]byte(fileContents)))
	if scanner.Scan() {
		line := scanner.Text()
		assert.Contains(t, line, "INFO")
	}
	if scanner.Scan() {
		assert.Contains(t, scanner.Text(), "ERRO")
	}
	assert.Contains(t, fileContents, "TestLog: InfoLog", "Expected file to contain log entry")
	assert.Contains(t, fileContents, "TestLog: ErrorLog", "Expected file to contain log entry")
	removeFile(t, filename)
}

func TestBlankLoggerImplementsLoggableInterface(t *testing.T) {
	blankLogger := logger.NewBlankLogger()

	var loggable logger.Loggable = blankLogger

	assert.NotNil(t, loggable, "Expected BlankLogger to implement Loggable interface")
}
