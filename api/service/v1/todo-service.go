package v1

import (
	"database/sql"

	"fmt"
	v1 "github.com/iliketss/goweb.git/api/proto/v1"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer {

	fmt.Println()
	return &toDoServiceServer{
		db: db,
	}
}
