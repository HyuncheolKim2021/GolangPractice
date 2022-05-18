package Switchs

import "testing"

var v interface{}

func TestSwitchString(t *testing.T) {
	t.Logf("Type of v (<nil>): %T\n", v)
	v = "Default value of string is empty string(\"\")"
	Switch(v, t)
}

func TestSwitchInt(t *testing.T) {
	t.Logf("Type of v (string): %T\n", v)
	v = 99
	Switch(v, t)
}

func TestSwitchFloat64(t *testing.T) {
	t.Parallel()
	t.Logf("Type of v (int): %T\n", v)
	v = 10.0
	Switch(v, t)
}

func Switch(v interface{}, t *testing.T) {
	switch v := v.(type) {
	case int:
		t.Logf("Type of v (int): %T\n", v)
	case float64:
		t.Logf("Type of v (float64): %T\n", v)
	case string:
		t.Logf("Type of v (string): %T\n", v)
	}
}
