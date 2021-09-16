package main

import "log"

// RotateMatrix rotates given matrix m to 90 degrees
func RotateMatrix(m [][]int) [][]int {
	n := len(m)

	mx := make([][]int, n)
	for i := 0; i < n; i++ {
		mx[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mx[j][n-1-i] = m[i][j]
		}
	}

	return mx
}

// RotateMatrixEfficient rotates given matrix m to 90 degrees
func RotateMatrixEfficient(m [][]int) [][]int {
	n := len(m)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			m[j][n-1-i], m[i][j], m[n-1-j][i], m[n-1-i][n-1-j] = m[i][j], m[n-1-j][i], m[n-1-i][n-1-j], m[j][n-1-i]
		}
	}

	return m
}


func test01Fail(){
      
      inicial := make([][]int, 2)
      for i:=range inicial  {
         inicial [i]=make([]int, 2)
      }
      inicial[0][0] = 4
      inicial[0][1] = 4
      inicial[1][0] = 3
      inicial[1][1] = 3
      final:= make([][]int, 2)
      final = RotateMatrix(inicial)

    if(    final[0][1] ==4){
        log.Print("RotateMatrix  true")
     }else{
         log.Print("RotateMatrix  false")
     }

}

func test01Sucess(){
    
      inicial := make([][]int, 2)
      for i:=range inicial  {
         inicial [i]=make([]int, 2)
      }
      inicial[0][0] = 4
      inicial[0][1] = 4
      inicial[1][0] = 3
      inicial[1][1] = 3
      final:= make([][]int, 2)
      final = RotateMatrix(inicial)

    if(    final[0][1] ==3){
        log.Print("RotateMatrix  true")
     }else{
         log.Print("RotateMatrix  false")
     }
}



func main() {
     test01Fail();
     test01Sucess()
}
