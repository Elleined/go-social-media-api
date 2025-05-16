package utils

import (
	"reflect"
	"slices"
	"strings"
)

func IsInDBTag(field string, structure any) bool {
	databaseTags := extractDatabaseTags(structure)
	return slices.Contains(databaseTags, field)
}

func IsInSortingOrder(sortBy string) bool {
	sortingOrders := []string{"ASC", "DESC"}
	return slices.Contains(sortingOrders, strings.ToUpper(sortBy))
}

func extractDatabaseTags(i interface{}) []string {
	val := reflect.TypeOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	var tags []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			tags = append(tags, dbTag)
		}
	}
	return tags
}
