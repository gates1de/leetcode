package main
import (
    "fmt"
)

type FreqStack struct {
    freqCounter map[int]int // key: num, value: frequency count
    freqGroup map[int][]int // key: frequency count, value: nums
    maxFreq int
}

func Constructor() FreqStack {
    stack := FreqStack{
        freqCounter: map[int]int{},
        freqGroup: map[int][]int{},
        maxFreq: 0,
    }
    return stack
}


func (this *FreqStack) Push(val int) {
    this.freqCounter[val]++
    f := this.freqCounter[val]
    this.freqGroup[f] = append(this.freqGroup[f], val)
    if this.maxFreq < f {
        this.maxFreq = f
    }
}


func (this *FreqStack) Pop() int {
    targetGroup := this.freqGroup[this.maxFreq]
    poppedNum := targetGroup[len(targetGroup) - 1]
    this.freqGroup[this.maxFreq] = targetGroup[:len(targetGroup) - 1]
    this.freqCounter[poppedNum]--
    if len(this.freqGroup[this.maxFreq]) == 0 {
        this.maxFreq--
    }
    return poppedNum
}
 
func main() {
  obj := Constructor()
  
  obj.Push(5)
  obj.Push(7)
  obj.Push(5)
  obj.Push(7)
  obj.Push(4)
  obj.Push(5)
  fmt.Printf("obj.Pop() = %v\n", obj.Pop())
  fmt.Printf("obj.Pop() = %v\n", obj.Pop())
  fmt.Printf("obj.Pop() = %v\n", obj.Pop())
  fmt.Printf("obj.Pop() = %v\n", obj.Pop()) 
 
  // obj.Push(0)
  // fmt.Printf("obj.Pop(%v)\n", obj.Pop(0))
}
 