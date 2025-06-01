## Problem Statement

Design a HashSet without using any built-in hash table libraries.</br>
Implement `MyHashSet` class:

-   `void add(key)`  Inserts the value  `key`  into the HashSet.
-   `bool contains(key)`  Returns whether the value  `key`  exists in the HashSet or not.
-   `void remove(key)`  Removes the value  `key`  in the HashSet. If  `key`  does not exist in the HashSet, do nothing.

**Example 1:**

**Input**</br>
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]</br>
[[], [1], [2], [1], [3], [2], [2], [2], [2]]

**Output**</br>
[null, null, null, true, false, null, true, null, false]

**Explanation**</br>
MyHashSet myHashSet = new MyHashSet();</br>
myHashSet.add(1);      // set = [1]</br>
myHashSet.add(2);      // set = [1, 2]</br>
myHashSet.contains(1); // return True</br>
myHashSet.contains(3); // return False, (not found)</br>
myHashSet.add(2);      // set = [1, 2]</br>
myHashSet.contains(2); // return True</br>
myHashSet.remove(2);   // set = [1]</br>
myHashSet.contains(2); // return False, (already removed)</br>

**Constraints:**

-   `0 <= key <= 10^6`
-   At most  `10^4`  calls will be made to  `add`,  `remove`, and  `contains`.