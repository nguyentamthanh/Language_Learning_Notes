/**
 * Java Basics - Classes và Objects
 */

// Class definition
class Person {
    // Fields (attributes)
    private String name;
    private int age;
    
    // Constructor
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
    
    // Getter methods
    public String getName() {
        return name;
    }
    
    public int getAge() {
        return age;
    }
    
    // Setter methods
    public void setName(String name) {
        this.name = name;
    }
    
    public void setAge(int age) {
        this.age = age;
    }
    
    // Method
    public void introduce() {
        System.out.println("Xin chào, tôi là " + name + ", " + age + " tuổi.");
    }
}

// Main class
public class Classes {
    public static void main(String[] args) {
        // Create objects
        Person person1 = new Person("Java", 25);
        Person person2 = new Person("Python", 30);
        
        // Call methods
        person1.introduce();
        person2.introduce();
        
        // Access fields through getters
        System.out.println("Tên: " + person1.getName());
        System.out.println("Tuổi: " + person1.getAge());
    }
}

