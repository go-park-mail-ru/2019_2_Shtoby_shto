package photo

import (
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/handle"
)

type HandlerBoardService interface {
}

type service struct {
	handle.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerBoardService {
	return &service{
		db: db,
	}
}
