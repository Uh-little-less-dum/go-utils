package file_handlers_json

import (
	file_handlers_base "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/baseFile"
	"github.com/tidwall/gjson"
)

type JsonFileHandler struct {
	base        file_handlers_base.BaseFileHandler
	data        gjson.Result
	hasReadData bool
}

func NewJsonFileHandler(filePath string) JsonFileHandler {
	return JsonFileHandler{
		base: file_handlers_base.BaseFileHandler{
			Path: filePath,
		},
	}
}

func (a *JsonFileHandler) Data() interface{} {
	if a.hasReadData {
		return a.data
	} else {
		d := a.base.Read()
		j := gjson.ParseBytes(d)
		a.data = j
		return j
	}
}
