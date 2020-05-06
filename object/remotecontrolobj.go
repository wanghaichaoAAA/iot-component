package object

import (
	"encoding/base64"
	"encoding/json"
)

type RemoteControlObj struct {
	ResponseType string
	ResponseData interface{}
}

func (obj *RemoteControlObj) GetBase64String() (string, error) {
	bytes, err := json.Marshal(*obj)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
