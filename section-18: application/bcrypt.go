package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func encript() {
	password := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(password)
	fmt.Println(string(bs))
	fmt.Println(bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println(err)
	}

	err = bcrypt.CompareHashAndPassword(bs, []byte(`password12`))
	if err != nil {
		fmt.Println(err)
	}

}
