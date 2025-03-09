package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson(firstName, lastName string, age int) Person {
	p1 := Person{
		firstName,
		lastName,
		age,
	}
	return p1
}

func makePersonPointer(firstName, lastName string, age int) *Person {
	p2 := Person{
		firstName,
		lastName,
		age,
	}
	return &p2
}

func main() {
	// pers := Person{"su", "bu", 61}
	// p := makePerson("su", "xu", 15)
	// p_ptr := makePersonPointer("su", "xu", 15)

	// fmt.Println(pers, p, *p_ptr)

	population := make([]Person, 10_000_000)

	for i, person := range population {
		person.firstName = "meh"
		person.lastName = "whoose"
		person.age = 10

		population[i] = person
	}
}
