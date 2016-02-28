package llist

import (
    "github.com/KTAtkinson/linked/node"
    "github.com/KTAtkinson/linked"
	"testing"
)


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

func generateLinks(numNodes int) (head linked.Noder, tail linked.Noder) {
	if numNodes > 0 {
		head = new(node.Dnode)
		head.SetData(0)
        tail = head
	}

	lastNode := head
	for i := 1; i < numNodes-1; i++ {
		newNode := new(node.Dnode)
		newNode.SetData(i)
		newNode.SetPrev(lastNode)
		lastNode.SetNext(newNode)

		lastNode = newNode
	}

	if numNodes > 1 {
		tail = new(node.Dnode)
		tail.SetData(numNodes-1)
		lastNode.SetNext(tail)
		tail.SetPrev(lastNode)
	}

	return head, tail
}

func TestAppend(t *testing.T) {
	cases := generateCases()
	for _, c := range cases {
		head, tail := generateLinks(c.numNodes)
		linkedList := &List{head: head, tail: tail}
		newNode := new(node.Dnode)
		newNode.SetData(c.numNodes+1)

		linkedList.Append(newNode)
		if linkedList.tail != newNode {
			t.Errorf("Expected new node to be the tail of the list when lists length was %#v.", c.numNodes)
		}

		if tail != nil {
			if prev, _, _ := newNode.Prev(); prev != tail {
				t.Errorf("Exected %#v to preceed the new node, found %#v.",
                         tail, prev)
			}

			if next, _, _ := tail.Next(); next != newNode {
				t.Errorf("Found %#v next to tail, expected %#v.", next, newNode)
			}
		} else {
            if linkedList.head != newNode {
                t.Errorf("When list was empty expected head to be %#v, found %#v",
                         newNode, linkedList.head)
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
		if linkedList.head != newNode {
			t.Error("The tail of the list was not changed.")
		}
		if head != nil {
			if next, _, _ := newNode.Next(); next != head {
				t.Error("The old tail does not refer to the new node.")
			}
			if prev, _, _ := head.Prev(); prev != newNode {
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
