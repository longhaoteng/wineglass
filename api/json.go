package api

import (
	"encoding/json"
	"net/http"

	"github.com/liip/sheriff"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type DiffGroupsJSON struct {
	Groups []string
	Data   interface{}
}

func (d DiffGroupsJSON) Render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, d.Groups, d.Data); err != nil {
		panic(err)
	}
	return
}

func (d DiffGroupsJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func WriteJSON(w http.ResponseWriter, groups []string, obj interface{}) error {
	writeContentType(w, jsonContentType)

	data, err := sheriff.Marshal(
		&sheriff.Options{
			Groups:          groups,
			ApiVersion:      nil,
			IncludeEmptyTag: true,
		},
		obj,
	)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
