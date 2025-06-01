package solutions

const (
	MAX_SIZE = 1000001
)

type MyHashSet struct {
	hashtable [MAX_SIZE]bool
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.hashtable[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.hashtable[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.hashtable[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
