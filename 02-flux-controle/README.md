# Exercice 02 : Flux de Contrôle

## Objectif
Maîtriser les structures conditionnelles (`if`, `else`, `switch`) et la boucle unique de Go (`for`).

## Instructions
Complétez les fonctions dans `main.go` :

1. `IsEven(n int) bool` : Retourne `true` si le nombre est pair, sinon `false`.
2. `GetGrade(score int) string` : Retourne une lettre selon le score :
    - 90+ : "A"
    - 80-89 : "B"
    - 70-79 : "C"
    - 60-69 : "D"
    - Moins de 60 : "F"
3. `SumUpTo(n int) int` : Calcule la somme de tous les entiers de 1 à `n` inclus en utilisant une boucle `for`.
4. `GetDayName(day int) string` : Retourne le nom du jour de la semaine (1 pour "Lundi", 2 pour "Mardi", etc.) en utilisant un `switch`. Retourne "Inconnu" si le chiffre n'est pas entre 1 et 7.

## Révision Leçon
- **If/Else** : Les parenthèses sont inutiles, les accolades `{}` sont obligatoires.
- **For** : C'est l'unique boucle en Go. Elle peut être utilisée comme un `while` : `for condition { ... }`.
- **Switch** : Pas besoin de `break` à la fin de chaque case. Go s'arrête automatiquement sauf si `fallthrough` est utilisé.

## Validation
```bash
go test ./02-flux-controle/...
```
