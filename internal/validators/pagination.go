package validators

import "strings"

func ValidateOrderBy(orderBy string, validProperties []string) bool {
	orderClauses := strings.Split(orderBy, " ")

	validPropertiesMap := make(map[string]struct{}, len(validProperties))
	for _, prop := range validProperties {
		validPropertiesMap[prop] = struct{}{}
	}

	for _, clause := range orderClauses {
		parts := strings.Split(clause, ":")
		if len(parts) != 2 {
			return false
		}

		property := parts[0]
		direction := parts[1]

		if _, exists := validPropertiesMap[property]; !exists {
			return false
		}

		if direction != "asc" && direction != "desc" {
			return false
		}
	}

	return true
}
