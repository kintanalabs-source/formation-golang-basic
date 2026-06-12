package errors_handling

import (
	"errors"
	"testing"
)

func TestValidateAge(t *testing.T) {
	t.Run("Negative Age", func(t *testing.T) {
		err := ValidateAge(-5)
		if !errors.Is(err, ErrNegativeValue) {
			t.Errorf("ValidateAge(-5) = %v; want %v", err, ErrNegativeValue)
		}
	})

	t.Run("Too Old", func(t *testing.T) {
		err := ValidateAge(150)
		if err == nil || err.Error() != "âge invalide: 150" {
			t.Errorf("ValidateAge(150) = %v; want \"âge invalide: 150\"", err)
		}
	})

	t.Run("Valid Age", func(t *testing.T) {
		err := ValidateAge(25)
		if err != nil {
			t.Errorf("ValidateAge(25) = %v; want nil", err)
		}
	})
}

func TestCustomError(t *testing.T) {
	err := CustomError{Code: 404, Message: "Non trouvé"}
	got := err.Error()
	want := "Erreur 404: Non trouvé"
	if got != want {
		t.Errorf("CustomError.Error() = %q; want %q", got, want)
	}
}
