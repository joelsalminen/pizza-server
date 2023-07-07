package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Address string `envconfig:"PIZZA_SERVER_ADDRESS" default:":8080"`
}

type RestaurantHandler interface {
	HandleGetPizza(w http.ResponseWriter, r *http.Request, pizzaId string)
	HandlePostPizza(w http.ResponseWriter, r *http.Request)
	HandleGetMenus(w http.ResponseWriter, r *http.Request)
	HandleGetMenu(w http.ResponseWriter, r *http.Request, menuId string)
}

func NewFromEnv(handler RestaurantHandler) *http.Server {
	var config ServerConfig
	err := envconfig.Process("", &config)
	if err != nil {
		panic("SOS: config failed to load")
	}

	return New(handler, config)
}

func New(handler RestaurantHandler, config ServerConfig) *http.Server {
	r := chi.NewRouter()
	r.Get("/pizza/{id}", func(w http.ResponseWriter, r *http.Request) {
		pizzaId := chi.URLParam(r, "id")
		handler.HandleGetPizza(w, r, pizzaId)
	})
	r.Post("/pizza", func(w http.ResponseWriter, r *http.Request) {
		handler.HandlePostPizza(w, r)
	})

	r.Get("/menus", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleGetMenus(w, r)
	})
	r.Get("/menu/{id}", func(w http.ResponseWriter, r *http.Request) {
		menuId := chi.URLParam(r, "id")
		handler.HandleGetMenu(w, r, menuId)
	})

	srv := &http.Server{
		Addr:    config.Address,
		Handler: r,
	}

	return srv
}
