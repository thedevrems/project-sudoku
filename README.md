# Raid Sudoku - Nantes Ynov Campus

## Description du Projet

Ce projet a pour objectif de résoudre un puzzle de Sudoku en utilisant l'algorithme de **retour sur trace** (backtracking), implémenté en **Golang**. Le retour sur trace est une méthode de recherche en profondeur, souvent employée dans la résolution de puzzles comme le Sudoku, où chaque cellule est testée pour une valeur valide. Ce projet vise à démontrer cette technique en trouvant la solution pour n'importe quelle grille de Sudoku valide.


## Fonctionnalités

- Possibilité de résoudre des grilles standard (9x9).
- Affichage du Sudoku s'il y a une solution
- Gestion des Erreurs
- Vérification des contraintes de lignes, colonnes et sous-grilles (régions 3x3).

## Algorithme Utilisé

Nous utilisons l'algorithme de **retour sur trace** (backtracking), qui suit ce processus :

1. Chaque cellule vide __(une cellule vide est représenté par un ".")__ est visitée une par une.
2. Un chiffre entre 1 et 9 est placé dans la cellule, et le programme vérifie si ce chiffre respecte les règles du Sudoku.
3. Si une contrainte est violée (ligne, colonne ou sous-grille), le programme essaye le chiffre suivant.
4. Si aucun chiffre ne fonctionne pour une cellule, le programme revient à la cellule précédente et incrémente la valeur de cette cellule.
5. Ce processus continue jusqu'à ce que la grille entière soit résolue.

L'algorithme garantit de trouver une solution sinon dans le cas contraire, la console affichera une erreur

## Prérequis

Avant d'exécuter le projet, assurez-vous que les éléments suivants sont installés sur votre machine :

- **Golang** (version 1.23 ou plus récente)

- **Git** pour cloner le dépôt

## Installation

Pour installer et exécuter le projet, suivez les étapes ci-dessous :

1. Clonez le dépôt du projet à l'aide de Git :
   ```bash
   git clone https://ytrack.learn.ynov.com/git/gremy/raid-sudoku.git
   ```
## Test notre projet:
Voici quelques commandes pour tester notre projet :

1. Sudoku Réalisable

    ```bash
    go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
    ```
2. Sudoku Non Réalisable
    ```bash
    go run . 1 2 3 4
    go run .
    go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
    ```

Attention les commandes doivent être réalisés dans le dossier main du projet

