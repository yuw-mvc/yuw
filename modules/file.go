package modules

import (
	"github.com/spf13/cast"
	E "github.com/yuw-mvc/yuw/exceptions"
	"os"
)

type File struct {

}

func NewFile() *File {
	return &File{}
}

func (fs *File) Open(pathname string) (f *os.File, err error) {
	return os.Open(pathname)
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

func (fs *File) Writer(pathname string, content ...interface{}) (err error) {
	if len(content) < 1 {
		err = E.Err("yuw^m_fs_a", E.ErrPosition())
		return
	}

	f, err := os.OpenFile(pathname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		return
	}

	for _, val := range content {
		_, err := f.WriteString(cast.ToString(val))
		if err != nil {
			f.WriteString(err.Error())
		}

		f.WriteString("\n")
	}

	return
}


