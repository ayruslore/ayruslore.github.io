# Python Basics
Welcome to the basics of Python Programming.

## Contents
* Hello, World!
* Indentation
* Variables and Types
* Assignments
* Operators

## Hello, World!
Python is a very simple language, and has a very straightforward syntax. The simplest directive in Python is the `print` directive - it simply prints out a line (and also includes a newline, unlike in C).

To print a String, just write:
```
>>> print("hello, world!!!")
hello, world!!!
```

## Indentation
Python uses indentation for blocks, instead of curly braces. Both tabs and spaces are supported, but standard indentation requires four spaces. For example:
```
>>> if True:
...     print("True")
True
```

## Variables and Types
As python is object oriented and not statically typed, you do not need to declare variables, or declare their type before using them. *Every variable in Python is an object.* Now let us cover the basic types of variables in python.

__Numbers__ Python supports three types of numbers - integers, floating point numbers, complex numbers.

* To define an integer, use the following syntax:
    ```
    >>> myint = 5
    >>> print(myint)
    5
    ```

* To define a floating point number, use the following syntax:
    ```
    >>> myfloat = 5.0
    >>> print(myfloat)
    5.0
    >>> myfloat = float(5)
    >>> print(myfloat)
    5.0
    ```

__Strings__ Strings are defined either with a single quote or a double quotes. Use the following syntax to define a string:
    ```
    >>> mystring = 'hello'
    >>> print(mystring)
    hello
    >>> mystring = "hello"
    >>> print(mystring)
    hello
    ```

    The difference between using the quotes is that using double quotes makes it easy to includes apostrophes.
    ```
    >>> mystring = "Don't worry about apostrophes"
    >>> print(mystring)
    Don't worry about apostrophes
    ```
    There are additional variations on defining strings that make it easier to include things such as carriage returns, backslashes and Unicode characters.

__Lists__ Lists are very similar to arrays. They can contain any type of variable, and they can contain as many variables as you wish. Lists can also be iterated over in a very simple manner. Here is an example of how to define and iterate a list.

## Assignments
* Assignments can be done on more than one variable "simultaneously" on the same line like this
    ```
    >>> a, b = 3, 4
    >>> print(a,b)
    ```

## Operators
* Simple operators can be executed on numbers and strings:
    ```
    >>> one = 1
    >>> two = 2
    >>> three = one + two
    >>> print(three)
    3

    >>> a = "hello"
    >>> b = "world"
    >>> c = a + b
    >>> print(c)
    hello world
    ```

* Mixing operators between numbers and strings is not supported:
    ```
    >>> a, b, c = 1, 2, "hello"
    >>> print(a + b + c)
    Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
    TypeError: unsupported operand type(s) for +: 'int' and 'str'
    ```