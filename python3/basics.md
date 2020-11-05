# Python Basics
Welcome to the basics of Python Programming.

## Contents
* Hello, World!
* Indentation
* Variables and Types
* Assignments
* Operators
* String Formatting
* String Operations

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

__Lists__ Lists are very similar to arrays. They can contain any type of variable, and they can contain as many variables as you wish. List index starts from 0. Lists can also be iterated over in a very simple manner. To define, add items and iterate through a list, use the following code as reference:
```
>>> mylist = [0]
>>> mylist.append(1)
>>> mylist.append("2")
>>> mylist.append(3.4)
>>> print(mylist[0])
1
>>> print(mylist[1])
2
>>> print(mylist[2])
3.4
>>> for x in mylist:
...     print(x)
0
1
2
3.4
```

__Tuples__ A tuple is a fixed size list.

## Assignments
* Assignments can be done on more than one variable "simultaneously" on the same line like this
    ```
    >>> a, b = 3, 4
    >>> print(a,b)
    ```

## Operators
* __Arithmetic Operators__ Just as any other programming languages, the addition, subtraction, multiplication, and division operators can be used with numbers.
    ```
    >>> number = 1 + 2 * 3 / 4.0 - 0.5
    >>> print(number)
    2.0
    ```

* Another operator available is the __modulo (%) operator__, which returns the integer remainder of the division `dividend % divisor = remainder`.
    ```
    >>> remainder = 11 % 3
    >>> print(remainder)
    2
    ```

* Use two consecutive multiplication symbol for a power operator.
    ```
    >>> square = 4 ** 2
    >>> cube = 4 ** 3
    >>> print(square)
    16
    >>> print(cube)
    64
    ```

* Simple operators can be executed on numbers and strings. Use addition to concatenate strings.
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

* Python also supports multiplying strings to form a string with a repeating sequence.
    ```
    >>> ones = "1" * 10
    >>> print(ones)
    1111111111
    ```

* Just as in strings, Python supports forming new lists with a repeating sequence using the multiplication operator.
    ```
    >>> print([1,2,3] * 3)
    [1, 2, 3, 1, 2, 3, 1, 2, 3]
    ```

* Lists can be joined with the addition operators.
    ```
    >>> first_list = [1,2,3,4,5]
    >>> second_list = [6,7,8,9,10]
    >>> full_list = first_list + second_list
    >>> print(full_list)
    [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    ```

* Mixing operators between numbers and strings is not supported:
    ```
    >>> a, b, c = 1, 2, "hello"
    >>> print(a + b + c)
    Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
    TypeError: unsupported operand type(s) for +: 'int' and 'str'
    ```

## String Formatting
Python uses C-style string formatting to create new, formatted strings. The "%" operator is used to format a set of variables, together with a format string, which contains normal text together with "argument specifiers", special symbols like "%s" and "%d".

* __repr__ method

* __string__ method

* 
    ```
    >>> world = "World!"
    >>> print("Hello, %s" % world)
    Hello, World!
    >>> name, age = "Bob", 23
    >>> print("%s is %d years old." % (name, age))
    Bob is 23 years old.
    ```

* Any object which is not a string can be formatted using the %s operator as well. The string which returns from the *repr* method of that object is formatted as the string. For example:
    ```
    >>> mylist = [1,2,3]
    >>> print("A list: %s" % mylist)
    A list: [1, 2, 3]
    ```

* Here are some basic argument specifiers to know about.
    ```
    %s - String (or any object with a string representation, like numbers)
    %d - Integers
    %f - Floating point numbers
    %.<number of digits>f - Floating point numbers with a fixed amount of digits to the right of the dot.
    %x/%X - Integers in hex representation (lowercase/uppercase)
    ```

## Basic String Operations