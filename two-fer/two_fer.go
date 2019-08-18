package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var you = "you"
	if name != "" {
		you = name
	}
	return fmt.Sprintf("One for %s, one for me.", you)
}
