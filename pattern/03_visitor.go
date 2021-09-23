package main

import "fmt"

type establishment interface {
	getType() string
	accept(visitor)
}

type cafe struct {
}

func (c *cafe) accept(v visitor) {
	v.visitForCafe(c)
}
func (c *cafe) getType() string {
	return "Cafe"
}

type restaurant struct {
}

func (r *restaurant) accept(v visitor) {
	v.visitForRestaurant(r)
}

func (r *restaurant) getType() string {
	return "Restaurant"
}

type pab struct {
}

func (p *pab) accept(v visitor) {
	v.visitForPab(p)
}

func (p *pab) getType() string {
	return "Pab"
}

type visitor interface {
	visitForCafe(*cafe)
	visitForRestaurant(*restaurant)
	visitForPab(*pab)
}

type drinker struct {
}

func (d *drinker) visitForCafe(c *cafe) {
	fmt.Println("Пьем кофе в кофейне")
}

func (d *drinker) visitForRestaurant(r *restaurant) {
	fmt.Println("Пьем дорогой виски в ресторане")
}

func (d *drinker) visitForPab(p *pab) {
	fmt.Println("Пьем пиво в пабе")
}

type eater struct {
}

func (e *eater) visitForCafe(c *cafe) {
	fmt.Println("Едим пирожное в кофейне")
}

func (e *eater) visitForRestaurant(r *restaurant) {
	fmt.Println("Едим изысканное блюдо в ресторане")
}

func (e *eater) visitForPab(p *pab) {
	fmt.Println("Едим луковые кольца в пабе")
}

func main() {
	cafe := &cafe{}
	restaurant := &restaurant{}
	pab := &pab{}

	drinker := &drinker{}

	cafe.accept(drinker)
	restaurant.accept(drinker)
	pab.accept(drinker)

	fmt.Println()
	eater := &eater{}
	cafe.accept(eater)
	restaurant.accept(eater)
	pab.accept(eater)
}
