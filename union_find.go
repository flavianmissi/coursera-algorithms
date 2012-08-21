package main

import "fmt"

var Size = 10
var id = make([]int, Size)

func init() {
    for i := 0; i < Size; i++ {
        id[i] = i
    }
}

func main() {
    union(3, 4)
    fmt.Printf("%d -- %d\n", id[3], id[4])
    // all ids should change to 8
    union(3, 8)
    fmt.Printf("3 id: %d -- 4 id: %d -- 8 id: %d\n", id[3], id[4], id[8])
}

/*
 * merge components containing p and q
 * change all entries whose id equals
 * id[p] to id[q]
 * * p and q are connected if they have
 * the same id
 */
func union(p, q int) {
    pId := id[p]
    qId := id[q]
    for i := 0; i < len(id); i++ {
        if id[i] == pId {
            id[i] = qId
        }
    }
}

/*
 * Check if p and q have the same id
 */
func find(p, q int) bool {
    return id[p] == id[q]

}
