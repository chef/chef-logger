// chef_log.ILogger is the base interface (set of methods and constants) for log implementations

// interface article https://go101.org/article/interface.html

package chef_logger

type ILogger interface {
	Configure(Configuration) error
	TestConfiguration()(bool, error)
	ChangeConfiguration(Configuration) error

	ClearMaskedFields() error
	AddMaskedField(MaskedField) error
	RemoveMaskedField(MaskedField) error

	// Log methods typically return the actual log line output (with substitutions made but no masking)
	LogTrace(string) (string, error)
	LogDebug(string) (string, error)
	LogInfo(string) (string, error)
	LogWarn(string) (string, error)
	LogError(string) (string, error)
	LogCritical(string) (string, error)
	
	LogClose() error
}

type Configuration struct {
	LogToStdout 			bool	`json:"stdout" example:"true"`
	LogToFile 				bool	`json:"fileout" example:"true"`
	LogFileDirectory 		string	`json:"filedirectory" example:"/var/log/chef"`
	LogFilePrefix	 		string	`json:"fileprefix" example:"automate"`
	LogFileExtension 		string	`json:"fileextension" example:".log"`
	LogFileMaxSizeMB		uint64	`json:"maxsizemb" example:"1024"`
	LogAsTelemetry			bool	`json:"telemetryout" example:"true"`
	LogTelemetryURI			string	`json:"telemetryuri" example:"telemetry.chef.io"`
	LogTelemetryProtocol	string	`json:"telemetryprotocol" example:"OTELv1"`
	LogTimestampInUTC		bool	`json:"utc" example:"true"`
	LogLineFormat			LogFormat	`json:"logformat" example:"json"`
	LogServerName			string	`json:"server" example:"myserver.myco.com"`
	LogServerIP				string	`json:"ip" example:"1.1.1.1"`
	LogProcessName			string 	`json:"process" example:"license-service"`
	LogLevel				LoggingLevel	`json:"loglevel" example:"Information"`
	LogImplementationHint	ImplementationHint	`json:"implhint" example:"Basic"`
}

type LogFormat string

const (
	Json 			LogFormat 	= "json"
	StructuredText	LogFormat	= "structured"
)

type LoggingLevel int

const (
	None 		LoggingLevel	= iota
	Trace 
	Debug 	
	Information 
	Warning	
	Error 
	Critical
)

func (ll LoggingLevel) ToString() string {
	switch ll {
	case None:
		return "none"
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Information:
		return "information"
	case Warning:
		return "warning"
	case Error:
		return "error"
	case Critical:
		return "critical"
	}
	return "unknown"
}
type ImplementationHint int

const (
	Basic = iota
	Logrus
	Zap
)

func (ih ImplementationHint) ToString() string {
	switch ih {
	case Basic:
		return "basic"
	case Logrus:
		return "logrus"
	case Zap:
		return "zap"
	}
	return "unknown"
}

type UserDefinedField struct {
	FieldName		string	`json:"field" example:"username"`
	FieldValue		string	`json:"value" example:"robert"`
}

type MaskedField struct {
	FieldName 		string	`json:"field" example:"Robert"`
	FieldRegex 		string	`json:"(.{1})(.*)(.{1})" example:"R****t"`
	ReplaceChar		rune	`json:"replace" example:"*"`
}

const PanicMessage string = "OMG"		// can have printf substitutions in implementation