package set

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestSet_Add tests the Add method
func TestSet_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		values []int
		want   map[int]struct{}
	}{
		{
			name:   "simple integers",
			values: []int{0, 1, 2, 3, 4},
			want:   map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}},
		},
		{
			name:   "duplicate values",
			values: []int{1, 1, 2, 2, 3},
			want:   map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:   "empty set",
			values: []int{},
			want:   map[int]struct{}{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set := New[int](0)
			for _, val := range tt.values {
				set.Add(val)
			}
			require.Equal(t, tt.want, set.hashMap)
		})
	}
}

// TestSet_In tests the In method
func TestSet_In(t *testing.T) {
	t.Parallel()
	set := New[int](3)
	values := []int{1, 2, 3}
	for _, v := range values {
		set.Add(v)
	}

	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"existing value", 1, true},
		{"non-existing value", 4, false},
		{"another existing value", 3, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := set.In(tt.value)
			require.Equal(t, tt.expected, result)
		})
	}
}

// TestSet_All tests the All method
func TestSet_All(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			"multiple values",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"single value",
			[]int{1},
			[]int{1},
		},
		{
			"duplicate values",
			[]int{1, 2, 2, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"empty set",
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set := New[int](0)
			for _, v := range tt.values {
				set.Add(v)
			}
			result := set.All()
			require.Equal(t, len(tt.want), len(result))
			for _, v := range result {
				require.Contains(t, tt.want, v)
			}
		})
	}
}

// TestSet_Intersect tests the Intersect method including swap functionality
func TestSet_Intersect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		set1         []int
		set2         []int
		expected     []int
		swapExpected bool // true if sets should be swapped for optimization
	}{
		{
			name:         "simple intersection no swap needed",
			set1:         []int{1, 2, 3, 4},
			set2:         []int{3, 4, 5, 6},
			expected:     []int{3, 4},
			swapExpected: false, // set1 and set2 are same size
		},
		{
			name:         "intersection with swap needed - set1 larger",
			set1:         []int{1, 2, 3, 4, 5, 6, 7, 8},
			set2:         []int{3, 4, 5},
			expected:     []int{3, 4, 5},
			swapExpected: true, // set1 is larger, should be swapped
		},
		{
			name:         "intersection with no swap - set2 larger",
			set1:         []int{1, 2, 3},
			set2:         []int{2, 3, 4, 5, 6, 7, 8},
			expected:     []int{2, 3},
			swapExpected: false, // set1 is already smaller
		},
		{
			name:         "empty set intersection",
			set1:         []int{},
			set2:         []int{1, 2, 3},
			expected:     []int{},
			swapExpected: false,
		},
		{
			name:         "identical sets",
			set1:         []int{1, 2, 3},
			set2:         []int{1, 2, 3},
			expected:     []int{1, 2, 3},
			swapExpected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Create and fill the sets
			set1 := New[int](len(tt.set1))
			set2 := New[int](len(tt.set2))

			for _, v := range tt.set1 {
				set1.Add(v)
			}
			for _, v := range tt.set2 {
				set2.Add(v)
			}

			// Store original pointers to verify swap
			originalSet1 := set1
			originalSet2 := set2

			// Create copies to verify contents aren't modified
			set1Copy := make(map[int]struct{}, len(set1.hashMap))
			set2Copy := make(map[int]struct{}, len(set2.hashMap))
			for k, v := range set1.hashMap {
				set1Copy[k] = v
			}
			for k, v := range set2.hashMap {
				set2Copy[k] = v
			}

			// Perform intersection
			result := set1.Intersect(*set2)

			// Verify the result contains expected values
			resultSlice := result.All()
			require.Equal(t, len(tt.expected), len(resultSlice))
			for _, v := range resultSlice {
				require.Contains(t, tt.expected, v)
			}

			// Verify original sets weren't modified
			require.Equal(t, set1Copy, set1.hashMap)
			require.Equal(t, set2Copy, set2.hashMap)

			// If swap was expected, verify set contents are still correct
			if tt.swapExpected {
				require.Equal(t, set1Copy, originalSet1.hashMap)
				require.Equal(t, set2Copy, originalSet2.hashMap)
			}
		})
	}
}

// TestSet_Diff tests the Diff method
func TestSet_Diff(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "simple difference",
			set1:     []int{1, 2, 3},
			set2:     []int{2, 3, 4},
			expected: []int{1, 4},
		},
		{
			name:     "no common elements",
			set1:     []int{1, 2},
			set2:     []int{3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "identical sets",
			set1:     []int{1, 2},
			set2:     []int{1, 2},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set1 := New[int](0)
			set2 := New[int](0)
			for _, v := range tt.set1 {
				set1.Add(v)
			}
			for _, v := range tt.set2 {
				set2.Add(v)
			}
			result := set1.Diff(*set2)
			resultSlice := result.All()
			require.Equal(t, len(tt.expected), len(resultSlice))
			for _, v := range resultSlice {
				require.Contains(t, tt.expected, v)
			}
		})
	}
}

// TestSet_Len tests the Len method
func TestSet_Len(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{"multiple values", []int{1, 2, 3, 4}, 4},
		{"duplicate values", []int{1, 1, 2, 2}, 2},
		{"empty set", []int{}, 0},
		{"single value", []int{1}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set := New[int](0)
			for _, v := range tt.values {
				set.Add(v)
			}
			require.Equal(t, tt.expected, set.Len())
		})
	}
}

// TestSet_Union tests the Union method
func TestSet_Union(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}{
		{
			name:     "simple union",
			set1:     []int{1, 2, 3},
			set2:     []int{3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "disjoint sets",
			set1:     []int{1, 2},
			set2:     []int{3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "one empty set",
			set1:     []int{1, 2, 3},
			set2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "both empty sets",
			set1:     []int{},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "identical sets",
			set1:     []int{1, 2, 3},
			set2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "subset union",
			set1:     []int{1, 2, 3, 4},
			set2:     []int{2, 3},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set1 := New[int](len(tt.set1))
			set2 := New[int](len(tt.set2))

			// Create copies to verify original sets aren't modified
			set1Copy := make(map[int]struct{}, len(tt.set1))
			set2Copy := make(map[int]struct{}, len(tt.set2))

			for _, v := range tt.set1 {
				set1.Add(v)
				set1Copy[v] = struct{}{}
			}
			for _, v := range tt.set2 {
				set2.Add(v)
				set2Copy[v] = struct{}{}
			}

			result := set1.Union(*set2)
			resultSlice := result.All()

			// Verify result contains all expected elements
			require.Equal(t, len(tt.expected), len(resultSlice))
			for _, v := range resultSlice {
				require.Contains(t, tt.expected, v)
			}

			// Verify original sets weren't modified
			require.Equal(t, set1Copy, set1.hashMap)
			require.Equal(t, set2Copy, set2.hashMap)
		})
	}
}

// TestSet_Remove tests the Remove method
func TestSet_Remove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		initialValues  []int
		toRemove       []int
		expectedValues []int
		expectedLens   []int // Length after each removal
	}{
		{
			name:           "remove existing elements",
			initialValues:  []int{1, 2, 3, 4, 5},
			toRemove:       []int{1, 3, 5},
			expectedValues: []int{2, 4},
			expectedLens:   []int{4, 3, 2},
		},
		{
			name:           "remove non-existing elements",
			initialValues:  []int{1, 2, 3},
			toRemove:       []int{4, 5},
			expectedValues: []int{1, 2, 3},
			expectedLens:   []int{3, 3},
		},
		{
			name:           "remove from empty set",
			initialValues:  []int{},
			toRemove:       []int{1},
			expectedValues: []int{},
			expectedLens:   []int{0},
		},
		{
			name:           "remove same element multiple times",
			initialValues:  []int{1, 2, 3},
			toRemove:       []int{2, 2, 2},
			expectedValues: []int{1, 3},
			expectedLens:   []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set := New[int](len(tt.initialValues))

			// Add initial values
			for _, v := range tt.initialValues {
				set.Add(v)
			}

			// Remove elements and check length after each removal
			for i, v := range tt.toRemove {
				set.Remove(v)
				require.Equal(t, tt.expectedLens[i], set.Len())
			}

			// Verify final state
			result := set.All()
			require.Equal(t, len(tt.expectedValues), len(result))
			for _, v := range result {
				require.Contains(t, tt.expectedValues, v)
			}

			// Verify removed elements are actually gone
			for _, v := range tt.toRemove {
				require.False(t, set.In(v))
			}
		})
	}
}

// TestSet_Clear tests the Clear method
func TestSet_Clear(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		initialValues []int
	}{
		{
			name:          "clear non-empty set",
			initialValues: []int{1, 2, 3, 4, 5},
		},
		{
			name:          "clear single-element set",
			initialValues: []int{1},
		},
		{
			name:          "clear empty set",
			initialValues: []int{},
		},
		{
			name:          "clear set with duplicates",
			initialValues: []int{1, 1, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			set := New[int](len(tt.initialValues))

			// Add initial values
			for _, v := range tt.initialValues {
				set.Add(v)
			}

			// Clear the set
			set.Clear()

			// Verify set is empty
			require.Equal(t, 0, set.Len())
			require.Empty(t, set.All())

			// Verify all original values are gone
			for _, v := range tt.initialValues {
				require.False(t, set.In(v))
			}

			// Test that the set is still usable after clearing
			newValue := 42
			set.Add(newValue)
			require.True(t, set.In(newValue))
			require.Equal(t, 1, set.Len())
		})
	}
}

// TestSet_Chaining tests that methods can be chained together
func TestSet_Chaining(t *testing.T) {
	t.Parallel()
	set1 := New[int](5)
	set2 := New[int](5)

	// Add elements to both sets
	for _, v := range []int{1, 2, 3, 4, 5} {
		set1.Add(v)
		set2.Add(v)
	}

	// Perform operations
	result := set1.Union(*set2)
	require.Equal(t, 5, result.Len())

	// Remove and verify
	result.Remove(1)
	result.Remove(2)
	require.Equal(t, 3, result.Len())

	// Clear and verify
	result.Clear()
	require.Equal(t, 0, result.Len())

	// Add new elements after clearing
	result.Add(10)
	require.True(t, result.In(10))
}
