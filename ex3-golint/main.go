package person

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, %s", p.Name)
}

func (p Person) CanEnterBar() bool {
	return p.isDrinkingAge()
}

func (p Person) isDrinkingAge() bool {
	return p.Age >= 21
}
