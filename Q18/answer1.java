class Employee {
    private String id;
    private String name;
    private double monthlySalary;
    public Employee() {
        this.id = "";
        this.name = "";
        this.monthlySalary = 0.0;
    }
    public Employee(String id, String name, double monthlySalary) {
        this.id = id;
        this.name = name;
        this.monthlySalary = monthlySalary;
    }
    public void setId(String id) {   this.id = id; }
    public void setName(String name) { this.name = name; }
    public void setMonthlySalary(double monthlySalary) {  this.monthlySalary = monthlySalary; }
    public String getId() { return id; }
    public String getName() {  return name; }
    public double getMonthlySalary() {   return monthlySalary; }
    public void displayEmployee() {
        System.out.println("Employee ID: " + id);
        System.out.println("Name: " + name);
        System.out.println("Monthly Salary: " + monthlySalary);
        System.out.println("Yearly Salary: " + (monthlySalary * 12));
        System.out.println("--------------------------------");
    }
}
public class EmployeeTest {
    public static void main(String[] args) {
        Employee emp1 = new Employee("E101", "Alice", 5000.0);
        Employee emp2 = new Employee("E102", "Bob", 6000.0);
        emp1.displayEmployee();
        emp2.displayEmployee();
    }
}
