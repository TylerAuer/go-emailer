package main

import (
	"fmt"

	"github.com/TwinProduction/go-color"
)

func main() {
	fmt.Println(color.Ize(color.Blue, "Initializing emailer"))

	e := Email{
		FromName:  "Super Bowl Props",
		FromEmail: "Tyler.Auer@gmail.com",
		ToName:    "Jessica",
		ToEmail:   "Jessica.L.Rosten@gmail.com",
		Subject:   "Testing",
		Message:   "Don't worry about this, it's just a test email",
	}

	sendMailFromEmail(e)
}
