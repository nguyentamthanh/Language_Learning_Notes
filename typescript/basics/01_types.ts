/**
 * TypeScript Basics - Types
 */

// 1. Basic Types
let name: string = "TypeScript";
let age: number = 30;
let isStudent: boolean = true;
let nothing: null = null;
let notDefined: undefined = undefined;

// 2. Arrays
let numbers: number[] = [1, 2, 3, 4, 5];
let names: Array<string> = ["Alice", "Bob", "Charlie"];

// 3. Tuples
let tuple: [string, number] = ["Alice", 25];

// 4. Enums
enum Color {
    Red,
    Green,
    Blue
}

let color: Color = Color.Red;

enum Status {
    Active = "ACTIVE",
    Inactive = "INACTIVE",
    Pending = "PENDING"
}

// 5. Any
let anything: any = "can be anything";
anything = 42;
anything = true;

// 6. Unknown (safer than any)
let unknownValue: unknown = "could be anything";
// unknownValue.toUpperCase(); // Error!
if (typeof unknownValue === "string") {
    unknownValue.toUpperCase(); // OK
}

// 7. Void
function logMessage(message: string): void {
    console.log(message);
}

// 8. Never
function throwError(message: string): never {
    throw new Error(message);
}

// 9. Union Types
let value: string | number = "hello";
value = 42;

// 10. Intersection Types
interface Person {
    name: string;
}

interface Employee {
    employeeId: number;
}

type EmployeePerson = Person & Employee;

let employee: EmployeePerson = {
    name: "John",
    employeeId: 123
};

// 11. Type Aliases
type ID = string | number;
let userId: ID = "123";
let productId: ID = 456;

// 12. Literal Types
type Direction = "up" | "down" | "left" | "right";
let direction: Direction = "up";

// 13. Optional Properties
interface User {
    name: string;
    age?: number; // Optional
    email?: string;
}

let user: User = {
    name: "John"
    // age và email là optional
};

// 14. Readonly
interface Config {
    readonly apiKey: string;
    readonly baseUrl: string;
}

let config: Config = {
    apiKey: "abc123",
    baseUrl: "https://api.example.com"
};
// config.apiKey = "new"; // Error!

