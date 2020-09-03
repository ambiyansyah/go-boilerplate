package utils

import (
	"encoding/json"
	"fmt"
)

// AnyToStruct use to bind map, interface, struct to the other type
func AnyToStruct(from interface{}, to interface{}) error {
	jsonRecords, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("Error encode records: %s", err.Error())
	}
	if err := json.Unmarshal(jsonRecords, to); err != nil {
		return fmt.Errorf("Error decode json to struct: %s", err.Error())
	}
	return nil
}
