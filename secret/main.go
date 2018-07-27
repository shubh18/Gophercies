package main

import (
	"fmt"
	vault "secret/vault"
)

func main() {
	v := vault.NewVault("key1")
	err := v.Set("demo-key", "Shubham-Malshikare")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo-key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
