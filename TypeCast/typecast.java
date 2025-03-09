//Type Casting 
// Type Castinf refers to converting one data type to another . It is an important concept in programming
// tHERE ARE two types of type casting 1. Widening (Autmoatic Type Casting) 2. Narrowing (Manual Type Casting)
// Widening:- This happedn automatically when a smaller data type is converted into a larger one . It is safe because there is no loss of data and precision Example like int -->long, float---> double 
// this type of type castinf automatically done by the java 
//Example: byte → short → int → long → float → double
public class WideningExample {
    public static void main(String[] args) {
        int number = 100;
        double result = number; // Implicit casting from int to double
        System.out.println("Widened Value: " + result);
    }
}
// Narrowing: THis does not happen automatically . We have to do it manually from a larger data type to smaller data type . There may be some amount of data loss during this data conversion 
//Example: double → float → long → int → short → byte
public class NarrowingExample {
    public static void main(String[] args) {
        double price = 99.99;
        int discountedPrice = (int) price; // Explicit casting from double to int
        System.out.println("Discounted Price: " + discountedPrice);
    }
}
