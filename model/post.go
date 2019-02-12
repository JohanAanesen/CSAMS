package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/database"
	"log"
	"time"
)

// Post struct is onyl for showcase
type Post struct {
	Id      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Created  time.Time `json:"created"`
}

// GetPosts shows how the database can be used without any global variables
func GetPosts() Post {
	// Database query
	rows, err := database.Get().Query("SELECT id, title, content, created FROM post")
	if err != nil {
		log.Println(err)
		return Post{}
	}
	// Close connection
	defer rows.Close()
	// Loop through rows
	for rows.Next() {
		// Declare empty Post
		var post = Post{}
		// Get data from rows
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Created)
		if err != nil {
			log.Println(err)
			return Post{}
		}
		// Return first result
		return post
	}

	return Post{}
}