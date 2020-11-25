---
layout: page
permalink: /python/advanced
---
[Go Back to Python Basics](/python/basics) &nbsp; &nbsp; &nbsp;

<br>

<h1><b>Python Advanced</b></h1>
Welcome to the advanced part of Python Programming

<h1><b>Table Of Contents</b></h1>

* TOC
{:toc}

### Map, Filter, Reduce
Map, Filter, and Reduce are paradigms of functional programming. They allow the programmer (you) to write simpler, shorter code, without neccessarily needing to bother about intricacies like loops and branching. Essentially, these three functions allow you to apply a function across a number of iterables, in one full swoop. map and filter come built-in with Python (in the `__builtins__` module) and require no importing. reduce, however, needs to be imported as it resides in the functools module.

__*map(func, iter)*__ : This function returns a map object(which is an iterator) of the results after applying the given function to each item of a given iterable (list, tuple etc.) `func` s a function to which map passes each element of given iterable. `iter` is an iterable which is to be mapped.
```python
>>> numbers = [1,2,3,4,5]
>>> result = map(double, numbers)
>>> print(result)
<map object at 0x7f2f2f47e3a0>
>>> type(result)
<class `map`>
>>> print(list(result))
[2, 4, 6, 8, 10]
>>> # we can also use lambda expressions to achieve the above result
>>> result = map(lambda n : 2*n, numbers)
>>> print(result)
<map object at 0x7f2f2fc79b50>
>>> type(result)
<class `map`>
>>> print(list(result))
[2, 4, 6, 8, 10]
```

__*filter(func, seq)*__ : This method filters the given sequence with the help of a function that tests each element in the sequence to be true or not. The `func` that tests if each element of a sequence true or not. The `seq` which needs to be filtered, it can be sets, lists, tuples, or containers of any iterators. The filter method returns an iterator that is already filtered.
```python
>>> def vowel(ch):
...     if ch in ['a', 'e', 'i', 'o', 'u']:
...             return True
...     return False
>>> seq = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o']
>>> filtered_seq = filter(vowel, seq)
>>> filtered_seq
<filter object at 0x7f2f2f47e400>
>>> type(filtered_seq)
<class `filter`>
>>> list(filtered_seq)
['a', 'e', 'i', 'o']
```

__*reduce(fun,seq)*__ : This function is used to apply a particular function passed in its argument to all of the list elements mentioned in the sequence passed along.This function is defined in `functools` module.
```python
>>> from functools import reduce
>>> numbers = [1,2,3,4,5]
>>> value = reduce(lambda a,b: a+b, numbers)
>>> print(value)
15
>>> value = reduce(lambda a,b: a*b, numbers)
>>> print(value)
120
```

### pass Statement
This is a null statement which does nothing when executed, useful as a placeholder
```python
>>> class New_Class:
...     pass
>>> 
>>> def New_func():
...     pass
>>> if True:
...     print(True)
...     pass
True
>>> New_func()
>>> New_Class()
<__main__.New_Class object at 0x7f637b442940>
```

### Object Oriented Python - Classes and Objects
* Objects are an encapsulation of variables and functions into a single entity. Objects get their variables and functions from classes. Classes are essentially a template to create your objects.
* `self` keyword is passed as a parameter to the methods of an class. This is passed in the declaration so as to access the current attibutes of the instance created.

Let's take a look at an example of a class. Use this code as reference to define a class.
```python
>>> # define a Circle class
>>> # with one attribute radius and two methods area and circumference
>>> class Circle:
...     radius = 5
...     def area(self):
...             return 3.14 * self.radius * self.radius
...     def circumference(self):
...             return 2 * 3.14 * self.radius
>>> # create two objects of Circle class c_1 and c_2
>>> c_1 = Circle()
>>> c_2 = Circle()
>>> c_1.area()
78.5
>>> c_1.circumference()
31.400000000000002
>>> c_2.radius = 4
>>> c_2.area()
50.24
>>> c_2.circumference()
25.12
>>> print(c_1.radius, " ", c_2.radius)
5   4
```
In the above example we declared a class named `Circle`. We have created two objects of the `Circle` class called `c_1` and `c_2`. To call the methods you can use `c_1.area()`. To access an attribute of the object, you can do that using `c_2.radius`.

### Inheritance
__*TODO*__

### *else* clause in loops
When the loop condition of `for` or `while` statement fails then code part in *else* is executed. If break statement is executed inside for loop then the *else* part is skipped. Note that *else* part is executed even if there is a continue statement. Here's an example
```python
>>> count = 0
>>> while(count < 5):
...     print(count)
...     count += 1
... else:
...     print("count value reached %d" % (count))
0
1
2
3
4
count value reached 5
>>> for i in range(1, 10):
...     if(i % 5 == 0):
...         break
...     print(i)
... else:
...     print("this is not printed because for loop is terminated because of break but not due to fail in condition")
1
2
3
4
```

### Generators
Generators are used to create iterators. They are simple functions which return an iterable set of items, one at a time, in a special way. When an iteration over a set of item starts using the `for` statement, the generator is run. Once the generator's function code reaches a `yield` statement, the generator yields its execution back to the `for` loop, returning a new value from the set. The generator function can generate as many values (possibly infinite) as it wants, yielding each one in its turn.

Let's create a generator which returns values of a list. To create generators use the following code for reference
```python
>>> def generator_func(arg_list):
...     for i in arg_list:
...             yield i
>>> gen = generator_func(list_1)
>>> gen
<generator object generator_func at 0x7f637a3cbe40>
>>> for ans in gen:
...     print(ans)
1
2
3
```

### List Comprehensions
List Comprehensions is a very powerful tool, which creates a new list based on another list, in a single, readable line. Let's  say we need to create a list of natural numbers from a list of integers, we'd normally do it the following way
```python
>>> integers = [3, -2, 1, 0, -1, 2, -3]
>>> natural = []
>>> for num in integers:
...     if num > 0:
...             natural.append(num)
>>> print(integers)
[3, -2, 1, 0, -1, 2, -3]
>>> print(natural)
[3, 1, 2]
```
Using a list comprehension, we could simplify the above code in to a single line.
```python
>>> integers = [3, -2, 1, 0, -1, 2, -3]
>>> natural = [num for num in integers if num > 0]
>>> print(natural)
[3, 1, 2]
```

### Exception Handling
When programming, errors happen. Python's solution to errors are exceptions. But sometimes you don't want exceptions to completely stop the program. You might want to do something special when an exception is raised. This is done in a *try/except* block. The `try` block lets you test a block of code for errors. The `except` block lets you handle the error. The `finally` block lets you execute code, regardless of the result of the try and except blocks.
```python
>>> x
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
NameError: name 'x' is not defined
>>> # exception handling
>>> try:
...     print(x)
... except:
...     print("An exception occurred")
An exception occurred
>>> # many exceptions : you can define as many exception blocks as you want,
>>> # e.g. if you want to execute a special block of code for a special
>>> # kind of error
>>> try:
...     print(x)
... except NameError:
...     print("Variable x is not defined")
... except:
...     print("Something else went wrong")
Variable x is not defined
>>> # else : you can use the else keyword to define a block of code to be
>>> # executed if no errors were raised:
>>> try:
...     print("else example")
... except:
...     print("something went wrong")
... else:
...     print("nothing went wrong")
else example
nothing went wrong
>>> # finally : The finally block, if specified, will be executed regardless
>>> # if the try block raises an error or not.
>>> try:
...     print(x)
... except:
...     print("something went wrong")
... finally:
...     print("the 'try except' is finished")
something went wrong
the 'try except' is finished
>>> # raise an exception : As a Python developer you can choose to throw an
>>> # exception if a condition occurs. To throw (or raise) an exception,
>>> # use the raise keyword.
>>> x = -1
>>> if x < 0:
...     raise Exception("Number below zero")
Traceback (most recent call last):
  File "<stdin>", line 2, in <module>
Exception: Number below zero
```

### Decorators
Decorators allow you to make simple modifications to callable objects like functions, methods, or classes. The syntax is as follows
```python
>>> @decorator
... def functions(arg):
...     return "value"
>>> # is equivalent to
>>> def function(arg):
...     return "value"
>>> function = decorator(function)
>>> # this passes the function to the decorator, and reassigns it
>>> # to the functions
```
Let's take a look at an example
```python
>>> def repeate_twice(input_func):
...     def output_func(*args, **kwds):
...         input_func(*args, **kwds) # we run the old function
...         input_func(*args, **kwds) # we do it twice
...     # we have to return the new_function, or it wouldn't
...     # reassign it to the value
...     return output_func
>>> @repeate_twice
... def add(a, b):
...     print(a + b)
>>> add(3,4)
7
7
```

<br><br>

[Go Back to Python Basics](/python/basics) &nbsp; &nbsp; &nbsp;