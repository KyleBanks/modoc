package document

// Sample demonstrates an example of a fleshed out Document.
var Sample = Document{
	Title: "A Modoc Tale",
	Body:  "This is the introduction to the  found below the title and before the Table of Contents.",

	Children: []Section{
		{Folder: "1-introduction", Document: Document{Title: "Introduction", Body: "This is the contents of the Introduction chapter.", Children: []Section{
			{Folder: "1-subsection-a", Document: Document{Title: "Subsection A", Body: "This is Introduction, Subsection A."}},
			{Folder: "2-subsection-b", Document: Document{Title: "Subsection B", Body: "This is Introduction, Subsection B."}},
			{Folder: "3-subsection-c", Document: Document{Title: "Subsection C", Body: "This is Introduction, Subsection C."}},
		}}},
		{Folder: "2-chapter-one", Document: Document{Title: "Chapter One", Body: "This is the contents of Chapter One.", Children: []Section{
			{Folder: "1-subsection-a", Document: Document{Title: "Subsection A", Body: "This is Chapter 1, Subsection A."}},
			{Folder: "2-subsection-b", Document: Document{Title: "Subsection B", Body: "This is Chapter 1, Subsection B."}},
			{Folder: "3-subsection-c", Document: Document{Title: "Subsection C", Body: "This is Chapter 1, Subsection C."}},
		}}},
		{Folder: "3-chapter-two", Document: Document{Title: "Chapter Two", Body: "This is the contents of Chapter Two.", Children: []Section{
			{Folder: "1-subsection-a", Document: Document{Title: "Subsection A", Body: "This is Chapter 2, Subsection A."}},
			{Folder: "2-subsection-b", Document: Document{Title: "Subsection B", Body: "This is Chapter 2, Subsection B."}},
			{Folder: "3-subsection-c", Document: Document{Title: "Subsection C", Body: "This is Chapter 2, Subsection C."}},
		}}},
	},
}
