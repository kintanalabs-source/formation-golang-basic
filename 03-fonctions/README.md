# Exercice 03 : Les Fonctions

## Objectif
Comprendre la signature des fonctions en Go, notamment les retours multiples et les retours nommés.

## Instructions
Complétez les fonctions dans `main.go` :

1. `Swap(a, b string) (string, string)` : Doit retourner les deux chaînes dans l'ordre inverse.
2. `Divide(a, b float64) (float64, error)` : Doit retourner le résultat de `a / b`. 
   - Si `b` est égal à `0`, retourner `0` et une erreur créée avec `errors.New("division par zéro")`.
3. `GetRectProps(width, height float64) (area, perimeter float64)` : Doit calculer la surface et le périmètre d'un rectangle en utilisant des **retours nommés**.

## Révision Leçon
- **Retours multiples** : Une fonction peut retourner plusieurs valeurs : `func f() (int, string)`.
- **Retours nommés** : On peut nommer les variables de retour dans la signature : `func f() (result int)`. Un simple `return` renverra la valeur actuelle de `result`.
- **Erreurs** : Par convention, l'erreur est souvent le dernier paramètre de retour.

## Validation
```bash
go test ./03-fonctions/...
```
