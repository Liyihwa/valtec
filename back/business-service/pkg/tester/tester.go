package tester

import (
	"reflect"
	"testing"
)

type Tester struct {
	T *testing.T
}

func New(t *testing.T) Tester {
	return Tester{T: t}
}

func (t Tester) shouleBe(a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.T.Errorf("Want %s, got %s", a, b)
	}
}
