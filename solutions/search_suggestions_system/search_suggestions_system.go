package search_suggestions_system

const MaxCharacterCount = 26

type Trie struct {
	isEnd      bool
	finalWord  string
	dictionary [MaxCharacterCount]*Trie
}

func (trie *Trie) insertWord(word string) {
	for _, char := range word {
		index := char - 'a'
		if trie.dictionary[index] == nil {
			trie.dictionary[index] = &Trie{}
		}
		trie = trie.dictionary[index]
	}
	trie.isEnd = true
	trie.finalWord = word
}

func (trie *Trie) getFinalWordsInPath() []string {
	var wordsInPath []string

	if trie.isEnd {
		wordsInPath = append(wordsInPath, trie.finalWord)
	}

	for index := range trie.dictionary {
		if trie.dictionary[index] != nil {
			wordsFromChildren := trie.dictionary[index].getFinalWordsInPath()
			for _, word := range wordsFromChildren {
				if len(wordsInPath) >= 3 {
					return wordsInPath
				} else {
					wordsInPath = append(wordsInPath, word)
				}
			}
		}
	}

	return wordsInPath
}

func (trie *Trie) getSuggestions(word string) []string {
	for _, ch := range word {
		index := ch - 'a'
		if trie.dictionary[index] == nil {
			return []string{}
		}
		trie = trie.dictionary[index]
	}
	return trie.getFinalWordsInPath()
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := &Trie{}
	for _, product := range products {
		trie.insertWord(product)
	}

	var suggestions [][]string

	for endIndex := range searchWord {
		subString := searchWord[0 : endIndex+1]
		currentSuggestions := trie.getSuggestions(subString)
		suggestions = append(suggestions, currentSuggestions)
	}

	return suggestions
}
