package server

import (
	"net/http"

	"github.com/saumyan/payment_gateway/internals/handler/health"
)


func RegisterRoutes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.Handler)
	return mux

}