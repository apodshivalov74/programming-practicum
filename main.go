package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const storageFile = "bookmarks.json"

var bookmarks = make(map[string]string)

func main() {
	if err := loadBookmarks(); err != nil {
		fmt.Printf("Ошибка загрузки данных: %v\n", err)
	}

	http.HandleFunc("/", startHandler)
	http.HandleFunc("/newBookmark", newBookmarkHandler)
	http.HandleFunc("/deleteBookmark", deleteBookmarkHandler)

	fmt.Println("Сервер запущен на http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func loadBookmarks() error {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			bookmarks["Google"] = "https://www.google.com"
			bookmarks["GitHub"] = "https://www.github.com"
			return saveBookmarks()
		}
		return err
	}

	return json.Unmarshal(data, &bookmarks)
}

func saveBookmarks() error {
	data, err := json.MarshalIndent(bookmarks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	query := r.URL.Query().Get("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<html>
		<head>
			<title>Мои закладки</title>
		</head>
		<body>
			<h1>Менеджер закладок</h1>
			<form method="GET" action="/" style="margin-bottom: 16px;">
				<input type="text" name="q" value="%s" placeholder="Поиск по названию или ссылке...">
				<button type="submit">Найти</button>
				<a href="/"><button type="button">Сбросить</button></a>
			</form>

			%s
			<br>
			<button onclick="window.location.href='/newBookmark'">Добавить закладку</button>
		</body>
	</html>
	`, query, showBookmarks(bookmarks, query))
}

func newBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		url := r.FormValue("url")

		if title != "" && url != "" {
			addBookmark(bookmarks, title, url)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<html>
		<head>
			<title>Новая закладка</title>
		</head>
		<body>
			<h1>Добавить новую закладку</h1>
			<form method="post" action="/newBookmark">
				<label for="title">Название:</label>
				<input type="text" id="title" name="title" required><br><br>
				<label for="url">URL:</label>
				<input type="url" id="url" name="url" required><br><br>
				<button type="submit">Добавить</button>
			</form>
		</body>
	</html>
	`)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		if title != "" {
			deleteBookmark(bookmarks, title)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addBookmark(b map[string]string, title string, url string) {
	b[title] = url
	_ = saveBookmarks()
}

func deleteBookmark(b map[string]string, title string) {
	delete(b, title)
	_ = saveBookmarks()
}

func showBookmarks(b map[string]string, query string) string {
	result := ""
	queryLower := strings.ToLower(strings.TrimSpace(query))
	count := 0

	for title, url := range b {
		if queryLower != "" {
			titleMatches := strings.Contains(strings.ToLower(title), queryLower)
			urlMatches := strings.Contains(strings.ToLower(url), queryLower)
			if !titleMatches && !urlMatches {
				continue
			}
		}

		count++
		result += fmt.Sprintf(`
			<div style="margin-bottom: 8px;">
				<span>%s: <a target="_blank" href="%s">%s</a></span>
				<form method="POST" action="/deleteBookmark" style="display:inline; margin-left: 6px;">
					<input type="hidden" name="title" value="%s">
					<button type="submit">X</button>
				</form>
			</div>`, title, url, url, title)
	}

	if count == 0 {
		return "<p>Ничего не найдено.</p>"
	}

	return result
}
