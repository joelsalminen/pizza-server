package models

import "fmt"

type Pizza struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p *Pizza) String() string {
	return fmt.Sprintf("%s-%d€", p.Name, p.Price)
}

func (p *Pizza) GetPrice() int {
	return p.Price
}

type MenuItem interface {
	GetPrice() int
	String() string
}

type Menu struct {
	Id    string              `json:"id"`
	Items map[string]MenuItem `json:"items"`
	Name  string              `json:"name"`
}

func (m *Menu) String() string {
	formatted := fmt.Sprintf("---%s---", m.Name)
	for _, item := range m.Items {
		formatted = fmt.Sprintf("%s\n%s", formatted, item)
	}
	return formatted
}
