/*
Implementation of a queue written with linked.llist.
*/
package queue

import (
    "github.com/KTAtkinson/linked/node"
    "github.com/KTAtkinson/linked/llist"
)

type Q struct {
    linkedList *llist.List
}

// Enqueue adds a value to the queue.
func (q *Q) Enqueue(v interface{}) (err error) {
    newNode := &node.Dnode{}
    newNode.SetData(v)
    if q.linkedList == nil {
        q.linkedList = new(llist.List)
    }
    q.linkedList.Append(newNode)
    return err
}

// Dequeue removes the next item from the queue and returns it.
func (q *Q) Dequeue() (v interface{}, err error) {
    popped, _:= q.linkedList.Pop()
    value, _ := popped.Data()
    return value, err
}
