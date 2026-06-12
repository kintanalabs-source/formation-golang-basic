# Exercice 09 : Algorithmique - FizzBuzz

## Objectif
Résoudre le classique "FizzBuzz". C'est un excellent test pour combiner boucles, modulo et conditions.

## Instructions
Dans `main.go`, implémentez la fonction `FizzBuzz(n int) []string`.
La fonction doit retourner un slice de chaînes de caractères de 1 à `n` avec les règles suivantes :
1. Pour les multiples de **3**, remplacez le nombre par `"Fizz"`.
2. Pour les multiples de **5**, remplacez le nombre par `"Buzz"`.
3. Pour les multiples de **3 et 5**, remplacez le nombre par `"FizzBuzz"`.
4. Sinon, retournez le nombre sous forme de chaîne (ex: `"1"`, `"2"`, `"4"`).

*Indice : Utilisez `strconv.Itoa(i)` pour convertir un entier en chaîne.*

## Révision Leçon (Algorithmique)
- **Modulo (`%`)** : Calcule le reste d'une division entière. `x % 3 == 0` signifie que `x` est un multiple de 3.
- **Ordre des conditions** : Attention à l'ordre de vos `if/else`. Si une condition est remplie, les suivantes sont ignorées.
- **Conversion** : Le package `strconv` est votre ami pour transformer des types numériques en texte.

## Validation
```bash
go test ./09-algo-fizzbuzz/...
```
