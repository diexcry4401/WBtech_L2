package main

import (
	"fmt"
	"log"
)

// Transport Интерфейс созданных транспортов фабрикой.
type Transport interface {
	Use()
}

type Car struct{}

func (c *Car) Use() {
	fmt.Println("Used car")
}

type Truck struct{}

func (t *Truck) Use() {
	fmt.Println("Used truck")
}

type Ship struct{}

func (s *Ship) Use() {
	fmt.Println("Used ship")
}

// Factory Интерфейс для создания транспорта фабрикой.
type Factory interface {
	CreateTransport(transport int) Transport
}

type concreteFactory struct{}

// NewFactory Конструктор фабрики.
func NewFactory() Factory {
	return &concreteFactory{}
}

func (c *concreteFactory) CreateTransport(create int) Transport {
	var transport Transport
	switch create {
	case 1:
		transport = &Car{}
	case 2:
		transport = &Truck{}
	case 3:
		transport = &Ship{}
	default:
		log.Fatal("Create product fail")
	}
	return transport
}

func main() {
	factory := NewFactory()
	transport := factory.CreateTransport(1)
	transport.Use()
	transport = factory.CreateTransport(2)
	transport.Use()
	transport = factory.CreateTransport(3)
	transport.Use()
}
