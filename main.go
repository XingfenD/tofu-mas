package main

import (
	"os"

	"github.com/XingfenD/tofu-mas/tofu_engine"
	"github.com/XingfenD/tofu-mas/tofu_gui"

	"github.com/sirupsen/logrus"
)

func main() {
	// 定义两个 channel
	println("Initializing channels...")
	guiToEngine := make(chan string)
	engineToGui := make(chan string)

	// 初始化 GUI 模块
	gui := NewGui(guiToEngine, engineToGui)

	// 初始化引擎模块
	engine := NewEngine(guiToEngine, engineToGui, initLogger(4))

	// 启动引擎
	go engine.Start()

	// 启动 GUI
	gui.Start()
}

func NewGui(guiToEngine chan<- string, engineToGui <-chan string) *tofu_gui.Gui {
	return &tofu_gui.Gui{
		GuiToEngine: guiToEngine,
		EngineToGui: engineToGui,
	}
}

// NewEngine 初始化引擎模块
func NewEngine(guiToEngine <-chan string, engineToGui chan<- string, logger logrus.Logger) *tofu_engine.Engine {
	return &tofu_engine.Engine{
		GuiToEngine: guiToEngine,
		EngineToGui: engineToGui,
		Logger:      logger,
	}
}

func initLogger(logLevel int) logrus.Logger {
	var logger logrus.Logger
	logger.Out = os.Stdout
	logger.Level = logrus.Level(logLevel)
	logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}
	logger.Info("Logger initialized")
	return logger
}
