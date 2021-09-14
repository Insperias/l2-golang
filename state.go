package main

import (
	"fmt"
	"log"
)

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state
	currentState  state
	itemCount     int
	itemPrice     int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}
func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Добавляем %d единиц товара\n", count)
	v.itemCount += count
}

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type noItemState struct {
	vendingMachine *vendingMachine
}

func (s *noItemState) requestItem() error {
	return fmt.Errorf("Товара нет в наличии")
}

func (s *noItemState) addItem(count int) error {
	s.vendingMachine.incrementItemCount(count)
	s.vendingMachine.setState(s.vendingMachine.hasItem)
	return nil
}

func (s *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Товара нет в наличии")
}

func (s *noItemState) dispenseItem() error {
	return fmt.Errorf("Товара нет в наличии")
}

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (s *hasItemState) requestItem() error {
	if s.vendingMachine.itemCount == 0 {
		s.vendingMachine.setState(s.vendingMachine.noItem)
		return fmt.Errorf("Товаров нет")
	}
	fmt.Println("Товар запрошен")
	s.vendingMachine.setState(s.vendingMachine.itemRequested)
	return nil
}

func (s *hasItemState) addItem(count int) error {
	fmt.Printf("%d товаров добавлено\n", count)
	s.vendingMachine.incrementItemCount(count)
	return nil
}

func (s *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Пожалуйста, сначала выберите товар")
}

func (s *hasItemState) dispenseItem() error {
	return fmt.Errorf("Пожалуйста, сначала выберите товар")
}

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (s *itemRequestedState) requestItem() error {
	return fmt.Errorf("Товар уже запрошен")
}

func (s *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Происходит выдача товара")
}

func (s *itemRequestedState) insertMoney(money int) error {
	if money < s.vendingMachine.itemPrice {
		return fmt.Errorf("Вложенных денег меньше, чем требуется. Пожалуйста внесите %d", s.vendingMachine.itemPrice)
	}
	fmt.Println("Оплата принята")
	s.vendingMachine.setState(s.vendingMachine.hasMoney)
	return nil
}
func (s *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Пожалуйста, произведите оплату")
}

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (s *hasMoneyState) requestItem() error {
	return fmt.Errorf("Выдача товара в процессе")
}

func (s *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Выдача товара в процессе")
}

func (s *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Товара нет в наличии")
}

func (s *hasMoneyState) dispenseItem() error {
	fmt.Println("Выдача товара")
	s.vendingMachine.itemCount -= 1
	if s.vendingMachine.itemCount == 0 {
		s.vendingMachine.setState(s.vendingMachine.noItem)
	} else {
		s.vendingMachine.setState(s.vendingMachine.hasItem)
	}
	return nil
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
