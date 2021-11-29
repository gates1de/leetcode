package main
import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	emailMap := make(map[string]int)
	indexConnections := make([]int, len(accounts))
	for i := range indexConnections {
		indexConnections[i] = i
	}

	for accountIndex, emails := range accounts {
		for i := 1; i < len(emails); i++ {
			if existIndex, exist := emailMap[emails[i]]; exist {
				for {
					if indexConnections[existIndex] == existIndex {
						indexConnections[existIndex] = accountIndex
						break
					}
					existIndex, indexConnections[existIndex] = indexConnections[existIndex], accountIndex
				}
			} else {
				emailMap[emails[i]] = accountIndex
			}
		}
	}

	for i := len(indexConnections)-1; i >= 0; i-- {
		indexConnections[i] = indexConnections[indexConnections[i]]
	}

	magic := make(map[int][]string)
	for email, index := range emailMap {
		if _, exist := magic[indexConnections[index]]; !exist {
			magic[indexConnections[index]] = []string{accounts[index][0]}
		}
		magic[indexConnections[index]] = append(magic[indexConnections[index]], email)
	}

	result := make([][]string, 0, len(magic))
	for _, v := range magic {
		sort.Strings(v[1:])
		result = append(result, v)
	}
	return result
}

// Wrong Answer
func ngSolution(accounts [][]string) [][]string {
	m := map[string][]int{}
	for i, addresses := range accounts {
		for _, address := range addresses[1:] {
			if m[address] == nil {
				m[address] = []int{}
			}
			m[address] = append(m[address], i)
		}
	}
	// fmt.Printf("m = %v\n", m)

	result := [][]string{}
	added := map[int]bool{}
	for _, indices := range m {
		if len(indices) == 1 {
			continue
		}

		allIndices := []int{}
		for _, i := range indices {
			addresses := accounts[i]
			for _, address := range addresses[1:] {
				if len(m[address]) <= 1 {
					continue
				}
				allIndices = append(allIndices, m[address]...)
			}
		}

		// fmt.Printf("allIndices = %v\n", allIndices)

		tmp := map[string]string{}
		for _, i := range allIndices {
			if added[i] {
				continue
			}
			added[i] = true
			addresses := accounts[i]
			name := addresses[0]
			tmp["name"] = name
			for _, address := range addresses[1:] {
				tmp[address] = ""
			}
		}

		if len(tmp) == 0 {
			continue
		}

		tmp2 := []string{tmp["name"]}
		for address, _ := range tmp {
			if address == "name" {
				continue
			}
			tmp2 = append(tmp2, address)
		}
		sort.Strings(tmp2)
		result = append(result, tmp2)
	}

	for i, addresses := range accounts {
		if added[i] {
			continue
		}
		name := addresses[0]
		tmp := addresses[1:]
		sort.Strings(tmp)
		result = append(result, append([]string{name}, tmp...))
	}

	return result
}

func main() {
	// result: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
	// accounts := [][]string{
	// 	{"John","johnsmith@mail.com","john_newyork@mail.com"},
	// 	{"John","johnsmith@mail.com","john00@mail.com"},
	// 	{"Mary","mary@mail.com"},
	// 	{"John","johnnybravo@mail.com"},
	// }

	// result: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]
	// accounts := [][]string{
	// 	{"Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"},
	// 	{"Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"},
	// 	{"Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"},
	// 	{"Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"},
	// 	{"Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"},
	// }

	// result: [["Alex","Alex0@m.co","Alex4@m.co","Alex5@m.co"],["Ethan","Ethan0@m.co","Ethan3@m.co"],["Gabe","Gabe0@m.co","Gabe2@m.co","Gabe3@m.co","Gabe4@m.co"],["Kevin","Kevin2@m.co","Kevin4@m.co"]]
	// accounts := [][]string{
	// 	{"Alex","Alex5@m.co","Alex4@m.co","Alex0@m.co"},
	// 	{"Ethan","Ethan3@m.co","Ethan3@m.co","Ethan0@m.co"},
	// 	{"Kevin","Kevin4@m.co","Kevin2@m.co","Kevin2@m.co"},
	// 	{"Gabe","Gabe0@m.co","Gabe3@m.co","Gabe2@m.co"},
	// 	{"Gabe","Gabe3@m.co","Gabe4@m.co","Gabe2@m.co"},
	// }

	// result: [["David","David0@m.co","David1@m.co","David3@m.co","David4@m.co","David5@m.co"]]
	// accounts := [][]string{
	// 	{"David","David0@m.co","David4@m.co","David3@m.co"},
	// 	{"David","David5@m.co","David5@m.co","David0@m.co"},
	// 	{"David","David1@m.co","David4@m.co","David0@m.co"},
	// 	{"David","David0@m.co","David1@m.co","David3@m.co"},
	// 	{"David","David4@m.co","David1@m.co","David3@m.co"},
	// }

	// result: [["David","David0@m.co","David1@m.co","David2@m.co","David3@m.co","David4@m.co","David5@m.co"]]
	accounts := [][]string{
		{"David","David0@m.co","David1@m.co"},
		{"David","David3@m.co","David4@m.co"},
		{"David","David4@m.co","David5@m.co"},
		{"David","David2@m.co","David3@m.co"},
		{"David","David1@m.co","David2@m.co"},
	}

	// result: 
	// accounts := [][]string{
	// }

	result := accountsMerge(accounts)
	fmt.Printf("result = %v\n", result)
}

