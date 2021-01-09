package LinkStack

type Node struct {
	data interface{}
	next *Node
}

type Stack interface {
	IsEmpty() bool            // 判空
	Push(element interface{}) // 入栈
	Pop() interface{}         // 出栈
	Top() interface{}         //查看栈顶元素
	Size() int                // 栈长
}

type LinkStack struct {
	top    *Node
	length int
}

func NewLinkStack() *LinkStack {
	stk := new(LinkStack)
	stk.top = new(Node)
	stk.length = 0
	return stk
}

func (stk *LinkStack) Size() int {
	return stk.length
}

func (stk *LinkStack) IsEmpty() bool {
	return stk.length == 0
}

func (stk *LinkStack) Push(element interface{}) {
	newNode := new(Node)
	newNode.data = element
	newNode.next = stk.top.next
	stk.top.next = newNode
	stk.length++
}

func (stk *LinkStack) Top() interface{} {
	if stk.IsEmpty() {
		return nil
	}
	return stk.top.data
}

func (stk *LinkStack) Pop() interface{} {
	if stk.IsEmpty() {
		return nil
	}
	res := stk.top.next.data
	stk.top.next = stk.top.next.next
	stk.length--
	return res
}
