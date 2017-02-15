package excercises

import ()

// Find the index of the first leftmost occurrence of k in XS.
// Time Complexity: log(n)
// Space Complexity: O(1)

// Example:
//  xs = 1, 2, 3, 4, 5, 6, 9
//  k = 5
func FirstK(xs []int, k int) int {
	f := -1
	l, r := 0, len(xs)-1
	for l <= r {
		m := l + (r-l)/2 // avoid integer overflow
		switch v := xs[m]; {
		case v > k:
			r = m - 1
		case v == k:
			f, r = m, m-1
		case v < k:
			l = m + 1
		}
	}
	return f
}
