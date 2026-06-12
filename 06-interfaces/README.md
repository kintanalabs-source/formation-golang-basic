# Exercice 06 : Interfaces

## Objectif
Comprendre le polymorphisme en Go via les interfaces. Apprendre comment une structure implémente une interface de manière implicite.

## Instructions
Dans `main.go` :

1. Définissez une interface `Shape` avec une méthode :
    - `Area() float64`
2. Définissez une structure `Rectangle` avec les champs `Width` et `Height` (float64). Implémentez la méthode `Area()` pour `Rectangle`.
3. Définissez une structure `Circle` avec le champ `Radius` (float64). Implémentez la méthode `Area()` pour `Circle`. (Utilisez `math.Pi` pour le calcul).
4. Implémentez la fonction `GetTotalArea(shapes []Shape) float64` qui calcule la somme des surfaces de toutes les formes passées en paramètre.

## Révision Leçon
- **Interface** : Un contrat définissant un ensemble de méthodes.
- **Implémentation implicite** : Pas de mot-clé `implements`. Une structure implémente une interface dès qu'elle possède les méthodes requises.
- **Interface vide** : `interface{}` ou `any` peut contenir n'importe quel type.

## Validation
```bash
go test ./06-interfaces/...
```
