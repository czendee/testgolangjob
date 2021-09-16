package main

import "log"

// ZeroMatrix will turn entire row and column into zeros if an element is zero
func ZeroMatrix(m [][]int) {
	rows := len(m)
	if rows < 1 {
		return
	}

	cols := len(m[0])
	if cols < 1 {
		return
	}

	var zrows, zcols []int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j] == 0 {
				zrows = append(zrows, i)
				zcols = append(zcols, j)
			}
		}
	}

	for _, row := range zrows {
		for j := 0; j < cols; j++ {
			m[row][j] = 0
		}
	}

	for _, col := range zcols {
		for i := 0; i < rows; i++ {
			m[i][col] = 0
		}
	}
}

// ZeroMatrixSpaceEfficient space efficient zero matrix
func ZeroMatrixSpaceEfficient(m [][]int) {
	rows := len(m)
	if rows < 1 {
		return
	}

	cols := len(m[0])
	if cols < 1 {
		return
	}

	zrow, zcol := false, false
	for j := 0; j < cols; j++ {
		if m[0][j] == 0 {
			zcol = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if m[i][0] == 0 {
			zrow = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if m[i][j] == 0 {
				m[0][j] = 0
				m[i][0] = 0
			}
		}
	}

	for j := 1; j < cols; j++ {
		if m[0][j] != 0 {
			continue
		}

		for i := 0; i < rows; i++ {
			m[i][j] = 0
		}
	}

	for i := 1; i < rows; i++ {
		if m[i][0] != 0 {
			continue
		}

		for j := 0; j < cols; j++ {
			m[i][j] = 0
		}
	}

	if zcol {
		for j := 0; j < cols; j++ {
			m[0][j] = 0
		}
	}

	if zrow {
		for i := 0; i < rows; i++ {
			m[i][0] = 0
		}
	}
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
//      final:= make([][]int, 2)
      ZeroMatrix(inicial)

    if(    inicial[0][1] ==0){ //not changed to zero, as this row does not have 0's
        log.Print("ZeroMatrix true")
     }else{
         log.Print("ZeroMatrix false")
     }

}

func test01Sucess(){
    
      inicial := make([][]int, 2)
      for i:=range inicial  {
         inicial [i]=make([]int, 2)
      }
      inicial[0][0] = 4
      inicial[0][1] = 0
      inicial[1][0] = 3
      inicial[1][1] = 3
//      final:= make([][]int, 2)
      ZeroMatrix(inicial)

    if(    inicial[0][0] ==0){
        log.Print("ZeroMatrix true")
     }else{
         log.Print("ZeroMatrix false")
     }
}



func main() {
     test01Fail();
     test01Sucess()
}
