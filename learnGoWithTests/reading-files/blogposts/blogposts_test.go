package blogposts_test

import (
	"learnGoWithTests/reading-files/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		//A MapFS is a simple in-memory file system for use in tests,
		//represented as a map from path names (arguments to Open) to information
		//about the files or directories they represent.
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Error("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
