package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

// 创建一个新的缓存实例，该实例包含一个队列和一个哈希表，用于存储缓存数据。
// 队列用于维护缓存数据的顺序，哈希表用于快速查找缓存数据。

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

//  函数是用于创建一个新的队列实例的函数。该队列实例包含一个头节点和一个尾节点，头节点的右侧指向尾节点，尾节点的左侧指向头节点。
// 这样的设计可以使得队列的插入和删除操作都可以在常数时间内完成

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail, Length: 0}
}

type Hash map[string]*Node

// Check() 函数用于检查缓存中是否存在指定的字符串。
// 如果存在，则将该字符串对应的节点从队列中删除，并将其插入到队列的头部；
// 如果不存在，则创建一个新的节点，并将其插入到队列的头部。

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}

	}
	c.Add(node)
	c.Hash[str] = node
}

// Remove函数接受一个节点 n 作为参数，
// 首先获取该节点的前驱节点 left 和后继节点 right，然后将 left 的右指针指向 right，将 right 的左指针指向 left，
// 从而将 n 从队列中删除。同时，该函数还会将哈希表中 n 对应的键值对删除，并将队列的长度减一。
// 最后，该函数返回被删除的节点 n。由于队列的实现方式是双向链表，因此在删除节点时只需要修改相邻节点的指针即可，时间复杂度为 O(1)。

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Length -= 1
	delete(c.Hash, n.Val)
	return n
}

// Add函数接受一个节点 n 作为参数，
// 首先获取队列头部的后继节点 tmp，
// 然后将队列头部的右指针指向 n，将 n 的左指针指向队列头部，
// 将 n 的右指针指向 tmp，将 tmp 的左指针指向 n，从而将 n 插入到队列的头部。
// 同时，该函数还会将队列的长度加一，如果队列的长度超过了预设的大小 SIZE，则会调用 Remove() 函数删除队列尾部的节点。
// 由于队列的实现方式是双向链表，因此在插入节点时只需要修改相邻节点的指针即可，时间复杂度为 O(1)。

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

// 该函数接受一个队列 q 作为参数，
// 首先获取队列头部的后继节点 node，然后遍历队列中的所有节点，打印节点的值。
// 最后，该函数会打印队列的长度，并在队列的两端添加方括号 []。
// 由于队列的实现方式是双向链表，因此在遍历队列时只需要沿着节点的右指针移动即可，时间复杂度为 O(n)。

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("start cache")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "dragon fruit", "tree", "potato", "tomato", "tree", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
