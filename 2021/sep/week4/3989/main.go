package main
import (
	"fmt"
	"regexp"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, e := range emails {
		a := strings.Split(e, "@")
		l := strings.Split(a[0], "+")[0]
		l = strings.ReplaceAll(l, ".", "")
		m[fmt.Sprintf("%s@%s", l, a[1])] = true
	}
	return len(m)
}

// Slow & Use more memory
func mySolution(emails []string) int {
	m := map[string]bool{}
	for _, email := range emails {
		localAndDomain := strings.Split(email, "@")
		newEmail := ""
		for i, s := range localAndDomain {
			if i == 0 {
				rep := regexp.MustCompile(`\.|\+[A-Za-z\.\+]*`)
				s = rep.ReplaceAllString(s, "")
			}
			newEmail += s
			if i == 0 {
				newEmail += "@"
			}
		}
		// fmt.Printf("email = %v\n", newEmail)
		m[newEmail] = true
	}
	return len(m)
}

func main() {
	// result: 2
	// emails := []string{
	// 	"test.email+alex@leetcode.com",
	// 	"test.e.mail+bob.cathy@leetcode.com",
	// 	"testemail+david@lee.tcode.com",
	// }

	// result: 3
	emails := []string{
		"a@leetcode.com",
		"b@leetcode.com",
		"c@leetcode.com",
	}

	// result: 
	// emails := []string{}

	result := numUniqueEmails(emails)
	fmt.Printf("result = %v\n", result)
}

