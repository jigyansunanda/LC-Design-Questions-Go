## Problem Statement

A  [**trie**](https://en.wikipedia.org/wiki/Trie) (pronounced as "try") or **prefix tree** is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:
-   `Trie()`  Initializes the trie object.
-   `void insert(String word)`  Inserts the string  `word`  into the trie.
-   `boolean search(String word)`  Returns  `true`  if the string  `word`  is in the trie (i.e., was inserted before), and  `false`  otherwise.
-   `boolean startsWith(String prefix)`  Returns  `true`  if there is a previously inserted string  `word`  that has the prefix  `prefix`, and  `false`  otherwise.

**Example 1:**

**Input**</br>
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]</br>
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

**Output**</br>
[null, null, true, false, true, null, true]

**Explanation**</br>
Trie trie = new Trie();</br>
trie.insert("apple");</br>
trie.search("apple");   // return True</br>
trie.search("app");     // return False</br>
trie.startsWith("app"); // return True</br>
trie.insert("app");</br>
trie.search("app");     // return True</br>