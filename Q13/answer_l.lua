
Rectangle = {}
Rectangle.__index = Rectangle
function Rectangle:new()
    local obj = setmetatable({}, self)
    obj.Length = 0
    obj.Width = 0
    return obj
end
function Rectangle:getArea()
    return self.Length * self.Width
end
local obj = Rectangle:new()
print(obj:getArea())  
