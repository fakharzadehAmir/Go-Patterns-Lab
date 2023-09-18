package Composite

// IEmployee is the base interface for all employees.
type IEmployee interface {
	GetSalary() float64
}

// Worker represents individual employees with a specific salary.
type Worker struct {
	salary float64
}

// NewWorker creates a new worker instance with the given salary.
func NewWorker(salary float64) *Worker {
	return &Worker{salary: salary}
}

// GetSalary returns the salary of the worker.
func (w *Worker) GetSalary() float64 {
	return w.salary
}

// Manager represents composite employees (managers) that can have subordinates.
type Manager struct {
	salary       float64
	subordinates []IEmployee
}

// NewManager creates a new manager instance with the given salary.
func NewManager(salary float64) *Manager {
	return &Manager{salary: salary, subordinates: nil}
}

// GetSalary calculates the total salary by summing up the salaries of subordinates and adding the manager's own salary.
func (m *Manager) GetSalary() float64 {
	totalSalary := 0.0
	for _, subordinate := range m.subordinates {
		totalSalary += subordinate.GetSalary()
	}
	return totalSalary + m.salary
}

// AddEmployee adds a subordinate employee to the manager's list of subordinates.
func (m *Manager) AddEmployee(employee IEmployee) {
	m.subordinates = append(m.subordinates, employee)
}
