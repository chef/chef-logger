package chef_logger

import (
    "testing"
    // "regexp"
)

// LogrusTestLogInit calls LogrusLogger.LogInit() which should return nothing
func LogrusTestLogInit(t *testing.T) {
	var log ILogger
	log = LogrusLogger{}
    err := log.LogInit("logrus")
    if err != nil {
         t.Fatalf(`LogrusLogger.LogInit("") = %v, want "", error`, err)
    }
}

// LogrusTestLogDebug calls LogrusLogger.LogDebug() which should return nothing
func LogrusTestLogDebug (t *testing.T) {
    var log ILogger
	log = LogrusLogger{}
    err := log.LogDebug("hello world")
    if err != nil {
         t.Fatalf(`LogrusLogger.LogDebug("") = %v, want "", error`, err)
    }
}

// LogrusTestLogWarn calls LogrusLogger.LogWarn() which should return nothing
func LogrusTestLogWarn (t *testing.T) {
    var log ILogger
	log = LogrusLogger{}
    err := log.LogWarn(" *** warn **** ")
    if err != nil {
         t.Fatalf(`LogrusLogger.LogWarn("") = %v, want "", error`, err)
    }
}

// LogrusTestLogError calls LogrusLogger.LogError() which should return nothing
func LogrusTestLogError (t *testing.T) {
    var log ILogger
	log = LogrusLogger{}
    err := log.LogError(" bye-bye ")
    if err != nil {
         t.Fatalf(`LogrusLogger.LogError("") = %v, want "", error`, err)
    }
}

// LogrusTestLogClose calls LogrusLogger.LogClose() which should return nothing
func LogrusTestLogClose (t *testing.T) {
    var log ILogger
	log = LogrusLogger{}
    err := log.LogClose()
    if err != nil {
         t.Fatalf(`LogrusLogger.LogClose("") = %v, want "", error`, err)
    }
}
