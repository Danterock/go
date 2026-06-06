package main

import (
	"fmt"
	"strings"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

func main() {
	books := []Book{
		{1, "Война и мир", "Лев Толстой", 1869},
		{2, "Анна Каренина", "Лев Толстой", 1877},
		{3, "Преступление и наказание", "Фёдор Достоевский", 1866},
		{4, "Идиот", "Фёдор Достоевский", 1869},
		{5, "Мастер и Маргарита", "Михаил Булгаков", 1967},
		{6, "Тихий Дон", "Михаил Шолохов", 1940},
		{7, "Отцы и дети", "Иван Тургенев", 1862},
		{8, "Герой нашего времени", "Михаил Лермонтов", 1840},
		{9, "Евгений Онегин", "Александр Пушкин", 1833},
		{10, "Доктор Живаго", "Борис Пастернак", 1957},
	}

	var author string
	fmt.Print("Please input name of author: ")
	fmt.Scan(&author)

	result := searchBooksByAuthor(books, author)

	if len(result) == 0 {
		fmt.Println("No books found.")
	}

	for i := 0; i < len(result); i++ {
		fmt.Printf("%v %v - %v (%v)\n", result[i].ID, result[i].Title, result[i].Author, result[i].Year)
	}

}

func searchBooksByAuthor(books []Book, query string) []Book {
	result := []Book{}
	query = strings.ToLower(query)

	for i := 0; i < len(books); i++ {
		author := strings.ToLower(books[i].Author)
		if strings.Contains(author, query) {
			result = append(result, books[i])
		}
	}
	return result
}
