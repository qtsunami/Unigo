package main

import "fmt"

type Pet interface {
	//SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1
	//dog := Dog{"Little Pig"}
	//_, ok := interface{}(dog).(Pet)
	//fmt.Printf("Dog implements interface Pet: %v\n", ok)
	//
	//_, ok = interface{}(&dog).(Pet)
	//fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	//
	//// 示例2
	//var pet Pet = &dog
	//fmt.Printf("This pet is a %s, the name is %q.\n", pet.Category(), pet.Name())

	dog := Dog{"Little Bit"}
	fmt.Printf("The dog's name is %q\n", dog.Name())

	var pet Pet = dog
	dog.SetName("Moasor")

	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

	dog1 := Dog{"Little Bit"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

}
