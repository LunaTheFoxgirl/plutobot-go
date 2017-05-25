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

// SliceToHumanListing converts a slice into a human listing of objects, eg. "Cat, dog and bird"
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

// SanitizeArgs Sanitizes the arguments of a slice, by combining indexes which are in quotes.
func SanitizeArgs(slice []string) []string {
	var fslice []string = make([]string, 0)
	var current string = ""
	var inquote bool = false

	for i := 0; i < len(slice); i++ {
		if strings.HasPrefix(slice[i], "\"") {
			current += slice[i][1:] + " "
			inquote = true
			continue
		}

		if inquote {
			if strings.HasSuffix(slice[i], "\"") {
				current += slice[i][:len(slice[i])-1]
				fslice = append(fslice, current)
				current = ""
				inquote = false
				continue
			}

			current += slice[i] + " "
			continue
		}

		fslice = append(fslice, slice[i])
	}

	if current != "" {
		return slice
	}
	return fslice
}

// StringToSlice turns a string into slices based of space character denotation..
func StringToSlice(text string) (slices []string) {
	return strings.Split(text, " ")
}

func StringToInt64(inp string) (int64, error) {
	return strconv.ParseInt(inp, 10, 64)
}
