package domain

type Queue struct {
	elements []string
}

func NewQueue(elements []string) Queue {
	return Queue{
		elements: elements,
	}
}

func (queue *Queue) Enqueue(element string) {
	queue.elements = append(queue.elements, element)
}

func (queue *Queue) Dequeue() string {
	element := queue.elements[0]
	queue.elements = queue.elements[1:]
	return element
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.elements) <= 0
}
