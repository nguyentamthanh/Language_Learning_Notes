/**
 * Java Basics - Variables và Data Types
 */
public class Variables {
    public static void main(String[] args) {
        // 1. Primitive Data Types
        
        // Integer types
        byte byteVar = 127;           // 8-bit, -128 to 127
        short shortVar = 32767;       // 16-bit
        int intVar = 2147483647;      // 32-bit
        long longVar = 9223372036854775807L; // 64-bit
        
        // Floating point types
        float floatVar = 3.14f;       // 32-bit
        double doubleVar = 3.14159;   // 64-bit
        
        // Character type
        char charVar = 'A';
        
        // Boolean type
        boolean boolVar = true;
        
        // 2. Reference Data Types
        String str = "Hello, Java!";
        String str2 = new String("Hello");
        
        // 3. Arrays
        int[] numbers = {1, 2, 3, 4, 5};
        int[] numbers2 = new int[5];
        
        // 4. Type conversion
        int x = 10;
        double y = x;  // Implicit conversion
        
        double d = 3.14;
        int i = (int) d;  // Explicit cast
        
        // 5. Constants
        final int MAX_SIZE = 100;
        final String APP_NAME = "MyApp";
        
        // 6. Print variables
        System.out.println("Integer: " + intVar);
        System.out.println("Double: " + doubleVar);
        System.out.println("String: " + str);
        System.out.println("Boolean: " + boolVar);
    }
}

