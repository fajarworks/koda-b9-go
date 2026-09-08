package model

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (p *Person) GetPersonData() string {
	return fmt.Sprintf("nama: %s, alamat: %s, no Hp: %s", p.Name, p.Address, p.Phone)

}

func (p *Person) SetPersonName(name string) {
	p.Name = name

}

func (p *Person) Greet() string {
	return fmt.Sprintf("halo %s", p.Name)
}
