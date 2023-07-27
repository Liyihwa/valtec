package jwt

import (
	"fmt"
	"testing"
)

func Test(tt *testing.T) {
	tok, err := New().Set("uid", 123).Set("uids", 345).Token()
	if err != nil {
		tt.Errorf(err.Error())
	}
	fmt.Println(tok)
	p, err := Parse(tok)
	if err != nil {
		tt.Errorf(err.Error())
	}
	fmt.Println(p.Get("uid"))
	fmt.Println(p.Get("uids"))
}
