package image

import "io"

type Manager interface {
	Create(name string, image io.ReadCloser) (*Image, error)
	List() ([]Image, error)
	Remove(name string) error
}

type Image struct {
	Name     string
	Location string
}
