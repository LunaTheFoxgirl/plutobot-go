package core

import (
	"strconv"
	"strings"
)

// SliceToString turns a slice of strings into a single string with space denotation.
func SliceToString(slices []string) string {
	var sl = ""
	for indx, val := range slices {
		if indx != len(slices)-1 {
			sl += val + " "
		} else {
			sl += val
		}
	}
	return sl
}

func SliceToHumanListing(slices []string) string {
	var sl = ""
	for indx, val := range slices {
		if indx == len(slices)-2 {
			sl += val + " and "
		} else if indx < len(slices)-2 {
			sl += val + ", "
		} else {
			sl += val
		}
	}
	return sl
}

// StringToSlice turns a string into slices based of space character denotation..
func StringToSlice(text string) (slices []string) {
	return strings.Split(text, " ")
}

func StringToInt64(inp string) (int64, error) {
	return strconv.ParseInt(inp, 10, 64)
}
