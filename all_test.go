package yaml

import (
	"testing"
)

func TestNewDecoder(t *testing.T) {
	//t.Log(getDecimalExponent(5.0e-05))

	test := map[string]float64{"a": 5.0e-05, "b": 1.0}
	out, err := Marshal(&test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
}
