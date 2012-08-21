package main

import "fmt"

var Size = 10
var id = make([]int, Size)
var sz = make([]int, Size)

func init() {
    for i := 0; i < Size; i++ {
        id[i] = i
    }
}

func main() {
    union(4, 3)
    fmt.Printf("id[3] = %d    id[4] = %d\n", id[3], id[4])
    union(3, 8)
    fmt.Printf("root of 8 should be 4, have: %d \n", root(8))
    union(2, 4)
    fmt.Printf("root of 4 should be 2, have: %d\n", root(2))
}

func union(p, q int) {
    i := root(p)
    j := root(q)
    fmt.Printf("i is %d -- j is %d\n", i, j)
    if sz[i] < sz[j] {
        id[i] = j
        sz[j] += sz[i]
    } else {
        id[j] = i
        sz[i] += sz[j]
    }
}

func root(i int) int {
    for i != id[i] {
        // path compression for weighted quick union
        id[i] = id[id[i]]
        i = id[i]
    }
    return i
}

func connected(p, q int) bool {
    return root(p) == root(q)
}
