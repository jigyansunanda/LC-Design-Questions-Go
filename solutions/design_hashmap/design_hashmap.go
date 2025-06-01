package design_hashmap

const (
	MaxHashTableSize = 1000001
	UnAssignedValue  = -1
)

type MyHashMap struct {
	hashTable [MaxHashTableSize]int
}

func Constructor() MyHashMap {
	this := MyHashMap{}
	for i, _ := range this.hashTable {
		this.hashTable[i] = UnAssignedValue
	}
	return this
}

func (hashMap *MyHashMap) Put(key int, value int) {
	hashMap.hashTable[key] = value
}

func (hashMap *MyHashMap) Get(key int) int {
	return hashMap.hashTable[key]
}

func (hashMap *MyHashMap) Remove(key int) {
	hashMap.hashTable[key] = UnAssignedValue
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
