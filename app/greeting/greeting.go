package greeting

import (
	"fmt"
)

//SayHello is a greeting function to welcome someone.
func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name);
}

