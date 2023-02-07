package binaryRepository

type BinaryRepository interface {
	Save(data []byte, filename string) (string, error)
	Load(link string) ([]byte, error)
}
