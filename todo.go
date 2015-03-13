package main

import (
	"time"
)

type User struct {
	Id 		 int    	`json:"id"`
	Username string 	`json:"username"`
	Email    string 	`json:"email"`
}

type Comment struct {
	Id        int 		`json:"id"`
	User      User 		`json:"user"`
	Text      string 	`json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type Post struct {
	Id        int		`json:"id"`
	User      User		`json:"user"`
	Topic     string	`json:"topic"`
	Text      string	`json:"text"`
	Comment   Comment   `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
}

type Posts    []Post
type Comments []Comment
type Users    []User


