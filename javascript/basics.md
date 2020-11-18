---
layout: page
permalink: /js/basics
---
[Go Next Javascript Advanced](/js/advanced)

<br>
<h1><b>Javascript Basics</b></h1>
Welcome to the basics part of Javascript Programming. Here we are gonna cover learning the language part.

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

### Variables and Types


<br>

[Go Next Javascript Advanced](/js/advanced)