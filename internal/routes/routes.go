package routes

import (
	"github.com/albertopformoso/inventory/internal/controller"
	"github.com/albertopformoso/inventory/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes interface {
	Init(e *echo.Echo)
}

type routes struct {
	controller controller.Controller
}

func New(controller controller.Controller) Routes {
	return &routes{controller: controller}
}

// Init configure server routes.
func (r *routes) Init(e *echo.Echo) {
	r.configMiddlewares(e)
	r.registerRoutes(e)
}

// Register the routes.
func (r *routes) registerRoutes(e *echo.Echo) {
	r.userRoutes(e)
	r.productRoutes(e)
}

// Setting up the middlewares for the server.
func (r *routes) configMiddlewares(e *echo.Echo) {
	e.Use(
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			Skipper:   middleware.DefaultSkipper,
			StackSize: 4 << 10, // 4KB,
			LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
				logger.New().Err(err).Msg("http: request recovered from panic")
				return nil
			},
		}),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:      true,
			LogStatus:   true,
			LogMethod:   true,
			LogRemoteIP: true,
			LogLatency:  true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logger.New().Info().
					Int("status", v.Status).
					Str("method", v.Method).
					Str("URI", v.URI).
					Str("IP", v.RemoteIP).
					Str("latency", v.Latency.String()).
					Msg("request")
				return nil
			},
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://127.0.0.1:8080"},
			AllowMethods:     []string{echo.POST},
			AllowHeaders:     []string{echo.HeaderContentType},
			AllowCredentials: true,
		}))
}
