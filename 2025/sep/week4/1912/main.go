package main

import (
	"container/heap"
	"fmt"
)

type IndexKey struct {
	Shop  int
	Movie int
}

func (k IndexKey) String() string {
	return fmt.Sprintf("%d_%d", k.Shop, k.Movie)
}

type RentingMovie struct {
	Shop     int
	Movie    int
	Price    int
	Offset   int
	IsRented bool
}

type Movies []*RentingMovie

func (m Movies) Len() int {
	return len(m)
}

func (m Movies) Less(i int, j int) bool {
	if !m[i].IsRented {
		if m[i].Price == m[j].Price {
			return m[i].Shop < m[j].Shop
		}
	} else if m[i].Price == m[j].Price {
		if m[i].Shop == m[j].Shop {
			return m[i].Movie < m[j].Movie
		}
		return m[i].Shop < m[j].Shop
	}

	return m[i].Price < m[j].Price
}

func (m *Movies) Pop() any {
	item := (*m)[m.Len() - 1]
	(*m) = (*m)[:m.Len() - 1]
	return item
}

func (m *Movies) Push(x any) {
	item := x.(*RentingMovie)
	item.Offset = m.Len()
	(*m) = append((*m), item)
}

func (m Movies) Swap(i int, j int) {
	m[i].Offset, m[j].Offset = j, i
	m[i], m[j] = m[j], m[i]
}

type MovieRentingSystem struct {
	IndexingKeyOffset map[string]*RentingMovie
	RentedMovieHeap   *Movies
	RentedMovies      map[string]*RentingMovie
	UnrentedMovies    map[int]*Movies
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	unrentedMovies := make(map[int]*Movies)
	indexingMovies := make(map[string]*RentingMovie)

	for _, entry := range entries {
		if _, exists := unrentedMovies[entry[1]]; !exists {
			unrentedMovies[entry[1]] = &Movies{}
		}

		m := &RentingMovie{
			Shop:  entry[0],
			Movie: entry[1],
			Price: entry[2],
		}
		heap.Push(unrentedMovies[entry[1]], m)
		indexingMovies[IndexKey{
			Shop:  entry[0],
			Movie: entry[1],
		}.String()] = m
	}

	return MovieRentingSystem{
		IndexingKeyOffset: indexingMovies,
		RentedMovieHeap:   &Movies{},
		RentedMovies:      make(map[string]*RentingMovie),
		UnrentedMovies:    unrentedMovies,
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	movies, exists := this.UnrentedMovies[movie]
	if !exists {
		return nil
	}

	result := make([]int, 0, 5)
	eligibleMovies := make([]*RentingMovie, 0, 5)
	for movies.Len() > 0 && len(eligibleMovies) < 5 {
		item := heap.Pop(movies)
		movie := item.(*RentingMovie)
		eligibleMovies = append(eligibleMovies, movie)
		result = append(result, movie.Shop)
	}

	if len(eligibleMovies) == 0 {
		return nil
	}

	for _, e := range eligibleMovies {
		heap.Push(movies, e)
	}

	return result
}

func (this *MovieRentingSystem) Rent(shop int, movie int)  {
	key := IndexKey{
		Shop:  shop,
		Movie: movie,
	}

	rentingMovie := this.IndexingKeyOffset[key.String()]
	if rentingMovie == nil {
		return
	}

	heap.Remove(this.UnrentedMovies[movie], rentingMovie.Offset)

	rentedMovie := &RentingMovie{
		Shop:     shop,
		Movie:    movie,
		Price:    rentingMovie.Price,
		IsRented: true,
	}

	this.RentedMovies[key.String()] = rentedMovie
	heap.Push(this.RentedMovieHeap, rentedMovie)
}

func (this *MovieRentingSystem) Drop(shop int, movie int)  {
	key := IndexKey{
		Shop:  shop,
		Movie: movie,
	}

	m, exists := this.RentedMovies[key.String()]
	if !exists {
		return
	}

	rentingMovie := &RentingMovie{
		Shop:  shop,
		Movie: movie,
		Price: m.Price,
	}

	this.IndexingKeyOffset[key.String()] = rentingMovie
	heap.Push(this.UnrentedMovies[movie], rentingMovie)
	heap.Remove(this.RentedMovieHeap, m.Offset)
	delete(this.RentedMovies, key.String())
}

func (this *MovieRentingSystem) Report() [][]int {
	result := make([][]int, 0, 5)
	histories := make([]*RentingMovie, 0, 5)

	for this.RentedMovieHeap.Len() > 0 && len(histories) < 5 {
		item := heap.Pop(this.RentedMovieHeap)
		m := item.(*RentingMovie)
		histories = append(histories, m)
		result = append(result, []int{m.Shop, m.Movie})
	}

	if len(histories) == 0 {
		return nil
	}

	for _, item := range histories {
		heap.Push(this.RentedMovieHeap, item)
	}

	return result
}

func main() {
	// result: [null, [1, 0, 2], null, null, [[0, 1], [1, 2]], null, [0, 1]]
	n := int(3)
	entries := [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}}
	operations := [][]int{
		{0, 1},
		{1, 0, 1},
		{1, 1, 2},
		{3},
		{2, 1, 2},
		{0, 2},
	}

	// result: []
	// n := int(0)
	// entries := [][]int{}
	// operations := [][]int{}

	obj := Constructor(n, entries)

	for _, ope := range operations {
		if ope[0] == 0 {
			movie := ope[1]

			fmt.Printf(
				"obj.Search(%v) = %v\n",
				movie,
				obj.Search(movie),
			)
		} else if ope[0] == 1 {
			shop := ope[1]
			movie := ope[2]
			obj.Rent(shop, movie)
		} else if ope[0] == 2 {
			shop := ope[1]
			movie := ope[2]
			obj.Drop(shop, movie)
		} else if ope[0] == 3 {
			fmt.Printf("obj.Report() = %v\n", obj.Report())
		}
	}
}

