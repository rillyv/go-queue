package goqueue

import "testing"

func TestQueueAndDequeue(t *testing.T) {
	q := NewQueue()
	id := q.Enqueue([]byte("payload"))
	job := q.Dequeue()

	if job == nil {
		t.Fatal("excepted job, got nil")
	}

	if job.ID != id {
		t.Fatalf("job ID mismatch: got %v, want %v", job.ID, id)
	}
}
