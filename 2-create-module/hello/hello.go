package main

import (
	"fmt"
	"log"

	greetings "example.com/greetings"
)

func main() {
	//message, err := greetings.HelloMultiple(os.Args[1], os.Args[2])

	var nameStruct []greetings.ComplexNameType

	nameToat := greetings.CreateComplexName("jk", "medikeol")
	nameEvans := greetings.CreateComplexName("white", "vans")
	nameDanbul := greetings.CreateComplexName("korea", "seoul")
	nameDanbul2 := greetings.CreateComplexName("korea", "seoul")

	nameStruct = append(nameStruct, nameToat)
	nameStruct = append(nameStruct, nameEvans)
	nameStruct = append(nameStruct, nameDanbul)

	//fmt.Println(nameStruct)

	message, err := greetings.HelloMultipleComplex(nameStruct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(message[nameDanbul2])

	simpleMessage, err := greetings.HelloMultiple([]string{"susan", "amoguis"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(simpleMessage)
}
