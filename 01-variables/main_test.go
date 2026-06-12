package variables

import "testing"

func TestGetInteger(t *testing.T) {
	got := GetInteger()
	want := 42
	if got != want {
		t.Errorf("GetInteger() = %d; want %d", got, want)
	}
}

func TestGetString(t *testing.T) {
	got := GetString()
	want := "Go est génial"
	if got != want {
		t.Errorf("GetString() = %q; want %q", got, want)
	}
}

func TestGetFloat(t *testing.T) {
	got := GetFloat()
	want := 3.14
	if got != want {
		t.Errorf("GetFloat() = %v; want %v", got, want)
	}
}

func TestGetBool(t *testing.T) {
	got := GetBool()
	want := true
	if got != want {
		t.Errorf("GetBool() = %v; want %v", got, want)
	}
}

func TestConvertIntToFloat(t *testing.T) {
	got := ConvertIntToFloat(10)
	want := 10.0
	if got != want {
		t.Errorf("ConvertIntToFloat(10) = %v; want %v", got, want)
	}
}
