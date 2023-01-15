package storage

// Storage for the input, output, and plugin files for both the coordinator and worker nodes
type Storage interface {
	Write(name string, content string) error
	Read(name string) (string, error)
}
