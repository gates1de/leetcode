package main
import (
	"fmt"
	"strconv"
	"sort"
)

type Tokenizer struct {
    Index int
    Str   string
}

func (t *Tokenizer) NextToken() (string, int) {
    if t.Index == len(t.Str) {
        return "EOT", 0
    }
    
    if t.Str[t.Index] == '(' {
        t.Index++
        return "(", 0
    }
    
    if t.Str[t.Index] == ')' {
        next := t.Index + 1
        for next < len(t.Str) && t.Str[next] >= '0' && t.Str[next] <= '9' {
            next++
        }

        num := t.Str[t.Index + 1:next]
        t.Index = next
        n, _ := strconv.Atoi(num)

        if n == 0 {
            n = 1
        }

        return ")", n
    }
    
    next := t.Index + 1
    for next < len(t.Str) && t.Str[next] >= 'a' && t.Str[next] <= 'z' {
        next++
    }

    elm := t.Str[t.Index:next]
    t.Index = next
    for next < len(t.Str) && t.Str[next] >= '0' && t.Str[next] <= '9' {
        next++
    }

    num := t.Str[t.Index:next]
    t.Index = next
    n, _ := strconv.Atoi(num)
    if n == 0 {
        n = 1
    }

    return elm, n
}

func countOfAtoms(formula string) string {
    var result string
    t := &Tokenizer{Str: formula}
    l, m := count(t)
    sort.Strings(l)
    
    for _, elem := range l {
        result += elem
        if m[elem] != 1 {
            result += strconv.Itoa(m[elem])
        }
    }
    return result
}

func count(t *Tokenizer) ([]string, map[string]int) {
    eList := make([]string, 0)
    m := make(map[string]int)
    
    for {
        token, freq := t.NextToken()
        if token == "EOT" {
            break
        }    
        
        if token == "(" {
            _, newMap := count(t)
            
            for key, val := range newMap {
                if m[key] == 0 {
                    eList = append(eList, key)
                }
                m[key] += val
            }
        } else if token == ")" {
            for key := range m {
                m[key] *= freq
            }
            break
        } else {
            if m[token] == 0 {
                eList = append(eList, token)
            }
            m[token] += freq
        }
    }

    return eList, m
}

func main() {
	// result: "H2O"
	formula := "H2O"

	// result: "H2MgO2"
	// formula := "Mg(OH)2"

	// result: "K4N2O14S4"
	// formula := "K4(ON(SO3)2)2"

	// result: ""
	// formula := ""

	result := countOfAtoms(formula)
	fmt.Printf("result = %v\n", result)
}

