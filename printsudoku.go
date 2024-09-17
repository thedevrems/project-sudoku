package sudoku

import (
    "fmt"
)

const (
    green  = "\033[32m" // Vert
    blue   = "\033[34m" // Bleu
    reset  = "\033[0m"  // Réinitialiser la couleur
)

// Cette fonction a pour but d'afficher la grille de sudoku à la fin de la résolution
func PrintSudoku(grille [][]string) {
    for y := 0; y < len(grille); y++ {
        for x := 0; x < len(grille[y]); x++ {
            // Affichage du caractère en vert
            char := rune(grille[y][x][0])
            fmt.Printf("%s%c%s", green, char, reset) // Utilisation de fmt pour afficher en couleur

            if x != len(grille[y])-1 {
                fmt.Print(" ") // Utilisation de fmt.Print pour l'espace
            }
            // Ajout d'un séparateur vertical en bleu après chaque cellule de 3x3
            if (x+1)%3 == 0 && x != len(grille[y])-1 {
                fmt.Printf("%s| %s", blue, reset) // Utilisation de fmt pour le séparateur
            }
        }
        // Retour à la ligne après chaque ligne
        fmt.Println() // Utilisation de fmt.Println pour le retour à la ligne

        // Ajout d'un séparateur horizontal en bleu après chaque groupe de 3 lignes
        if (y+1)%3 == 0 && y != len(grille)-1 {
            fmt.Printf("%s", blue) // Commencer la couleur bleue
            for i := 0; i < 21; i++ { // 21 tirets pour séparer les groupes de 3 lignes
                fmt.Print("-")
            }
            fmt.Printf("%s\n", reset) // Réinitialiser la couleur et retour à la ligne
        }
    }
}