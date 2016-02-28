/*
Provides an interface and class for making linked lists.
*/
package llist

import (
	"github.com/KTAtkinson/linked"
)

type List struct {
	head linked.Noder
	tail linked.Noder
}

// Append adds the given node as the tail of the list.
func (l *List) Append(n linked.Noder) (err error) {
	if l.head != nil {
		l.head.SetPrev(n)
		n.SetNext(l.head)
	} else {
		l.tail = n
	}

	l.head = n
	return err
}

// Push adds the given node as the head of the list.
func (l *List) Push(n linked.Noder) (err error) {
	return err
}

// Pop returns the node at the head of the list.
func (l *List) Pop() (n linked.Noder, err error) {
	return n, err
}
