package hello

import "fmt"

func ReturnGreeting(g string) string{
	return fmt.Sprintf("%s yourself!", g)
}