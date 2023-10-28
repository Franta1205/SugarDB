package utils

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func NewNode[T comparable](value T, next *Node[T]) *Node[T] {
	return &Node[T]{
		value: value,
		next:  next,
	}
}

func (n *Node[T]) GetValue() T {
	return n.value
}

type LinkedListOptions struct {
	Ring bool
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	ring bool
}

func NewLinkedList[T comparable](options *LinkedListOptions) *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		ring: options.Ring,
	}
}

func (l *LinkedList[T]) Add(value T) {
	if l.head == nil {
		l.head = NewNode(value, nil)
		return
	}

	if l.tail == nil && !l.ring {
		l.tail = NewNode(value, nil)
		l.head.next = l.tail
		return
	}

	if l.tail == nil && l.ring {
		l.tail = NewNode(value, l.head)
		l.head.next = l.tail
		return
	}

	if !l.ring {
		n := NewNode(value, nil)
		l.tail.next = n
		l.tail = n
		return
	}

	n := NewNode(value, l.head)
	l.tail.next = n
	l.tail = n
}

func (l *LinkedList[T]) Remove(value T) {
	if !l.Contains(value) {
		return
	}

	n := l.head

	if n == nil || (n.next == nil && n.value != value) {
		return
	} else if n.next == nil && n.value == value {
		l.head = nil
		return
	}

	if !l.ring {
		if l.head.value == value {
			l.head = l.head.next
			return
		}
		for {
			if n.next.value == value {
				n.next = n.next.next
				if n.next == nil {
					l.tail = n
				}
				return
			}
			n = n.next
		}
	}

	// Linked list is ring
	if l.head.value == value {
		l.head = l.head.next
		l.tail.next = l.head
		if l.head == l.tail {
			l.tail = nil
		}
		return
	}

	if l.head.next.value == value {
		if l.head.next != l.tail {
			l.head.next = l.head.next.next
		} else {
			l.head.next = nil
		}
		return
	}

	n = n.next
	for {
		if n == l.head {
			return
		}
		if n.next.value == value {
			if n.next == l.tail {
				l.tail = n
				l.tail.next = l.head
				return
			}
			n.next = n.next.next
		}
		n = n.next
	}
}

func (l *LinkedList[T]) Contains(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head == nil {
		return false
	}

	if l.head.value == value {
		return true
	}

	n := l.head

	if n == nil || (n.next == nil && n.value != value) {
		return false
	}

	n = n.next
	for {
		if n == nil || n == l.head {
			return false
		}
		if n.value == value {
			return true
		}
		n = n.next
	}
}

func (l *LinkedList[T]) Iter() *chan *Node[T] {
	c := make(chan *Node[T])
	go func() {
		n := l.head
		for {
			if n == nil {
				break
			}
			c <- n
			if n == l.head && n.next == nil && l.ring {
				continue
			}
			n = n.next
		}
		close(c)
	}()
	return &c
}

func (l *LinkedList[T]) Print() {
	fmt.Println("HEAD: ", l.head)
	fmt.Println("TAIL: ", l.tail)

	n := l.head

	fmt.Println(n)

	if n == nil {
		return
	}
	n = n.next

	for {
		if n == nil || n == l.head {
			return
		}
		fmt.Println(n)
		n = n.next
	}
}
