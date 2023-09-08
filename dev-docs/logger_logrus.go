// LogrusLogger implements ILogger with the logrus module

package chef_logger

/* import (
  log "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	// distance int - internal variables
}

func (l LogrusLogger) LogInit(s string) error {
	log.WithFields(log.Fields{
		"animal": "walrus",
	  }).Info("A walrus appears", s)
	  
	  return nil
}

func (l LogrusLogger) LogDebug(s string) error {
	log.Println(s)
 
	return nil
}

func (l LogrusLogger) LogWarn(s string) error {
	log.Panicln("panic!")
 
	return nil
}

func (l LogrusLogger) LogError(s string) error {
	log.Fatalln("bye, bye!")
 
	return nil
}

func (l LogrusLogger) LogClose() error {
	return nil
}
*/