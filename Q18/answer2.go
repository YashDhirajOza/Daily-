package main
import "fmt"
type Employee struct {
	id            string
	name          string
	monthlySalary float64
}
func NewEmployee(id string, name string, monthlySalary float64) *Employee {
	return &Employee{id: id, name: name, monthlySalary: monthlySalary} 
}
func (e *Employee) SetId(id string) {	e.id = id }
func (e *Employee) SetName(name string) { e.name = name }
func (e *Employee) SetMonthlySalary(salary float64) {	e.monthlySalary = salary}
func (e *Employee) GetId() string {return e.id}
func (e *Employee) GetName() string {return e.name}
func (e *Employee) GetMonthlySalary() float64 {	return e.monthlySalary}

func (e *Employee) DisplayEmployee() {
	fmt.Println("Employee ID:", e.id)
	fmt.Println("Name:", e.name)
	fmt.Println("Monthly Salary:", e.monthlySalary)
	fmt.Println("Yearly Salary:", e.monthlySalary*12)
	fmt.Println("--------------------------------")
}

func main() {
	emp1 := NewEmployee("E101", "Alice", 5000.0)
	emp2 := NewEmployee("E102", "Bob", 6000.0)
	emp1.DisplayEmployee()
	emp2.DisplayEmployee()
}
