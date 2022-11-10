package parser

import (
	"fmt"
	"testing"
)

func TestTryParsePassportRF(t *testing.T) {
	res := TryParsePassportRF("5421 425555")
	fmt.Printf("%v\n", res)
}
