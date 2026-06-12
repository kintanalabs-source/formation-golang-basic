# Exercice 07 : Gestion des Erreurs

## Objectif
Apprendre à gérer les erreurs de manière "idiomatic" en Go. Comprendre comment créer des erreurs personnalisées et comment les retourner.

## Instructions
Dans `main.go` :

1. Définissez une variable d'erreur globale `ErrNegativeValue` en utilisant `errors.New("valeur négative non autorisée")`.
2. Implémentez la fonction `ValidateAge(age int) error` :
   - Si `age` est inférieur à 0, retourner `ErrNegativeValue`.
   - Si `age` est supérieur à 120, retourner une nouvelle erreur : `"âge invalide: [age]"`.
   - Sinon, retourner `nil`.
3. Implémentez une structure `CustomError` qui implémente l'interface `error`. Elle doit avoir un champ `Code int` et `Message string`. La méthode `Error() string` doit retourner `"Erreur [Code]: [Message]"`.

## Révision Leçon
- **Interface error** : Tout type ayant la méthode `Error() string` est une erreur.
- **Panic vs Error** : En Go, on utilise `panic` uniquement pour des erreurs fatales irrécupérables. Le reste est géré via des retours de type `error`.
- **Fmt.Errorf** : Permet de créer une erreur formatée avec des variables.

## Validation
```bash
go test ./07-gestion-erreurs/...
```
