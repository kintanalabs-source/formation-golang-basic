package models

import "fmt"

// TODO: Définir la structure User ici
type User struct {
}

// NewUser crée et retourne un nouvel utilisateur
func NewUser(id int, username, email string) User {
	// TODO: Retourner une instance de User
	return User{}
}

// GetInfo retourne les informations de l'utilisateur
func (u User) GetInfo() string {
	// TODO: Retourner la chaîne formatée "ID: %d, Username: %s"
	return ""
}

// UpdateEmail met à jour l'email de l'utilisateur
func (u *User) UpdateEmail(newEmail string) {
	// TODO: Mettre à jour le champ Email
}
