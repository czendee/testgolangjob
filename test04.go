package main

import "log"

// IsPalindromePermutation returns true if a srting is a palindrome permutation
func IsPalindromePermutation(s string) bool {
	set := make(map[rune]int)

	for _, c := range s {
		if c == ' ' {
			continue
		}

		set[c]++
	}

	odd := false
	for _, n := range set {
		if n%2 != 0 {
			if odd {
				return false
			}

			odd = true
		}
	}

	return true
}

// IsPalindromePermutationSimplified returns true if a srting is a palindrome permutation
// It uses a single for loop to check the permutations
func IsPalindromePermutationSimplified(s string) bool {
	set := make(map[rune]int)
	odd := 0

	for _, c := range s {
		if c == ' ' {
			continue
		}

		set[c]++

		if set[c]%2 != 0 {
			odd++
		} else {
			odd--
		}
	}

	return !(odd > 1)
}

func test01Fail(){
    if(     IsPalindromePermutation(" wert")){
        log.Print("IsPalindromePermutation true")
     }else{
         log.Print("IsPalindromePermutation false")
     }

}

func test01Sucess(){
    if(     IsPalindromePermutation(" wertrew ")){
        log.Print("IsPalindromePermutation true")
     }else{
         log.Print("IsPalindromePermutation false")
     }

}

func main() {
     test01Fail();
     test01Sucess()
}
