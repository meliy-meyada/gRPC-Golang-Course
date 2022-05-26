package main

import (
	pb "github.com/meliy-meyada/gRPC-Golang-Course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID 				primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId 	string 						 `bson:"author_id"`
	Title 		string 						 `bson:"title"`
	Content 	string						 `bson:"content"`
}

func documemtToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id: data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title: data.Title,
		Content: data.Content,
	}
}