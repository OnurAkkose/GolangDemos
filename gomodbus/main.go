package main

import (
	"fmt"

	modbus "github.com/goburrow/modbus"
)

func main() {
	handler := modbus.NewTCPClientHandler("localhost:502")
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer handler.Close()
	client := modbus.NewClient(handler)

	//	_, err = client.WriteSingleRegister(0, 10)
	//results, err := client.WriteMultipleCoils(0, 5, []byte{3, 2})
	//results, err := client.ReadCoils(2, 4)
	//	results, err := client.ReadHoldingRegisters(0, 3)
	//client.ReadInputRegisters()
	//results, err := client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	results, err := client.ReadHoldingRegisters(0, 9)
	// go func() {
	// 	for i := 0; i < 5000; i++ {
	// 		results, err := client.ReadHoldingRegisters(0, 3)
	// 		if err != nil {
	// 			fmt.Printf("%v\n", err)
	// 		}
	// 		fmt.Printf("results %v\n", results)
	// 	}

	// }()
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }

	// results, err := client.ReadHoldingRegisters(0, 3)
	// if err != nil {
	// 	fmt.Println("gefer")
	// 	fmt.Printf("%v\n", err)
	// }

}
