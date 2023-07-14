package ds

type Link struct {
	head *LinkNode
}

func (l *Link) Insert(data int) {
	if l.head == nil {
		l.head = &LinkNode{Data: data}
		return
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &LinkNode{
		Data: data,
	}
}

func (l *Link) Traverse() []int {
	rsp := []int{}
	if l.head == nil {
		return rsp
	}
	cur := l.head
	for cur != nil {
		rsp = append(rsp, cur.Data)
		cur = cur.Next
	}
	return rsp
}

type LinkNode struct {
	Data int
	Next *LinkNode
}
