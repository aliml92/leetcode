package ladderlength127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(wordList, endWord) {
		return 0
	}
	nei := make(map[string][]string)
	wordList = append(wordList, beginWord)
	for _, word := range wordList {
		for j := 0; j < len(word); j++ {
			pattern := word[:j] + "*" + word[j+1:]
			nei[pattern] = append(nei[pattern], word)
		}
	}
	visit := make(map[string]struct{})
	visit[beginWord] = struct{}{}
	res := 1
	q := &Queque{}
	q.enQueque(beginWord)
	for q.size() > 0 {
		n := q.size()
		for i := 0; i < n; i++ {
			word := q.deQueque()
			if word == endWord {
				return res
			}
			for j := 0; j < len(word); j++ {
				pattern := word[:j] + "*" + word[j+1:]
				for _, neiWord := range nei[pattern] {
					_, ok := visit[neiWord]
					if !ok {
						visit[neiWord] = struct{}{}
						q.enQueque(neiWord)
					}
				}
			}
		}
		res++
	}
	return 0
}

type Queque struct {
	len  int
	head *Node
	tail *Node
}

type Node struct {
	value string
	next  *Node
}

func (q *Queque) enQueque(item string) {
	n := &Node{
		value: item,
		next:  nil,
	}
	if q.len == 0 {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.len++
}

func (q *Queque) deQueque() string {
	val := q.head.value
	nextPtr := q.head.next
	q.head = nil
	q.head = nextPtr
	q.len--
	return val
}

func (q *Queque) size() int {
	return q.len
}

//  helper functions

func contains(wordList []string, endWord string) bool {
	for _, s := range wordList {
		if endWord == s {
			return true
		}
	}
	return false
}
