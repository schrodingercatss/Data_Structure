package LinkQueue

type Queue interface {
	Front() interface{}       // 返回队首元素
	Push(element interface{}) // 入队
	Pop() interface{}         // 出队
	Size() int                // 返回队列长度
	IsEmpty() bool            // 判空
}

type Node struct {
	data interface{}
	next *Node
}

type LinkQueue struct {
	length      int
	front, rear *Node
}

func NewLinkQueue() *LinkQueue {
	lqueue := new(LinkQueue)
	lqueue.length = 0
	newNode := new(Node)
	lqueue.front, lqueue.rear = newNode, newNode
	return lqueue
}

func (lq *LinkQueue) Size() int {
	return lq.length
}

func (lq *LinkQueue) IsEmpty() bool {
	return lq.length == 0
}

func (lq *LinkQueue) Push(element interface{}) {
	newNode := new(Node)
	newNode.data = element
	lq.rear.next = newNode
	lq.rear = newNode
	lq.length++
}

func (lq *LinkQueue) Pop() interface{} {
	if lq.IsEmpty() {
		return nil
	}
	res := lq.front.next
	lq.front.next = res.next
	if res == lq.rear {
		lq.rear = lq.front
	}
	lq.length--
	return res.data
}

func (lq *LinkQueue) Front() interface{} {
	if lq.IsEmpty() {
		return nil
	}
	return lq.front.next.data
}
