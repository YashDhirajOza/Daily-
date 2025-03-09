Employee = {}
Employee.__index = Employee
function Employee:new(id, name, monthlySalary)
    local obj = {
        id = id or "",
        name = name or "",
        monthlySalary = monthlySalary or 0.0
    }
    setmetatable(obj, Employee)
    return obj
end
function Employee:setId(id)
    self.id = id
end

function Employee:setName(name)
    self.name = name
end

function Employee:setMonthlySalary(salary)
    self.monthlySalary = salary
end

function Employee:getId()
    return self.id
end

function Employee:getName()
    return self.name
end

function Employee:getMonthlySalary()
    return self.monthlySalary
end

function Employee:displayEmployee()
    print("Employee ID:", self.id)
    print("Name:", self.name)
    print("Monthly Salary:", self.monthlySalary)
    print("Yearly Salary:", self.monthlySalary * 12)
    print("--------------------------------")
end
emp1 = Employee:new("E101", "Alice", 5000.0)
emp2 = Employee:new("E102", "Bob", 6000.0)

emp1:displayEmployee()
emp2:displayEmployee()
