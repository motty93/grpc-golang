package main

import (
	"fmt"

	"github.com/motty93/grpc-golang/pb"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "motty",
		Email: "motty@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123-4567", Type: pb.Person_HOME},
		},
	}

	// Wirte the new address book back to disk.
	// book := pb.AddressBook{}
	// out, err := proto.Marshal(&book)
	// if err != nil {
	// 	log.Fatalln("Failed to encode address book:", err)
	// }
	// if err := ioutil.WriteFile(fname, out, 0644); err != nil {
	// 	log.Fatalln()
	// }

	// Read the existing address book.
	// book := pb.AddressBook{}
	// in, err := ioutil.ReadFile(fname)
	// if err != nil {
	// 	log.Fatalln("Error reading file:", err)
	// }
	// book := pb.AddressBook{}
	// if err := proto.Unmarshal(in, book); err != nil {
	// 	log.Fatalln("Failed to parse address book:", err)
	// }

	fmt.Printf("p: %+v\n", p)
}
