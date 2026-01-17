package main

import (
	"fmt"
)

const ZERO = byte('0')

func maximalRectangle(matrix [][]byte) int {
    result := int(0)
    for i, rows := range matrix {
        for j, cell := range rows {
            if cell == ZERO {
                continue
            }

            maxHeight := int(0)
            for down := i; down < len(matrix); down++ {
                if matrix[down][j] == ZERO {
                    break
                }
                maxHeight++
            }

            minWidth := int(200)
            maxWidth := int(0)
            height := int(0)
            for col := i; col < i + maxHeight; col++ {
                height++
                width := int(0)
                for row := j; row < len(rows); row++ {
                    if matrix[col][row] == ZERO {
                        break
                    }
                    width++
                }
                if minWidth > width {
                    minWidth = width
                }
                if maxWidth < width {
                    maxWidth = width
                }

                rect := height * minWidth
                if maxWidth > rect {
                    rect = maxWidth
                }
                if result < rect {
                    result = rect
                }
            }

        }
    }
    return result
}

func main() {
	// result: 6
	// matrix := [][]byte{
	// 	{'1','0','1','0','0'},
	// 	{'1','0','1','1','1'},
	// 	{'1','1','1','1','1'},
	// 	{'1','0','0','1','0'},
	// }

	// result: 0
	// matrix := [][]byte{}

	// result: 0
	// matrix := [][]byte{{'0'}}

	// result: 1
	matrix := [][]byte{{'1'}}

	// result: 
	// matrix := [][]byte{
	// }

	result := maximalRectangle(matrix)
	fmt.Printf("result = %v\n", result)
}

