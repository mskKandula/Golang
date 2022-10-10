package main

import (
	"errors"
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
	end    *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length += 1
}

func (f *Feed) Remove(publishedDate int64) {
	if f.length == 0 {
		panic(errors.New("Feed is Empty"))
	}

	var previousPost *Post
	currentPost := f.start

	for currentPost.publishedDate != publishedDate {
		if currentPost.next == nil {
			panic(errors.New("No such post found"))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}

	previousPost.next = currentPost.next
	f.length--
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
