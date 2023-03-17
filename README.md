### This project implements seven data storage structures:
1. Stack
2. Queue
3. Binary search tree
4. Prefix tree
5. Doubly linked list
6. Hash table. 
7. Directed weighted graph

### Stack
Description: a stack is a data storage structure that follows the rule: "First in, last out". This implementation of stack allows you to store elements of the same data type. The data type can be anything.

### Queue
Description: a queue is a data storage structure that follows the rule: "First in, first out". This implementation of queue allows you to store elements of the same data type. The data type can be anything.

### Binary search tree
Description: a binary search tree is a hierarchical data structure in which each node has no more than two children and for which the following conditions:
1. both subtrees - left and right - are binary search trees;
2. for all nodes of the left subtree of an arbitrary node X, the values of keys are less than or equal to the value of key of the node X itself;
3. all nodes of the right subtree of an arbitrary node X have key values greater than the value of key of the node X itself.

Features of implementation: this implementation of a binary search tree allows only integer elements to be stored.

### Prefix tree
Description: a prefix tree is a special kind of search tree that typically uses strings for node keys. Since this structure is designed to search for keys with a specific prefix, it is often used to implement autocompletion. This implementation of the prefix tree will allow you to store only the string data type as a key and value.

### Doubly linked list
Description: doubly linked list is a basic dynamic data structure consisting of nodes containing data and links to the next and previous nodes in the list.

Features of implementation: this implementation of doubly linked list allows you to store elements of the same data type. The data type can be anything. 

### Hash table
Description: a hash table is a data structure that is an efficient data structure for implementing dictionaries, namely, it allows you to store pairs (key, value) and quickly perform three operations: the operation of adding a new pair, the operation of searching, and the operation of deleting a pair by key.

Features of implementation: this implementation of the hash table allows using only a string data type as a key, any data type can be used as a value. To eliminate collisions, the method with open addressing is used.

### Directed weight graph
Description: a graph is a data structure that is a collection of nodes that have data and are connected to other nodes using edges. A bidirectional graph is one whose edges have a direction. A weighted graph is a graph whose edges have a weight.
