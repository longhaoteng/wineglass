// @author mr.long

package wineglass

import (
	"fmt"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	routers []Router
)

type Wineglass struct {
	Config *Config
	router *gin.Engine
}

func New(conf *Config) *Wineglass {
	w := &Wineglass{Config: conf}
	return w
}

func Default() *Wineglass {
	c := cors.DefaultConfig()
	c.AllowCredentials = true
	c.AllowOrigins = []string{"*"}

	return New(&Config{
		RunMode: ReleaseMode,
		Pprof:   &Pprof{Open: false},
		Middleware: &Middleware{
			Cors: &c,
		},
	})
}

func Routers(rs ...Router) {
	routers = append(routers, rs...)
}

func (w *Wineglass) ApiEngine() *gin.Engine {
	return w.router
}

// SetMode Used for default wineglass update run mode.
func (w *Wineglass) SetMode(mode string) {
	w.Config.RunMode = mode
}

func (w *Wineglass) setupRouter() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	w.router = router
}

func (w *Wineglass) init() *gin.Engine {
	w.setupRouter()

	apiCtl := new(API)
	apiCtl.Validator()

	// middleware
	if w.Config.Middleware != nil {
		middleware := w.Config.Middleware

		// cors
		if middleware.Cors != nil {
			w.router.Use(cors.New(*middleware.Cors))
		}
		// rate limit
		if middleware.Limiter != nil && middleware.Limiter.Enable {
			w.router.Use(middleware.RateLimiter())
		}
		// session
		if middleware.Session != nil && middleware.Session.Enable {
			w.router.Use(middleware.Sessions())
		}
		// auth
		if middleware.Session != nil && middleware.Session.Enable &&
			middleware.Authorize != nil && middleware.Authorize.Enable {
			w.router.Use(middleware.Authorizer())
		}

		if len(middleware.Middlewares) > 0 {
			w.router.Use(middleware.Middlewares...)
		}
	}
	// pprof
	if w.Config.Pprof != nil && w.Config.Pprof.Open {
		pprof.Register(w.router, w.Config.Pprof.Prefix)
	}

	// 404
	w.router.NoRoute(apiCtl.API404)

	fmt.Println(len(routers))
	for _, r := range routers {
		r.Router(w.router)
	}

	return w.router
}

func (w *Wineglass) Run(addr ...string) error {
	// set the number of CPU processor will be used
	runtime.GOMAXPROCS(runtime.NumCPU())

	gin.ForceConsoleColor()
	gin.SetMode(w.Config.RunMode)

	if w.router == nil {
		w.init()
	}

	if err := w.router.Run(addr...); err != nil {
		return err
	}

	return nil
}
