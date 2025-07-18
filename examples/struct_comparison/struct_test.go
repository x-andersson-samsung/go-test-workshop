package struct_comparison

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestAddress() Address {
	return Address{
		Street:  "Street",
		City:    "City",
		State:   "State",
		ZipCode: "Zip Code",
		Country: "Country",
	}
}

// Standard library `reflect` method
func TestCompareReflect(t *testing.T) {
	a, b := newTestAddress(), newTestAddress()
	b.State = "Other State"

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("TestCompareReflect\ngot:\n%#v\nwant:\n%#v", a, b)
	}
}

// stretchr/testify `require.Equal`
func TestCompareAssert(t *testing.T) {
	a, b := newTestAddress(), newTestAddress()
	b.State = "Other State"

	require.Equal(t, a, b)
}
