package main
import (
	"fmt"
	heap "container/heap"
)

type Food struct {
    Name string
    Rating int
}

type maxheap []Food

func (h maxheap) Peek() Food {
    return h[0]
}

func (h maxheap) Len() int {
    return len(h)
}

func (h maxheap) Less(i int, j int) bool {
    if h[i].Rating == h[j].Rating {
        return h[i].Name < h[j].Name
    } else {
        return h[i].Rating > h[j].Rating
    }
}

func (h maxheap) Swap(i int, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Push(a interface{}) {
    *h = append(*h, a.(Food))
}

func (h *maxheap) Pop() interface{} {
    l := len(*h)
    res := (*h)[l - 1]
    *h = (*h)[:l - 1]
    return res
}

type FoodRatings struct {
	FoodCuisines map[string]string
	FoodRatings map[string]int
	Cuisines map[string]*maxheap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    result := FoodRatings{
        FoodCuisines: make(map[string]string),
        FoodRatings: make(map[string]int),
        Cuisines: make(map[string]*maxheap),
    }

    for i, v := range foods {
        result.FoodCuisines[v] = cuisines[i]
        result.FoodRatings[v] = ratings[i]
        if _, ok := result.Cuisines[cuisines[i]]; !ok {
            result.Cuisines[cuisines[i]] = &maxheap{}
        }
        heap.Push(result.Cuisines[cuisines[i]], Food{
            Name: v,
            Rating: ratings[i],
        })
    }

    return result
}

func (this *FoodRatings) ChangeRating(food string, newRating int)  {
    this.FoodRatings[food] = newRating
    heap.Push(this.Cuisines[this.FoodCuisines[food]], Food{
        Name: food,
        Rating: newRating,
    })
}

func (this *FoodRatings) HighestRated(cuisine string) string {
    for {
        food := this.Cuisines[cuisine].Peek()
        if this.FoodRatings[food.Name] != food.Rating {
            heap.Pop(this.Cuisines[cuisine])
            continue
        } else {
            return food.Name
        }
    }
    return ""
}

func main() {
	// result: [null, "kimchi", "ramen", null, "sushi", null, "ramen"]
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}
	operations := [][]int{{0,0},{0,1},{1,2,16},{0,1},{1,4,16},{0,1}}

	// result: 
	// operations := [][]int{}

	obj := Constructor(foods, cuisines, ratings)
	for _, ope := range operations {
		if ope[0] == 0 {
			cuisine := ope[1]

			if cuisine == 0 {
				fmt.Printf("obj.HighestRated(%v) = %v\n", "korean", obj.HighestRated("korean"))
			} else if cuisine == 1 {
				fmt.Printf("obj.HighestRated(%v) = %v\n", "japanese", obj.HighestRated("japanese"))
			} else if cuisine == 2 {
				fmt.Printf("obj.HighestRated(%v) = %v\n", "greek", obj.HighestRated("greek"))
			}
		} else if ope[0] == 1 {
			food := ope[1]
			newRating := ope[2]

			if food == 0 {
				obj.ChangeRating("kimchi", newRating)
			} else if food == 1 {
				obj.ChangeRating("miso", newRating)
			} else if food == 2 {
				obj.ChangeRating("sushi", newRating)
			} else if food == 3 {
				obj.ChangeRating("mmoussaka", newRating)
			} else if food == 4 {
				obj.ChangeRating("ramen", newRating)
			} else if food == 5 {
				obj.ChangeRating("bulgogi", newRating)
			}
		}
	}
}

