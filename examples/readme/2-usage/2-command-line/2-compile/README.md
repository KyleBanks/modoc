With a project initialized, you can compile it into a single Markdown file:

```
$ modoc compile [--source ./optional/path] [--output ./OPTIONAL.md]
```

This will compile the current directory (or optional `--source`) into a single file called `COMPILED.md` - of course, you can name it whatever you like using the `--output` flag. 

By default the file will contain:

- **Header**: the root `TITLE` followed by the root `README.md`
- **ToC**: the table of contents
- **Body**: the titles and contents of each section, subsection, etc.

You can disable generating any of these components using the following flags:

```
$ modoc compile --header false --toc false --body false
```