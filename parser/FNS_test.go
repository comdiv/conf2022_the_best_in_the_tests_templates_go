package parser

import (
	"fmt"
	"testing"
)

func TestTryParseInnFl(t *testing.T) {
	inn := TryParseInnFl("513670907640")
	fmt.Printf("%v\n", inn)
}
