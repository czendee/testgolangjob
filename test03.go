package main

import "log"

// URLifyn returns the string with HTML encoded spacese
// Expects real string length as a second parameter
func URLifyn(s string, n int) string {
	spaces := 0
	for _, c := range s {
		if c == ' ' {
			spaces++
		}
	}

	index := n + spaces*2
	runes := make([]rune, n+spaces*2)

	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			runes[index-1] = '0'
			runes[index-2] = '2'
			runes[index-3] = '%'
			index = index - 3
		} else {
			runes[index-1] = rune(s[i])
			index--
		}
	}

	return string(runes)
}

// URLify returns the string with HTML encoded spacese
// Expects trimmed URL string as a parameter
func URLify(s string) string {
	spaces := 0
	for _, c := range s {
		if c == ' ' {
			spaces++
		}
	}

	index := len(s) + spaces*2
	runes := make([]rune, index)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			runes[index-1] = '0'
			runes[index-2] = '2'
			runes[index-3] = '%'
			index = index - 3
		} else {
			runes[index-1] = rune(s[i])
			index--
		}
	}

	return string(runes)
}

func test01Fail(){
    if(     URLifyn(" wert ",4)=="%20wert"){
        log.Print("URLifyn  true")
     }else{
         log.Print("URLifyn  false")
     }

}

func test01Sucess(){
    if(     URLifyn(" wert ",6)=="%20wert%20"){
        log.Print("URLifyn true")
     }else{
         log.Print("URLifyn  false")
     }

}

func main() {
     test01Fail();
     test01Sucess()
}
