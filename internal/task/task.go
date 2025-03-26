package task

type Task struct {
	ID          int
	Name        string
	Schedule    int
	Description string
	TaskFunc    func()
	DoneChan    chan bool
}

func (t *Task) Execute() {
	t.TaskFunc()
}
