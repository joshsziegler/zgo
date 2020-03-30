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
		// If the error is due to syntax, try to print a useful error message
		if jsonError, ok := err.(*json.SyntaxError); ok {
			line, character, lcErr := getLineAndCharacter(string(file), int(jsonError.Offset))
			// TODO: Return errors, don't print them! - JZ
			fmt.Fprintf(os.Stderr, "error: Cannot parse JSON schema due to a syntax error at line %d, character %d: %v\n", line, character, jsonError.Error())
			if lcErr != nil {
				fmt.Fprintf(os.Stderr, "Couldn't find the line and character position of the error due to error: %v\n", lcErr)
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

// Get the line and character number referenced by an integer offset.
// This is useful for JSON errors which return an offset, indicating where in
// the input the error occured.
func getLineAndCharacter(input string, offset int) (line int, character int, err error) {
	lineFeed := rune(0x0A)
	inputLen := len(input) // avoid doing this operation twice in an error
	if offset > inputLen {
		return 0, 0, fmt.Errorf("offset (%d) larger than input length (%d)", offset, inputLen)
	} else if offset < 0 {
		return 0, 0, fmt.Errorf("offset is less than zero (%d)", offset)
	}
	line = 1 // Count from one, like a non-CS human :)
	for i, b := range input {
		if b == lineFeed {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}
	return line, character, nil
}
