// BasicLogger implements ILogger iwt the Go-standard module "log"

package chef_logger

import (
	"log"
)

type BasicLogger struct {
	// distance int - any internal variables
}

func (l BasicLogger) LogInit(s string) error {
	log.Println("new log entry ", s)
	return nil
}

func (l BasicLogger) LogDebug(s string) error {
	log.Println(s)
	return nil
}

func (l BasicLogger) LogWarn(s string) error {
	log.Panicln("panic!")
	return nil
}

func (l BasicLogger) LogError(s string) error {
	log.Fatalln("bye, bye!")
	return nil
}

func (l BasicLogger) LogClose() error {
	return nil
}