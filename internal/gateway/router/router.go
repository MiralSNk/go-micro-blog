package router

import (
	"log/slog"

	"github.com/MiralSNk/go-micro-blog/internal/gateway/handler/hello"
	loggermiddleware "github.com/MiralSNk/go-micro-blog/internal/gateway/middlewares/loggerMiddleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	chi.Router
	logger *slog.Logger
}

// New собирает и возвращает роутер для сервера
func New(log *slog.Logger) *Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(loggermiddleware.LoggerMiddleware(log))

	rt := &Router{Router: r}
	rt.registerRouters()

	return rt
}

// registerRouters регистрация всех маршрутов (приватный)
func (r *Router) registerRouters() {
	r.Get("/", hello.Handler)
}
