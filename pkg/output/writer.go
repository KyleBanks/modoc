package output

// Writer defines a type that can write a string to a particular location, such
// as a file.
type Writer interface {
	Write(location string, contents string) error
}
