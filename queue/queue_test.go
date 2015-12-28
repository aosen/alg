package queue

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	q := NewQueue(2)

	if !q.Empty() ||
		q.len != 0 ||
		q.Len() != 0 {
		t.Error()
	}

	var er error
	er = q.EnQueue(1)
	if er != nil {
		log.Println("1", er.Error())
	}
	er = q.EnQueue(2)
	if er != nil {
		log.Println("2", er.Error())
	}
	er = q.EnQueue(3)
	if er != nil {
		log.Println("3", er.Error())
	}

	a, e := q.OutQueue()
	if e != nil {
		log.Println(e.Error())
	} else {
		log.Println("out:", a)
	}

	b, err := q.Peek()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("peek:", b)
	}

	log.Println("len:", q.Len())
}
