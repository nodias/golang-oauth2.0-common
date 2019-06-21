package response

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nodias/go-ApmCommon/model"
)

var ErrUserNotExist = errors.New("user not exist")

type ID string

type ResponseErr struct {
	Err error
}

type Response struct {
	Id   ID
	User *model.User
	Err  ResponseErr
}

func (r ResponseErr) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, r.Err)), nil
}

func (r *ResponseErr) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v == nil {
		r.Err = nil
		return nil
	}
	switch p := v.(type) {
	case string:
		if r.Err == ErrUserNotExist {
			r.Err = ErrUserNotExist
			return nil
		}
		r.Err = errors.New(p)
		return nil
	default:
		return errors.New("unexpected response error")
	}
	return nil
}
