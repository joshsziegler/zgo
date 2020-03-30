package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LoadJSON reads and unmarshals the given JSON file to a map.
//
// This is all done in memory, so large JSON files should use other methods.
func LoadJSON(path string) (data map[string]interface{}, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		if jsonError, ok := err.(*json.SyntaxError); ok {
			line, character, lcErr := lineAndCharacter(string(file), int(jsonError.Offset))
			fmt.Fprintf(os.Stderr, "error: Cannot parse JSON schema due to a syntax error at line %d, character %d: %v\n", line, character, jsonError.Error())
			if lcErr != nil {
				fmt.Fprintf(os.Stderr, "Couldn't find the line and character position of the error due to error %v\n", lcErr)
			}
		}
		return
	}
	return
}

// SaveJSON marshals a map to a JSON string, and saves it to the file path.
//
// This is all done in memory, so large JSON files should use other methods.
func SaveJSON(path string, data map[string]interface{}) (err error) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return
	}
	return
}
