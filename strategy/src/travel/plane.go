package travel

import (
	"fmt"
)

type Plane struct {
}

func (o *Plane) Move(from, to string) {
	fmt.Printf("travel by plane from %s to %s\n", from, to)
}
