/*
Provides an interface and class for making linked lists.
*/
package llist

import (
	"github.com/KTAtkinson/linked"
    "errors"
)
type List struct {
	head linked.Noder
	tail linked.Noder
}

// Append adds the given node as the tail of the list.
func (l *List) Append(n linked.Noder) (err error) {
    if l.head == nil {
        l.head = n
    }
    if l.tail != nil {
        n.SetPrev(l.tail)
        l.tail.SetNext(n)
    }

	l.tail = n
	return err
}
// Push adds the given node as the head of the list.
func (l *List) Push(n linked.Noder) (err error) {
    if l.tail == nil {
        l.tail = n
    }
    if l.head != nil {
        n.SetNext(l.head)
        l.head.SetPrev(n)
    }

    l.head = n
	return err
}

// Pop returns the node at the head of the list.
func (l *List) Pop() (n linked.Noder, err error) {
    if l.head == nil {
        return n, errors.New("Can't pop from an empty list.")
    }
    newHead, isLast, _ := l.head.Next()
    oldHead := l.head

    if !isLast {
        newHead.SetPrev(n)
    }
    l.head = newHead
	return oldHead, err
}
