package main

import (
	"fmt"
	secret "secret/vault"
)

func main() {
	v := secret.File("key1", ".secrets")
	err := v.Set("demo_key1", "Shubham Malshikare")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo_key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
