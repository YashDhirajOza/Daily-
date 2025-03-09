public class Rectangle{
  public int Length;
  public int Width;
  public Rectangle(){
    this.Length=0; 
    this.Width=0;
    
  }
  public Rectangle(int Length, int Width){
    this.Length= Length;
    this.Width= Width;
  }
  
  public int getArea(){
    return Length*Width ;
  }
}
public class Main{
  public static void main(String[] lol){
    Rectangle obj = new Rectangle;
    Rectangle obj2= new Rectangle(5,10);
    System.out.println(obj2.getArea());
    System.out.println(obj.getArea());
    
  }
}
