package markdown

// Writer defines a type that can output a string to a particular location.
type Writer interface {
	Write(path, content string) error
}
