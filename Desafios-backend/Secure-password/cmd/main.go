package main

import (
	"fmt"
	"secure-password/internal/validator"
)

func main() {

	password := ""
	result, err := validator.IsValidPassword(&password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Is valid password!", result)

}
