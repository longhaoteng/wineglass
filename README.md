[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/longhaoteng/wineglass/blob/master/LICENSE)

## Wineglass

🍸🍹 Wineglass is minimalist scaffolding based on [gin](https://github.com/gin-gonic/gin) .

## Install
```shell
go get github.com/longhaoteng/wineglass
```

## Getting Started
```go
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
```

## [More examples](https://github.com/longhaoteng/wineglass/blob/master/_examples)

## License
[MIT License](https://github.com/longhaoteng/wineglass/blob/master/LICENSE)