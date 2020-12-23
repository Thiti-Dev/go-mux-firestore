package entity

// Post is a DTO for Post collection
type Post struct{
	ID 		int64 	`json:"id"`
	Title 	string	`json:"title"`
	Text	string	`json:"text"`
}