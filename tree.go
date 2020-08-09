package main

import (
    "fmt"
)

type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}

type Stack struct {
    Value *Tree
    Next *Stack
}

func Push(s *Stack, t *Tree) *Stack {
    ns := &Stack{t, s}
    return ns
}

func Pop(s *Stack) *Stack {
    return s.Next
}


func (t *Tree) WalkRec(ch chan int) {
    if t.Left != nil {
        t.Left.WalkRec(ch)
    }
    ch <- t.Value
    if t.Right != nil {
        t.Right.WalkRec(ch)
    }
}

func (t *Tree) Walk(ch chan int) {
    t.WalkRec(ch)
    close(ch)
}

func (t *Tree) SizeRec(l int) int {
    if t.Left != nil {
        l = t.Left.SizeRec(l)
    }
    l++
    if t.Right != nil {
        l = t.Right.SizeRec(l)
    }
    return l
}

func (t *Tree) Size() int {
    return t.SizeRec(0)
}

func (t *Tree) Add(e int) {
    a := t
    var b *Tree
    for a != nil {
        b = a
        if e > a.Value {
            a = a.Right
        } else {
            a = a.Left
        }
    }
    nt := &Tree{Value: e}
    if e > b.Value {
        b.Right = nt
    } else {
        b.Left = nt
    }
}

func Same(t1, t2 *Tree) bool {
    if t1.Size() != t2.Size() {
        return false
    }
    ch1 := make(chan int)
    ch2 := make(chan int)
    go t1.Walk(ch1)
    go t2.Walk(ch2)
    for v := range ch1 {
        if v != <-ch2 {
            return false
        }
    }
    return true
}

func main() {
    t := &Tree{Value: 50}
    t.Add(25)
    t.Add(60)
    t.Add(30)
    t.Add(15)
    t.Add(40)
    t.Add(60)
    t.Add(45)
    r := &Tree{Value: 50}
    r.Add(30)
    r.Add(25)
    r.Add(40)
    r.Add(15)
    r.Add(45)
    r.Add(70)
    r.Add(60)
    fmt.Println(Same(t, r))
}
