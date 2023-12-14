package main

import (
	"fmt"
	"os"
	"strconv"
	
	encryption "go-one-time-pad/encryption"
	utils "go-one-time-pad/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No function specified")
		return
	}

	switch args[0] {
	case "generateRandomKey":
		if len(args) != 2 {
			fmt.Println("Usage: generateRandomKey [length]")
			return
		}
		length, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid length:", args[1])
			return
		}
		fmt.Println(utils.GenerateRandomKey(length))
	case "EncodeTo64":
		if len(args) != 2 {
			fmt.Println("Usage: EncodeTo64 [string]")
			return
		}
		fmt.Println(utils.EncodeTo64(args[1]))
	case "DecodeFrom64":
		if len(args) != 2 {
			fmt.Println("Usage: DecodeFrom64 [string]")
			return
		}
		fmt.Println(utils.DecodeFrom64(args[1]))
	case "encrypt":
		if len(args) != 2 {
			fmt.Println("Usage: encrypt [message]")
			return
		}
		encrypted, key := encryption.Encrypt(args[1])
		fmt.Printf("Encrypted: %s, Key: %s\n", encrypted, key)
	case "decrypt":
		if len(args) != 3 {
			fmt.Println("Usage: decrypt [message] [key]")
			return
		}
		fmt.Println(encryption.Decrypt(args[1], args[2]))
	default:
		fmt.Println("Unknown function:", args[0])
	}
}
