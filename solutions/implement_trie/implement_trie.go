package implement_trie

const MaxCharacterCount = 26

type Trie struct {
	isEnd      bool
	dictionary [MaxCharacterCount]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	for _, c := range word {
		index := c - 'a'
		if trie.dictionary[index] == nil {
			trie.dictionary[index] = &Trie{}
		}
		trie = trie.dictionary[index]
	}
	trie.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	for _, c := range word {
		index := c - 'a'
		if trie.dictionary[index] == nil {
			return false
		}
		trie = trie.dictionary[index]
	}
	return trie.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		index := c - 'a'
		if trie.dictionary[index] == nil {
			return false
		}
		trie = trie.dictionary[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
