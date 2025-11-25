/**
 * JavaScript Basics - Functions
 */

// 1. Function declaration
function greet() {
    console.log("Hello!");
}

greet();

// 2. Function with parameters
function greetPerson(name) {
    console.log(`Hello, ${name}!`);
}

greetPerson("JavaScript");

// 3. Function with default parameters (ES6)
function greetWithDefault(name = "Guest") {
    console.log(`Hello, ${name}!`);
}

greetWithDefault();
greetWithDefault("JavaScript");

// 4. Function with return value
function add(a, b) {
    return a + b;
}

let result = add(5, 3);
console.log(result);

// 5. Function expression
const multiply = function(a, b) {
    return a * b;
};

// 6. Arrow functions (ES6)
const subtract = (a, b) => {
    return a - b;
};

// Arrow function với single expression
const divide = (a, b) => a / b;

// Arrow function với single parameter
const square = x => x * x;

// Arrow function không có parameters
const sayHello = () => console.log("Hello!");

// 7. Rest parameters (ES6)
function sum(...numbers) {
    return numbers.reduce((total, num) => total + num, 0);
}

console.log(sum(1, 2, 3, 4, 5)); // 15

// 8. Higher-order functions
function createMultiplier(multiplier) {
    return function(number) {
        return number * multiplier;
    };
}

const double = createMultiplier(2);
console.log(double(5)); // 10

// 9. Array methods với functions
const numbers = [1, 2, 3, 4, 5];

// map
const doubled = numbers.map(n => n * 2);

// filter
const evens = numbers.filter(n => n % 2 === 0);

// reduce
const sum = numbers.reduce((acc, n) => acc + n, 0);

// forEach
numbers.forEach(n => console.log(n));

// 10. Immediately Invoked Function Expression (IIFE)
(function() {
    console.log("IIFE executed!");
})();

// 11. Generator functions (ES6)
function* numberGenerator() {
    yield 1;
    yield 2;
    yield 3;
}

const gen = numberGenerator();
console.log(gen.next().value); // 1
console.log(gen.next().value); // 2

// 12. Async functions (ES2017)
async function fetchData() {
    try {
        const response = await fetch('https://api.example.com/data');
        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error:", error);
    }
}

