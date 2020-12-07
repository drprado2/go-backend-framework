package binarytree

import (
	"errors"
	"fmt"
)

type BinarySearchTree struct {
	Root *Node
}

type Node struct {
	Left  *Node
	Right *Node
	ID    int
	Data  interface{}
}

func NewBinarySearchTree(root *Node) (*BinarySearchTree, error) {
	if root == nil {
		return nil, errors.New("the root node must not be null")
	}
	return &BinarySearchTree{
		Root: root,
	}, nil
}

func (t *BinarySearchTree) Add(node *Node) error {
	currentNode := t.Root
	for currentNode != nil {
		if currentNode.ID < node.ID {
			if currentNode.Right == nil {
				currentNode.Right = node
				break
			} else {
				currentNode = currentNode.Right
			}
		} else if currentNode.ID > node.ID {
			if currentNode.Left == nil {
				currentNode.Left = node
				break
			} else {
				currentNode = currentNode.Left
			}
		} else {
			return fmt.Errorf("The element %v already exists in the tree", node.ID)
		}
	}
	return nil
}

func (t *BinarySearchTree) FindNodeAndFather(nodeId int) (*Node, *Node) {
	var father *Node
	currentNode := t.Root
	for currentNode != nil {
		if currentNode.ID < nodeId {
			father = currentNode
			currentNode = currentNode.Right
			continue
		}
		if currentNode.ID > nodeId {
			father = currentNode
			currentNode = currentNode.Left
			continue
		}
		return currentNode, father
	}
	return nil, nil
}

func (t *BinarySearchTree) Print() string {
	result := print(t.Root)
	fmt.Println(result)
	return result
}

func print(currentNode *Node) string {
	if currentNode == nil {
		return ""
	}
	currentPrint := fmt.Sprintf("%v(", currentNode.ID)
	currentPrint += print(currentNode.Left)
	currentPrint += print(currentNode.Right)
	currentPrint += ")"
	return currentPrint
}

func (t *BinarySearchTree) Delete(nodeId int) (bool, error) {
	if t.Root.ID == nodeId {
		return false, errors.New("you can`t delete the root node")
	}
	node, father := t.FindNodeAndFather(nodeId)
	if node == nil {
		return false, nil
	}

	var newNode *Node = nil

	if node.Left != nil && node.Right != nil {
		if node.Left.Right == nil {
			newNode = node.Left
		} else {
			fatherNode := node.Left
			currentNode := node.Left.Right
			for currentNode.Right != nil {
				fatherNode = currentNode
				currentNode = currentNode.Right
			}
			fatherNode.Right = currentNode.Left
			newNode = currentNode
			newNode.Left = node.Left
			newNode.Right = node.Right
		}
	} else if node.Right != nil {
		newNode = node.Right
	} else if node.Left != nil {
		newNode = node.Left
	}

	if father.ID > node.ID {
		father.Left = newNode
	} else {
		father.Right = newNode
	}
	return true, nil
}

func (t *BinarySearchTree) Count() int {
	return count(t.Root)
}

func count(currentNode *Node) int {
	if currentNode == nil {
		return 0
	}
	result := 1
	result += count(currentNode.Left)
	result += count(currentNode.Right)
	return result
}
