package chef_logger

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	// "regexp"
)


func TestLogInit(t *testing.T) {
	var log = Logger{}

	c := new(Configuration)
	c.LogToStdout = true
	c.LogToFile = true
	c.LogFileDirectory = "."
	c.LogFilePrefix = "sample"
	c.LogFileExtension = ".log"
	// c.LogFileMaxSizeMB= 10
	c.LogAsTelemetry = false
	// c.LogTelemetryURI= ""
	// c.LogTelemetryProtocol= ""
	c.LogTimestampInUTC = true
	c.LogLineFormat = StructuredText
	c.LogServerName, _ = os.Hostname()
	// c.LogServerIP= "1.1.1.1"
	c.LogProcessName = "courier"
	c.LogLevel = Information
	// c.LogImplementationHint= Zap

	err := log.Configure(*c)
	assert.NoError(t, err)
}

func TestLogDebug(t *testing.T) {
	var log = Logger{}
	err := log.Configure(Configuration{
		LogToStdout: true,
		LogLevel:    Debug,
	})
	assert.NoError(t, err)
	_, err = log.LogDebug("hello world")
	assert.NoError(t, err)
}

func TestLogWarn(t *testing.T) {
	var log = Logger{}
	err := log.Configure(Configuration{
		LogToStdout: true,
		LogLevel:    Warning,
	})
	assert.NoError(t, err)
	_, err = log.LogWarn(" *** warn **** ")
	assert.NoError(t, err)
}

func TestLogError(t *testing.T) {
	var log = Logger{}
	err := log.Configure(Configuration{
		LogToStdout: true,
		LogLevel:    Error,
	})
	assert.NoError(t, err)
	_, err = log.LogError(" bye-bye ")
	assert.NoError(t, err)
}

func TestLogClose(t *testing.T) {
	var log = Logger{}
	err := log.LogClose()
	assert.NoError(t, err)
}
