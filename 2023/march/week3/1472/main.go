package main
import (
	"fmt"
)

type BrowserHistory struct {
	Index int
	Urls []string
}

func Constructor(homepage string) BrowserHistory {
	result := BrowserHistory{Index: 0, Urls: make([]string, 0, 5000)}
	result.Urls = append(result.Urls, homepage)
	return result
}

func (this *BrowserHistory) Visit(url string)  {
	if this.Index == len(this.Urls) - 1 {
		this.Urls = append(this.Urls, url)
		this.Index++
		return
	}

	this.Urls = this.Urls[:this.Index + 1]
	this.Urls = append(this.Urls, url)
	this.Index = len(this.Urls) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	nextIndex := max(0, this.Index - steps)
	this.Index = nextIndex
	return this.Urls[nextIndex]
}

func (this *BrowserHistory) Forward(steps int) string {
	nextIndex := min(len(this.Urls) - 1, this.Index + steps)
	this.Index = nextIndex
	return this.Urls[nextIndex]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [null,null,null,null,"facebook.com","google.com","facebook.com",null,"linkedin.com","google.com","leetcode.com"]
	homepage := "leetcode.com"
	operations := [][]interface{}{
		{0, "google.com", 0},
		{0, "facebook.com", 0},
		{0, "youtube.com", 0},
		{1, "back", 1},
		{1, "back", 1},
		{2, "forward", 1},
		{0, "linkedin.com", 0},
		{2, "forward", 2},
		{1, "back", 2},
		{1, "back", 7},
	}

	// result: 
	// homepage := ""
	// operations := [][]interface{}{
	// 	{0, "visit", 0},
	// 	{1, "back", 0},
	// 	{2, "forward", 0},
	// }

	obj := Constructor(homepage)

	for _, val := range operations {
		operation := val[0].(int)
		url := val[1].(string)
		steps := val[2].(int)

		if operation == 0 {
			obj.Visit(url)
		} else if operation == 1 {
			fmt.Printf("obj.Back(%v) = %v\n", steps, obj.Back(steps))
		} else if operation == 2 {
			fmt.Printf("obj.Forward(%v) = %v\n", steps, obj.Forward(steps))
		}
	}
}

