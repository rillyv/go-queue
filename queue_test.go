package goqueue

import "testing"

func TestQueueAndDequeue(t *testing.T) {
	q := NewQueue()
	id := q.Enqueue([]byte("payload"))
	job := q.Dequeue()

	if job == nil || job.ID != id {
		t.Fatal("nah fam, this shit ain't good tho")
	}
}
