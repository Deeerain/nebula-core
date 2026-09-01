package service

import (
	"log"
	"net/http"
)

type Service interface {
	Run()
}

type HTTPService interface {
	Service
	UseRoute(pattern string, handler http.HandlerFunc)
	UseHandler(pattern string, handler http.Handler)
}

type DefaultHTTPService struct {
	mux *http.ServeMux
}

func (dhs *DefaultHTTPService) Run() {
	if err := http.ListenAndServe("127.0.0.1:8080", dhs.mux); err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}
}

func (dhs *DefaultHTTPService) UseHandler(pattern string, handler http.Handler) {
	dhs.mux.Handle(pattern, handler)
}

func (dhs *DefaultHTTPService) UseRoute(pattern string, handler http.HandlerFunc) {
	dhs.mux.HandleFunc(pattern, handler)
}

func NewHTTPService() HTTPService {
	return &DefaultHTTPService{
		mux: http.NewServeMux(),
	}
}
