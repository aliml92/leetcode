package designhashmap706

type MyHashMap struct {
	ht map[int]int
}

func Constructor() MyHashMap {
	h := MyHashMap{
		ht: make(map[int]int, 0),
	}
	return h
}

func (this *MyHashMap) Put(key int, value int) {
	this.ht[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if val, ok := this.ht[key]; ok {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.ht, key)
}
