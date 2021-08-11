package webserver

import (
	"blueBot_go_webserver_echo/src/router"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	initWebServer sync.Once
	ws            *WebServer
)

type WebServer struct {
	router *echo.Echo
}

func Instance() *WebServer {
	initWebServer.Do(func() {
		ws = &WebServer{}
		ws.router = router.New()
	})
	return ws
}

func (w *WebServer) Start(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		log.Println(Starting HTTP server on port :8000)
		log.Fatal(w.router.Start(":8000"))
		
	}()
	wg.Wait()
}

