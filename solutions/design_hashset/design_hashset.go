package design_hashset

const (
	MaxHashSetSize = 1000001
)

type MyHashSet struct {
	hashtable [MaxHashSetSize]bool
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (hashSet *MyHashSet) Add(key int) {
	hashSet.hashtable[key] = true
}

func (hashSet *MyHashSet) Remove(key int) {
	hashSet.hashtable[key] = false
}

func (hashSet *MyHashSet) Contains(key int) bool {
	return hashSet.hashtable[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
