package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func startHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
	`<html>
		<head>
			<title>Сервер запущен</title>
		</head>
		<body>
			<h1>Сервер успешно запущен!</h1>
			<button onclick="window.location.href='/about'">Страница about</button>
			<button onclick="window.location.href='/generator'">Генератор чисел</button>
			<button onclick="window.location.href='/converter'">Конвертер</button>
		</body>
	</html>
	`)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
			<head>
				<title>About</title>
			</head>
			<body>
				<h1>О сайте</h1>
				<p>Это страница About.</p>
				<p>Сайт создан на Go 1.27.1 с использованием библиотеки net/http.</p>
				<p>Подшивалов Андрей, ПИ-332</p>
				<button onclick="window.location.href='/'">На главную</button>
			</body>
		</html>
	`)
}

func generatorHandler(w http.ResponseWriter, r *http.Request) {
	randomNumber1 := rand.Intn(100)
	randomNumber2 := rand.Intn(100)
	randomSumm := randomNumber1 + randomNumber2
	randomMult := randomNumber1 * randomNumber2
	fmt.Fprintf(w,
	`<html>
		<head>
			<title>Генератор чисел</title>
		</head>
		<body>
			<h1>Генератор чисел</h1>
			<p>Первое случайное число: %d</p>
			<p>Второе случайное число: %d</p>
			<p>Сумма чисел: %d</p>
			<p>Произведение чисел: %d</p>
			<button onclick="window.location.href='/generator'">Сгенерировать новое число</button>
			<button onclick="window.location.href='/'">На главную</button>
		</body>
	</html>
	`, randomNumber1, randomNumber2, randomSumm, randomMult)
}

func converterHandler(w http.ResponseWriter, r *http.Request) {
    const usdRate = 82.90
    const eurRate = 93.03

    if r.Method != http.MethodGet && r.Method != http.MethodPost {
        w.Header().Set("Allow", "GET, POST")
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }

    var result string
    if r.Method == http.MethodPost {
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Не удалось прочитать сумму", http.StatusBadRequest)
            return
        }

        rubles, err := strconv.ParseFloat(r.FormValue("rub"), 64)
        if err != nil || math.IsNaN(rubles) || math.IsInf(rubles, 0) || rubles < 0 {
            http.Error(w, "Введите корректную сумму в рублях", http.StatusBadRequest)
            return
        }

        result = fmt.Sprintf(
            "<p>Доллары: %.2f</p><p>Евро: %.2f</p>",
            rubles/usdRate,
            rubles/eurRate,
        )
    }

    fmt.Fprintf(w, `<!DOCTYPE html>
    <html>
        <head>
            <meta charset="UTF-8">
            <title>Конвертер</title>
        </head>
        <body>
            <h1>Конвертер валют</h1>
            <form method="POST" action="/converter">
                <label for="rub">Сумма в рублях:</label>
                <input id="rub" name="rub" type="number" min="0" step="0.01" required>
                <button type="submit">Сконвертировать</button>
            </form>
            %s
            <button onclick="window.location.href='/'">На главную</button>
        </body>
    </html>`, result)
}

func main() {
    http.HandleFunc("/", startHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/generator", generatorHandler)
	http.HandleFunc("/converter", converterHandler)
	fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
