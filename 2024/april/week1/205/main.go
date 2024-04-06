package main
import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
    // 32 is space (start character) in ascii
    sAsciiByteNum := int(32)
    tAsciiByteNum := int(32)
    sMemo := map[byte]byte{}
    tMemo := map[byte]byte{}
    for i := 0; i < len(s); i++ {
        sByte := s[i]
        nextS := ""
        if sMemo[sByte] == 0 {
            sMemo[sByte] = byte(sAsciiByteNum)
            nextS = string(byte(sAsciiByteNum))
            sAsciiByteNum++
        } else {
            b := sMemo[sByte]
            nextS = string(b)
        }

        tByte := t[i]
        nextT := ""
        if tMemo[tByte] == 0 {
            tMemo[tByte] = byte(tAsciiByteNum)
            nextT = string(byte(tAsciiByteNum))
            tAsciiByteNum++
        } else {
            b := tMemo[tByte]
            nextT = string(b)
        }
        if nextS != nextT {
            return false
        }
    }
    return true
}

func main() {
	// result: true
	// s := "egg"
	// t := "add"

	// result: false
	// s := "foo"
	// t := "bar"

	// result: true
	s := "paper"
	t := "title"

	// result: 
	// s := ""
	// t := ""

	result := isIsomorphic(s, t)
	fmt.Printf("result = %v\n", result)
}

