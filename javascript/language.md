---
layout: page
permalink: /js/language
---

<h1><b>Javascript Language</b></h1>

Welcome to the language of Javascript Programming.

<h1><b>Table Of Contents</b></h1>

* TOC
{:toc}

### Hello, World!
To print something, use the `console.log` function.
```js
> console.log("Hello, World!");
Hello, World!
undefined
```

### console.log says undefined
The console will print the result of evaluating an expression. The result of evaluating `console.log()` is `undefined` since `console.log` does not explicitly return something. It has the side effect of printing to the console. You can observe the same behaviour with many expressions:
```js
> var a = 5;
undefined;
```
A variable declaration does not produce a value so again undefined is printed to the console. As a counter-example, expressions containing mathematical operators do produce a value which is printed to the console instead of undefined:
```js
> 3 + 5;
8
```
Another example - define two functions, one returns a value and the other does not return anything. Call them and see what happens.
```js
> function funcA() {
... return 1;
... }
undefined
> function funcB() {
... return;
... }
undefined
> funcA();
1
> funcB();
undefined
```

### Naming Variables in JS
A variable name should accurately identify your variable. When you create good variable names, your JavaScript code becomes easier to understand and easier to work with. Properly naming variables is really important! Here are rules JavaScript has for naming variables:
* Variable names cannot contain spaces.
* Variable names must begin with a letter, an underscore (_) or a dollar sign ($).
* Variable names can only contain letters, numbers, underscores, or dollar signs.
* Variable names are case-sensitive.
* Certain words may not be used as variable names, because they have other meanings within JavaScript.

There are many choices available. Many professional programmers agree that there are best practices to keep in mind when naming your variables:
* Don’t use names that are too short.
* Use more than one word to name your variable.
* Pick a style for names with more than one word, and stick to it (use camelCase or and underscore for joining the wrods).

I am going to use camelCase through out this tutorial.

### Variables and Types
Types in Javascripts
* number
* string
* boolean
* bigint
* undefined
* object
* function
* symbol

The `typeof` is the function used to find out the type of variable. Syntax of the typeof function is as follows
```js
> typeof(variableName);
> // or
> typeof variableName;
```
Every variable is defined using the `var` keyword.
```js
> var myNumber = 1; // number can be both integer or floating point
undefined
> var myFloatNumber = 1.5;
undefined
> var myString = "this is a string";
undefined
> var myBoolean = true; // boolean can be true or false
undefined
> typeof(myNumber);
'number'
> typeof myFloatNumber;
'number'
> typeof myString;
'string'
> typeof(myBoolean);
'boolean'
```
Bigint is a built-in object that provides a way to represent whole numbers larger than 2<sup>53</sup> - 1. It is created by appending n to the end of integer literal - `10n`.
```js
> var myBigInt = BigInt(10);
undefined
```
A variable type is undefined if no value has been assigned to it yet.
```js
> var myVariable;
undefined
> typeof myVariable;
'undefined'
> console.log(myVariable);
undefined
undefined
```
The `null` value is a different type of value, and is used when a variable should be marked as empty.
```js
> var nullVariable = null;
undefined
> typeof(nullVariable);
'object'
> console.log(nullVariable);
null
undefined
```
Unlike the lists and dictionaries in Python, in JavaScript we have a common type called `object`.
```js
> var myObject = [];
undefined
> typeof(myObject);
'object'
> var myObject = {};
undefined
> typeof(myObject);
'object'
```
We are gonna discuss in the later sections about `function` and `symbol`.

### Array Object
JS can hold an array of variables in an array object, the values can be of any type. To define an array, use the following syntax
```js
> var myArray = [1, "a", true, false];
undefined
> var sameArray = new Array(1, "a", true, false);
undefined
> // indexing start at 0. // to access just use myvariable[index]
undefined
> console.log(myArray[1], myArray[2]);
a true
undefined
```
Arrays in JavaScript are sparse, meaning that we can also assign variables to random locations even though previous cells were undefined. For example:
```js
> var newArray = [];
undefined
> newArray[2] = 1;
1
> console.log(newArray);
[ <2 empty items>, 1 ]
undefined
> newArray[0] = true;
true
> newArray[4] = false;
false
> console.log(newArray);
[ true, <1 empty item>, 1, <1 empty item>, false ]
undefined
```
__Methods on Arrays__
* __*splice :*__ Meaning to remove a certain part from an array to create a new array.
    ```js
    > var myArray = [0,1,2,3,4,5,6,7,8,9,10];
    undefined
    > // let use remove 4 numbers from second index
    > var mySplice = myarray.splice(2,4);
    undefined
    > console.log(mySplice);
    [ 2, 3, 4, 5 ]
    undefined
    > console.log(myArray);
    [ 0, 1, 6, 7, 8, 9, 10 ]
    undefined
    ```
* __*push and pop :*__ Arrays in js can function as a stack using the push and pop methods to insert and remove the values at the end of an array respectively.
    ```js
    > var myStackArray = [];
    undefined
    > myStackArray.push(0);
    1
    > myStackArray.push('a');
    2
    > myStackArray.push(1);
    3
    > console.log(myStackArray);
    [ 0, 'a', 1 ]
    undefined
    > myStackArray.push(2);
    4
    > myStackArray.push('b');
    5
    > console.log(myStackArray);
    [ 0, 'a', 1, 2, 'b' ]
    undefined
    > myStackArray.pop();
    'b'
    > console.log(myStackArray);
    [ 0, 'a', 1, 2 ]
    undefined
    ```
* __*unshift and shift :*__ unshift method on arrays inserts elements at the beginning of the array. shift method removed elements from the begining of the array.
    ```js
    > var myArray = [];
    undefined
    > myArray.unshift(0);
    1
    > myArray.unshift(1);
    2
    > myArray.unshift(2);
    3
    > console.log(myArray);
    [ 2, 1, 0 ]
    undefined
    > myArray.shift();
    2
    > console.log(myArray);
    [ 1, 0 ]
    undefined
    ```

### Dictionary Object
Dictionary Object in JS is similar to the one in Python but the way we access the dictionary is a little different. Let's see with an example.
```js
> var emptyDict = {};
undefined
> console.log(emptyDict);
{}
undefined
> var studentDict = {
... firstName:"first",
... lastName:"last"
... }
undefined
> console.log(studentDict);
{ firstName: 'first', lastName: 'last' }
undefined
> // accessing the object
undefined
> studentDict['firstName']
'first'
> studentDict.lastName
'last'
> // add new keys
undefined
> studentDict.rollNum = 540
540
> console.log(studentDict);
{ firstName: 'first', lastname: 'last', rollNum: 540 }
undefined
> // iterating over the object
undefined
> for (var key in studentDict)
... {
... console.log(key, studentDict[key]);
... }
firstName first
lastname last
rollNum 540
undefined
> // delete key
undefined
> delete studentDict.rollNum
true
> console.log(studentDict);
{ firstName: 'first', lastname: 'last' }
undefined
```

### Math Operators

__*Typecasting*__ or coercion in simple term means to change the data type of a value to another data type like for example, integer to a string or a string to boolean, etc.

Every variable in Javascript is casted automatically so any operator between two variables will always give some kind of result.

Let's see with some examples how each operator behaves.
* __Addition Operator (+)__
    ```js
    > // number + number = number
    undefined
    > 1 + 2;
    3
    > typeof(1+2);
    'number'
    > // number + string = string
    undefined
    > 1 + '2';
    '12'
    > typeof(1+'2');
    'string'
    > 1 + 'a';
    '1a'
    > typeof(1+'a');
    'string'
    > // string + string = string
    undefined
    > 'a' + 'b';
    'ab'
    > // number + boolean = number
    undefined
    > 1 + true;
    2
    > 1 + false;
    1
    > typeof(1 + false);
    'number'
    > // string + boolean = string
    undefined
    > '1' + true;
    '1true'
    > // boolean + boolean = number
    undefined
    > true + true;
    2
    > typeof(true + true);
    'number'
    ```
* __Multiplication Operator (*)__
    ```js
    > // number*number = number
    undefined
    > 2*3;
    6
    > typeof(2*3);
    'number'
    > // number * string = number
    undefined
    > 2 * '3';
    6
    > typeof(2 * '3');
    'number'
    > 2 * 'a';
    NaN
    > typeof(2 * 'a');
    'number'
    > // string * string = number
    undefined
    > 'a'*'b';
    NaN
    > typeof('a'*'b');
    'number'
    > // number * boolean = number
    undefined
    > 2 * true;
    2
    > 2 * false;
    0
    > typeof(2*false);
    'number'
    > // string * boolean = number
    undefined
    > '1' * true;
    1
    > '1' * false;
    0
    > typeof('1'*false);
    'number'
    > // boolean * boolean = number
    undefined
    > true * true;
    1
    > true * false;
    0
    ```
* __Subtraction Operator (-)__
    ```js
    > // number - number = number
    undefined
    > 1-3;
    -2
    > typeof(1-3);
    'number'
    > // number - string = number
    undefined
    > 2 - 'a';
    NaN
    > 'a' - 2;
    NaN
    > 2 - '3';
    -1
    > typeof(2-'3');
    'number'
    > // string - string = number
    undefined
    > 'a' - 'b';
    NaN
    > '1' - '3';
    -2
    > typeof('1'-'3');
    'number'
    > // number - boolean = number
    undefined
    > 2 - true;
    1
    > true - 3;
    -2
    > typeof(true-3);
    'number'
    > // string - boolean = number
    undefined
    > 'a' - true;
    NaN
    > '2' - true;
    1
    > typeof('2'-true);
    'number'
    > // boolean - boolean = number
    undefined
    > true - true;
    0
    > false - true;
    -1
    > typeof(false - true);
    'number'
    ```
* __Division Operator (/)__
    ```js
    > // number/number = number
    undefined
    > 6/3;
    2
    > typeof(6/3);
    'number'
    > // number/string = number (or) string/number = number
    undefined
    > 6/'3';
    2
    > '6'/3;
    2
    > typeof('6'/3);
    'number'
    > // string/string = number
    undefined
    > '6'/'3';
    2
    > typeof('6'/'3');
    'number'
    > 'a'/'3'
    NaN
    > // number/boolean = number
    undefined
    > 6/true;
    6
    > 6/false;
    Infinity
    > true/2;
    0.5
    > typeof(true/2);
    'number'
    > // string/boolean = number
    undefined
    > '6'/true;
    6
    > 'a'/true;
    NaN
    > typeof('6'/true);
    'number'
    > // boolean/boolean = number
    undefined
    > true/false;
    Infinity
    > false/true;
    0
    > false/false;
    NaN
    > true/true;
    1
    > typeof(true/true);
    'number'
    ```
* __Modulo Operator (%)__
    ```js
    > // number%number = number
    undefined
    > 6%3;
    0
    > 8%3;
    2
    > typeof(8%3);
    'number'
    > // number%string = number (or) string%number = number
    undefined
    > 8%'3';
    2
    > '8'%3;
    2
    > typeof('8'%3);
    'number'
    > // string%string = number
    undefined
    > '8'%'3';
    2
    > // number%boolean = number
    undefined
    > 8%true;
    0
    > 8%false;
    NaN
    > // string%boolean = number
    undefined
    > '8'%true;
    0
    > '8'%false;
    NaN
    > // boolean%boolean = number
    undefined
    > true%true;
    0
    > true%false;
    NaN
    > false%true;
    0
    > false%true;
    0
    ```
*Note :* JS, similar to python supports combined assignment and operation operators. So, instead of typing `x = x * 2`, you can type `x *= 2`.

JS also has a `Math` module which contains more advanced functions
* `Math.abs(x)` calculates the absolute value of number `x`
    ```js
    > Math.abs;
    [Function: abs]
    > Math.abs(-3.5);
    3.5
    > Math.abs(3.5);
    3.5
    ```
* `Math.exp(a)` calculates `e` to the power of number `a`, which is e<sup>a</sup>.
    ```js
    > Math.exp;
    [Function: exp]
    > Math.exp(2);
    7.38905609893065
    > Math.exp(3);
    20.085536923187668
    ```
* `Math.pow(x,y)` calculates the result of x to the power of y
    ```js
    > Math.pow;
    [Function: pow]
    > Math.pow(2,3);
    8
    ```
* `Math.floor(x)` calculates the floor of number x
    ```js
    > Math.floor;
    [Function: floor]
    > Math.floor(3.4);
    3
    > Math.floor(3.8);
    3
    > Math.floor(-3.8);
    -4
    > Math.floor(-3.4);
    -4
    ```
* `Math.random()` will give a random number x where `0 <= x < 1`
    ```js
    > Math.random;
    [Function: random]
    > Math.random();
    0.3170583707148862
    > Math.random();
    0.7260007199323444
    > Math.random();
    0.8801701750277489
    > Math.random();
    0.7888717146073436
    ```
and many more functions are offered by this module. Use the [Mozilla Developer Docs](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math) for reference.

### Logical Operators
Logical AND (&&) operator, Logical OR (||) operator and NOT (!) operator
```js
> true && true;
true
> true && false;
false
> false && false;
false
> true || true;
true
> true || false;
true
> false || false;
false
> !true;
false
```

### Equals(==) and Equality(===) Operator
To evaluate whether two variables are equal, the == operator is used. The equality operator in JS, ===, does a strict comparison. This means that it will be true only if the two things you are comparing are the same type as well as same content.
```js
> 1 == 1;
true
> 1 === 1;
true
> 1 === '1';
false
> 'a' == 'b';
false
> 'a' === 'a';
true
```

### Conditions
* __*if*__ statement - syntax is as follows.
    ```js
    if condition {
        ....
    } else {
        ....
    }
    ```
    Let's look at an example.
    ```js
    > var a = 5;
    undefined
    > if(a == 5) {
    ... console.log("Value is 5");
    ... }
    Value is 5
    undefined
    > 
    > if(a == 5) {
    ... console.log("Value is 5");
    ... } else {
    ... console.log("Value is not 5");
    ... }
    Value is 5
    undefined
    ```
* __*switch*__ statement - syntax is as follows
    ```js
    switch(variable)
    {
        case value1:
        case value2:
            console.log('variable is value2');
            break;
        default:
            console.log('default block');
            break;
    }
    ```
    Let's look at a switch example
    ```js
    > var num = 5;
    undefined
    > switch(num)
    ... {
    ...   case 1:
    ...   case 2:
    ...         console.log("num value is 2");
    ...         break;
    ...   case 5:
    ...         console.log("num value is 5");
    ...   case 6:
    ...         console.log("num value is 6");
    ...         break;
    ...   default:
    ...         console.log("num value is unknown");
    ...         break;
    ... }
    num value is 5
    num value is 6
    undefined
    > switch(num)
    ... {
    ...   case 5:
    ...         console.log("num value is 5");
    ...         break;
    ...   case 6:
    ...         console.log("num value is 6");
    ...         break;
    ... }
    num value is 5
    undefined
    ```
    Using the switch statement in general is not recommended, because forgetting the break keyword causes very confusing results. If break keyword is missing from a case then it continues to execute the next block too, like in the above example.

### Loops
* __*for*__ loop : The `for` statement has the same syntax as in Java, C and Python. The `for` loop syntax is as follows
    ```js
    for(initialize; condition; increment)
    {
        ....
    }
    ```
    Let's look at an example
    ```js
    > for(var i = 0; i < 5; i++)
    ... {
    ... console.log(i);
    ... }
    0
    1
    2
    3
    4
    undefined
    > // let us iterate over an array of numbers
    undefined
    > var myArray = [1,2,3,4,5];
    undefined
    > for(var i = 0; i < myArray.length; i++)
    ... {
    ...  console.log(myArray[i]);
    ... }
    1
    2
    3
    4
    5
    undefined
    ```
* __*while*__ loop : The `while` statement has the same syntax as in Java, C and Python. The `while` loop syntax is as follows
    ```js
    while(condition)
    {
        ....
    }
    ```
    Let's look at an example
    ```js
    > var i = 0;
    undefined
    > while(i < 5)
    ... {
    ...  console.log(i);
    ...  i++;
    ... }
    0
    1
    2
    3
    4
    > var myArray = [1,2,3,4,5];
    undefined
    > var i = 0;
    undefined
    > while(i < myArray.length)
    ... {
    ...  console.log(myArray[i]);
    ...  i++;
    ... }
    1
    2
    3
    4
    5
    ```
* __*break*__ and __*continue*__ statement : The `break` statement allows to stop the execution of the loop, whereas the `continue` statement skips the rest of the loop and jumps back to the begining of the loop. Let's take a look at two examples using the `for` loop.
    ```js
    > var myArray = [1,2,3,4,5,6,7,8,9,10];
    undefined
    > for(var i = 0; i < myArray.length; i++)
    ... {
    ...  console.log(myArray[i]);
    ...  if(i == myArray.length/2 - 1)
    ...  {
    ..... break;
    ..... }
    ... }
    1
    2
    3
    4
    5
    undefined
    > for(var i = 0; i < myArray.length; i++)
    ... {
    ... if(i%2 == 0)
    ... {
    ..... continue;
    ..... }
    ... console.log(myArray[i]);
    ... }
    2
    4
    6
    8
    10
    undefined
    ```

### Functions
Functions in JS are similar to functions in Python, Java. They are code blocks that can have arguments, and have their own scope. Syntax of a function in JS is as follows
```js
// defining a function
function functionName(arg1, arg2, and, soOn)
{
    ....
    return something;
}
// calling the function
functionName(arg_1, arg_2, aND, soON)
```
There are two types of functions in JS - name functions and anonymous functions. The named function is defined as above, whereas in an anonymous function there is no functionName in the definition and we assign it to a variable.
```js
var functionVariableName = function(arg1, arg2, and, soOn)
{
    ....
    return something;
}
// calling the function
functionVariableName(arg_1, arg_2, aND, soON)
```
Let's look at an example for both the types with a sum function that takes two variables `x` and `y` and returns `x+y`.
```js
> function sumNamed(x, y)
... {
... return x + y;
... }
undefined
> sumNamed(3, 4);
7
> var sumAnonymous = function(x, y)
... {
... return x + y;
... }
undefined
> sumAnonymous(3,4);
7
> typeof(sumNamed);
'function'
> typeof(sumAnonymous)
'function'
```
The only difference between a named and an anonymous function is the way we define them.

__*Binding a Function :*__ Functions in JavaScript run in a specific context, and using the this variable we have access to it. Functions defined under an object or a class (another function) will use the context of the object it was created in. However, we can also change the context of a function at runtime, either before or while executing the function. To bind a function to an object and make it an object method, we can use the `bind` function. Let's look at an example
```js
> var person = {name : {first: "John", last: "Doe"}, age : 23};
undefined
> function printName() {
... console.log(this.name);
... }
undefined
> // we need to associate the printName function to person and then call it.
undefined
> var boundMethod = printName.bind(person);
undefined
> boundMethod();
{ first: 'John', last: 'Doe' }
undefined
> // or we can also use the call function to bind it temporarily
undefined
> printName.call(person);
{ first: 'John', last: 'Doe' }
undefined
```
Think of `call` as executing the return value of `bind`. Hence
```js
> printName.call(person); // is the same as
{ first: 'John', last: 'Doe' }
undefined
> printName.bind(person)(); // executes the function returned by bind
{ first: 'John', last: 'Doe' }
undefined
```

### Object Oriented Javascript
Javascript Classes are templates for JS Objects. The `class` keyword is used to create a class. We need to add a method named `constructor()` to instatiate the objects. The syntax to create a class is as follows
```js
class ClassName {
    constructor() {....}

    // add any number of methods
    method1(arg1..) { ... }
}
```
Let's create a Class Person with two properties name and year of birth and a method called age.
```js
> class Person {
... constructor(name, birthYear) {
..... this.name = name;
..... this.birthYear = birthYear;
..... }
... age() {
..... var date = new Date();
..... return date.getFullYear() - this.birthYear;
..... }
... }
undefined
> 
> var person1 = new Person("John", 1997);
undefined
> person1.name;
'John'
> person1.birthYear;
1997
> person1.age();
23
```
In JS we can also use functions as classes to create objects using `new` keyword. Let's take a look at an example
```js
> function Person(name, birthYear) {
... this.name = name;
... this.birthYear = birthYear;
... this.age = function() {
..... var date = new Date();
..... return date.getFullYear() - this.birthYear;
..... }
... }
undefined
> var person1 = new Person("John", 1997);
undefined
> person1.name;
'John'
> person1.birthYear;
1997
> person1.age();
23
> // creating an object using the new keyword is similar to
undefined
> var person2 = {
... name : "John",
... birthYear : 1997,
... age : function() {
..... var today = new Date();
..... return today.getFullYear() - this.birthYear;
..... }
... }
undefined
> 
> person2.name;
'John'
> person2.birthYear;
1997
> person2.age();
23
```
The difference between the two methods of creating objects is that the first method uses a class to define the object and then the `new` keyword to instantiate it, and the second method immediately creates an instance of the object. But the classses in JS are not as powerfull as the functions being used as classes. It is recommended to use functions for classes.

__*Inheritance :*__ JS uses the concept of prototype and prototype chaining for inheritance. Prototype chaining means an object’s proto property will point to another object instead of pointing to the constructor function prototype. If the other object’s proto property points to another object it will result in the chain.
```js
> function Person(fName, lName, age) {
... this.name = {firstName : fName, lastName : lName};
... this.age = age;
... this.describe1 = function() {
..... return "This is describe function";
..... }
... }
undefined
> // let's define a method describe2 using the prototype method
undefined
> // this method is not inherited
undefined
> Person.prototype.describe2 = function() {
... return "Hi. My name is " + this.name.firstName + " " + this.name.lastName + ". I am " + this.age + " years old.";
... }
[Function]
> var person1 = new Person('John', "Doe", 19);
undefined
> person1.name;
{ firstName: 'John', lastName: 'Doe' }
> person1.age;
19
> person1.describe1();
'This is describe function'
> person1.describe2();
'Hi. My name is John Doe. I am 19 years old.'
> // say we wanted to create a Teacher class which inherits all the members
undefined
> // from Person, also includes a new attribute subject and a method teach
undefined
> // and an updated describe method.
undefined
> function Teacher(fName, lName, age, subject) {
... this.subject = subject;
... // the call() function - basically allows you to call a function defined
... // somewhere else, but in the current context. The first parameter
... // specifies the value of this that you want to use when running the
... // function, and the other parameters are those that should be passed
... // to the function when it is invoked.
... Person.call(this, fName, lName, age);
... }
undefined
> Teacher.prototype.teach = function() {
... return "I teach " + this.subject;
... }
[Function]
> var teacher1 = new Teacher('John', "Doe", 19, "Maths");
undefined
> teacher1.name;
{ firstName: 'John', lastName: 'Doe' }
> teacher1.age;
19
> teacher1.subject;
'Maths'
> teacher1.teach();
'I teach Maths'
> teacher1.describe1();
'This is describe function'
> teacher1.describe2();
Thrown:
TypeError: teacher1.describe2 is not a function
> // now let's update the describe function
undefined
> Teacher.prototype.describe2 = function() {
... return "Hi. My name is " + this.name.firstName + " " + this.name.lastName + ". I am " + this.age + " years old. I teach " + this.subject;
... }
[Function]
> teacher1.describe2();
'Hi. My name is John Doe. I am 19 years old. I teach Maths'
```