package util

import (
	"encoding/json"
	"log"
)

func MapToStruct[T any](dict *map[string]any, obj *T) {
	jsonbody, err := json.Marshal(dict)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.Unmarshal(jsonbody, obj); err != nil {
		// do error check
		log.Panic(err)
		return
	}

}
