// Construtors in java are speical kinds of methods that are called when an object of class is created 
// Contrutors have no return value and they share the same name as class 
// Contrutors are used to assign initial calues to the variable when an object is created 
public class myClass {
  int x;
  public MyClass(){ this.x=0;}
  public void display(){ System.out.println(x);}
}
public static void main(String[] args){
Myclass okay = new MyClass();
  okay.display();
}}
// Just like the method over loading in java overloading of  construtors is also possiable 
// we can declare multiple construtoes within a calss and each of them can have different types of parameters with the help of contructor overloading
// Construtor Overloading help in creating miltiple contrictors within a class each accepting different parameters thi can be usefull in many ways such as 
// 
//it provides flexibility in object creation.
//it enables you to initialize objects in different ways depending on the information available or needed at the time of creation.
//it also helps in making the code more simple.
//

   public class MyClass {
    public MyClass(){}
    public MyClass(int parameter) {}
    public MyClass(int parameter1, String parameter2) {
    }
}

    
  
