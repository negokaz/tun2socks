package tun2socks

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FlowerWrong/tun2socks/util"
)

func (app *App) SignalHandler() *App {
	// signal handler
	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func(app *App) {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("Exit", s)
				util.Exit()
			default:
				log.Println("Other", s)
			}
		}
	}(app)

	return app
}
