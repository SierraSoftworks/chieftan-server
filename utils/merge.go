package utils

// Merge will merge a number of map[string]string objects into
// a single object in much the same way that Object.assign does
// in JavaScript.
func Merge(maps ...map[string]string) map[string]string {
	result := map[string]string{}

	for _, m := range maps {
		if m == nil {
			continue
		}

		for key, value := range m {
			result[key] = value
		}
	}

	return result
}
