package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pizza-api/models"
)

type RestaurantHandler struct {
	restaurant Restaurant
}

type Restaurant interface {
	FindPizza(pizzaId string) (models.Pizza, error)
	AddPizza(models.Pizza) error
	GetMenus() map[string]models.Menu
	FindMenu(menuId string) (models.Menu, error)
}

func New(rest Restaurant) *RestaurantHandler {
	pSrv := &RestaurantHandler{restaurant: rest}

	return pSrv
}

func (h *RestaurantHandler) HandleGetPizza(w http.ResponseWriter, r *http.Request, pizzaId string) {
	foundPizza, err := h.restaurant.FindPizza(pizzaId)
	if err != nil {
		writeError(w, 404, err.Error())

		return
	}

	response, err := json.Marshal(foundPizza)

	if err != nil {
		writeError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *RestaurantHandler) HandlePostPizza(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, 500, err.Error())
		return
	}

	pizza := &models.Pizza{}

	err = json.Unmarshal(b, pizza)
	if err != nil {
		writeError(w, 500, err.Error())
		return
	}

	h.restaurant.AddPizza(*pizza)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (h *RestaurantHandler) HandleGetMenus(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(h.restaurant.GetMenus())

	if err != nil {
		writeError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *RestaurantHandler) HandleGetMenu(w http.ResponseWriter, r *http.Request, menuId string) {
	foundMenu, err := h.restaurant.FindMenu(menuId)
	if err != nil {
		writeError(w, 404, err.Error())

		return
	}

	response, err := json.Marshal(foundMenu)

	if err != nil {
		writeError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, msg string) {
	resp := errorResponse{Error: msg}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("log: marshaling error failed")
	}

	w.WriteHeader(status)
	w.Write(respJSON)
	return
}
