// chef_logger provides an interface and several implementations of logging frameworks 
// to support server, CLI and agent requirements while emitting debug information in a 
// common format for processing by downstream tools
//
// chef_logger.ILogger is the interface set of methods all loggers need to adhere to
// the following implementations are provided:
// * basic - uses the Go standard "log" implementation
// * logrus - uses the XYZ implementation
// * zap - uses the XYZ implementation
//
// Code example:
// 
// references: https://gobyexample.com/interfaces 

package chef_logger

type ILogger interface {
	LogInit(string) error
	LogDebug(string) error
	LogWarn(string) error
	LogError(string) error
	LogClose() error
}