package customType

import "2019_2_Shtoby_shto/src/utils"

type StringUUID string

func (s StringUUID) IsEmpty() bool {
	return len(string(s)) == 0
}

func (s StringUUID) IsUUID() bool {
	return utils.IsUUID(string(s))
}

func (s StringUUID) String() string {
	return string(s)
}

func (StringUUID) Empty() StringUUID {
	return StringUUID("")
}
