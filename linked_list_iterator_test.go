package collections

import "testing"

func TestLinkedListIterator_Next(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)

	itr := NewLinkedListIterator[int](*l)

	if itr.Next() != 1 {
		t.Errorf("iterator next element must be 1")
		t.FailNow()
	}

	if itr.Next() != 2 {
		t.Errorf("iterator next element must be 2")
		t.FailNow()
	}
}

func TestLinkedListIterator_HasNext(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)

	itr := NewLinkedListIterator[int](*l)

	if itr.HasNext() != true {
		t.Errorf("iterator must have next element")
		t.FailNow()
	}

	itr.Next()

	if itr.HasNext() != true {
		t.Errorf("iterator must have next element")
		t.FailNow()
	}

	itr.Next()

	if itr.HasNext() != false {
		t.Errorf("iterator must not have next element")
		t.FailNow()
	}
}
