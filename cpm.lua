person ={}
fucntion person:new(name,age)
  local obj = {Name= name, Age= age}
  setmetatable(obj,self)
  return obj
end

p1= person:new("Alice",25)
p2=person:new("Bob",30)

print(p1.Name,p1.Age)
print(p2.Name,p2.Age)
