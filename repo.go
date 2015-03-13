package main

import (
	"fmt"
	"time"
)

var currentId int

var posts Posts

// Give us some seed data
func init() {
	/*
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	*/
	CreatePost(Post{
		User: User{
			Username: "pieohpah",
			Email: "jo.chasinga@gmail.com",
		},
		Topic: "My First Post",
		Text:  "Hello everyone! This is awesome.",
		//Comment: Comment{},
		Timestamp: time.Now(),
	})
	
	CreatePost(Post{
		User: User{
			Username: "IronMan",
			Email: "iron_mann@hotmale.com",
		},
		Topic: "Greeting",
		Text: "Greetings from Ironman",
		Comment: Comment{},
		Timestamp: time.Now(),
	})
}

func FindPost(id int) Post {
	for _, t := range posts {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Post{}
}

func CreatePost(p Post) Post {
	currentId += 1
	p.Id = currentId
	posts = append(posts, p)
	return p
}

func DestroyPost(id int) error {
	for i, p := range posts {
		if p.Id == id {
			posts = append(posts[:i], posts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Couldn't find Todo with id of %d to delete", id)
}
