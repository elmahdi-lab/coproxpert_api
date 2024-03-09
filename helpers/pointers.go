package helpers

import "time"

func IntSlicePointer(s []int) *[]int {
	return &s
}

func StringMapPointer(m map[string]string) *map[string]string {
	return &m
}

func StringPointer(s string) *string {
	return &s
}

func IntPointer(s int) *int {
	return &s
}

func BoolPointer(s bool) *bool {
	return &s
}

func Float64Pointer(s float64) *float64 {
	return &s
}

func Int64Pointer(s int64) *int64 {
	return &s
}

func TimePointer(s time.Time) *time.Time {
	return &s
}
