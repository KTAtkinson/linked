package node

import "testing"

func TestNodeData(t *testing.T) {
	cases := []struct {
		value interface{}
	}{
		{5},
		{"this"},
	}
	for _, c := range cases {
		node := &Node{data: c.value}
		result, _ := node.Data()
		if result != c.value {
			t.Errorf("Expected data to be %q, got %q", c.value, result)
		}
	}
}

func TestNodeSetData(t *testing.T) {
	cases := []struct {
		initData interface{}
		newData  interface{}
	}{
		{"current", "new"},
		{2, 5},
	}
	for _, c := range cases {
		node := &Node{data: c.initData}
		node.SetData(c.newData)
		if node.data != c.newData {
			t.Errorf("Expected data to be reset to %q, found %q.", c.newData, node.data)
		}
	}
}

func TestNodeNext(t *testing.T) {
	nextNode := new(Node)
	node := &Node{next: nextNode}
	retrievedNode, _, _ := node.Next()
	if retrievedNode != nextNode { 
		t.Errorf("Expected next node to be returned, got %q", retrievedNode)
	}
}

func TestNodeSetNext(t *testing.T) {
	nextNode := new(Node)
	node := new(Node)
	node.SetNext(nextNode)
	if node.next != nextNode {
		t.Errorf("Expected new node SetNext was called, found %q.", node.next)
	}
}
