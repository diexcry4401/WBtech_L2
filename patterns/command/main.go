package main

import "fmt"

// Command - интерфейс для команд
type Command interface {
	Execute() // Метод выполнения команды
	Undo()    // Метод отмены команды
}

// RemoteControl - структура, представляющая пульт управления
type RemoteControl struct {
	com Command
}

// SetCommand - метод для установки команды на пульт
func (r *RemoteControl) SetCommand(command Command) {
	r.com = command
}

// PressButton - метод для выполнения команды
func (r *RemoteControl) PressButton() {
	r.com.Execute()
}

// PressUndo - метод для отмены последней команды
func (r *RemoteControl) PressUndo() {
	r.com.Undo()
}

// Инициатор, записывающий команды в стек и провоцирует их выполнение.
type invoker struct {
	commands []Command
}

func (i *invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *invoker) DeleteCommand() {
	if len(i.commands) > 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}

// func (i *invoker) Undo() {
// 	if i.commands[len(i.commands)-1] == command.Execute()
// }

// TV - структура, представляющая телевизор
type TV struct {
	isOn bool
}

// On - метод для включения телевизора
func (tv *TV) On() {
	tv.isOn = true
	fmt.Println("TV is On")
}

// Off - метод для выключения телевизора
func (tv *TV) Off() {
	tv.isOn = false
	fmt.Println("TV is Off")
}

// TVOnCommand - команда для включения телевизора
type TVOnCommand struct {
	device *TV
}

// Execute - выполнение команды включения телевизора
func (c *TVOnCommand) Execute() {
	c.device.On()
}

// Undo - отмена команды включения телевизора
func (c *TVOnCommand) Undo() {
	c.device.Off()
}

// TVOffCommand - команда для выключения телевизора
type TVOffCommand struct {
	device *TV
}

// Execute - выполнение команды выключения телевизора
func (c *TVOffCommand) Execute() {
	c.device.Off()
}

// Undo - отмена команды выключения телевизора
func (c *TVOffCommand) Undo() {
	c.device.On()
}

// PC - структура, представляющая ПК
type PC struct {
	isOn bool
}

// On - метод для включения ПК
func (pc *PC) On() {
	pc.isOn = true
	fmt.Println("PC is On")
}

// Off - метод для выключения ПК
func (pc *PC) Off() {
	pc.isOn = false
	fmt.Println("PC is Off")
}

// PCOnCommand - команда для включения ПК
type PCOnCommand struct {
	device *PC
}

// Execute - выполнение команды включения ПК
func (c *PCOnCommand) Execute() {
	c.device.On()
}

// Undo - отмена команды включения ПК
func (c *PCOnCommand) Undo() {
	c.device.Off()
}

// PCOffCommand - команда для выключения ПК
type PCOffCommand struct {
	device *PC
}

// Execute - выполнение команды выключения ПК
func (c *PCOffCommand) Execute() {
	c.device.Off()
}

// Undo - отмена команды выключения ПК
func (c *PCOffCommand) Undo() {
	c.device.On()
}

func main() {
	// Создаем телевизор и ПК
	tv := &TV{}
	pc := &PC{}

	// Создаем команды для телевизова
	onTV := &TVOnCommand{device: tv}
	offTV := &TVOffCommand{device: tv}

	// Создаем команды для телевизова
	onPC := &PCOnCommand{device: pc}
	offPC := &PCOffCommand{device: pc}

	invoker := &invoker{}

	invoker.AddCommand(onTV)
	invoker.AddCommand(onPC)
	invoker.AddCommand(offTV)
	invoker.AddCommand(offPC)

	invoker.Execute()

	fmt.Println("Пример с использованием пульта:")

	// Создаем пульт
	remote := &RemoteControl{}

	// Включаем телевизор
	remote.SetCommand(onTV)
	remote.PressButton() // Вывод: "TV is On"

	// Включаем ПК
	remote.SetCommand(onPC)
	remote.PressButton() // Вывод: "PC is On"

	// Выключаем телевизор
	remote.SetCommand(offTV)
	remote.PressButton() // Вывод: "TV is Off"

	// Выключаем телевизор
	remote.SetCommand(offPC)
	remote.PressButton() // Вывод: "PC is Off"

	// Отменяем последнюю команду (включаем ПК)
	remote.PressUndo() // Вывод: "PC is On"
}
