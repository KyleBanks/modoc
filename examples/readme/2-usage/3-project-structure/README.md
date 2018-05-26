Each directory represents a section or chapter of the document, with child directories being subsections, and so on.

- `TITLE` *(plaintext, required)*: contains the section title
- `README.md` *(markdown, optional)*: contains the section content

`modoc` traverses your section hierarchy in alphabetical order, so it's recommended that you prefix your folders with an incrementing counter at each level to preserve ordering, like so:

```
examples/basic/
├── 1-introduction
│   ├── 1-subsection-a
│   ├── 2-subsection-b
│   ├── 3-subsection-c
├── 2-chapter-one
│   ├── 1-subsection-a
│   ├── 2-subsection-b
│   ├── 3-subsection-c
├── 3-chapter-two
│   ├── 1-subsection-a
│   ├── 2-subsection-b
│   ├── 3-subsection-c
```