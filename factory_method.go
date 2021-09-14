package main

import "fmt"

type iSword interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type sword struct {
	name  string
	power int
}

func (s *sword) setName(name string) {
	s.name = name
}

func (s *sword) getName() string {
	return s.name
}

func (s *sword) setPower(power int) {
	s.power = power
}

func (s *sword) getPower() int {
	return s.power
}

type oneHandedSword struct {
	sword
}

func newOneHandedSword() iSword {
	return &oneHandedSword{
		sword: sword{
			name:  "one-handed sword",
			power: 3,
		},
	}
}

type twoHandedSword struct {
	sword
}

func newTwoHandedSword() iSword {
	return &twoHandedSword{
		sword: sword{
			name:  "two-handed sword",
			power: 7,
		},
	}
}

func getSword(swordType string) (iSword, error) {
	if swordType == "one-handed" {
		return newOneHandedSword(), nil
	}
	if swordType == "two-handed" {
		return newTwoHandedSword(), nil
	}
	return nil, fmt.Errorf("Выбран неправильный тип")
}

func printInfo(s iSword) {
	fmt.Println(s.getName())
	fmt.Println("Мощность: ", s.getPower())
}

func main() {
	oneHanded, _ := getSword("one-handed")
	twoHanded, _ := getSword("two-handed")

	printInfo(oneHanded)
	printInfo(twoHanded)
}
