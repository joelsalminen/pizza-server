package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"pizza-api/api/handler"
	"pizza-api/api/server"
	"pizza-api/models"
	"pizza-api/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createServer(t *testing.T, port string) (*http.Server, net.Listener) {
	storage := storage.New()
	h := handler.New(storage)

	srv := server.New(h, server.ServerConfig{Address: port})

	l, err := net.Listen("tcp", port)
	assert.NoError(t, err)
	return srv, l
}

func Test_API(t *testing.T) {
	port := ":8085"

	srv, l := createServer(t, port)

	var (
		c    http.Client
		resp *http.Response
		err  error
	)

	go func() {
		resp, err = c.Get(fmt.Sprintf("http://localhost%s/pizza/1", port))
		assert.NoError(t, err)
		srv.Close()
	}()

	srv.Serve(l)

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	assert.NoError(t, err)

	actual := &models.Pizza{}
	assert.NoError(t, json.Unmarshal(data, actual))

	assert.Equal(t, &models.Pizza{Id: "1", Name: "Hawaii"}, actual)
}
