package dao_test // /!\ don't change this! create a new test file instead.

import (
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	"testing"
)

func TestDAOMock(t *testing.T) {

	daoMock, err := dao.GetTaskDAO("", "", dao.DAOMock)
	if err != nil {
		t.Error(err)
	}

	// TODO create a new task called "toSave" for testing purpose
	//taskPrivate := model.NewTaskPrivate()
	//fmt.Println(taskPrivate.PrintMe())

	toSave := model.NewTask()
	toSave.Title = "do it now!"
	// TODO save the "toSave" task
	// TODO check the error
	daoMock.Save(toSave)

	tasks, err := daoMock.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	// TODO get "toSave" task by ID and verify that it is successfully retrieved
	tById, _ := daoMock.GetByID(toSave.ID)
	if tById == nil || tById.ID != toSave.ID {
		t.Error("GetByID error", tById)
	}
	// TODO delete the "toSave" task and verify with a get by ID that it is removed
	tsByTitle, _ := daoMock.GetByTitle("do it now!")
	if tsByTitle == nil || len(tsByTitle) != 1 {
		t.Error("GetByTitle error", tsByTitle)
	} else {
		println("len", len(tsByTitle))
		if tsByTitle[0].ID != toSave.ID {
			t.Error("GetByTitle bad id")
		}
	}

}
