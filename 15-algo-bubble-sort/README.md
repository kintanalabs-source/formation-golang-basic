# Exercice 15 : Algorithmique - Tri à Bulles (Bubble Sort)

## Objectif
Implémenter l'algorithme de tri à bulles. C'est l'un des algorithmes de tri les plus simples à comprendre, bien qu'il ne soit pas le plus performant pour de grands tableaux.

## Instructions
Dans `main.go`, implémentez la fonction `BubbleSort(nums []int) []int`.
- La fonction doit trier le slice `nums` par ordre croissant **en place** (elle modifie le slice original) et le retourner.

## Révision Leçon (Bubble Sort)
- **Principe** : On parcourt le tableau plusieurs fois. À chaque passage, on compare chaque élément avec le suivant. S'ils sont dans le mauvais ordre, on les échange.
- **Optimisation** : À chaque passage complet, le plus grand élément "remonte" à sa position finale (comme une bulle d'air). On peut donc réduire la zone de recherche à chaque itération.
- **Complexité** : Sa complexité est de **O(n²)**, ce qui le rend lent pour de grandes listes.

## Validation
```bash
go test ./15-algo-bubble-sort/...
```
