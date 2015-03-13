package main

import (
	"fmt"
	"time"
	"encoding/json"
	"strconv"
	
	"github.com/garyburd/redigo/redis"
)

var currentId int

var posts Posts

// Give us some seed data
func init() {
	/*
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	*/
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	
	p1 := CreatePost(Post{
			User: User{
					Username: "pieohpah",
					Email: "jo.chasinga@gmail.com",
			},
			Topic: "My First Post",
			Text:  "Hello everyone! This is awesome.",
			//Comment: Comment{},
			Timestamp: time.Now(),
	})
	
	p2 := CreatePost(Post{
			User: User{
					Username: "IronMan",
					Email: "iron_mann@hotmale.com",
			},
			Topic: "Greeting",
			Text: "Greetings from Ironman",
			Comment: Comment{},
			Timestamp: time.Now(),
	})
	
	SetPost(c, p1)
	SetPost(c, p2)
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

func SetPost(c redis.Conn, p Post) {
	
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	
	// Save a whole blob of JSON
	reply, err := c.Do("SET", "post:" + strconv.Itoa(p.Id), b)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
	
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
