package main

import (
    "fmt"
    "math/rand"
)

// size should the size of the square
var Range int = 199999
var id = make([]int, Range)
var sz = make([]int, Range)
var times = 0

func init() {
    for i := 0; i < Range; i++ {
        id[i] = i
    }
}

func main() {
    xRoot := 0// pseudo root for x
    yRoot := 9// pseudo root for y
    for !connected(xRoot, yRoot) {
        union(xRoot, rand.Intn(Range))
    }
    fmt.Printf("total of unions performed: %d\n", times)
}

func union(p, q int) {
    i := root(p)
    j := root(q)
    if sz[i] < sz[j] {
        id[i] = j
        sz[j] += sz[i]
    } else {
        id[j] = i
        sz[i] += sz[j]
    }
    times += 1
}

func root(i int) int {
    for i != id[i] {
        id[i] = id[id[i]]
        i = id[i]
    }
    return i
}

func connected(p, q int) bool {
    return root(p) == root(q)
}
