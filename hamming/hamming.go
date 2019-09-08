// Package hamming takes 2 DNA strands as strings and calculates the point difference between them.
package hamming

import (
    "errors"
)

// Distance takes 2 strings as input and returns the distance for equal length strands.
func Distance(a, b string) (int, error) {
	distance := 0
	err := errors.New("Strands are not of same length")

	if len(a) == len(b) {
		err = nil
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				distance++
			}
		}
	}
	return distance, err
}
