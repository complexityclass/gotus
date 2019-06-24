package doubly_linked_list

// DListItem is an element of doubly linked list.
type DListItem struct {
	next  *DListItem
	prev  *DListItem
	list  *DList
	value interface{}
}

// Value returns a  value of element.
func (item *DListItem) Value() interface{} {
	return item.value
}

// Next returns the next list element or nil
func (item *DListItem) Next() *DListItem {
	if n := item.next; item.list != nil && n != &item.list.root {
		return n
	}
	return nil
}

// Prev returns the previous list element or nil.
func (item *DListItem) Prev() *DListItem {
	if p := item.prev; item.list != nil && p != &item.list.root {
		return p
	}
	return nil
}

// Remove removes item from list.
func (item *DListItem) Remove() {
	if item.list == nil {
		return
	}

	item.prev.next = item.next
	item.next.prev = item.prev
	item.next = nil
	item.prev = nil
	item.list.len -= 1
	item.list = nil
}

// DList represents a doubly linked list.
type DList struct {
	root DListItem
	len  int
}

// Init is a ctor
func (list *DList) Init() *DList {
	list.root.next = &list.root
	list.root.prev = &list.root
	list.len = 0
	return list
}

func (list *DList) lazyInit() {
	if list.root.next == nil {
		list.Init()
	}
}

// Len returns length of the list
func (list *DList) Len() int {
	return list.len
}

// First returns first element of the list
func (list *DList) First() *DListItem {
	if list.len == 0 {
		return nil
	}
	return list.root.next
}

// Last returns last element of the list
func (list *DList) Last() *DListItem {
	if list.len == 0 {
		return nil
	}
	return list.root.prev
}

// PushFront pushes element to front
func (list *DList) PushFront(value interface{}) {
	list.lazyInit()
	_ = list.insert(value, &list.root)
}

// PushBack pushes element to back
func (list *DList) PushBack(value interface{}) {
	list.lazyInit()
	_ = list.insert(value, list.root.prev)
}

func (list *DList) insert(value interface{}, at *DListItem) *DListItem {
	var elem = DListItem{value: value}
	var next = at.next
	at.next = &elem
	elem.prev = at
	elem.next = next
	next.prev = &elem
	elem.list = list
	list.len += 1
	return &elem
}
