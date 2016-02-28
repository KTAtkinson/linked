/*
Linked has basic interfaces for interacting with linked lists.
*/
package linked

type Lister interface {
	Append(Noder, error)
	Push(Noder, error)
	Pop() Noder
}

type Noder interface {
	Data() interface{}
	Next() (Noder, bool, error)
	Prev() (Noder, bool, error)

	SetData(interface{})
	SetNext(Noder, error)
	SetPrev(Noder, error)
}
