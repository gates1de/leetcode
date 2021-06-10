package main
import (
	"fmt"
)

type MyCalendar struct {
	Books [][]int
	Activity [][]int
}


func Constructor() MyCalendar {
	return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _, v := range this.Activity {
        if v[0] < end && start < v[1] {
            return false
        }
    }
    this.Activity = append(this.Activity, []int{start, end})
    return true
}

// Slow & Use more memory
func (this *MyCalendar) MyAnswer(start int, end int) bool {
	for _, book := range this.Books {
		if (start <= book[0] && book[0] < end) || (start < book[1] && book[1] <= end) || (book[0] <= start && end <= book[1]) {
			return false
		}
	}
	this.Books = append(this.Books, []int{start, end})
	return true
}


func main() {
	obj := Constructor()
	// fmt.Printf("obj.Book(10, 20) = %v\n", obj.Book(10, 20))
	// fmt.Printf("books = %v\n", obj.Books)
	// fmt.Printf("obj.Book(15, 25) = %v\n", obj.Book(15, 25))
	// fmt.Printf("books = %v\n", obj.Books)
	// fmt.Printf("obj.Book(20, 30) = %v\n", obj.Book(20, 30))
	// fmt.Printf("books = %v\n", obj.Books)

	// books := [][]int{{20,29},{13,22},{44,50},{1,7},{2,10},{14,20},{19,25},{36,42},{45,50},{47,50},{39,45},{44,50},{16,25},{45,50},{45,50},{12,20},{21,29},{11,20},{12,17},{34,40},{10,18},{38,44},{23,32},{38,44},{15,20},{27,33},{34,42},{44,50},{35,40},{24,31}}
	books := [][]int{{23,32},{42,50},{6,14},{0,7},{21,30},{26,31},{46,50},{28,36},{0,6}}
	for i, book := range books {
		fmt.Printf("%v: obj.Book(book[0], book[1]) = %v\n", i, obj.Book(book[0], book[1]))
		// fmt.Printf("books = %v\n", obj.Books)
	}
}

