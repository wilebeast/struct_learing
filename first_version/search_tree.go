package first_version

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func (t *TreeNode) Search(value int) *TreeNode {
	if t == nil {
		return t
	}

	if t.value == value {
		return t
	}

	if t.left != nil && value < t.left.value {
		return t.left.Search(value)
	}
	if t.right != nil && value > t.right.value {
		return t.right.Search(value)
	}
	return nil
}

func (t *TreeNode) Insert(value int) {
	if t == nil {
		panic("")
	}

	if t.value == value {
		return
	}

	if t.value > value {
		if t.left != nil {
			t.left.Insert(value)
		} else {
			t.left = &TreeNode{
				left:  nil,
				right: nil,
				value: value,
			}
		}
	} else {
		if t.right != nil {
			t.right.Insert(value)
		} else {
			t.right = &TreeNode{
				left:  nil,
				right: nil,
				value: value,
			}
		}
	}
}

func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		panic("")
	}

	if t.value > value {
		t.left = t.left.Delete(value)
	} else if t.value < value {
		t.right = t.right.Delete(value)
	}

	if t.left == nil {
		return t.right
	}
	if t.right == nil {
		return t.left
	}
	minNode := t.right
	for {
		if minNode.left != nil {
			minNode = minNode.left
		}
	}

	t.value = minNode.value
	t.right = t.right.Delete(t.value)
	return t
}
