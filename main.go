package main

import "fmt"

//problem-1
// Input:  []int{3, 1, 4, 1, 5, 9, 2, 6}
// Output: evens = [4 2 6], odds = [3 1 1 5 9]

// problem-2
// Input:  []string{"error", "info", "error", "debug", "error", "info"}
// Output: map[debug:1 error:3 info:2]

//problem-3

type task struct {
	id   int
	name string
	done bool
}

type queue struct {
	tasks []task
}

func (q *queue) Add(name string) {
	newTask := task{
		id:   len(q.tasks) + 1,
		name: name,
	}
	q.tasks = append(q.tasks, newTask)
}

func (q *queue) Complete(id int) {
	for i := range q.tasks {
		if q.tasks[i].id == id {
			q.tasks[i].done = true
		}
	}
}

func (q *queue) Pending() []task {
	pending := make([]task, 0, len(q.tasks))
	for _, v := range q.tasks {
		if !v.done {
			pending = append(pending, v)
		}
	}
	return pending
}

func main() {
	queue := queue{}
	queue.Add("add migration")
	queue.Add("add new task")
	queue.Add("create")
	queue.Complete(1)
	result := queue.Pending()
	fmt.Println(result)
}

func problem2(n []string) map[string]int {

	r := make(map[string]int)

	for _, v := range n {
		r[v]++
	}
	return r
}

func problem1(n []int) ([]int, []int) {
	even := make([]int, 0, len(n))
	odd := make([]int, 0, len(n))

	for _, v := range n {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return even, odd
}
