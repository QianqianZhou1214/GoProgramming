package blogposts

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

// NewPostsFromFS cannot use fstest.MapFS here as param because its a concrete impl not an abstract interface
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	//fs.ReadDir reads a dir inside given fs.FS returning []DirEntry

	if err != nil {
		return nil, err //todo: needs clarification, should we totally fail if one file fails
	}
	var posts []Post

	//iterating over the entries, create a post for each one and return the slice
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	// fs.FS gives us a way(.Open()) of opening file within it by name
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

// not coupling to an fs.File, but using io.Reader
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descriptionSeparator):]

	return Post{Title: title, Description: description}, nil
	/*
		scanner.Scan()
		titleLine := scanner.Text()

		scanner.Scan()
		descriptionLine := scanner.Text()

		return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil

	*/
	/*
		postData, err := io.ReadAll(postFile)
		if err != nil {
			return Post{}, err
		}

		post := Post{Title: string(postData)[7:]}
		return post, nil

	*/
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no, I always fail")
}
