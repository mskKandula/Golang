package main

import (
	"fmt"
	"time"
)

type Post struct {
	publishedDate int64
	body          string
	next          *Post
}

type Feed struct {
	length int8
	start  *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		currentPost := f.start

		for currentPost.next != nil {
			currentPost = currentPost.next
		}

		currentPost.next = newPost
	}
	f.length += 1
}

func main() {
	f := &Feed{}

	p1 := Post{
		body:          "First Post()",
		publishedDate: time.Now().Unix(),
	}
	f.Append(&p1)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)

	p2 := Post{
		body:          "Second Post",
		publishedDate: time.Now().Unix(),
	}
	f.Append(&p2)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	fmt.Printf("Second: %v\n", f.start.next)
}
