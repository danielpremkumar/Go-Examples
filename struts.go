package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	cont := contactInfo{"alex@gmail.com", 41234}
	alex := person{"Alex", "Anderson", cont}
	rebecca := person{
		firstName: "Rebecca",
		lastName:  "Cox",
		contactInfo: contactInfo{
			"rebecca@yahoo.com",
			10001,
		},
	}
	var jenifer person
	jenifer.firstName = "Jenifer"
	jenifer.lastName = "Werb"
	jenifer.contactInfo.email = "jenifer@outlook.com"
	jenifer.contactInfo.zipCode = 75010
	fmt.Println(alex, rebecca, jenifer)
	alex.print()
	fmt.Println(jenifer.updateName("jeni"))
	jenifer.print()
	jenifer.updateNamePointer("Jenny")
	jenifer.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) person {
	p.firstName = newFirstName
	return p
}

func (p *person) updateNamePointer(newFirstName string) {
	p.firstName = newFirstName
}
