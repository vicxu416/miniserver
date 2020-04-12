package handler

import (
	"github.com/XuVic/miniserver/server"
)

func Routes() *server.Mux {
	mux := server.NewMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/products", ProductsHandler)
	return mux
}
