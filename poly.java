public class Animal {
    String name = "";
    int hunger = 50;
    public Animal() {    }
    public Animal(String name) {this.name = name;}
    public void makeNoise() { System.out.println("Makes noise"); }
    public void eat() {  hunger -= 10; }
    public void move() {  hunger += 10;}
    public boolean isHungry() {    return hunger > 50; }
    @Override
    public String toString() { return "Animal, name: " + name;}
  
}
public class Chicken extends Animal {
    int numberOfFeathers = 5_000;
    public Chicken() { }
    public Chicken(String name) { super(name); }
    public Chicken(String name, int numberOfFeathers) {  super(name);   this.numberOfFeathers = numberOfFeathers;}
    @Override
    public void makeNoise() {  System.out.println("Cluck!");    }
    @Override
    public String toString() { return "Chicken, name: " + name; }
}
public class Lion extends Animal {
    int numberOfTeeth = 30;
    public Lion() {}
    public Lion(String name) { super(name);}
    public Lion(String name, int numberOfTeeth) { super(name); this.numberOfTeeth = numberOfTeeth; }
    @Override
    public void makeNoise() { System.out.println("ROAR!"); }
    @Override
    public String toString() { return "Lion, name: " + name; }
}
public class TestAnimals {
    public static void main(String[] args) {
        Animal[] animals = { new Lion(), new Animal(), new Chicken(),new Animal("Kaz"), new Lion("Matthias", 50),
            new Chicken("Jordie"),
            new Chicken("Inej", 3000),
            new Lion("Nina"),
        };
        for (Animal a : animals) {
            System.out.println(a);
        }
    }
}

