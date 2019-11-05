package photo

import (
	"2019_2_Shtoby_shto/src/dicts"
	"time"
)

const photoTableName = "photos"

//easyjson:json
type Photo struct {
	dicts.BaseInfo
	TimeLoad time.Time `json:"time_load" sql:"type:time"`
	Path     string    `json:"path" sql:"type:varchar(50)"`
}

func (p Photo) GetTableName() string {
	return photoTableName
}

func (p Photo) IsValid() bool {
	return p.Path != ""
}
