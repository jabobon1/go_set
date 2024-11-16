package main

import (
	"fmt"
	"go-set/pkg/set"
	"strings"
)

func main() {
	// Basic Integer Set Operations
	fmt.Println("=== Integer Set Examples ===")
	demoIntegerSet()

	fmt.Println("\n=== String Set Examples ===")
	demoStringSet()

	//fmt.Println("\n=== Custom Type Example ===")
	//demoCustomType()

	fmt.Println("\n=== Set Operations Example ===")
	demoSetOperations()

	fmt.Println("\n=== Practical Example: Text Analysis ===")
	demoTextAnalysis()
}

func demoIntegerSet() {
	// Create a new integer set
	numbers := set.New[int](10)

	// Add some numbers
	for i := 0; i < 5; i++ {
		numbers.Add(i)
	}

	// Add duplicates (will be ignored)
	numbers.Add(3)
	numbers.Add(4)

	fmt.Printf("Numbers in set: %v\n", numbers.All())
	fmt.Printf("Set size: %d\n", numbers.Len())
	fmt.Printf("Contains 3? %v\n", numbers.In(3))
	fmt.Printf("Contains 7? %v\n", numbers.In(7))

	// Remove an element
	numbers.Remove(3)
	fmt.Printf("After removing 3: %v\n", numbers.All())
}

func demoStringSet() {
	// Create sets of fruits
	fruits1 := set.New[string](5)
	fruits1.Add("apple")
	fruits1.Add("banana")
	fruits1.Add("orange")

	fruits2 := set.New[string](5)
	fruits2.Add("banana")
	fruits2.Add("grape")
	fruits2.Add("kiwi")

	fmt.Printf("Fruits set 1: %v\n", fruits1.All())
	fmt.Printf("Fruits set 2: %v\n", fruits2.All())

	// Demonstrate Union
	union := fruits1.Union(*fruits2)
	fmt.Printf("All fruits (union): %v\n", union.All())

	// Clear the union set
	union.Clear()
	fmt.Printf("After clearing union set: %v\n", union.All())
}

// User represents a custom type for demonstration
type User struct {
	ID   int
	Name string
}

//func demoCustomType() {
//	users := set.New[User](5)
//
//	// Add some users
//	users.Add(User{1, "Alice"})
//	users.Add(User{2, "Bob"})
//	users.Add(User{1, "Alice"}) // Duplicate will be ignored
//
//	fmt.Printf("Number of unique users: %d\n", users.Len())
//
//	for _, user := range users.All() {
//		fmt.Printf("User: ID=%d, Name=%s\n", user.ID, user.Name)
//	}
//}

func demoSetOperations() {
	// Create two sets of numbers
	set1 := set.New[int](10)
	set2 := set.New[int](10)

	// Add numbers to first set
	for i := 0; i < 5; i++ {
		set1.Add(i)
	}

	// Add numbers to second set
	for i := 3; i < 8; i++ {
		set2.Add(i)
	}

	fmt.Printf("Set 1: %v\n", set1.All())
	fmt.Printf("Set 2: %v\n", set2.All())

	// Demonstrate various set operations
	intersection := set1.Intersect(*set2)
	fmt.Printf("Intersection: %v\n", intersection.All())

	difference := set1.Diff(*set2)
	fmt.Printf("Symmetric Difference: %v\n", difference.All())

	union := set1.Union(*set2)
	fmt.Printf("Union: %v\n", union.All())
}

func demoTextAnalysis() {
	// Example text analysis using sets
	text1 := "the quick brown fox jumps over the lazy dog"
	text2 := "the brown dog sleeps in the garden"

	// Create sets of unique words
	words1 := textToSet(text1)
	words2 := textToSet(text2)

	// Find common words
	commonWords := words1.Intersect(*words2)
	fmt.Printf("Common words: %v\n", commonWords.All())

	// Find unique words in each text
	uniqueWords := words1.Diff(*words2)
	fmt.Printf("Words unique to each text: %v\n", uniqueWords.All())

	// Total unique words across both texts
	allWords := words1.Union(*words2)
	fmt.Printf("Total unique words: %v\n", allWords.All())
}

func textToSet(text string) *set.Set[string] {
	words := strings.Fields(strings.ToLower(text))
	wordSet := set.New[string](len(words))
	for _, word := range words {
		wordSet.Add(word)
	}
	return wordSet
}
