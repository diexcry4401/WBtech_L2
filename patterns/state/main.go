package main

import "fmt"

/*
Состояние — это поведенческий паттерн проектирования,
который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.
*/

// State - интерфейс, определяющий методы для различных состояний автомата.
type State interface {
	InsertCoin()
	EjectCoin()
	DispenseTicket()
}

// TicketMachine - контекст, который содержит текущее состояние автомата.
type TicketMachine struct {
	state       State
	ticketCount int
}

// NewTicketMachine - конструктор, создающий новый автомат с определенным количеством билетов.
func NewTicketMachine(ticketCount int) *TicketMachine {
	machine := &TicketMachine{ticketCount: ticketCount}
	if ticketCount > 0 {
		machine.state = &NoCoinState{machine}
	} else {
		machine.state = &NoTicketState{machine}
	}
	return machine
}

// SetState - метод для изменения текущего состояния автомата.
func (m *TicketMachine) SetState(state State) {
	m.state = state
}

// InsertCoin - метод для вставки монеты.
func (m *TicketMachine) InsertCoin() {
	m.state.InsertCoin()
}

// EjectCoin - метод для возврата монеты.
func (m *TicketMachine) EjectCoin() {
	m.state.EjectCoin()
}

// DispenseTicket - метод для выдачи билета.
func (m *TicketMachine) DispenseTicket() {
	m.state.DispenseTicket()
}

// NoCoinState - состояние, когда в автомате нет монеты.
type NoCoinState struct {
	machine *TicketMachine
}

func (s *NoCoinState) InsertCoin() {
	fmt.Println("Coin inserted.")
	s.machine.SetState(&HasCoinState{s.machine})
}

func (s *NoCoinState) EjectCoin() {
	fmt.Println("No coin to eject.")
}

func (s *NoCoinState) DispenseTicket() {
	fmt.Println("Insert a coin first.")
}

// HasCoinState - состояние, когда в автомате есть монета.
type HasCoinState struct {
	machine *TicketMachine
}

func (s *HasCoinState) InsertCoin() {
	fmt.Println("Coin already inserted.")
}

func (s *HasCoinState) EjectCoin() {
	fmt.Println("Coin ejected.")
	s.machine.SetState(&NoCoinState{s.machine})
}

func (s *HasCoinState) DispenseTicket() {
	fmt.Println("Ticket dispensed.")
	s.machine.ticketCount--
	if s.machine.ticketCount > 0 {
		s.machine.SetState(&NoCoinState{s.machine})
	} else {
		s.machine.SetState(&NoTicketState{s.machine})
	}
}

// NoTicketState - состояние, когда в автомате нет билетов.
type NoTicketState struct {
	machine *TicketMachine
}

func (s *NoTicketState) InsertCoin() {
	fmt.Println("No tickets available.")
}

func (s *NoTicketState) EjectCoin() {
	fmt.Println("No coin to eject.")
}

func (s *NoTicketState) DispenseTicket() {
	fmt.Println("No tickets available.")
}

func main() {
	// Создаем автомат с 1 билетом.
	machine1 := NewTicketMachine(1)

	// Создаем автомат с 2 билетами
	machine2 := NewTicketMachine(2)

	fmt.Printf("Machine with 1 ticket:\n")
	// Пробуем вставить монету и получить билет.
	machine1.InsertCoin()
	machine1.DispenseTicket()

	// Пробуем вставить монету и получить билет снова, когда билеты закончились.
	machine1.InsertCoin()
	machine1.DispenseTicket()

	fmt.Printf("\n")
	fmt.Printf("Machine with 2 tickets:\n")
	// Пробуем вставить монету и получить билет.
	machine2.InsertCoin()
	machine2.DispenseTicket()

	// Пробуем вставить монету и получить билет снова.
	machine2.InsertCoin()
	machine2.DispenseTicket()
}
