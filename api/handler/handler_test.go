package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"pizza-api/api/handler"
	"pizza-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestaurantHandler_HandleGetPizza(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/pizza/1", nil)
	w := httptest.NewRecorder()

	expected := models.Pizza{Id: "666", Name: "Hellfire"}
	s := &mockStorage{expected: expected}

	h := handler.New(s)
	pizzaId := "1"

	h.HandleGetPizza(w, req, pizzaId)

	resp := w.Result()
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	actual := &models.Pizza{}
	assert.NoError(t, json.Unmarshal(data, actual))

	assert.Equal(t, expected, *actual)
}

type mockStorage struct {
	expected models.Pizza
}

func (s *mockStorage) FindPizza(pizzaId string) (models.Pizza, error) {
	return s.expected, nil
}

func (s *mockStorage) AddPizza(p models.Pizza) error {
	return nil
}

func (s *mockStorage) GetMenus() map[string]models.Menu {
	return map[string]models.Menu{}
}

func (s *mockStorage) FindMenu(menuId string) (models.Menu, error) {
	return models.Menu{}, nil
}
