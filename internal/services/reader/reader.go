package reader

type Reader interface {
	Read(filePath string) error
}
