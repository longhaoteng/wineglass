// @author mr.long

package main

import (
	"github.com/longhaoteng/wineglass"
	"github.com/longhaoteng/wineglass/_examples/api"
	"log"
)

func main() {
	w := wineglass.Default()
	w.SetMode(wineglass.DebugMode)

	w.Routers(
		&api.Ping{},
		&api.Hello{},
		&api.User{},
	)

	// defined port
	// w.Run(fmt.Sprintf(":%d", 9999))

	if err := w.Run(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
