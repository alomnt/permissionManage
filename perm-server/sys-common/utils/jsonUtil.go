package utils

import (
	"encoding/json"
	"github.com/kataras/golog"
)

//
func StructToString(stru interface{}) string {

	jsons, err := json.Marshal(stru) //转换成JSON返回的是byte[]
	if err != nil {
		err.Error()
		golog.Error(err)
		return ""
	}
	return string(jsons)
}

func ByteToStruct(byt []byte, stru interface{}) interface{} {
	err := json.Unmarshal(byt, &stru)
	if err != nil {
		err.Error()
		golog.Error(err)
		return stru
	}
	return stru
}
