package utils

import "go.mongodb.org/mongo-driver/bson"

func Filter(filterKey []string, data bson.M) bson.M {
	filteredData := make(bson.M)

	for key, value := range data {
		// Check if the key is in the list of keys to filter out
		filtered := false
		for _, filterKey := range filterKey {
			if key == filterKey {
				filtered = true
				break
			}
		}

		// If the key is not in the list of keys to filter out, add it to the filteredData map
		if !filtered {
			filteredData[key] = value
		}
	}
	return filteredData
}
