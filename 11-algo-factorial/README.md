# Exercice 11 : Algorithmique - Factorielle

## Objectif
Calculer la factorielle d'un nombre entier positif.
Rappel : `n! = n * (n-1) * (n-2) * ... * 1`. Par définition, `0! = 1`.

## Instructions
Dans `main.go`, implémentez la fonction `Factorial(n int) int`.
- Si `n` est négatif, retournez `0` (pour simplifier).
- Vous pouvez choisir une approche **itérative** (boucle for) ou **récursive** (la fonction s'appelle elle-même).

## Révision Leçon (Récursivité)
- **Condition d'arrêt** : C'est le point le plus important d'une fonction récursive pour éviter une boucle infinie (stack overflow). Ici, c'est quand `n == 0`.
- **Performance** : En Go, les boucles `for` sont souvent plus performantes et moins gourmandes en mémoire que la récursivité pure, car Go n'optimise pas systématiquement la "tail recursion".

## Validation
```bash
go test ./11-algo-factorial/...
```
