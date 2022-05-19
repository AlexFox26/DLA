package main

import (
	"fmt"
	"reflect"

	"hw1/converter"
	"hw1/counter"
)

type unknownstruct struct{}

func main() {
	var TesterFString string = "abcd"
	getString1, enyerror1 := converter.ConvertToString(TesterFString)
	if enyerror1 != nil {
		fmt.Println("Error: ", enyerror1)
	} else {
		fmt.Println(getString1)
	}

	var TestStruct unknownstruct
	getString2, enyerror2 := converter.ConvertToString(TestStruct)
	if enyerror2 != nil {
		fmt.Println("Error: ", enyerror2)
	} else {
		fmt.Println(getString2)
	}

	var TesterIString string = "-54"
	getType, enyerror3 := converter.ConvertFromString(TesterIString)
	if enyerror3 != nil{
		fmt.Println("Error: ", enyerror3)
	}else{
		fmt.Println("Type: ", reflect.TypeOf(getType), "\tValue: ", getType)
	}

	fmt.Println(counter.Count("How much wood could a woodchuck chuck if a woodchuck could chuck wood? Wood!", 2))
}
