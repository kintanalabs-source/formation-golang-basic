# Exercice 05 : Structs et Méthodes

## Objectif
Apprendre à définir des structures de données (`struct`) et à leur attacher des comportements via des méthodes. Comprendre la différence entre un receveur par valeur et un receveur par pointeur.

## Instructions
Dans `main.go` :

1. Définissez une structure `User` avec les champs :
    - `ID` (int)
    - `Username` (string)
    - `Email` (string)
2. Implémentez la fonction `NewUser(id int, username, email string) User` qui retourne une nouvelle instance de `User`.
3. Implémentez une méthode `GetInfo() string` sur `User` (receveur par valeur) qui retourne une chaîne au format : `"ID: 1, Username: admin"`.
4. Implémentez une méthode `UpdateEmail(newEmail string)` sur `*User` (**receveur par pointeur**) qui met à jour le champ `Email` de l'utilisateur.

## Révision Leçon
- **Struct** : Un regroupement de champs typés.
- **Méthodes** : Fonctions attachées à un type.
- **Receveur par valeur** : `func (u User) Method()` -> Travaille sur une copie.
- **Receveur par pointeur** : `func (u *User) Method()` -> Permet de modifier l'original et évite de copier de grosses structures.

## Validation
```bash
go test ./05-structs-methodes/...
```
