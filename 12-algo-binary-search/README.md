# Exercice 12 : Algorithmique - Recherche Binaire

## Objectif
Implémenter l'algorithme de recherche binaire (Binary Search). C'est un algorithme de recherche très efficace qui fonctionne sur des tableaux **déjà triés**.

## Instructions
Dans `main.go`, implémentez la fonction `BinarySearch(nums []int, target int) int`.
- La fonction doit retourner l'**index** de `target` s'il est présent dans le slice `nums`.
- Si `target` n'est pas trouvé, retournez `-1`.
- L'algorithme doit avoir une complexité de **O(log n)**.

## Révision Leçon (Recherche Binaire)
- **Principe** : On compare l'élément au milieu du tableau avec la cible.
- Si l'élément du milieu est la cible, on a fini.
- Si la cible est plus petite, on cherche dans la moitié gauche.
- Si la cible est plus grande, on cherche dans la moitié droite.
- **Condition de boucle** : On continue tant que l'indice de gauche est inférieur ou égal à l'indice de droite.

## Validation
```bash
go test ./12-algo-binary-search/...
```
