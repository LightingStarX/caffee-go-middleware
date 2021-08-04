package stringx

import (
	"fmt"
	"testing"
)

func TestStringsTransferWord2Tree(t *testing.T) {
	word := "I am a person"
	nd := &node{}

	nd.parseWordToTree(word)
}

func TestGenerateRandString(t *testing.T) {
	s := RandStringWithLen(10)
	fmt.Println(s)
}

func TestFilterString(t *testing.T) {
	f := func(r rune) bool {
		if r == 'a' {
			return true
		} else {
			return false
		}
	}

	s := "I am a person with a cat"
	fmt.Println(Filter(s, f))
}

func TestRemoveStrings(t *testing.T) {
	strings := []string{"Cat", "Pig", "Dog", "Fish", "Bird"}
	ss := []string{"Pig", "Cat", "Xxx"}

	res := Remove(strings, ss...)
	fmt.Println(res)
}

func TestReverse(t *testing.T) {
	s := "asdasfasfasasdasdy"
	fmt.Println(Reverse(s))
}

func TestUnion(t *testing.T) {
	ss1 := []string{"A", "B", "C", "D", "EE", "Cat"}
	ss2 := []string{"C", "B", "D", "XXX", "FFF"}

	fmt.Println(Union(ss1, ss2))
}
