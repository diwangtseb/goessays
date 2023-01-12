package example

import "fmt"

func main() {
	Engine(&Person{})
	Engine(&Order{})

	// Output
	// *Person.BeforeSave
	// Save person data
	// *Person.AfterSave
	// Save order data
}

func Engine(valPtr interface{}) error {
	if v, ok := valPtr.(interface{ BeforeSave() }); ok {
		v.BeforeSave()
	}
	if v, ok := valPtr.(interface{ Save() }); ok {
		v.Save()
	}
	if v, ok := valPtr.(interface{ AfterSave() }); ok {
		v.AfterSave()
	}
	return nil
}

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) Save() {
	// call beforeSave()
	fmt.Println("Save person data")
	// call afterSave()
}
func (p *Person) BeforeSave() { fmt.Println("*Person.BeforeSave") }
func (p *Person) AfterSave()  { fmt.Println("*Person.AfterSave") }

type Order struct {
	Number ObjectId
	Items  []Item
}

func (o *Order) Save() {
	// call beforeSave()
	fmt.Println("Save order data")
	// call afterSave()
}

type Item struct{}     // for demonstration purpose
type ObjectId struct{} // for demonstration purpose
