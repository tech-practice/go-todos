package models

type Todo struct {
	UserID    int         `json:"userId" bson:"userId"`
	ID        interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string      `json:"title" bson:"title"`
	Completed bool        `json:"completed" bson:"completed"`
}

type TodoUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Todo
}

type TodoDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
