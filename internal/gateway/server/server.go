package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MiralSNk/go-micro-blog/internal/gateway/config"
)

const (
	read     = 10
	write    = 30
	idle     = 60
	shutdown = 10
)

type Server struct {
	cfg          *config.Config
	log          *slog.Logger
	handler      http.Handler
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	srv          *http.Server
}

// New собирает и возвращает сервер
func New(h http.Handler, cfg *config.Config, log *slog.Logger) *Server {
	return &Server{
		cfg:          cfg,
		log:          log,
		handler:      h,
		readTimeout:  read * time.Second,
		writeTimeout: write * time.Second,
		idleTimeout:  idle * time.Second,
	}
}

// Start запускает сервер
func (s *Server) Start() {
	s.srv = &http.Server{
		Addr:         s.cfg.Port,
		Handler:      s.handler,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		IdleTimeout:  s.idleTimeout,
	}

	s.graceShutdown()
}

// graceShutdown безопасно подготавливает к закрытию сервера
func (s *Server) graceShutdown() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s.log.Info("Запуск сервиса", slog.String("addr", s.cfg.Port))
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Error("Ошибка прослушивания", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	<-quit
	s.log.Info("Подготавливаем сервис к закрытию")

	ctx, cancel := context.WithTimeout(context.Background(), shutdown*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Error("shutdown error", slog.String("error", err.Error()))
	}
}
