package blogposts

import (
	"errors"
	"io/fs"
)

type Post struct {
	Title string
}

// NewPostsFromFS cannot use fstest.MapFS here as param because its a concrete impl not an abstract interface
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	//fs.ReadDir reads a dir inside given fs.FS returning []DirEntry

	if err != nil {
		return nil, err
	}
	var posts []Post

	//iterating over the entries, create a post for each one and return the slice
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no, I always fail")
}
