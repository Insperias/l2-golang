package main

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type device interface {
	on()
	off()
}

type modem struct {
	isRunning bool
}

func (m *modem) on() {
	m.isRunning = true
	fmt.Println("Включаем модем")
}

func (m *modem) off() {
	m.isRunning = false
	fmt.Println("Выключаем модем")
}

func main() {
	modem := &modem{}

	onCommand := &onCommand{
		device: modem,
	}

	offCommand := &offCommand{
		device: modem,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
