package tree

// create node

type Node struct {
	data  int
	left  *Node
	right *Node
}

// create tree

type BinaryTree struct {
	root *Node
}

// constructor

func New() *BinaryTree {
	return &BinaryTree{}
}

// tree insert

func (b *BinaryTree) Insert(data int) {
	if b.root != nil {
		b.root.insert(data)
	} else {
		b.root = &Node{data: data}
	}
}

// root insert

func (root *Node) insert(data int) {
	if root != nil {
		if data <= root.data {
			if root.left != nil {
				root.left.insert(data)
			} else {
				root.left = &Node{data: data}
			}
		} else {
			if root.right != nil {
				root.right.insert(data)
			} else {
				root.right = &Node{data: data}
			}
		}
	}
	return
}

func (b *BinaryTree) InOrder() []int {
	// left -> root -> right
	out := make([]int, 0)
	fn := func(data int) {
		out = append(out, data)
	}
	inOrderTraverse(b.root, fn)
	return out
}

func inOrderTraverse(root *Node, fn func(out int)) {
	if root != nil {
		inOrderTraverse(root.left, fn)
		fn(root.data)
		inOrderTraverse(root.right, fn)
	}
	return
}

func (b *BinaryTree) PreOrder() []int {
	// root -> left -> right
	out := make([]int, 0)
	fn := func(in int) {
		out = append(out, in)
	}
	preOrderTraverse(b.root, fn)
	return out
}
func preOrderTraverse(root *Node, fn func(int)) {
	if root != nil {
		fn(root.data)
		preOrderTraverse(root.left, fn)
		preOrderTraverse(root.right, fn)
	}
	return
}

func (b *BinaryTree) PostOrder() []int {
	// left -> right -> root
	out := make([]int, 0)
	fn := func(in int) {
		out = append(out, in)
	}
	postOrderTraverse(b.root, fn)
	return out
}

func postOrderTraverse(root *Node, fn func(int)) {
	if root != nil {
		postOrderTraverse(root.left, fn)
		postOrderTraverse(root.right, fn)
		fn(root.data)
	}
	return
}
