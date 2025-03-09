-- Lua has no built-in parseInt function, but you can use tonumber() to convert a string to a number.

str = "123"
num = tonumber(str)
print(num) 
-- Lua does not have a built-in toUpperCase function, but you can convert a string to uppercase using the string.upper function.
str = "hello"
str = string.upper(str)
print(str)
 -- Lua uses the string.compare function, which is not built-in but can be implemented or compared using <, >, or == for string comparison.str1 = "Hello"
str2 = "World"
if str1 < str2 then
    print("-1") 
end
str1 = "Hello"
str2 = "Hello"
print(str1 == str2)  
myTable = {1, 2, 3}
table.insert(myTable, 4)
for _, v in ipairs(myTable) do
    print(v)  
end

