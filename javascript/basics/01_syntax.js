/**
 * JavaScript Basics - Cú pháp cơ bản
 */

// 1. Comments (Ghi chú)
// Single line comment
/* Multi-line comment */

// 2. Console output
console.log("Hello, World!");
console.log("Xin chào", "JavaScript");

// 3. Variables
// var (old way, function-scoped)
var oldVar = "old";

// let (block-scoped, can be reassigned)
let name = "JavaScript";
name = "JS";

// const (block-scoped, cannot be reassigned)
const PI = 3.14159;
// PI = 3.14; // Error!

// 4. Data Types
// String
let text = "Hello";
let text2 = 'World';
let template = `Template literal ${text}`;

// Number
let integer = 42;
let float = 3.14;
let infinity = Infinity;
let notANumber = NaN;

// Boolean
let isTrue = true;
let isFalse = false;

// Null and Undefined
let nothing = null;
let notDefined = undefined;

// Symbol (ES6)
let sym = Symbol("unique");

// BigInt (ES2020)
let bigNumber = 9007199254740991n;

// 5. Type checking
console.log(typeof text);        // "string"
console.log(typeof integer);     // "number"
console.log(typeof isTrue);      // "boolean"
console.log(typeof nothing);     // "object" (bug in JS)
console.log(typeof notDefined);  // "undefined"

// 6. Type conversion
let strNumber = "123";
let intNumber = parseInt(strNumber);
let floatNumber = parseFloat("3.14");
let numberFromString = Number("42");
let stringFromNumber = String(42);
let boolFromString = Boolean("true");

// 7. Template literals (ES6)
let firstName = "John";
let lastName = "Doe";
let fullName = `${firstName} ${lastName}`;
console.log(fullName); // "John Doe"

// 8. Operators
// Arithmetic
let sum = 5 + 3;
let difference = 5 - 3;
let product = 5 * 3;
let quotient = 5 / 3;
let remainder = 5 % 3;
let power = 5 ** 3; // ES2016

// Comparison
let isEqual = 5 === 5;      // Strict equality
let isNotEqual = 5 !== 3;   // Strict inequality
let isLooseEqual = 5 == "5"; // Loose equality (avoid!)

// Logical
let and = true && false;
let or = true || false;
let not = !true;

// 9. Nullish coalescing (ES2020)
let value = null ?? "default";
let value2 = undefined ?? "default";

// 10. Optional chaining (ES2020)
let obj = { name: "John" };
let name = obj?.name;
let age = obj?.age?.value;

