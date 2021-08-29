package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Book 1",
		Author:        "Author 1",
		YearPublished: 1985,
	},
	{
		ID:            2,
		Title:         "Book 2",
		Author:        "Author 2",
		YearPublished: 1985,
	},
	{
		ID:            3,
		Title:         "Book 3",
		Author:        "Author 3",
		YearPublished: 1985,
	},
	{
		ID:            4,
		Title:         "Book 4",
		Author:        "Author 4",
		YearPublished: 1985,
	},
	{
		ID:            5,
		Title:         "Book 5",
		Author:        "Author 5",
		YearPublished: 1985,
	},
	{
		ID:            6,
		Title:         "Book 6",
		Author:        "Author 6",
		YearPublished: 1985,
	},
	{
		ID:            7,
		Title:         "Book 7",
		Author:        "Author 7",
		YearPublished: 1985,
	},
	{
		ID:            8,
		Title:         "Book 8",
		Author:        "Author 8",
		YearPublished: 1985,
	},
	{
		ID:            9,
		Title:         "Book 9",
		Author:        "Author 9",
		YearPublished: 1985,
	},
	{
		ID:            10,
		Title:         "Book 10",
		Author:        "Author 10",
		YearPublished: 1985,
	},
}
