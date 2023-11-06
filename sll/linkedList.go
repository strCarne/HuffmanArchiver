package sll

type node struct {
	val  byte
	next *node
}

type List struct {
	head   *node
	tail   *node
	length int
}

func (l *List) Append(value byte) {
	second := &node{val: value}

	defer l.incLen()

	if l.length == 0 {
		l.head = second
		l.tail = second
		return
	}

	l.tail.next = second
	l.tail = l.tail.next
}

func (l *List) ToArr() []byte {
	// arr := []byte{}
	arr := make([]byte, l.length)
	ptr := l.head
	i := 0
	for ptr != nil {
		arr[i] = ptr.val
		ptr = ptr.next
		i++
	}
	return arr
}

func (l *List) incLen() {
	l.length++
}
