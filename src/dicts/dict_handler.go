package dicts

type Dict interface {
	SetId(id string)
	GetId() string
	GetTableName() string
}

type BaseInfo struct {
	ID string `json:"id" sql:"id"`
}

func (b BaseInfo) GetId() string {
	return b.ID
}

func (d *BaseInfo) SetId(id string) {
	d.ID = id
}
