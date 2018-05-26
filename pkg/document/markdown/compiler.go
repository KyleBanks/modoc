package markdown

import (
	"regexp"
	"strings"

	"github.com/KyleBanks/modoc/pkg/document"
)

const tableOfContentsTitle = "Table of Contents"

var linkRegexp = regexp.MustCompile("[^a-z0-9\\-]+")

// Compiler handles compiling Documents.
type Compiler struct {
	Header bool
	TOC    bool
	Body   bool
}

// Compile generates and returns the compiled contents of a Document.
func (c Compiler) Compile(d document.Document) (string, error) {
	var sb strings.Builder

	if c.Header {
		if err := c.compileHeader(&sb, d); err != nil {
			return "", err
		}
	}

	if c.TOC {
		if err := c.compileTOC(&sb, d); err != nil {
			return "", err
		}
	}

	if c.Body {
		if err := c.compileBody(&sb, d); err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func (c Compiler) compileHeader(sb *strings.Builder, d document.Document) error {
	sb.WriteString("# ")
	sb.WriteString(d.Title)
	sb.WriteString("\n\n")
	sb.WriteString(d.Body)
	sb.WriteString("\n\n")
	return nil
}

func (c Compiler) compileTOC(sb *strings.Builder, d document.Document) error {
	var write document.WalkSectionFn
	write = func(s document.Section, depth int) error {
		for i := 0; i < depth-1; i++ {
			sb.WriteString("   ")
		}
		sb.WriteString("- [")
		sb.WriteString(s.Title)
		sb.WriteString("](#")
		sb.WriteString(c.toLink(s.Title))
		sb.WriteString(")\n")

		return s.ForEachChild(write, depth+1)
	}

	sb.WriteString("## ")
	sb.WriteString(tableOfContentsTitle)
	sb.WriteString("\n\n")

	if err := d.ForEachChild(write, 1); err != nil {
		return err
	}
	sb.WriteString("\n")
	return nil
}

func (c Compiler) compileBody(sb *strings.Builder, d document.Document) error {
	var write document.WalkSectionFn
	write = func(s document.Section, depth int) error {
		for i := 0; i <= depth; i++ {
			sb.WriteString("#")
		}
		sb.WriteString(" ")
		sb.WriteString(s.Title)
		sb.WriteString("\n\n")
		sb.WriteString(s.Body)
		sb.WriteString("\n\n")

		return s.ForEachChild(write, depth+1)
	}

	return d.ForEachChild(write, 1)
}

func (Compiler) toLink(title string) string {
	link := title
	link = strings.ToLower(link)
	link = strings.Replace(link, " ", "-", -1)
	link = linkRegexp.ReplaceAllString(link, "")
	return link
}
