package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const usd = 83.92
const eur = 94.04

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        var rubles float64
        fmt.Print("Ввести рубли: ")
        _, err := fmt.Scanln(&rubles)
        if err != nil || rubles <= 0 {
            fmt.Println("Ошибка: введите корректное число!")
            reader.ReadString('\n')
            continue
        }

        fmt.Printf("Доллары: %.2f\n", RubToUsd(rubles))
        fmt.Printf("Евро: %.2f\n", RubToEur(rubles))

        var convertAgain string
        fmt.Print("Посчитать заново? (yes/no): ")
        fmt.Scanln(&convertAgain)

        convertAgain = strings.ToLower(convertAgain)
        if convertAgain == "yes" || convertAgain == "y" {
            continue
        } else if convertAgain == "no" || convertAgain == "n" {
            break
        }
    }
}

func RubToUsd(rubles float64) float64 {
    return rubles / usd
}

func RubToEur(rubles float64) float64 {
    return rubles / eur
}