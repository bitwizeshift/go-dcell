/*
Package levenshtein provides a function to calculate the Damerau-Levenshtein
distance between two strings. This package is primarily used to provide hints
on error-conditions in the go-dcell package.
*/
package levenshtein

import "iter"

// DamerauDistance calculates the Damerau-Levenshtein distance between two
// strings. The Damerau-Levenshtein distance is a measure of the similarity
// between two strings, defined as the minimum number of operations required to
// transform one string into the other. The allowed operations are insertion,
// deletion, substitution, and transposition of two adjacent characters.
func DamerauDistance(from, to string) int {
	if len(from) == 0 {
		return len(to)
	}
	if len(to) == 0 {
		return len(from)
	}

	fromLen := len(from)
	toLen := len(to)

	distance := make([][]int, fromLen+1)
	for i := range distance {
		distance[i] = make([]int, toLen+1)
	}

	for i := 0; i <= fromLen; i++ {
		distance[i][0] = i
	}
	for j := 0; j <= toLen; j++ {
		distance[0][j] = j
	}

	for i := 1; i <= fromLen; i++ {
		for j := 1; j <= toLen; j++ {
			cost := 0
			if from[i-1] != to[j-1] {
				cost = 1
			}
			distance[i][j] = min(
				distance[i-1][j]+1,
				distance[i][j-1]+1,
				distance[i-1][j-1]+cost,
			)

			if i > 1 && j > 1 && from[i-1] == to[j-2] && from[i-2] == to[j-1] {
				distance[i][j] = min(
					distance[i][j],
					distance[i-2][j-2]+cost,
				)
			}
		}
	}

	return distance[fromLen][toLen]
}

// Suggestions generates a list of suggestions for a given input string based
// on the Damerau-Levenshtein distance. It returns a slice of strings that are
// within the optimal distance from the input string. The optimal distance is
// determined based on the length of the input string.
func Suggestions(input string, candidates iter.Seq[string]) []string {
	var suggestions []string
	lowest := optimalDistance(input)
	for candidate := range candidates {
		distance := DamerauDistance(input, candidate)
		if lowest < distance {
			continue
		}

		if distance < lowest {
			suggestions = []string{candidate}
			lowest = distance
		} else if distance == lowest {
			suggestions = append(suggestions, candidate)
		}
	}
	return suggestions
}

func optimalDistance(s string) int {
	l := len(s)
	switch {
	case l <= 4:
		return 1
	case l <= 8:
		return 2
	case l <= 12:
		return 3
	}
	return (l / 4) + 1
}
