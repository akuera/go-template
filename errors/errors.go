package errors

import "encoding/json"

type err struct {
	ErrCode    string `json:"errorCode"`
	ErrMessage string `json:"errorMessage"`
}

var (
	ErrNotFound    = err{"DOM404", "not found"}
	ErrInternal    = err{"DOM500", "internal error"}
	ErrInvalidCode = err{"DOM405", "invalid code"}
	ErrDuplicate   = err{"DOM409", "entity with provided attributes exists"}
	ErrNoChanges   = err{"DOM202", "no changes"}
)

func (e err) Error() string {
	marshalledErr, err := json.Marshal(e)
	if err != nil {
		return e.ErrMessage
	}
	return string(marshalledErr)
}

func Err(e error) error {
	return err{
		ErrCode:    "UKN0001",
		ErrMessage: e.Error(),
	}
}
