package convert

import "encoding/json"

// ToObj marshals data to JSON, then unmarshals into the target struct.
func ToObj(data interface{}, into interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, into)
}
