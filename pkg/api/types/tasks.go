package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TaskID      int32              `json:"task_id" bson:"task_id"`
	TaskTitle   string             `json:"task_title" bson:"task_title"`
	UserID      string             `json:"user_id" bson:"user_id"`
	ProjectID   int64              `json:"project_id" bson:"project_id"`
	TaskContent string             `json:"task_content" bson:"task_content"`
	Tags        []string           `json:"tags" bson:"tags"`
	Status      string             `json:"status" bson:"status"`
	DueDate     string             `json:"due_date" bson:"due_date"`
	CreateTime  string             `json:"create_time" bson:"create_time"`
	UpdateTime  string             `json:"update_time" bson:"update_time"`
}
