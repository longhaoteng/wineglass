package main

import (
	"log"

	"github.com/longhaoteng/wineglass"
	_ "github.com/longhaoteng/wineglass/_examples/api"
)

func main() {
	w := wineglass.Default()
	w.SetMode(wineglass.DebugMode)

	// defined port
	// w.Run(fmt.Sprintf(":%d", 9999))

	if err := w.Run(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
