# Exercice 10 : Algorithmique - Palindrome

## Objectif
Déterminer si une chaîne de caractères est un palindrome (se lit de la même façon dans les deux sens, ex: "radar").

## Instructions
Dans `main.go`, implémentez la fonction `IsPalindrome(s string) bool`.
- La fonction doit ignorer la casse (majuscules/minuscules).
- Elle doit retourner `true` si c'est un palindrome, sinon `false`.

*Indice : Vous pouvez inverser la chaîne ou comparer les caractères deux à deux en partant des extrémités.*

## Révision Leçon (Strings)
- **Conversion de casse** : Utilisez `strings.ToLower(s)` du package `strings`.
- **Runes** : En Go, une chaîne est un slice d'octets. Pour gérer correctement tous les caractères (Unicode), il est préférable de convertir la string en un slice de `rune` : `r := []rune(s)`.
- **Inversion** : Il n'y a pas de fonction `reverse` native pour les strings en Go, il faut souvent la construire soi-même.

## Validation
```bash
go test ./10-algo-palindrome/...
```
