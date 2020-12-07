package tree

import (
	"errors"
	"fmt"
)

type Tree struct {
	Root *Node
}

type Node struct {
	ID           int
	Data         interface{}
	FirstChild   *Node
	FirstSibling *Node
}

func NewTree(root *Node) (*Tree, error) {
	if root == nil {
		return nil, errors.New("Root node must be not null")
	}
	return &Tree{
		Root: root,
	}, nil
}

func (t *Tree) Add(fatherId int, newNode *Node) bool {
	father, _ := t.FindNodeAndFather(fatherId)
	if father == nil {
		return false
	}
	if father.FirstChild == nil {
		father.FirstChild = newNode
		return true
	}

	currentNode := father.FirstChild
	for currentNode.FirstSibling != nil {
		currentNode = currentNode.FirstSibling
	}
	currentNode.FirstSibling = newNode
	return true
}

func (t *Tree) Delete(nodeId int) (bool, error) {
	if t.Root.ID == nodeId {
		return false, errors.New("can`t delete the root node")
	}
	node, fatherNode := t.FindNodeAndFather(nodeId)
	if node == nil || fatherNode == nil {
		return false, nil
	}

	var currentSibling *Node
	if node.ID == fatherNode.FirstChild.ID {
		fatherNode.FirstChild = node.FirstSibling
		currentSibling = fatherNode.FirstChild
	} else {
		lastSibling := fatherNode.FirstChild
		currentSibling = lastSibling.FirstSibling
		for currentSibling.ID != node.ID {
			lastSibling = currentSibling
			currentSibling = currentSibling.FirstSibling
		}
		lastSibling.FirstSibling = currentSibling.FirstSibling
	}

	for currentSibling.FirstSibling != nil {
		currentSibling = currentSibling.FirstSibling
	}

	currentSibling.FirstSibling = node.FirstChild
	return true, nil
}

func (t *Tree) FindNodeAndFather(nodeId int) (*Node, *Node) {
	return findNodeAndFather(nil, t.Root, nodeId)
}

func findNodeAndFather(fatherNode *Node, currentNode *Node, idToFind int) (*Node, *Node) {
	if currentNode == nil {
		return nil, fatherNode
	}
	if currentNode.ID == idToFind {
		return currentNode, fatherNode
	}
	if node, father := findNodeAndFather(currentNode, currentNode.FirstChild, idToFind); node != nil {
		return node, father
	}
	return findNodeAndFather(fatherNode, currentNode.FirstSibling, idToFind)
}

func (t *Tree) Print() string {
	printResult := print(t.Root, "")
	fmt.Println(printResult)
	return printResult
}

func print(node *Node, currentPrint string) string {
	if node == nil {
		return currentPrint
	}
	result := fmt.Sprintf("%s%v(", currentPrint, node.ID)
	result = print(node.FirstChild, result)
	result = result + ")"
	result = print(node.FirstSibling, result)
	return result
}

func (t *Tree) Count() int {
	return count(t.Root)
}

func count(node *Node) int {
	if node == nil {
		return 0
	}
	currentCount := 1 + count(node.FirstChild)

	currentCount += count(node.FirstSibling)
	//
	//currentSibling := node.FirstSibling
	//for currentSibling != nil {
	//	currentCount += count(currentSibling)
	//	currentSibling = currentSibling.FirstSibling
	//}

	return currentCount
}
