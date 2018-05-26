# modoc

[![GoDoc](https://godoc.org/github.com/KyleBanks/modoc?status.svg)](https://godoc.org/github.com/KyleBanks/modoc)&nbsp; 
[![Build Status](https://travis-ci.org/KyleBanks/modoc.svg?branch=master)](https://travis-ci.org/KyleBanks/modoc)&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/modoc)](https://goreportcard.com/report/github.com/KyleBanks/modoc)&nbsp;
[![Coverage Status](https://coveralls.io/repos/github/KyleBanks/modoc/badge.svg?branch=master)](https://coveralls.io/github/KyleBanks/modoc?branch=master)

`modoc` is the **M**aster **O**f **D**ocument **O**rganization and **C**ompilation *(and definitely not just a play on the name [MODOK](https://en.wikipedia.org/wiki/MODOK))*.

What does that mean? It means `modoc` allows you to compile a large Markdown project from an organized folder structure. In fact, this documentation you're currently reading was compiled by modoc from the [examples/readme/](./examples/readme) directory:

```
.
├── 1-features
│   ├── README.md
│   └── TITLE
├── 2-usage
│   ├── 1-install
│   │   ├── README.md
│   │   └── TITLE
│   ├── 2-command-line
│   │   ├── 1-init
│   │   │   ├── README.md
│   │   │   └── TITLE
│   │   ├── 2-compile
│   │   │   ├── README.md
│   │   │   └── TITLE
│   │   ├── README.md
│   │   └── TITLE
│   ├── 3-project-structure
│   │   ├── README.md
│   │   └── TITLE
│   ├── 4-examples
│   │   ├── README.md
│   │   └── TITLE
│   └── TITLE
├── 3-authors
│   ├── 1-contributing
│   │   ├── README.md
│   │   └── TITLE
│   ├── README.md
│   └── TITLE
├── 4-license
│   ├── README.md
│   └── TITLE
├── README.md
└── TITLE
```


## Table of Contents

- [Features](#features)
- [Usage](#usage)
   - [Install](#install)
   - [Command Line](#command-line)
      - [init](#init)
      - [compile](#compile)
   - [Project Structure](#project-structure)
   - [Examples](#examples)
- [Authors](#authors)
   - [Contributing](#contributing)
- [License](#license)

## Features

- **organize** your projects into smaller, digestable sections
- nest sections and subsections as deep as you like
- generate a **table of contents** for your document

## Usage



### Install

To install `modoc` from source you'll need a working Go environment:

```
$ go get -u github.com/KyleBanks/modoc
```

### Command Line

First up, definitely check `modoc` usage like so: 

```
$ modoc --help
```


#### init

If you're starting a new project, try the `init` command:

```
$ modoc init [--path ./optional/path]
```

This will generate a sample project in the current directory (or you can provide an optional `--path`) to help you get started right away. You can also reference the [examples](./examples) directory if you want to create your own project manually.

#### compile

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

### Project Structure

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

### Examples

You can generate this README like so:

```
$ modoc compile --source ./examples/readme --output ./README.md
```

Compiling the [basic example](./examples/basic) will produce the following:

```
$ modoc compile --source ./examples/basic --output ./examples/BASIC.md
$ cat ./examples/BASIC.md
# A Modoc Tale

This is the introduction to the document, found below the title and before the Table of Contents.

## Table of Contents

- [Introduction](#introduction)
   - [Subsection A](#subsection-a)
   - [Subsection B](#subsection-b)
   - [Subsection C](#subsection-c)
- [Chapter One](#chapter-one)
   - [Subsection A](#subsection-a)
   - [Subsection B](#subsection-b)
   - [Subsection C](#subsection-c)
- [Chapter Two](#chapter-two)
   - [Subsection A](#subsection-a)
   - [Subsection B](#subsection-b)
   - [Subsection C](#subsection-c)

## Introduction

This is the contents of the Introduction chapter.

### Subsection A

This is Introduction, Subsection A.

### Subsection B

This is Introduction, Subsection B.

### Subsection C

This is Introduction, Subsection C.

## Chapter One

This is the contents of Chapter One.

### Subsection A

This is Chapter 1, Subsection A.

### Subsection B

This is Chapter 1, Subsection B.

### Subsection C

This is Chapter 1, Subsection C.

## Chapter Two

This is the contents of Chapter Two.

### Subsection A

This is Chapter 2, Subsection A.

### Subsection B

This is Chapter 2, Subsection B.

### Subsection C

This is Chapter 2, Subsection C.
```

## Authors

`modoc` was developed by:

- [Kyle Banks](https://twitter.com/kylewbanks)

### Contributing

If you'd like to contribute to `modoc` then take a look at the [issues list](https://github.com/KyleBanks/modoc/issues). If you have an idea/suggestion/bug/question/etc. then feel free to create an issue to discuss it - you can also submit a pull request but it's probably better to chat before you spend time making big changes. 

Also, feel free to add your name and twitter to the authors list above if you'd like a little recognition for your effort!

## License

`modoc` is made available under the [MIT License](./LICENSE).

