package chef_logger

import (
    "testing"
    // "regexp"
)

// BasicTestLogInit calls BasicLogger.LogInit() which should return nothing
func BasicTestLogInit(t *testing.T) {
	var log ILogger
	log = BasicLogger{}
    err := log.LogInit("basic")
    if err != nil {
         t.Fatalf(`BasicLogger.LogInit("") = %v, want "", error`, err)
    }
}

// BasicTestLogDebug calls BasicLogger.LogDebug() which should return nothing
func BasicTestLogDebug (t *testing.T) {
    var log ILogger
	log = BasicLogger{}
    err := log.LogDebug("hello world")
    if err != nil {
         t.Fatalf(`BasicLogger.LogDebug("") = %v, want "", error`, err)
    }
}

// BasicTestLogWarn calls BasicLogger.LogWarn() which should return nothing
func BasicTestLogWarn (t *testing.T) {
    var log ILogger
	log = BasicLogger{}
    err := log.LogWarn(" *** warn **** ")
    if err != nil {
         t.Fatalf(`BasicLogger.LogWarn("") = %v, want "", error`, err)
    }
}

// BasicTestLogError calls BasicLogger.LogError() which should return nothing
func BasicTestLogError (t *testing.T) {
    var log ILogger
	log = BasicLogger{}
    err := log.LogError(" bye-bye ")
    if err != nil {
         t.Fatalf(`BasicLogger.LogError("") = %v, want "", error`, err)
    }
}

// BasicTestLogClose calls BasicLogger.LogClose() which should return nothing
func BasicTestLogClose (t *testing.T) {
    var log ILogger
	log = BasicLogger{}
    err := log.LogClose()
    if err != nil {
         t.Fatalf(`BasicLogger.LogClose("") = %v, want "", error`, err)
    }
}
