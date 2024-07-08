package main

import "fmt"

// House представляет собой сложный объект, который мы хотим построить.
type House struct {
	Foundation string
	Walls      int
	Roof       string
	Doors      int
}

// HouseBuilder определяет интерфейс для построения частей House.
type HouseBuilder interface {
	SetFoundation()
	SetWalls()
	SetRoof()
	SetDoors()
	GetHouse() House
}

// ConcreteHouseBuilder представляет конкретного строителя, который строит дом из бетона.
type ConcreteHouseBuilder struct {
	house House
}

// NewConcreteHouseBuilder создает новый экземпляр ConcreteHouseBuilder.
func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{}
}

// SetFoundation устанавливает бетонный фундамент дома.
func (b *ConcreteHouseBuilder) SetFoundation() {
	b.house.Foundation = "Concrete Foundation"
}

// SetWalls устанавливает стены из бетона.
func (b *ConcreteHouseBuilder) SetWalls() {
	b.house.Walls = 4
}

// SetRoof устанавливает крышу из черепицы.
func (b *ConcreteHouseBuilder) SetRoof() {
	b.house.Roof = "Tile Roof"
}

// SetDoors устанавливает двери.
func (b *ConcreteHouseBuilder) SetDoors() {
	b.house.Doors = 1
}

// GetHouse возвращает построенный бетонный дом.
func (b *ConcreteHouseBuilder) GetHouse() House {
	return b.house
}

// WoodenHouseBuilder представляет конкретного строителя, который строит дом из дерева.
type WoodenHouseBuilder struct {
	house House
}

// NewWoodenHouseBuilder создает новый экземпляр WoodenHouseBuilder.
func NewWoodenHouseBuilder() *WoodenHouseBuilder {
	return &WoodenHouseBuilder{}
}

// SetFoundation устанавливает деревянный фундамент дома.
func (b *WoodenHouseBuilder) SetFoundation() {
	b.house.Foundation = "Wooden Foundation"
}

// SetWalls устанавливает стены.
func (b *WoodenHouseBuilder) SetWalls() {
	b.house.Walls = 6
}

// SetRoof устанавливает деревянную крышу.
func (b *WoodenHouseBuilder) SetRoof() {
	b.house.Roof = "Wooden Roof"
}

// SetDoors устанавливает двери.
func (b *WoodenHouseBuilder) SetDoors() {
	b.house.Doors = 2
}

// GetHouse возвращает построенный деревянный дом.
func (b *WoodenHouseBuilder) GetHouse() House {
	return b.house
}

// Director определяет порядок построения дома.
type Director struct {
	builder HouseBuilder
}

// NewDirector создает новый экземпляр Director.
func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

// Construct строит дом, используя шаги, определенные в HouseBuilder.
func (d *Director) Construct() {
	d.builder.SetFoundation()
	d.builder.SetWalls()
	d.builder.SetRoof()
	d.builder.SetDoors()
}

func main() {
	// Создаем строителя для бетонного дома.
	concreteBuilder := NewConcreteHouseBuilder()

	// Создаем директора, передавая ему строителя бетонного дома.
	director1 := NewDirector(concreteBuilder)

	// Директор строит дом.
	director1.Construct()

	// Получаем построенный бетонный дом.
	concreteHouse := concreteBuilder.GetHouse()

	// Выводим детали бетонного дома.
	fmt.Printf("Foundation: %s\n", concreteHouse.Foundation)
	fmt.Printf("Walls: %d\n", concreteHouse.Walls)
	fmt.Printf("Roof: %s\n", concreteHouse.Roof)
	fmt.Printf("Doors: %d\n\n", concreteHouse.Doors)

	woodenBuilder := NewWoodenHouseBuilder()

	// Создаем директора, передавая ему строителя для деревянного дома.
	director2 := NewDirector(woodenBuilder)

	// Директор строит дом.
	director2.Construct()

	// Получаем построенный деревянный дом.
	woodenHouse := woodenBuilder.GetHouse()

	// Выводим детали деревянного дома.
	fmt.Printf("Foundation: %s\n", woodenHouse.Foundation)
	fmt.Printf("Walls: %d\n", woodenHouse.Walls)
	fmt.Printf("Roof: %s\n", woodenHouse.Roof)
	fmt.Printf("Doors: %d\n", woodenHouse.Doors)
}
