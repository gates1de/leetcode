package main
import (
	"fmt"
	"reflect"
)

type Node struct {
	visited bool
	depth int
	val string
	nextNodeList []*Node
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(endWord, wordList) ||
		reflect.DeepEqual(wordList, []string{beginWord, endWord}) {
		return 0
	}
	if diff(beginWord, endWord) == 1 {
		return 2
	}

	graph := makeGraph(beginWord, endWord, wordList)
	// fmt.Printf("graph val = %v, nextNodeList = %v\n", graph.val, graph.nextNodeList[0])
	result := bfs(beginWord, endWord, graph)
	return result
}

func bfs(beginWord string, endWord string, node *Node) int {
	queue := []*Node{}
	queue = append(queue, node)
	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]
		// fmt.Printf("currentNode.val = %v, depth = %v\n", currentNode.val, currentNode.depth)
		if currentNode.visited {
			continue
		}
		currentNode.visited = true
		if currentNode.val == endWord {
			// fmt.Printf("finish depth = %v\n", currentNode.depth + 1)
			return currentNode.depth + 1
		}
		for _, n := range currentNode.nextNodeList {
			if n.depth != 0 {
				continue
			}
			n.depth = currentNode.depth + 1
			// fmt.Printf("pushed node = %v\n", n.val)
			queue = append(queue, n)
		}
	}
	return 0
}

func appendNextNodeList(targetNode *Node, nodeList []*Node, visited map[string]bool) {
	nextNodeList := []*Node{}
	targetNodeVal := targetNode.val
	// fmt.Printf("targetNodeVal = %v\n", targetNodeVal)
	for _, node := range nodeList {
		if diff(targetNodeVal, node.val) != 1 {
			continue
		}
		// fmt.Printf("append node.val = %v\n", node.val)
		nextNodeList = append(nextNodeList, node)
	}

	for _, nextNode := range nextNodeList {
		if visited[nextNode.val] {
			// fmt.Printf("already visited = %v\n", nextNode.val)
			continue
		}
		visited[nextNode.val] = true
		appendNextNodeList(nextNode, nodeList, visited)
	}
	targetNode.nextNodeList = nextNodeList
}

func makeGraph(beginWord string, endWord string, wordList []string) *Node {
	nodeList := make([]*Node, 0, len(wordList))
	beginNode := &Node{visited: false, depth: 0, val: beginWord}
	nodeList = append(nodeList, beginNode)
	for _, word := range wordList {
		node := &Node{visited: false, depth: 0, val: word}
		nodeList = append(nodeList, node)
	}

	appendNextNodeList(beginNode, nodeList, map[string]bool{})
	return beginNode
}

func copySlice(target []string) []string {
	result := make([]string, len(target))
	copy(result, target)
	return result
}

func contains(word string, wordList []string) bool {
	for _, w := range wordList {
		if word == w {
			return true
		}
	}
	return false
}

func diff(word1 string, word2 string) int {
	count := int(0)
	for i, _ := range word1 {
		if word1[i] != word2[i] {
			count++
		}
	}
	return count
}

func main() {
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	// beginWord := "hot"
	// endWord := "dog"
	// wordList := []string{"hot", "dog", "dot"}
	// wordList := []string{"hot","cog","dog","tot","hog","hop","pot","dot"}

	// beginWord := "a"
	// endWord := "c"
	// wordList := []string{"a", "b", "c"}

	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","tog","cog"}

	// beginWord := "teach"
	// endWord := "place"
	// wordList := []string{"peale","wilts","place","fetch","purer","pooch","peace","poach","berra","teach","rheum","peach"}

	beginWord := "cet"
	endWord := "ism"
	wordList := []string{"kid","tag","pup","ail","tun","woo","erg","luz","brr","gay","sip","kay","per","val","mes","ohs","now","boa","cet","pal","bar","die","war","hay","eco","pub","lob","rue","fry","lit","rex","jan","cot","bid","ali","pay","col","gum","ger","row","won","dan","rum","fad","tut","sag","yip","sui","ark","has","zip","fez","own","ump","dis","ads","max","jaw","out","btu","ana","gap","cry","led","abe","box","ore","pig","fie","toy","fat","cal","lie","noh","sew","ono","tam","flu","mgm","ply","awe","pry","tit","tie","yet","too","tax","jim","san","pan","map","ski","ova","wed","non","wac","nut","why","bye","lye","oct","old","fin","feb","chi","sap","owl","log","tod","dot","bow","fob","for","joe","ivy","fan","age","fax","hip","jib","mel","hus","sob","ifs","tab","ara","dab","jag","jar","arm","lot","tom","sax","tex","yum","pei","wen","wry","ire","irk","far","mew","wit","doe","gas","rte","ian","pot","ask","wag","hag","amy","nag","ron","soy","gin","don","tug","fay","vic","boo","nam","ave","buy","sop","but","orb","fen","paw","his","sub","bob","yea","oft","inn","rod","yam","pew","web","hod","hun","gyp","wei","wis","rob","gad","pie","mon","dog","bib","rub","ere","dig","era","cat","fox","bee","mod","day","apr","vie","nev","jam","pam","new","aye","ani","and","ibm","yap","can","pyx","tar","kin","fog","hum","pip","cup","dye","lyx","jog","nun","par","wan","fey","bus","oak","bad","ats","set","qom","vat","eat","pus","rev","axe","ion","six","ila","lao","mom","mas","pro","few","opt","poe","art","ash","oar","cap","lop","may","shy","rid","bat","sum","rim","fee","bmw","sky","maj","hue","thy","ava","rap","den","fla","auk","cox","ibo","hey","saw","vim","sec","ltd","you","its","tat","dew","eva","tog","ram","let","see","zit","maw","nix","ate","gig","rep","owe","ind","hog","eve","sam","zoo","any","dow","cod","bed","vet","ham","sis","hex","via","fir","nod","mao","aug","mum","hoe","bah","hal","keg","hew","zed","tow","gog","ass","dem","who","bet","gos","son","ear","spy","kit","boy","due","sen","oaf","mix","hep","fur","ada","bin","nil","mia","ewe","hit","fix","sad","rib","eye","hop","haw","wax","mid","tad","ken","wad","rye","pap","bog","gut","ito","woe","our","ado","sin","mad","ray","hon","roy","dip","hen","iva","lug","asp","hui","yak","bay","poi","yep","bun","try","lad","elm","nat","wyo","gym","dug","toe","dee","wig","sly","rip","geo","cog","pas","zen","odd","nan","lay","pod","fit","hem","joy","bum","rio","yon","dec","leg","put","sue","dim","pet","yaw","nub","bit","bur","sid","sun","oil","red","doc","moe","caw","eel","dix","cub","end","gem","off","yew","hug","pop","tub","sgt","lid","pun","ton","sol","din","yup","jab","pea","bug","gag","mil","jig","hub","low","did","tin","get","gte","sox","lei","mig","fig","lon","use","ban","flo","nov","jut","bag","mir","sty","lap","two","ins","con","ant","net","tux","ode","stu","mug","cad","nap","gun","fop","tot","sow","sal","sic","ted","wot","del","imp","cob","way","ann","tan","mci","job","wet","ism","err","him","all","pad","hah","hie","aim","ike","jed","ego","mac","baa","min","com","ill","was","cab","ago","ina","big","ilk","gal","tap","duh","ola","ran","lab","top","gob","hot","ora","tia","kip","han","met","hut","she","sac","fed","goo","tee","ell","not","act","gil","rut","ala","ape","rig","cid","god","duo","lin","aid","gel","awl","lag","elf","liz","ref","aha","fib","oho","tho","her","nor","ace","adz","fun","ned","coo","win","tao","coy","van","man","pit","guy","foe","hid","mai","sup","jay","hob","mow","jot","are","pol","arc","lax","aft","alb","len","air","pug","pox","vow","got","meg","zoe","amp","ale","bud","gee","pin","dun","pat","ten","mob"}
	result := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("result = %v\n", result)
}

