/*
平衡二叉树又称 AVL 树
* 特点：一个根节点的左右个子树的高度差不超过1

红黑树是一种平衡的二叉查找树.
* 节点是红色或黑色。
* 根节点是黑色。
* 每个叶子节点都是黑色的空节点（NIL节点）。
* 每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)
* 从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点。
*/

package rbtree

const (
	//RED  color
	RED = true
	//BLACK color
	BLACK = false
)

//RedBlackTree tree
type RedBlackTree struct {
	root *Node //no forward declartion needed in go??
}

//Item any type can be compared
type Item interface {
	Less(than Item) bool
	More(than Item) bool
}

//Node type
type Node struct {
	size        uint64
	left, right *Node
	item        Item
	value       string
	color       bool
}

//New return a pointer to the empty strcut
func New() *RedBlackTree {
	return &RedBlackTree{}
}

//Size  return the size of the tree
func (t *RedBlackTree) Size() uint64 {
	return t.root.Size()
}

//Put   insert item into the tree
func (t *RedBlackTree) Put(key Item) {
	t.root = put(t.root, key)
	t.root.color = BLACK
}

//Find  find the itme in the tree
func (t *RedBlackTree) Find(key Item) (Item, bool) {
	return find(t.root, key)
}

//Size  return size of node
func (n *Node) Size() uint64 {
	return n.size
}

//Item   return item of the node
func (n *Node) Item() Item {
	return n.item
}

func size(n *Node) uint64 {
	if n == nil {
		return 0
	}

	return n.size
}

/*
每次插入元素的时候会将 元素 着色为红色。其目的为了快的满足红黑树的4个条件

红黑树结构不会旋转变化情况：。
1）当插入的节点为的父亲为黑色节点。【什么都不用做】
2）被插入的节点是根节点。【直接把此节点涂为黑色】


*/

func put(n *Node, item Item) *Node {
	switch {
	case n == nil:
		return &Node{item: item, size: 1, color: RED}
	case item.Less(n.item):
		n.left = put(n.left, item)
	case n.item.Less(item):
		n.right = put(n.right, item)
	default:
		n.item = item
	}

	/*parent 黑色，uncle红色
	 */
	if !isRed(n.left) && isRed(n.right) {
		n = rotateLeft(n)
	}

	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	/*
		如果parent 和 uncle都是红色
		*将 parent 和 uncle 标记为黑色
		* 将 grand parent (祖父) 标记为红色
	*/
	if isRed(n.left) && isRed(n.right) {
		n.left.color = BLACK
		n.right.color = BLACK
		n.color = RED
	}

	n.size = size(n.left) + size(n.right) + 1
	return n
}

func find(n *Node, item Item) (Item, bool) {
	switch {
	case n == nil:
		return nil, false
	case item.Less(n.item):
		return find(n.left, item)
	case item.More(n.item):
		return find(n.right, item)
	default:
		return n.Item(), true
	}

}

/* n 成为x 的左子树, 下降一级
 */
func rotateLeft(n *Node) *Node {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = RED

	x.size = n.size
	n.size = size(n.left) + size(n.right) + 1

	return x
}

/* n 成为x 的右子树, 下降一级
 */
func rotateRight(n *Node) *Node {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = RED

	x.size = n.size
	n.size = size(n.left) + size(n.right) + 1

	return x
}

func isRed(n *Node) bool {
	if n == nil {
		return false
	}

	return (n.color == RED)
}
