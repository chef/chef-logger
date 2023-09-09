// BasicLogger implements ILogger with the Go-standard module "log"

// standard format is described at https://chefio.atlassian.net/wiki/spaces/AR/pages/2631925766/Logging+and+telemetry+and+performance+metrics+standard

// multierror return
// 	var err = errors.New("not implemented")
//	err = errors.Join(err, errors.New("a second config problem"))

// Multiple errors can be wrapped with Go 1.20 errors.Join; avoid extra dependency on
// + https://github.com/hashicorp/go-multierror
// + https://github.com/uber-go/multierr

/*
machine-name (FQDN)

Application name will be optional (multiple apps may run on a single machine and tail to same 
	log instance), and version will only be noted in startup (e.g., “App XYZ version 3.5.1 starting up”)
 Changing log level dynamically may be supported but is not required.
Logs below X criticality will include file name and line number being used and may 
include stack tracing; they also may include un-masked data if the customer has indicated that is allowed.  These lower level logs will be used for troubleshooting specific scenarios.

how to:
method entry as debug/trace with parameters
parsed inputs reflected as informational
error catching - warn on null lists, error within try/catch panic/recover/return error value/defer
method exit as debug/trace with return value

Specific fields automatically put in the log by the Go component
user id GUID
tenant GUID
UTC timestamp

Tools to read logs:
Jaeger, Prometheus,
Datadog, splunk
NewRelic, Dynatrace

Logging is secure (not tamperable), will retain N events by default, 
can overwrite or split log files at certain size
*/

// use Go 1.20 Join() instead of .Wrap unless needed for backlevel Go apps - https://stackoverflow.com/questions/33470649/how-do-we-combine-multiple-error-strings-in-golang

// consider ImplementationHint passed into a factory to create different base logger types:
// + log

// TODO: need to add printf style variadics like https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085
// func Printf(format string, a ...interface{}) (n int, err error) {
//	or ..any
//	func variadicExample(s ...string) {
//		fmt.Printf(s) - not println
//	}
//

package chef_logger

import (
	"log"
	"errors"
	"fmt"
	"os"
)

type Logger struct {
	// internal copies
	logToStdout 			bool	
	logStdoutWriter			*log.Logger

	logToFile 				bool	
	logFileDirectory 		string	
	logFilePrefix	 		string	
	logFileExtension 		string	
	logFileIncrement 		int
	logFileName				string		// full path and filename, computed
	logFileHandle			*os.File
	logWriter				*log.Logger
	logFileMaxSizeMB		uint64

	logAsTelemetry			bool	
	logTelemetryURI			string	
	logTelemetryProtocol	string	

	logTimestampInUTC		bool
	logLineFormat			LogFormat

	logServerName			string
	logServerIP				string
	logProcessName			string 
	
	logLevel				LoggingLevel
	logImplementationHint	ImplementationHint

	// add [] of all UDF fields, and fields which if non-null get output
	// add [] of masked fields
	// add field for line level [warning] [info], etc...
	// pass in correlation/transaction id as a UDF metadata field

	unableToContinue			bool
}

const PrependString string = ""		// could be "INFO: " or other unstructured text

func (l Logger) Configure(c Configuration) error {
	l.unableToContinue = false
	defer l.checkAbleToContinue()
	
	// copy parameters to internal struct
	l.logToStdout 			= c.LogToStdout
	l.logStdoutWriter		= nil
	l.logToFile  			= c.LogToFile
	l.logFileDirectory 		= c.LogFileDirectory
	l.logFilePrefix 		= c.LogFilePrefix
	l.logFileExtension 		= c.LogFileExtension
	l.logFileIncrement 		= 1
	l.logFileName 			= ""
	l.logWriter 			= nil
	l.logFileHandle			= nil
	l.logFileMaxSizeMB 		= c.LogFileMaxSizeMB
	l.logAsTelemetry 		= c.LogAsTelemetry	
	l.logTelemetryURI 		= c.LogTelemetryURI	
	l.logTelemetryProtocol 	= c.LogTelemetryProtocol	
	l.logTimestampInUTC 	= c.LogTimestampInUTC
	l.logLineFormat 		= c.LogLineFormat
	l.logServerName 		= c.LogServerName
	l.logServerIP 			= c.LogServerIP
	l.logProcessName 		= c.LogProcessName
	l.logLevel 				= c.LogLevel
	l.logImplementationHint = c.LogImplementationHint
	
	// log can go to multiple outputs - any of stdout, file or telemetry
	if(c.LogToStdout) {
		if(c.LogTimestampInUTC) {
			l.logStdoutWriter = log.New(os.Stdout, PrependString, log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
		} else {
			l.logStdoutWriter = log.New(os.Stdout, PrependString, log.Ldate|log.Ltime|log.Lmicroseconds)
		}
		/* options: 
			log.Ldate
			log.Ltime
			log.Lmicroseconds
			log.Llongfile - filename where log comes from, will always be this file so do not use
			log.Lshortfile 
			log.LUTC - use UTC
			log.Lmsgprefix - use the prefix
			log.LstdFlags
		*/
		if(l.logStdoutWriter == nil) {
			log.Fatal(fmt.Errorf("%s", "could not create logStdoutWriter"))
		}
	}

	// build the filename
	l.logFileName = l.constructFilename()
	if(c.LogToFile && l.logFileName != "") {
		logFileHandle, err := os.OpenFile(l.logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		// do not do os.Create
		if err != nil {
			log.Fatal(err)
		}
	
		l.logFileHandle = logFileHandle
		l.logWriter = log.New(l.logFileHandle, PrependString, log.Ldate|log.Ltime)
		if(l.logWriter == nil) {
			log.Fatal(fmt.Errorf("%s at %s", "could not create logWriter", l.logFileName))
		}
		// other loggers do not have a specific prepend field, for basic Go log, we can create 
		// different stremas to same file to differentiate warnings, errors, etc... per
		// 		WarningLogger *log.Logger
		//		InfoLogger    *log.Logger
		//		ErrorLogger   *log.Logger
		//
		// 		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		//		WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		// 		...
		
	}

	if(c.LogAsTelemetry) {
		return errors.New("telemetry configuration not implemented")
	}

	//if(err != nil) {
	//	l.unableToContinue = true
//}

	return nil
}

func (l Logger) TestConfiguration() (bool, error) {
	// validate configuration... can write log directory, reach telem endpoint, etc.
	//defer checkAbleToContinue()
	
	// if (err != nil) {
	// 	imminentPanic = true
	// }

	return false, errors.New("not implemented")
}

func (l Logger) ChangeConfiguration(Configuration) error {
	// might reset log.SetOutput(logfile)
	return errors.New("not implemented")
}

func (l Logger) ClearMaskedFields() error {
	return errors.New("not implemented")
}

func (l Logger) AddMaskedField(MaskedField) error {
	return errors.New("not implemented")
}

func (l Logger) RemoveMaskedField(MaskedField) error {
	return errors.New("not implemented")
}

// Log methods typically return the actual log line output (with substitutions made but no masking)
func (l Logger) Log(s string, lvl LoggingLevel) (string, error) {
	out := l.constructOutputLine(s, true)

	if(l.logLevel >= lvl) {
		if(l.logToStdout && l.logStdoutWriter != nil) {
			l.logStdoutWriter.Println(out)
			// do we do colors for {error, critical}
		}
		
		if(l.logToFile && l.logWriter != nil) {
			l.logWriter.Println(out)
			// do we Panicln or Fatalln for {error, critical}?
		}

		if(l.logAsTelemetry) {
			return "", errors.New("telemetry output not implemented")
		}
	}
	return "", nil
}

func (l Logger) LogTrace(s string) (string, error) {
		return l.Log(s, Trace)
}

func (l Logger) LogDebug(s string) (string, error) {
	return l.Log(s, Debug)
}

func (l Logger) LogInfo(s string) (string, error) {
	return l.Log(s, Information)
}

func (l Logger) LogWarn(s string) (string, error) {
	return l.Log(s, Warning)
}

func (l Logger) LogError(s string) (string, error) {
	return l.Log(s, Error)
}

func (l Logger) LogCritical(s string) (string, error) {
	return l.Log(s, Critical)
}

func (l Logger) LogClose() error {
	var err error
	if(l.logFileHandle != nil) {
		err = l.logFileHandle.Close()
	}
	return err
}

func (l Logger) constructFilename() string {
	return fmt.Sprintf("%s%s%s%04d%s%s", l.logFileDirectory, 
		"\\", l.logFilePrefix, l.logFileIncrement, ".", l.logFileExtension)
}

// exists returns whether the given file or directory exists
/*
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}
*/

// if a method hits an error, it's likely we can't output anything;
// we will try to printf and then return an error 
func (l Logger) checkAbleToContinue() {
	if(l.unableToContinue) {
		fmt.Println(PanicMessage)
	}
}

func (l Logger) constructOutputLine(msg string, performMasking bool, UDF ...UserDefinedField) (string) {
	var s string
	// prepend stuff
	// message
	// add user-defined fields, masked or not
	// add required fields
	return s
}