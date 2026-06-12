# Exercice 14 : Algorithmique - Anagramme

## Objectif
Déterminer si deux chaînes de caractères sont des anagrammes (elles contiennent exactement les mêmes caractères avec la même fréquence, mais dans un ordre différent, ex: "gare" et "rage").

## Instructions
Dans `main.go`, implémentez la fonction `IsAnagram(s1, s2 string) bool`.
- La fonction doit ignorer la casse.
- Elle doit ignorer les espaces (pour simplifier, vous pouvez supposer qu'il n'y en a pas ou utiliser `strings.ReplaceAll`).
- Retournez `true` si ce sont des anagrammes, sinon `false`.

## Révision Leçon (Anagrammes)
- **Approche 1 : Tri** : Si on trie les deux chaînes par ordre alphabétique, elles doivent devenir identiques.
- **Approche 2 : Map de fréquences** : On compte l'occurrence de chaque caractère dans une map. Pour `s1`, on incrémente ; pour `s2`, on décrémente. À la fin, toutes les valeurs de la map doivent être à zéro. C'est l'approche la plus performante (**O(n)**).

## Validation
```bash
go test ./14-algo-anagram/...
```
