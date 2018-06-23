package travel

import (
	"fmt"
)

type Car struct {
}

func (o *Car) Move(from, to string) {
	fmt.Printf("travel by car from %s to %s\n", from, to)
}
