package web

import (
	"fmt"
	"github.com/Sfeir/golang-200/dao"
	"github.com/gorilla/mux"
)


// NewRouter creates a new router instance
func NewRouter() *mux.Router {
	// new router
	router := mux.NewRouter()

	// default JSON not found handler
	router.NotFoundHandler = NotFoundHandler()

	// no strict slash
	router.StrictSlash(false)


	// TODO fail fast, try to get the DAO of the required type and return (nil,error) if it fails
	// TODO don't forget to log the error and the parameters
	// task dao
	taskDao, err := dao.GetTaskDAO(db, "", dao.DAOMongo)
	if err != nil {
		panic("Initialization error:" + fmt.Sprintf("%v", err))
	}
	NewTaskController(taskDao, router)

	return &router
}
