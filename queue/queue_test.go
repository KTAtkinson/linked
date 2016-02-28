package queue

import (
    "testing"
    "github.com/KTAtkinson/linked/node"
    "github.com/KTAtkinson/linked/llist"
)

func TestEnqueue(t *testing.T) {
    value := 5
    q := new(Q)
    q.Enqueue(value)
    popped, _ := q.linkedList.Pop()

    if popped == nil {
        t.Error("Item was not added to queue.")
    } else if popped, _ := popped.Data(); popped != value {
        t.Error("New node did not have the right value.")
    }
}

func TestDequeue(t *testing.T) {
    value := 5
    head := new(node.Dnode)
    head.SetData(value)
    tail := new(node.Dnode)
    node1 := new(node.Dnode)
    node2 := new(node.Dnode)
    node3 := new(node.Dnode)

    ll := new(llist.List)
    ll.Push(tail)
    ll.Push(node3)
    ll.Push(node2)
    ll.Push(node1)
    ll.Push(head)

    q := &Q{linkedList: ll}

    dequeued, _ := q.Dequeue()
    if dequeued != value {
        t.Errorf("Expected dequeue to return %#v, return %#v.", value, dequeued)
    }
}

