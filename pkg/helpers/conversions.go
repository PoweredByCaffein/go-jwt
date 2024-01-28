package helpers

import "strconv"

func ConvertStringToInt(input string, defaultValue int) (int, error) {

	output, err := strconv.Atoi(input)
	if err != nil {
		return defaultValue, err
	}

	return output, nil

}
