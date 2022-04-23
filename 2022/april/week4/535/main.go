package main
import (
	"fmt"
	"math/rand"
)

type Codec struct {
	Map map[string]string
}

func Constructor() Codec {
	return Codec{Map: map[string]string{}}
}

func (this *Codec) encode(longUrl string) string {
	shortUrl := fmt.Sprintf("https://tinyurl.com/%v", generateRandomString(6))
	this.Map[shortUrl] = longUrl
	this.Map[longUrl] = shortUrl
	return shortUrl
}

func (this *Codec) decode(shortUrl string) string {
	if longUrl, ok := this.Map[shortUrl]; ok {
		return longUrl
	}
	return shortUrl
}

func generateRandomString(n int) string {
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func main() {
	// result: "https://leetcode.com/problems/design-tinyurl"
	// longUrl := "https://leetcode.com/problems/design-tinyurl"

	// result: "http://exmaple.com/about"
	longUrl := "http://exmaple.com/about"

	// result: 
	// longUrl := ""

	obj := Constructor()
	shortUrl := obj.encode(longUrl)
	fmt.Printf("short url = %v\n", shortUrl)
	result := obj.decode(shortUrl)
	fmt.Printf("result = %v\n", result)
}

