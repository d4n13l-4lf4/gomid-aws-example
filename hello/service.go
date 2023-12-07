package hello

import (
	"context"
	"fmt"
	"log"
)

type Greeter func(context.Context, string) (string, error)

func Greet(_ context.Context, name string) (string, error) {
	log.Printf("Saying hello to %s\n", name)
	greeting := fmt.Sprintf("Hello %s!", name)

	return greeting, nil
}
