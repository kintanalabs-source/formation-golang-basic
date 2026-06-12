# Exercice 04 : Slices et Maps

## Objectif
Manipuler les collections de données les plus utilisées en Go : les Slices (tableaux dynamiques) et les Maps (tables de hachage).

## Instructions
Complétez les fonctions dans `main.go` :

1. `GetFruits() []string` : Doit retourner un slice contenant `"pomme"`, `"banane"`, `"orange"`.
2. `AddFruit(fruits []string, fruit string) []string` : Doit ajouter `fruit` à la liste `fruits` et retourner le nouveau slice.
3. `GetScores() map[string]int` : Doit retourner une map avec les entrées `"Alice": 10` et `"Bob": 15`.
4. `CheckScore(scores map[string]int, name string) (int, bool)` : Doit vérifier si `name` existe dans la map. Retourner le score et `true` s'il existe, sinon `0` et `false`.

## Révision Leçon
- **Slices** : Listes dynamiques. On utilise `append(slice, element)` pour ajouter.
- **Maps** : Dictionnaires. Initialisation avec `make(map[TypeCle]TypeValeur)`.
- **Comma ok** : Pour vérifier si une clé existe : `val, ok := maMap["cle"]`. Si `ok` est `false`, la clé n'existe pas.

## Validation
```bash
go test ./04-slices-maps/...
```
