package datastruc

import "fmt"

type Linked_list struct {
	Data int
	Link *Linked_list
}


func (l *Linked_list) GetData() (int, error) {
	if (l != nil) {
		return l.Data, nil
	}
	return -1, fmt.Errorf("Linked List is empty")
}

func (l *Linked_list) GetNext() (*Linked_list, error) {
	if (l != nil) {
		return l.Link, nil
	}
	return nil, fmt.Errorf("Linked List is empty")
	
}

func (l *Linked_list) AddToTail(data int) {
    current := l
    for current.Link != nil {
        current = current.Link
    }
    current.Link = &Linked_list{
        Data: data,
        Link: nil, 
    }
}


func (l *Linked_list) AddToHead(data int) {
	l = &Linked_list {
		Data : data,
		Link : l,
	}
}


func (l *Linked_list) Print() {
	for l != nil {
		fmt.Printf("%d -> ", l.Data)
		l = l.Link 
	}
	fmt.Println("nil")
}
