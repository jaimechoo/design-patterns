package support

import (
	"strconv"
)

type TroubleStu struct {
	Number int
}

func (t *TroubleStu) GetNumber() int {
	return t.Number
}

func (t *TroubleStu) String() string {
	return "[Trouble " + strconv.Itoa(t.Number) + "]"
}
