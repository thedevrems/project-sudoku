package sudoku

import "fmt"

// liste des couleurs disponibles
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// Fonction qui s'occupe d'afficher une erreur, avec le nom de l'erreur : typeOfError
func PrintError(typeOfError string) {
	fmt.Println(Red, "Erreur :", Reset, typeOfError)
}
