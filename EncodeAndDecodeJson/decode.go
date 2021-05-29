package EncodeAndDecodeJson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []person
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
}
