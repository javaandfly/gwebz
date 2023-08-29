package gwebz

import (
	"fmt"
	"testing"

	"go.uber.org/fx"
)

type HTTPServer struct {
	App string
	Cfg string
}

type Cfg struct {
	Name  string `mapstructure:"name"`
	Param struct {
		Age int64 `mapstructure:"age"`
	} `mapstructure:"param"`
}

func TestIOC(t *testing.T) {
	cfgPojo := &Cfg{}

	err := ReadConfig("config_test.yaml", cfgPojo)
	if err != nil {
		panic(err)
	}
	// RunContainerIOC(func() {}, NewHTTPServer, NewServer)
	fx.New(fx.Provide(NewHTTPServer,)).Run()
}

func NewHTTPServer(cfg *Cfg) *HTTPServer {
	fmt.Println("Excuting NewHTTPServer")
	ret := &HTTPServer{
		App: cfg.Name,
		Cfg: cfg.Name,
	}
	return ret
}

type Server struct {
	HttpS *HTTPServer
}

func NewServer(app *HTTPServer, lc fx.Lifecycle) *Server {
	fmt.Println("Excuting NewServer")
	srv := &Server{
		HttpS: app,
	}
	// lc.Append(fx.Hook{
	// 	OnStart: func(ctx context.Context) error {
	// 		go func() {
	// 			fmt.Println(111)
	// 		}()
	// 		return nil
	// 	},
	// 	OnStop: func(ctx context.Context) error {

	// 		fmt.Println(1)
	// 		return nil
	// 	},
	// })
	return srv
}
