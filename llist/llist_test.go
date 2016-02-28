package llist

import (
	"testing"
)

import "github.com/KTAtkinson/linked/node"

func generateCases() []struct{ numNodes int } {
	cases := []struct {
		numNodes int
	}{
		{0},
		{1},
		{2},
		{5},
	}

	return cases
}

func generateLinks(numNodes int) (head *node.Dnode, tail *node.Dnode) {
	if numNodes > 0 {
		head = new(node.Dnode)
		head.SetData(0)
	}

	lastNode := head
	for i := 1; i < numNodes-1; i++ {
		newNode := new(node.Dnode)
		newNode.SetData(0)
		newNode.SetPrev(lastNode)
		lastNode.SetNext(newNode)

		lastNode = newNode
	}

	if numNodes > 1 {
		tail = new(node.Dnode)
		tail.SetData(0)
		lastNode.SetNext(tail)
		tail.SetPrev(lastNode)
	}

	return head, tail
}

func TestAppend(t *testing.T) {
	cases := generateCases()
	for _, c := range cases {
		head, tail := generateLinks(c.numNodes)
		linkedList := new(List)
		if head != nil {
			linkedList.head = head
		}
		if tail != nil {
			linkedList.tail = tail
		}
		newNode := new(node.Dnode)
		newNode.SetData(0)

		linkedList.Append(newNode)
		if linkedList.head != newNode {
			t.Errorf("Expected new node to be the head of the list when lists length was %#v.", c.numNodes)
		}

		if head != nil {
			nextToNew, _, _ := newNode.Next()
			if nextToNew != head {
				t.Errorf("New head was not linked to list, when list length was %#v.",
					c.numNodes)
			}

			prevToHead, _, _ := head.Prev()
			if prevToHead != newNode {
				t.Error("Old head does not referance the new node.")
			}
		}
	}
}

func TestPush(t *testing.T) {
	cases := generateCases()
	for _, c := range cases {
		head, tail := generateLinks(c.numNodes)
		newNode := new(node.Dnode)
		newNode.SetData(0)
		linkedList := &List{head: head, tail: tail}

		linkedList.Push(newNode)
		if linkedList.tail != newNode {
			t.Error("The tail of the list was not changed.")
		}
		if tail != nil {
			if next, _, _ := tail.Next(); next != newNode {
				t.Error("The old tail does not refer to the new node.")
			}
			if prev, _, _ := newNode.Prev(); prev != tail {
				t.Error("The new node does not refer to the tail.")
			}
		}
	}
}

func TestPop(t *testing.T) {
	cases := generateCases()
	for _, c := range cases {
		head, tail := generateLinks(c.numNodes)
		linkedList := &List{head: head, tail: tail}

		popped, err := linkedList.Pop()
		if c.numNodes < 1 {
			if popped != nil {
				t.Error("Non-nil value popped, when there are no nodes.")
			}
			if err == nil {
				t.Error("No error returned when value was popped from an empyty list.")
			}
		} else {
			if popped != head {
				t.Errorf("Expected %d to be popped off the list, found %d",
					head, popped)
			}
		}
	}
}
