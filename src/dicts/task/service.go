package task

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/handle"
	"github.com/microcosm-cc/bluemonday"
)

type HandlerTaskService interface {
	FindTaskByID(id customType.StringUUID) (*Task, error)
	FetchTasksByCardIDs(cardIDs []string) (tasks []Task, err error)
	CreateTask(data []byte) (*Task, error)
	UpdateTask(data []byte, id customType.StringUUID) (*Task, error)
	DeleteTask(id customType.StringUUID) error
	FetchTasks(limit, offset int) (tasks []Task, err error)
}

type service struct {
	handle.HandlerImpl
	db        database.IDataManager
	sanitizer *bluemonday.Policy
}

func CreateInstance(db database.IDataManager) HandlerTaskService {
	sanitizer := bluemonday.UGCPolicy()
	return &service{
		db:        db,
		sanitizer: sanitizer,
	}
}

func (s service) FindTaskByID(id customType.StringUUID) (*Task, error) {
	task := &Task{
		BaseInfo: dicts.BaseInfo{
			ID: id,
		},
	}
	err := s.db.FindDictById(task)
	return task, err
}

func (s service) CreateTask(data []byte) (*Task, error) {
	task := &Task{}
	if err := task.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	task.Text = s.sanitizer.Sanitize(task.Text)
	err := s.db.CreateRecord(task)
	return task, err
}

func (s service) UpdateTask(data []byte, id customType.StringUUID) (*Task, error) {
	task := &Task{}
	if err := task.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	task.Text = s.sanitizer.Sanitize(task.Text)
	err := s.db.UpdateRecord(task, id)
	return task, err
}

func (s service) DeleteTask(id customType.StringUUID) error {
	task := Task{}
	return s.db.DeleteRecord(&task, id)
}

func (s service) FetchTasks(limit, offset int) (tasks []Task, err error) {
	_, err = s.db.FetchDict(&tasks, "tasks", limit, offset, nil, nil)
	return tasks, err
}

func (s service) FetchTasksByCardIDs(cardIDs []string) (tasks []Task, err error) {
	where := []string{"card_id = ?"}
	whereArgs := cardIDs
	_, err = s.db.FetchDict(&tasks, "tasks", 10000, 0, where, whereArgs)
	return tasks, err
}
