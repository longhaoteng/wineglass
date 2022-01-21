package main

import (
	_ "github.com/longhaoteng/wineglass/_examples/api"
	_ "github.com/longhaoteng/wineglass/_examples/cron"
	_ "github.com/longhaoteng/wineglass/_examples/db"
	_ "github.com/longhaoteng/wineglass/_examples/handler"
	"github.com/longhaoteng/wineglass/server"
)

func main() {
	// Init server
	server.Init(
		server.Name("helloworld"),
		server.EnablePprof(),
		server.DisableDB(),
		server.DisableAuth(),
		server.DisableRedis(),
	)

	// Run server
	server.Run()
}
