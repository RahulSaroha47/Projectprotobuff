package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuff/protobean"
	"google.golang.org/protobuf/proto"
)

func main() {

	person := &protobean.Person{
		Name:    "Rahul",
		Age:     23,
		Pincode: 125001,
		City:    "Hisar",
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	if err := ioutil.WriteFile("proto", data, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	fmt.Println(data)

	in, err := ioutil.ReadFile("proto")

	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	person1 := &protobean.Person{}

	if err := proto.Unmarshal(in, person1); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	//  err1 := proto.Unmarshal(data,person)

	//  if err1 != nil {
	// 	 fmt.Println("Error occured during unmarshaling",err1)
	//  }

	fmt.Println("Age", person1.GetAge())
	fmt.Println("Name", person1.GetName())
	fmt.Println("City",person1.GetCity())
	fmt.Println("Pincode",person1.GetPincode())

}
