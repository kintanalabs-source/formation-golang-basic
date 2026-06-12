# Exercice 13 : Algorithmique - Suite de Fibonacci

## Objectif
Générer la suite de Fibonacci jusqu'à un certain rang `n`.
Rappel : `F(0) = 0`, `F(1) = 1`, et pour `n > 1`, `F(n) = F(n-1) + F(n-2)`.

## Instructions
Dans `main.go`, implémentez la fonction `Fibonacci(n int) int`.
- Retournez le `n`-ième nombre de la suite de Fibonacci.
- Si `n` est négatif, retournez `0`.

## Révision Leçon (Fibonacci)
- **Approche récursive** : Très simple à écrire (`return Fibonacci(n-1) + Fibonacci(n-2)`), mais extrêmement lente pour les grands nombres (complexité exponentielle).
- **Approche itérative** : Beaucoup plus rapide. On utilise deux variables pour stocker les deux derniers nombres calculés et on avance boucle après boucle.
- **Dépassement (Overflow)** : Les nombres de Fibonacci grandissent très vite. Pour `n` très grand, un `int` pourrait ne pas suffire (mais ici nous resterons sur des valeurs raisonnables).

## Validation
```bash
go test ./13-algo-fibonacci/...
```
