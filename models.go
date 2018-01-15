package main

import (
	"time"
)

// User represents a blog user.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Comment represents a comment a user can post in a blog post.
type Comment struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

// Post represents a blog post.
type Post struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Topic     string    `json:"topic"`
	Text      string    `json:"text"`
	Comment   Comment   `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
}

// Posts represents a list of Posts of a User.
type Posts []Post

// Comments represents a list of Comments in a Post.
type Comments []Comment

// Users represents a list of Users of this blog.
type Users []User
