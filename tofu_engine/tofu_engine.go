package tofu_engine

import (
    "fmt"
    "time"
)

type Engine struct {
    GuiToEngine <-chan string
    EngineToGui chan<- string
}

func (e *Engine) Start() {
    for input := range e.GuiToEngine {
        /* simulate the engine logic */
        fmt.Println("Engine received:", input)
        time.Sleep(1 * time.Second) /* simulate the time consuming operation */
        result := fmt.Sprintf("Processed: %s", input)
        e.EngineToGui <- result
    }
}
