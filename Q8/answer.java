
public class Animal {  
    final int speed = 400;

    public void displaySpeed() {
        System.out.println("Speed: " + speed);
    }
}

public class Cat extends Animal {  
    public void showInfo() {
        System.out.println("The cat's speed is inherited from Animal.");
        super.displaySpeed(); 
    }
}

public class Okay {
    public static void main(String[] args) {  // 
        Cat c = new Cat();
        c.showInfo();  
    }
}
