package tree

import "testing"

func getTestTree() *Tree {
	root := &Node{
		Id: 1,
	}
	node2 := &Node{
		Id: 2,
	}
	node3 := &Node{
		Id: 3,
	}
	node4 := &Node{
		Id: 4,
	}
	node5 := &Node{
		Id: 5,
	}
	node6 := &Node{
		Id: 6,
	}
	node7 := &Node{
		Id: 7,
	}
	node8 := &Node{
		Id: 8,
	}
	node9 := &Node{
		Id: 9,
	}
	node10 := &Node{
		Id: 10,
	}
	node11 := &Node{
		Id: 11,
	}
	node12 := &Node{
		Id: 12,
	}
	node13 := &Node{
		Id: 13,
	}
	node14 := &Node{
		Id: 14,
	}
	node15 := &Node{
		Id: 15,
	}
	node16 := &Node{
		Id: 16,
	}
	node17 := &Node{
		Id: 17,
	}
	node18 := &Node{
		Id: 18,
	}

	tree, _ := NewTree(root)
	tree.Add(root.Id, node2)
	tree.Add(root.Id, node3)
	tree.Add(root.Id, node4)
	tree.Add(root.Id, node5)
	tree.Add(node2.Id, node6)
	tree.Add(node2.Id, node7)
	tree.Add(node3.Id, node8)
	tree.Add(node3.Id, node9)
	tree.Add(node3.Id, node10)
	tree.Add(node4.Id, node11)
	tree.Add(node4.Id, node12)
	tree.Add(node4.Id, node13)
	tree.Add(node6.Id, node14)
	tree.Add(node6.Id, node15)
	tree.Add(node6.Id, node16)
	tree.Add(node7.Id, node17)
	tree.Add(node15.Id, node18)
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
		Id:           1,
		Data:         nil,
		FirstChild:   nil,
		FirstSibling: nil,
	}
	tree, err := NewTree(root)
	if err != nil {
		t.Errorf("Error must be null got %v", err)
	}
	if tree.Root.Id != root.Id {
		t.Errorf("The root node id must be %v got %v", root.Id, tree.Root.Id)
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
