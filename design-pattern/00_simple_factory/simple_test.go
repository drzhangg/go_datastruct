package simplefactory

import (
	"fmt"
	"testing"
)

func TestHi(t *testing.T) {
	api := NewAPI(1)
	jerry := api.Say("jerry")
	fmt.Println(jerry)
}
