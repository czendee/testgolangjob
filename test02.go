package main

import "log"

///this should be in another program in the same package  sort.go
//start

import "sort"

// SortableString is a representation of a sortable string
type SortableString []rune

func (s SortableString) Len() int           { return len(s) }
func (s SortableString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortableString) Less(i, j int) bool { return s[i] < s[j] }

// SortString returns the sorted string
func SortString(s string) string {
	sr := SortableString(s)
	sort.Sort(sr)

	return string(sr)
}

///this should be in another program in the same package  sort.go
///end



// CheckPermutationBySort returns true if one string is a permutation of the other
// Sorts both strings and checks if they are equal
func CheckPermutationBySort(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if SortString(s1) != SortString(s2) {
		return false
	}

	return true
}

// CheckPermutationBySet returns true if one string is a permutation of the other
// Use set of runes to reduce complexity
func CheckPermutationBySet(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	set := make(map[rune]int)

	for _, c := range s1 {
		set[c]++
	}

	for _, c := range s2 {
		set[c]--

		if set[c] < 0 {
			return false
		}
	}

	return true
}


func test01Fail(){
    if(     CheckPermutationBySort("wert","notrew")){
        log.Print("CheckPermutationBySort  true")
     }else{
         log.Print("CheckPermutationBySort  false")
     }

}

func test01Sucess(){
    if(     CheckPermutationBySort("werttt","tttrew")){
        log.Print("CheckPermutationBySort  true")
     }else{
         log.Print("CheckPermutationBySort  false")
     }

}

func main() {
     test01Fail();
     test01Sucess()
}
