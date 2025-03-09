// JAva is does not have append as a in-built fucntion , so we have to use StringBuilder 
StringBuilder sb = new StringBuilder("Hello");
sb.append(" World");
System.out.println(sb);
// Java has replace built in so we can use it. It is used to replace all occurrences of a specified character or string with another character or string 
String str= "Gnd";
str= str.replace("Gnd","Gandu");
System.out.println(str);
// equals() equals is a method in java which checks if two objects are equal based on thir content 
String str1="Andu";
String str2= "Gandu";
System.out.println(str1.equals(str2));
//compareTo() is a method used to compare two strings lexicographically.
String str1 = "Hello";
String str2 = "World";
System.out.println(str1.compareTo(str2));
//toUpperCase() converts all characters in a string to uppercase.
String str = "hello";
str = str.toUpperCase();
System.out.println(str);  
//Integer.parseInt() is used to parse a string and convert it into an integer.
String str = "123";
int num = Integer.parseInt(str);
System.out.println(num);
