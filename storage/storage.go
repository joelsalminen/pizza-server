package storage

import (
	"fmt"
	"pizza-api/models"
	"sync"
)

type Storage struct {
	pizzas map[string]models.Pizza
	menus  map[string]models.Menu
	mu     sync.RWMutex
}

func New() *Storage {
	initialMenus := map[string]models.Menu{}
	menu1Pizzas := map[string]models.MenuItem{"1": &models.Pizza{Id: "1", Name: "Hawaii"}, "2": &models.Pizza{Id: "2", Name: "Dillinger"}, "3": &models.Pizza{Id: "3", Name: "Julia"}, "4": &models.Pizza{Id: "4", Name: "Tuna"}, "5": &models.Pizza{Id: "5", Name: "Potato"}}
	menu2Pizzas := map[string]models.MenuItem{"1": &models.Pizza{Id: "1", Name: "Hawaii"}, "2": &models.Pizza{Id: "2", Name: "Dillinger"}, "3": &models.Pizza{Id: "3", Name: "Julia"}, "4": &models.Pizza{Id: "4", Name: "Tuna"}}
	initialMenus["Lunch"] = models.Menu{Id: "Lunch", Items: menu1Pizzas, Name: "Lunch"}
	initialMenus["Dinner"] = models.Menu{Id: "Dinner", Items: menu2Pizzas, Name: "Dinner"}

	initialPizzas := map[string]models.Pizza{"1": {Id: "1", Name: "Hawaii"}, "2": {Id: "2", Name: "Dillinger"}, "3": {Id: "3", Name: "Julia"}, "4": {Id: "4", Name: "Tuna"}, "5": {Id: "5", Name: "Potato"}}

	storage := &Storage{menus: initialMenus, pizzas: initialPizzas}
	return storage
}

func (s *Storage) FindPizza(pizzaId string) (models.Pizza, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.pizzas[pizzaId]
	if !ok {
		return models.Pizza{}, fmt.Errorf("Not found :((")
	}

	return v, nil
}

func (s *Storage) AddPizza(p models.Pizza) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pizzas[p.Id] = p

	return nil
}

func (s *Storage) GetMenus() map[string]models.Menu {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.menus
}

func (s *Storage) FindMenu(menuId string) (models.Menu, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.menus[menuId]
	if !ok {
		return models.Menu{}, fmt.Errorf("Not found :((")
	}

	return v, nil
}
