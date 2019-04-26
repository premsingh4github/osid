package main

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"log"
)

func main() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OS Id:" + id)
	id_new, err_new := machineid.ProtectedID("myAppName")
	if err_new != nil {
		log.Fatal(err)
	}
	fmt.Println("Secure Id: " + id_new)
}
