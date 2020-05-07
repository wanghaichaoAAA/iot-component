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

type ResendDataObj struct {
	RtuMN     string `json:"rtuMN"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	DataType  int    `json:"dataType"`
}

type RemoteControlResendDataObj struct {
	ResponseType string
	ResponseData ResendDataObj
}
