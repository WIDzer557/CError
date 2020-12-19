package cerror

import (
	"fmt"
	"testing"

	"./catcher"
)

func TestEBuffer_Catch(t *testing.T) {
	Throw("AAAAAaaaa").
		Catch(func(err interface{}) {
			if err != nil {
				fmt.Printf("Hello, World!")
			}
	})

	Throw("Suka blyat!!", "OH BLYA", "POSHLI VSE NAHUI!!", 404).
		Catch(catcher.Print)
}