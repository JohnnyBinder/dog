package dog

import "strings"

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}

func WhenBigAndGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s) + "!!!"
}

func Tag12() string {
	return "This is version v1.2.0 of dog."
}