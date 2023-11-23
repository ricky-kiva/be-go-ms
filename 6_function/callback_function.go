package main

func filterCallback(arr []string, callback func(string) bool) []string {
	var result []string
	for _, item := range arr {
		if filtered := callback(item); filtered {
			result = append(result, item)
		}
	}

	return result
}
