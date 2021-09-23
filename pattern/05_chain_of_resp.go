package main

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Пациент прошел регистрацию")
		r.next.execute(p)
		return
	}
	fmt.Println("Пациент в регистратуре")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Пациент уже был у доктора")
		d.next.execute(p)
		return
	}
	fmt.Println("Пациент на осмотре у доктора")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Пациент выбрал нужные лекарства")
		m.next.execute(p)
		return
	}
	fmt.Println("Пациент выбирает  лекарства в аптеке")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Пациент оплатил лекарства")
		return
	}
	fmt.Println("Пациент оплачивает выбранные лекарства")
	p.paymentDone = true
}

func (c *cashier) setNext(next department) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "Tomas"}
	reception.execute(patient)
	fmt.Println()
	reception.execute(patient)

}
