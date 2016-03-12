/*
Node contains node classes for linked lists.
*/
package node

import ( 
    "github.com/KTAtkinson/linked"
    "errors"
)

/*
Snode is a singly linked node.

A singly linked node keeps track of the node in front of itself.
*/
type Node struct {
	data interface{}
	next linked.Noder
}

/*
Value returns the struct representing the value of the node.
*/
func (n *Node) Data() (i interface{}, err error) {
	return n.data, err
}

/*
SetData sets the value of the node to the given struct.
*/
func (n *Node) SetData(i interface{}) (err error) {
	n.data = i
	return err
}

/*
Next returns to next node in the list, and a boolean indicating if this is the
last thing in the list.
*/
func (n *Node) Next() (next linked.Noder, is_last bool, err error) {
	next = n.next
	if next == nil {
		is_last = true
		return next, is_last, err
	}
	return next, is_last, err
}

/*
SetNext sets the next node in the sequence to the given Dnode.
*/
func (n *Node) SetNext(next linked.Noder) (err error) {
	n.next = next
	return err
}

/*
Returns an error because finding the previous node is not supported for on a singly
linked node.
*/
func (n *Node) Prev() (pn linked.Noder, is_first bool, err error) {
  return pn, is_first, errors.New("Can't find previous node.")     
}

/*
SetPrev sets the previous element to the given node interface.
*/
func (n *Node) SetPrev(prev linked.Noder) (err error) {
    return errors.New("Can't set previous node.")
}
