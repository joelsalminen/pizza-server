package main

import (
	"pizza-api/api/handler"
	"pizza-api/api/server"
	"pizza-api/storage"
)

func main() {
	storage := storage.New()
	h := handler.New(storage)

	srv := server.NewFromEnv(h)
	srv.ListenAndServe()
}
