package parser

import (
	"fmt"
	"testing"
)

func TestTryParseGrz(t *testing.T) {
	res := TryParseGrz("A123AA96")
	fmt.Printf("%v\n", res)
}
