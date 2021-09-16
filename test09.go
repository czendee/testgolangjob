package main

import "log"

import "strings"

// IsRotation returns true if the first string is a rotation of the scond one
func IsRotation(s1 string, s2 string) bool {
	s1len, s2len := len(s1), len(s2)
	if s1len != s2len {
		return false
	}

	s2 += s2
	         log.Print("IsRotation step 1"+s2)
	return strings.Contains(s2, s1)
}

func test01Fail(){
    

    if(    IsRotation("dollar","llodar")){ //not rotation
        log.Print("IsRotation true")
     }else{
         log.Print("IsRotation false")
     }

}

func test01Sucess(){
    if(    IsRotation("dollar","lardol")){ // rotation  lardol   dollar
        log.Print("IsRotation true")
     }else{
         log.Print("IsRotation false")
     }

}



func main() {
     test01Fail();
     test01Sucess()
}
