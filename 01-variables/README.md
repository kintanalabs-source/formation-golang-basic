# Exercice 01 : Les Variables et Types

## Objectif
Apprendre à déclarer des variables avec les différentes syntaxes de Go et comprendre les types de base.

## Instructions
Dans le fichier `main.go`, complétez les fonctions suivantes :

1. `GetInteger()` : Doit retourner un entier (`int`) valant `42`.
2. `GetString()` : Doit retourner une chaîne (`string`) valant `"Go est génial"`.
3. `GetFloat()` : Doit retourner un flottant (`float64`) valant `3.14`.
4. `GetBool()` : Doit retourner un booléen (`bool`) valant `true`.
5. `ConvertIntToFloat(n int)` : Doit prendre un entier en paramètre et le retourner sous forme de `float64`.

## Révision Leçon
- **Déclaration standard** : `var nom type = valeur`
- **Inférence de type** : `nom := valeur` (uniquement à l'intérieur d'une fonction).
- **Types numériques** : `int`, `float64`.
- **Zéro-valeur** : En Go, toute variable non initialisée reçoit une valeur par défaut (`0` pour int, `""` pour string, `false` pour bool).
- **Conversion** : Elle est toujours explicite, ex: `float64(monInt)`.

## Validation
Pour vérifier votre exercice, lancez la commande suivante dans le terminal :
```bash
go test ./01-variables/...
```
