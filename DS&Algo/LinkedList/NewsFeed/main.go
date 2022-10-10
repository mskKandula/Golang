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

func (f *Feed) Insert(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		var previousPost *Post
		currentPost := f.start

		for currentPost.publishedDate < newPost.publishedDate {
			previousPost = currentPost
			currentPost = currentPost.next
		}

		previousPost.next = newPost
		newPost.next = currentPost

	}
	f.length += 1
}

func (f *Feed) Inspect() {
	if f.length == 0 {
		fmt.Println("Feed is empty")
	}
	fmt.Println("========================")
	fmt.Println("Feed Length: ", f.length)

	var currentIndex int8 = 0
	currentPost := f.start

	for currentIndex < f.length {
		fmt.Printf("Item: %v - %v\n", currentIndex, currentPost)
		currentPost = currentPost.next
		currentIndex++
	}
	fmt.Println("========================")
}

func main() {
	rightNow := time.Now().Unix()

	f := &Feed{}

	p1 := Post{
		body:          "First Post()",
		publishedDate: rightNow,
	}
	f.Append(&p1)

	p2 := Post{
		body:          "Second Post",
		publishedDate: rightNow + 20,
	}
	f.Append(&p2)

	f.Inspect()

	p3 := &Post{
		body:          "This is a new post",
		publishedDate: rightNow + 15,
	}
	f.Insert(p3)

	f.Inspect()
}
