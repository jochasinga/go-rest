package main

import (
	"fmt"
	"time"
	"encoding/json"
	"strconv"
	
	"github.com/garyburd/redigo/redis"
)

var currentPostId int
var currentUserId int

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

// Give us some seed data
func init() {
	
	CreatePost(Post{
			User: User{
					Username: "pieohpah",
					Email: "jo.chasinga@gmail.com",
			},
			Topic: "My First Post",
			Text:  "Hello everyone! This is awesome.",
	})
	
	CreatePost(Post{
			User: User{
					Username: "IronMan",
					Email: "iron_mann@hotmale.com",
			},
			Topic: "Greeting",
			Text: "Greetings from Ironman",
	})
}

func FindAll() Posts {

	c := RedisConnect()
	defer c.Close()
	
	keys, err := c.Do("KEYS", "post:*")

	HandleError(err)
	
	var posts Posts
	
	for _, k := range keys.([]interface{}) {
		
		var post Post
		
		reply, err := c.Do("GET", k.([]byte))

		HandleError(err)
		
		if err := json.Unmarshal(reply.([]byte), &post); err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	return posts
}

func FindPost(id int) Post {
	
	var post Post

	c := RedisConnect()
	defer c.Close()
	
	reply, err := c.Do("GET", "post:" + strconv.Itoa(id))

	HandleError(err)
	
	fmt.Println("GET OK")
	
	if err = json.Unmarshal(reply.([]byte), &post); err != nil {
		panic(err)
	}
	return post
}

func CreatePost(p Post) {
	
	currentPostId += 1
	currentUserId += 1
	
	p.Id = currentPostId
	p.User.Id = currentUserId
	p.Timestamp = time.Now()
	
	c := RedisConnect()
	defer c.Close()
	
	b, err := json.Marshal(p)

	HandleError(err)
	
	// Save JSON blob to Redis
	reply, err := c.Do("SET", "post:" + strconv.Itoa(p.Id), b)

	HandleError(err)
	
	fmt.Println("GET ", reply)
	
}

func DeletePost(id int) {
	
	c := RedisConnect()
	defer c.Close()
	
	reply, err := c.Do("DEL", "post:" + strconv.Itoa(id))
	HandleError(err)
	
	if reply.(int) != 1 {
		fmt.Println("No post removed")
	} else {
		fmt.Println("Post removed")
	}
}
