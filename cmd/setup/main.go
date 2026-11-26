package main

import (
	"fmt"

	"github.com/1Vewton/harper-go/cmd/setup/auth"
)

const currentVersion string = "0.1.0"

func main() {
	fmt.Println(
		"Start setting up harper-go...",
	)
	fmt.Println(
		"Start validating current version",
	)
	auth.GetToken()
}
