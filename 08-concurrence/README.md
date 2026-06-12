# Exercice 08 : Concurrence

## Objectif
Découvrir la puissance de Go pour la programmation concurrente en utilisant les **Goroutines** et les **Channels**.

## Instructions
Dans `main.go` :

1. `SendValue(ch chan int, value int)` : Doit envoyer la valeur `value` dans le channel `ch`.
2. `SumParallel(a, b int) int` : 
   - Créez un channel d'entiers.
   - Lancez une goroutine qui calcule `a + b` et envoie le résultat dans le channel.
   - Recevez le résultat depuis le channel et retournez-le.
3. `PingPong(count int) ([]string, []string)` :
   - Créez deux channels de strings : `pings` et `pongs`.
   - Lancez une goroutine qui envoie `"ping"` dans le channel `pings`, `count` fois.
   - Lancez une autre goroutine qui reçoit de `pings` et renvoie `"pong"` dans `pongs`.
   - Collectez tous les résultats et retournez les deux slices de strings.

## Révision Leçon
- **Goroutine** : `go maFonction()` lance l'exécution en arrière-plan (thread léger).
- **Channels** : Tuyaux pour faire circuler des données : `ch := make(chan int)`.
- **Bloquage** : La lecture (`<-ch`) et l'écriture (`ch <- val`) sur un channel non bufférisé sont bloquantes tant que l'autre côté n'est pas prêt.
- **Close** : Fermer un channel indique qu'aucune autre valeur ne sera envoyée.

## Validation
```bash
go test ./08-concurrence/...
```
