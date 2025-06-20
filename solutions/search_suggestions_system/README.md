## Problem Statement

You are given an array of strings `products` and a string `searchWord`.

Design a system that suggests at most three product names from `products` after each character of `searchWord` is typed. Suggested products should have common prefix with `searchWord`. If there are more than three products with a common prefix return the three lexicographically minimums products.

Return _a list of lists of the suggested products after each character of_ `searchWord` _is typed_.

**Example 1:**

**Input:**
```text
products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
```

**Output:**
```text
[["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],["mouse","mousepad"],["mouse","mousepad"]]
```

**Explanation:**
```text
products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"].
After typing m and mo all products match and we show user ["mobile","moneypot","monitor"].
After typing mou, mous and mouse the system suggests ["mouse","mousepad"].
```

**Example 2:**

**Input:**
```text
products = ["havana"], searchWord = "havana"
```

**Output:**
```text
[["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
```

**Explanation:**
```text
The only word "havana" will be always suggested while typing the search word.
```

**Constraints:**
-   `1 <= products.length <= 1000`
-   `1 <= products[i].length <= 3000`
-   `1 <= sum(products[i].length) <= 2 * 10^4`
-   All the strings of `products` are **unique**.
-   `products[i]` consists of lowercase English letters.
-   `1 <= searchWord.length <= 1000`
-   `searchWord` consists of lowercase English letters.