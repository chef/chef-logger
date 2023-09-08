package chef_logger

import (
	"os"
	"testing"
	// "regexp"
)

// BasicTestLogInit calls BasicLogger.LogInit() which should return nothing
func BasicTestLogInit(t *testing.T) {
	var log = Logger{}

    c := new(Configuration)
    c.LogToStdout = true
    c.LogToFile= true
    c.LogFileDirectory= "."
    c.LogFilePrefix= "sample"
    c.LogFileExtension= ".log"
    // c.LogFileMaxSizeMB= 10
    c.LogAsTelemetry= false
    // c.LogTelemetryURI= ""
    // c.LogTelemetryProtocol= ""
    c.LogTimestampInUTC= true
    c.LogLineFormat= StructuredText
    c.LogServerName, _ = os.Hostname()
    // c.LogServerIP= "1.1.1.1"
    c.LogProcessName= "courier"
    c.LogLevel= Information
    // c.LogImplementationHint= Zap

    // other way to define
    /*
    d := &Configuration{
        LogServerName:  "",
      } */
    err := log.Configure(*c)
        
        // IP - // could https://www.golinuxcloud.com/golang-get-ip-address/
    if err != nil {
        t.Fatalf(`BasicLogger.LogInit("") = %v, want "", error`, err)
    }
}

// BasicTestLogDebug calls BasicLogger.LogDebug() which should return nothing
func BasicTestLogDebug (t *testing.T) {
    var log = Logger{}
    _, err := log.LogDebug("hello world")
    if err != nil {
         t.Fatalf(`BasicLogger.LogDebug("") = %v, want "", error`, err)
    }
}

// BasicTestLogWarn calls BasicLogger.LogWarn() which should return nothing
func BasicTestLogWarn (t *testing.T) {
    var log = Logger{}
    _, err := log.LogWarn(" *** warn **** ")
    if err != nil {
         t.Fatalf(`BasicLogger.LogWarn("") = %v, want "", error`, err)
    }
}

// BasicTestLogError calls BasicLogger.LogError() which should return nothing
func BasicTestLogError (t *testing.T) {
	var log = Logger{}
    _, err := log.LogError(" bye-bye ")
    if err != nil {
         t.Fatalf(`BasicLogger.LogError("") = %v, want "", error`, err)
    }
}

// BasicTestLogClose calls BasicLogger.LogClose() which should return nothing
func BasicTestLogClose (t *testing.T) {
    var log = Logger{}
    err := log.LogClose()
    if err != nil {
         t.Fatalf(`BasicLogger.LogClose("") = %v, want "", error`, err)
    }
}
