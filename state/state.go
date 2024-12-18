package state

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// function deletes the whole object if it found deletingKey as value in interface and if found deletingKey as key in interface then delete single key value pair
func deleteObjectContainingKeyOrValueFromMap(m interface{}, keysToDelete []string) interface{} {
	switch m := m.(type) {
	case map[string]interface{}:
		for k, v := range m {
			if arrayContainsString(keysToDelete, k) {
				delete(m, k)
				continue
			}
			switch v.(type) {
			case string:
				if arrayContainsString(keysToDelete, v.(string)) {
					return nil
				}
			default:
				modifiedValue := deleteObjectContainingKeyOrValueFromMap(v, keysToDelete)
				switch value := modifiedValue.(type) {
				case []interface{}:
					if len(value) == 0 {
						delete(m, k)
					} else {
						m[k] = modifiedValue
					}
				case nil:
					if k == "stream_descriptor" {
						delete(m, k)
						delete(m, "stream_state")
					}
				default:
					m[k] = modifiedValue
				}
			}
		}
	case []interface{}:
		var arr []interface{}
		for i := 0; i < len(m); i++ {
			m[i] = deleteObjectContainingKeyOrValueFromMap(m[i], keysToDelete)
			if m[i] != nil {
				arr = append(arr, m[i])
			}
		}
		return arr
	}

	return m
}

func arrayContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ParseState(st string, keyToDelete ...string) error {
	logrus.Infof("Parsing state: %s", st)
	var m interface{}
	if err := json.Unmarshal([]byte(st), &m); err != nil {
		return err
	}
	if bytes, err := json.Marshal(deleteObjectContainingKeyOrValueFromMap(m, keyToDelete)); err != nil {
		return err
	} else {
		logrus.Infof("State After Parsing: %s", string(bytes))
	}
	return nil
}
