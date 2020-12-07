package tree

import "testing"

func getTestTree() *Tree {
	root := &Node{
		ID: 1,
	}
	node2 := &Node{
		ID: 2,
	}
	node3 := &Node{
		ID: 3,
	}
	node4 := &Node{
		ID: 4,
	}
	node5 := &Node{
		ID: 5,
	}
	node6 := &Node{
		ID: 6,
	}
	node7 := &Node{
		ID: 7,
	}
	node8 := &Node{
		ID: 8,
	}
	node9 := &Node{
		ID: 9,
	}
	node10 := &Node{
		ID: 10,
	}
	node11 := &Node{
		ID: 11,
	}
	node12 := &Node{
		ID: 12,
	}
	node13 := &Node{
		ID: 13,
	}
	node14 := &Node{
		ID: 14,
	}
	node15 := &Node{
		ID: 15,
	}
	node16 := &Node{
		ID: 16,
	}
	node17 := &Node{
		ID: 17,
	}
	node18 := &Node{
		ID: 18,
	}

	tree, _ := NewTree(root)
	tree.Add(root.ID, node2)
	tree.Add(root.ID, node3)
	tree.Add(root.ID, node4)
	tree.Add(root.ID, node5)
	tree.Add(node2.ID, node6)
	tree.Add(node2.ID, node7)
	tree.Add(node3.ID, node8)
	tree.Add(node3.ID, node9)
	tree.Add(node3.ID, node10)
	tree.Add(node4.ID, node11)
	tree.Add(node4.ID, node12)
	tree.Add(node4.ID, node13)
	tree.Add(node6.ID, node14)
	tree.Add(node6.ID, node15)
	tree.Add(node6.ID, node16)
	tree.Add(node7.ID, node17)
	tree.Add(node15.ID, node18)
	return tree
}

func TestTree_NewTreeWithNilRoot(t *testing.T) {
	tree, err := NewTree(nil)
	if tree != nil || err == nil {
		t.Errorf("Tree must be null got %v error must be not null got null", tree)
	}
}

func TestTree_NewTree(t *testing.T) {
	root := &Node{
		ID:           1,
		Data:         nil,
		FirstChild:   nil,
		FirstSibling: nil,
	}
	tree, err := NewTree(root)
	if err != nil {
		t.Errorf("Error must be null got %v", err)
	}
	if tree.Root.ID != root.ID {
		t.Errorf("The root node id must be %v got %v", root.ID, tree.Root.ID)
	}
}

func TestTree_BuildTreeAndPrint(t *testing.T) {
	tree := getTestTree()
	expectedPrint := "1(2(6(14()15(18())16())7(17()))3(8()9()10())4(11()12()13())5())"
	print := tree.Print()
	if print != expectedPrint {
		t.Errorf("The tree must be \n%s\ngot\n%s", expectedPrint, print)
	}
}

func TestTree_Count(t *testing.T) {
	tree := getTestTree()
	if count := tree.Count(); count != 18 {
		t.Errorf("Count must be 18 got %v", count)
	}
}

func TestTree_DeleteRoot(t *testing.T) {
	tree := getTestTree()
	if ok, err := tree.Delete(1); ok || err == nil {
		t.Errorf("Delete must be not ok got %v, error must be not null got null", ok)
	}
}

func TestTree_DeleteInexistentNode(t *testing.T) {
	tree := getTestTree()
	if ok, err := tree.Delete(20); ok || err != nil {
		t.Errorf("Delete must be not ok got %v, error must be null got %v", ok, err)
	}
}

func TestTree_Delete(t *testing.T) {
	tree := getTestTree()
	if ok, err := tree.Delete(12); !ok || err != nil {
		t.Errorf("Expected delete ok but got %v", err)
	}
	if ok, err := tree.Delete(3); !ok || err != nil {
		t.Errorf("Expected delete ok but got %v", err)
	}
	if ok, err := tree.Delete(14); !ok || err != nil {
		t.Errorf("Expected delete ok but got %v", err)
	}
	if ok, err := tree.Delete(6); !ok || err != nil {
		t.Errorf("Expected delete ok but got %v", err)
	}

	if count := tree.Count(); count != 14 {
		t.Errorf("Count must be 14 got %v", count)
	}
	expectedPrint := "1(2(7(17())15(18())16())4(11()13())5()8()9()10())"
	if print := tree.Print(); print != expectedPrint {
		t.Errorf("Print must be\n%v\ngot\n%v", expectedPrint, print)
	}
}
