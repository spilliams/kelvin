package endpoints

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	obj := map[string]string{"message": "Hello world"}
	return writeJson(w, obj)
}
