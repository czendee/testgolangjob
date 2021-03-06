package main

import "log"


// IsOneAway returns true if given strings are one (or zero) edits away
func IsOneAway(s1 string, s2 string) bool {
	s1len := len(s1)
	s2len := len(s2)
	diff := s1len - s2len

	if diff > 0 {
		return isOneInsertAway(s2, s1)
	}

	if diff < 0 {
		return isOneInsertAway(s1, s2)
	}

	return isOneReplceAway(s1, s2)
}

// isOneInsertAway returns true if given strings are one insert away
// s1 should be the sortest string always
func isOneInsertAway(s1 string, s2 string) bool {
	edits := 0

	for i, c := range s1 {
		if c != rune(s2[i+edits]) {
			edits++
		}

		if edits > 1 {
			return false
		}
	}

	return true
}

// isOneReplceAway returns true if given strings are one replace away
func isOneReplceAway(s1 string, s2 string) bool {
	edits := 0

	for i, c := range s1 {
		if c != rune(s2[i]) {
			edits++
		}

		if edits > 1 {
			return false
		}
	}

	return true
}

// IsOneAwayLoop returns true if given strings are one (or zero) edits away
func IsOneAwayLoop(s1 string, s2 string) bool {
	s1len := len(s1)
	s2len := len(s2)
	diff := s1len - s2len

	var i, j, edits int
	for i < s1len && j < s2len {
		if s1[i] != s2[j] {
			edits++

			if edits > 1 {
				return false
			}

			if diff > 0 {
				// s1 is longer than s2, move s2 index to a previous char
				// so in the next iteration we will check s2[j] against s1[i++]
				j--
			}

			if diff < 0 {
				// s2 is longer than s1, move s1 index to a previous char
				// so in the next iteration we will check s1[i] against s2[j++]
				i--
			}
		}

		i++ // move s1 index to the next char
		j++ // move s2 index to the next char
	}

	return true
}

// IsOneAwayClosure returns true if given strings are one (or zero) edits away
func IsOneAwayClosure(s1 string, s2 string) bool {
	s1len := len(s1)
	s2len := len(s2)
	diff := s1len - s2len

	// next rune closure for insert edit cases
	next := func(s string, i int, n int) rune {
		return rune(s[i+n])
	}

	if diff > 0 {
		return isOneEditAway(s2, s1, next)
	}

	if diff < 0 {
		return isOneEditAway(s1, s2, next)
	}

	// next rune closure for replace edit case
	next = func(s string, i int, n int) rune {
		return rune(s[i])
	}
	return isOneEditAway(s1, s2, next)
}

// isOneEditAway returns true if given strings are one (or zero) edits away
// s1 should be the sortest string always
func isOneEditAway(s1 string, s2 string, next func(s string, i int, n int) rune) bool {
	edits := 0

	for i, c := range s1 {
		if c != next(s2, i, edits) {
			edits++
		}

		if edits > 1 {
			return false
		}
	}

	return true
}


func test01Fail(){
    if(     IsOneAway("wertrew","wertrewhgkjh") ){
        log.Print("IsOneAway true")
     }else{
         log.Print("IsOneAway false")
     }

}

func test01Sucess(){
    if(     IsOneAway("wertrew","wertrewA")){
        log.Print("IsOneAway true")
     }else{
         log.Print("IsOneAway false")
     }

}



func main() {
     test01Fail();
     test01Sucess()
}
