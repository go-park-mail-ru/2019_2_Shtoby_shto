package dicts

import . "2019_2_Shtoby_shto/src/customType"

type Dict interface {
	SetId(id StringUUID)
	GetId() StringUUID
	IsValid() bool
	GetTableName() string
}

type BaseInfo struct {
	ID StringUUID `json:"id" sql:"type:uuid;not null;unique"`
}

func (b BaseInfo) GetId() StringUUID {
	return b.ID
}

func (b *BaseInfo) SetId(id StringUUID) {
	b.ID = id
}

func (b BaseInfo) GetTableName() string {
	return "default table name"
}

func (b BaseInfo) IsValid() bool {
	return true
}
