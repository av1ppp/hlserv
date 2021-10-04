package hlserv

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type File struct {
	name  string
	mtime time.Time
	data  []byte

	offset int
}

func (file *File) Size() int {
	return len(file.data)
}

func (file *File) ModTime() time.Time {
	return file.mtime
}

func (file *File) Read(p []byte) (int, error) {
	if file.offset >= file.Size() {
		return 0, io.EOF
	}

	n := copy(p, file.data[file.offset:])

	file.offset += n
	return n, nil
}

func (file *File) Reset() {
	file.offset = 0
}

// func (file *File) Seek(offset int64, whence int) (int64, error) {
// }

type Store struct {
	files map[string]*File

	sync.Mutex
}

func newStore() *Store {
	return &Store{files: make(map[string]*File)}
}

var store = newStore()

func (store *Store) File(name string) (*File, error) {
	store.Lock()
	defer store.Unlock()

	file, find := store.files[name]
	if !find {
		return nil, os.ErrNotExist
	}

	return file, nil
}

func (store *Store) Create(name string, data []byte) error {
	store.Lock()
	defer store.Unlock()

	if _, find := store.files[name]; find {
		return os.ErrExist
	}

	if name == "playlist.m3u8" {
		fmt.Println(string(data))
	}

	store.files[name] = &File{
		name:   name,
		mtime:  time.Now(),
		data:   data,
		offset: 0,
	}

	return nil
}

func (store *Store) Write(name string, data []byte) error {
	file, err := store.File(name)

	// create
	if os.IsNotExist(err) {
		if err := store.Create(name, data); err != nil {
			return err
		}
		return nil
	}

	// or change mtime/data/offset
	file.mtime = time.Now()
	file.data = data
	file.offset = 0

	return nil
}

func (store *Store) Remove(name string) error {
	store.Lock()
	defer store.Unlock()

	if _, find := store.files[name]; !find {
		return os.ErrNotExist
	}

	delete(store.files, name)
	return nil
}
