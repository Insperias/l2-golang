package main

import "fmt"

type iBuilder interface {
	setMeat()
	setSauce()
	setCheese()
	setPepper()
	setTomatoes()
	setOlives()
	setOnion()
	setSalat()
	setFreshCucumbers()
	setPickles()
	getSandwich() sandwich
}

func getBuilder(builderType string) iBuilder {
	if builderType == "standart" {
		return &standartBuilder{}
	}
	if builderType == "vegetarian" {
		return &vegetarianBuilder{}
	}
	return nil
}

type standartBuilder struct {
	meat           string
	sauce          string
	cheese         string
	pepper         string
	tomatoes       string
	olives         string
	onion          string
	salat          string
	freshCucumbers string
	pickles        string
}

func newStandartBuilder() *standartBuilder {
	return &standartBuilder{}
}
func (b *standartBuilder) setMeat() {
	b.meat = "индейка"
}
func (b *standartBuilder) setSauce() {
	b.sauce = "соус BBQ"
}
func (b *standartBuilder) setCheese() {
	b.cheese = "сыр"
}
func (b *standartBuilder) setPepper() {
	b.pepper = ""
}
func (b *standartBuilder) setTomatoes() {
	b.tomatoes = "томаты"
}
func (b *standartBuilder) setOlives() {
	b.olives = ""
}
func (b *standartBuilder) setOnion() {
	b.onion = ""
}
func (b *standartBuilder) setSalat() {
	b.salat = "листья салата"
}
func (b *standartBuilder) setFreshCucumbers() {
	b.freshCucumbers = ""
}
func (b *standartBuilder) setPickles() {
	b.pickles = "маринованные огурцы"
}
func (b *standartBuilder) getSandwich() sandwich {
	return sandwich{
		meat:           b.meat,
		sauce:          b.sauce,
		cheese:         b.cheese,
		pepper:         b.pepper,
		tomatoes:       b.tomatoes,
		olives:         b.olives,
		onion:          b.onion,
		salat:          b.salat,
		freshCucumbers: b.freshCucumbers,
		pickles:        b.pickles,
	}
}

type vegetarianBuilder struct {
	meat           string
	sauce          string
	cheese         string
	pepper         string
	tomatoes       string
	olives         string
	onion          string
	salat          string
	freshCucumbers string
	pickles        string
}

func newVegetarianBuilder() *vegetarianBuilder {
	return &vegetarianBuilder{}
}

func (b *vegetarianBuilder) setMeat() {
	b.meat = "без мяса"
}
func (b *vegetarianBuilder) setSauce() {
	b.sauce = "соус чесночный"
}
func (b *vegetarianBuilder) setCheese() {
	b.cheese = ""
}
func (b *vegetarianBuilder) setPepper() {
	b.pepper = "перец"
}
func (b *vegetarianBuilder) setTomatoes() {
	b.tomatoes = "томаты"
}
func (b *vegetarianBuilder) setOlives() {
	b.olives = "оливки"
}
func (b *vegetarianBuilder) setOnion() {
	b.onion = "лук"
}
func (b *vegetarianBuilder) setSalat() {
	b.salat = "листья салата"
}
func (b *vegetarianBuilder) setFreshCucumbers() {
	b.freshCucumbers = "свежие огурцы"
}
func (b *vegetarianBuilder) setPickles() {
	b.pickles = ""
}
func (b *vegetarianBuilder) getSandwich() sandwich {
	return sandwich{
		meat:           b.meat,
		sauce:          b.sauce,
		cheese:         b.cheese,
		pepper:         b.pepper,
		tomatoes:       b.tomatoes,
		olives:         b.olives,
		onion:          b.onion,
		salat:          b.salat,
		freshCucumbers: b.freshCucumbers,
		pickles:        b.pickles,
	}
}

type sandwich struct {
	meat           string
	sauce          string
	cheese         string
	pepper         string
	tomatoes       string
	olives         string
	onion          string
	salat          string
	freshCucumbers string
	pickles        string
}
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) makeSandwich() sandwich {
	d.builder.setMeat()
	d.builder.setSauce()
	d.builder.setCheese()
	d.builder.setPepper()
	d.builder.setTomatoes()
	d.builder.setOlives()
	d.builder.setOnion()
	d.builder.setSalat()
	d.builder.setFreshCucumbers()
	d.builder.setPickles()
	return d.builder.getSandwich()
}

func main() {
	standartBuilder := getBuilder("standart")
	vegetarianBuilder := getBuilder("vegetarian")

	director := newDirector(standartBuilder)
	standartSandwich := director.makeSandwich()
	fmt.Println(standartSandwich)

	director.setBuilder(vegetarianBuilder)
	vegetarianSandwich := director.makeSandwich()
	fmt.Println(vegetarianSandwich)
}
