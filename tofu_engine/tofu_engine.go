package tofu_engine

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Engine struct {
	GuiToEngine <-chan string
	EngineToGui chan<- string
	Logger      logrus.Logger
}

func (e *Engine) Start() {
	for input := range e.GuiToEngine {
		/* simulate the engine logic */
		e.Logger.Info("Engine received:", input)
		time.Sleep(1 * time.Second) /* simulate the time consuming operation */
		result := fmt.Sprintf("Processed: %s", input)
		e.EngineToGui <- result
	}
}
