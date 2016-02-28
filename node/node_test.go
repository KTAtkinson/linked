package node

import "testing"

func TestData(t *testing.T) {
	cases := []struct {
		value interface{}
	}{
		{5},
		{"this"},
	}
	for _, c := range cases {
		node := &Dnode{data: c.value}
		result, _ := node.Data()
		if result != c.value {
			t.Errorf("Expected data to be %q, got %q", c.value, result)
		}
	}
}

func TestSetData(t *testing.T) {
	cases := []struct {
		initData interface{}
		newData  interface{}
	}{
		{"current", "new"},
		{2, 5},
	}
	for _, c := range cases {
		node := &Dnode{data: c.initData}
		node.SetData(c.newData)
		if node.data != c.newData {
			t.Errorf("Expected data to be reset to %q, found %q.", c.newData, node.data)
		}
	}
}

func TestNext(t *testing.T) {
	nextNode := new(Dnode)
	node := &Dnode{next: nextNode}
	retrievedNode, _, _ := node.Next()
	if retrievedNode != nextNode {
		t.Errorf("Expected next node to be returned, got %q", retrievedNode)
	}
}

func TestSetNext(t *testing.T) {
	nextNode := new(Dnode)
	node := new(Dnode)
	node.SetNext(nextNode)
	if node.next != nextNode {
		t.Errorf("Expected new node SetNext was called, found %q.", node.next)
	}
}
