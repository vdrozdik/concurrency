package book

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
			"Published:\t%v", b.Title, b.Author, b.YearPublished,
	)
}

var Books = []Book{
	Book{
		ID:            1,
		Title:         "Title1",
		Author:        "Author1",
		YearPublished: 1991,
	},
	Book{
		ID:            2,
		Title:         "Title2",
		Author:        "Author2",
		YearPublished: 1992,
	},
	Book{
		ID:            2,
		Title:         "Title2",
		Author:        "Author2",
		YearPublished: 1992,
	},
	Book{
		ID:            3,
		Title:         "Title3",
		Author:        "Author3",
		YearPublished: 1993,
	},
	Book{
		ID:            4,
		Title:         "Title4",
		Author:        "Author4",
		YearPublished: 1994,
	},
	Book{
		ID:            4,
		Title:         "Title4",
		Author:        "Author4",
		YearPublished: 1994,
	},
	Book{
		ID:            5,
		Title:         "Title5",
		Author:        "Author5",
		YearPublished: 1995,
	},
	Book{
		ID:            6,
		Title:         "Title6",
		Author:        "Author6",
		YearPublished: 1996,
	},
	Book{
		ID:            7,
		Title:         "Title7",
		Author:        "Author7",
		YearPublished: 1997,
	},
	Book{
		ID:            8,
		Title:         "Title8",
		Author:        "Author8",
		YearPublished: 1998,
	},
	Book{
		ID:            9,
		Title:         "Title9",
		Author:        "Author9",
		YearPublished: 1999,
	},
	Book{
		ID:            10,
		Title:         "Title10",
		Author:        "Author10",
		YearPublished: 1990,
	},
}
