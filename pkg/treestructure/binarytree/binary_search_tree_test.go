package binarytree

import (
	"testing"
)

func buildTestTree() *BinarySearchTree {
	root := &Node{
		ID: 7,
	}
	tree, _ := NewBinarySearchTree(root)
	tree.Add(&Node{ID: 3})
	tree.Add(&Node{ID: 12})
	tree.Add(&Node{ID: 1})
	tree.Add(&Node{ID: 6})
	tree.Add(&Node{ID: 9})
	tree.Add(&Node{ID: 13})
	tree.Add(&Node{ID: 0})
	tree.Add(&Node{ID: 2})
	tree.Add(&Node{ID: 4})
	tree.Add(&Node{ID: 8})
	tree.Add(&Node{ID: 11})
	tree.Add(&Node{ID: 15})
	tree.Add(&Node{ID: 5})
	tree.Add(&Node{ID: 10})
	tree.Add(&Node{ID: 14})
	return tree
}

func TestBinarySearchTree_NewWithNullRoot(t *testing.T) {
	tree, err := NewBinarySearchTree(nil)
	if tree != nil || err == nil {
		t.Errorf("Tree must be null got %v, error must be not null got %v", tree, err)
	}
}

func TestBinarySearchTree_New(t *testing.T) {
	root := &Node{
		Left:  nil,
		Right: nil,
		ID:    1,
		Data:  nil,
	}
	tree, err := NewBinarySearchTree(root)
	if tree == nil || tree.Root.ID != root.ID || err != nil {
		t.Errorf("Invalid tree %v", tree)
	}
	if count := tree.Count(); count != 1 {
		t.Errorf("Count must be 1 got %v", count)
	}
}

func TestBinarySearchTree_Count(t *testing.T) {
	tree := buildTestTree()
	if count := tree.Count(); count != 16 {
		t.Errorf("Count must be 16 got %v", count)
	}
}

func TestBinarySearchTree_FindNodeAndFather(t *testing.T) {
	tree := buildTestTree()
	if node, father := tree.FindNodeAndFather(11); node == nil || node.ID != 11 || father == nil || father.ID != 9 {
		t.Errorf("Search result of node 11 is wrong\nNode: %v\nFather: %v", node, father)
	}
	if node, father := tree.FindNodeAndFather(5); node == nil || node.ID != 5 || father == nil || father.ID != 4 {
		t.Errorf("Search result of node 5 is wrong\nNode: %v\nFather: %v", node, father)
	}
	if node, father := tree.FindNodeAndFather(23); node != nil || father != nil {
		t.Errorf("Search result of node 23 is wrong\nNode: %v\nFather: %v", node, father)
	}
	if node, father := tree.FindNodeAndFather(7); node == nil || node.ID != 7 || father != nil {
		t.Errorf("Search result of root node 7 is wrong\nNode: %v\nFather: %v", node, father)
	}
}

func TestBinarySearchTree_Print(t *testing.T) {
	tree := buildTestTree()
	expectedPrint := "7(3(1(0()2())6(4(5())))12(9(8()11(10()))13(15(14()))))"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Invalid print expetected\n%v\ngot\n%v", expectedPrint, print)
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	tree := buildTestTree()

	tree.Delete(8)
	if count := tree.Count(); count != 15 {
		t.Errorf("Count must be 15 got %v", count)
	}
	expectedPrint := "7(3(1(0()2())6(4(5())))12(9(11(10()))13(15(14()))))"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Invalid print expetected\n%v\ngot\n%v", expectedPrint, print)
	}

	tree.Delete(9)
	if count := tree.Count(); count != 14 {
		t.Errorf("Count must be 14 got %v", count)
	}
	expectedPrint = "7(3(1(0()2())6(4(5())))12(11(10())13(15(14()))))"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Invalid print expetected\n%v\ngot\n%v", expectedPrint, print)
	}

	tree.Delete(15)
	if count := tree.Count(); count != 13 {
		t.Errorf("Count must be 13 got %v", count)
	}
	expectedPrint = "7(3(1(0()2())6(4(5())))12(11(10())13(14())))"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Invalid print expetected\n%v\ngot\n%v", expectedPrint, print)
	}

	tree.Delete(3)
	if count := tree.Count(); count != 12 {
		t.Errorf("Count must be 12 got %v", count)
	}
	expectedPrint = "7(2(1(0())6(4(5())))12(11(10())13(14())))"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Invalid print expetected\n%v\ngot\n%v", expectedPrint, print)
	}

	tree = buildTestTree()
	tree.Delete(12)
	if count := tree.Count(); count != 15 {
		t.Errorf("Count must be 15 got %v", count)
	}
	expectedPrint = "7(3(1(0()2())6(4(5())))11(9(8()10())13(15(14()))))"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Invalid print expetected\n%v\ngot\n%v", expectedPrint, print)
	}
}
