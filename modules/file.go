package modules

import (
	"os"
)

type File struct {

}

func NewFile() *File {
	return &File{}
}

func (fs *File) IsExists(pathname string) (ok bool, err error) {
	_, err = os.Stat(pathname)
	if err == nil {
		ok = true
		return
	}

	if os.IsNotExist(err) {
		ok, err = false, nil
		return
	}

	ok = false
	return
}

func (fs *File) Create(pathname string) (err error) {
	f, err := os.Create(pathname)
	defer f.Close()

	return
}


