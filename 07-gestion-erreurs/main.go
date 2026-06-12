package errors_handling

import (
	"errors"
	"fmt"
)

// TODO: Définir ErrNegativeValue ici

// ValidateAge vérifie si l'âge est valide
func ValidateAge(age int) error {
	// TODO: Implémenter la logique de validation
	return nil
}

// TODO: Définir la structure CustomError et sa méthode Error()
type CustomError struct {
}

func (e CustomError) Error() string {
	return ""
}
