package blogposts

import (
	"io/fs"
)

type Post struct {
}

// NewPostsFromFS cannot use fstest.MapFS here as param because its a concrete impl not an abstract interface
func NewPostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	//fs.ReadDir reads a dir inside given fs.FS returning []DirEntry
	var posts []Post

	//iterating over the entries, create a post for each one and return the slice
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
