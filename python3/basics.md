---
layout: page
permalink: /python/basics
---
[Go Next to Python Advanced](/python/advanced)

<br>

<h1><b>Python Basics</b></h1>
Welcome to the basics part of Python Programming

<h1><b>Table Of Contents</b></h1>

* TOC
{:toc}

### Hello, World!
Python is a very simple language, and has a very straightforward syntax. The simplest directive in Python is the `print` directive - it simply prints out a line (and also includes a newline, unlike in C). To print a String, just write:
```python
>>> print("hello, world")
hello, world
```

### Indentation
Python uses indentation for blocks, instead of curly braces. Both tabs and spaces are supported, but standard indentation requires four spaces. For example:
```python
>>> if True:
...     print("True")
True
```

### Naming Variables in Python
Rules for naming Python variables:
* A variable name must start with a letter or the underscore character
* A variable name cannot start with a number
* A variable name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
* Variable names are case-sensitive (age, Age and AGE are three different variables)

```python
>>> #Legal variable names:
>>> myvar = "John"
>>> my_var = "John" # we are going to stick to this pattern throughout the tutorial
>>> _my_var = "John"
>>> myVar = "John"
>>> MYVAR = "John"
>>> myvar2 = "John"
>>> 
>>> #Illegal variable names:
>>> 2myvar = "John"
  File "<stdin>", line 1
    2myvar = "John"
     ^
SyntaxError: invalid syntax
>>> my-var = "John"
  File "<stdin>", line 1
SyntaxError: cannot assign to operator
>>> my var = "John"
  File "<stdin>", line 1
    my var = "John"
       ^
SyntaxError: invalid syntax
```

### Variables and Types
Python has no command for declaring a variable, it is created the moment you first assign a value to it. Variables do not need to be declared with any particular type, and can even change type after they have been set.
```python
>>> x = 6
>>> type(x)
<class 'int'>
>>> x = 'John'
>>> type(x)
<class 'str'>
```
The `type` function is used to find out the type of a variable. Let's take a look at the different types of variables in Python.
* #### Numbers
    Python supports three types of numbers - integers, floating point numbers, complex numbers.
    * To define an integer, use the following code as reference. It's type is known as an `int`
        ```python
        >>> my_int = 5
        >>> print(my_int)
        5
        >>> type(my_int)
        <class `int`>
        ```
    * To define a floating point number, use the following code as reference. It's type is known as a `float`
        ```python
        >>> my_float = 5.0
        >>> print(my_float)
        5.0
        >>> type(my_float)
        <class `float`>
        >>> my_float = float(5)
        >>> print(my_float)
        5.0
        >>> type(my_float)
        <class `float`>
        ```
    * To define a complex number, use the following code as reference. It's type is know as `complex`
        ```python
        >>> c_1 = 3 + 6j
        >>> c_2 = complex(4,7)
        >>> type(c_1)
        <class `complex`>
        >>> type(c_2)
        <class `complex`>
        >>> c_1
        (3+6j)
        >>> c_2
        (4+7j)
        >>> # Each complex has two parts real and imaginary.
        >>> # Hence the type complex offers two attributes - real, imag
        >>> # It also offers a method - conjugate()
        >>> # It returns the conjugate value
        >>> c_1.real
        3.0
        >>> c_1.imag
        6.0
        >>> c_2.real
        4.0
        >>> c_2.imag
        7.0
        >>> c_1.conjugate()
        (3-6j)
        >>> c_2.conjugate()
        (4-7j)
        ```
        __Note :__ Normal addition, subtraction, multiplication and division of complex numbers is supported, but they do not support comparison operators like `<`, `>`, `<=`, `>=` and it will throw a *TypeError*.

* #### Strings
    They are defined either with a single quote or a double quotes. Use the following code as reference to define a string. It's type is known as `str`
    ```python
    >>> my_string = 'hello'
    >>> print(my_string)
    hello
    >>> type(my_string)
    <class `str`>
    >>> my_string = "hello"
    >>> print(my_string)
    hello
    ```
    The difference between using the quotes is that using double quotes makes it easy to includes apostrophes.
    ```python
    >>> my_string = "Don't worry about apostrophes"
    >>> print(my_string)
    Don't worry about apostrophes
    ```
    There are additional variations on defining strings that make it easier to include things such as carriage returns, backslashes and Unicode characters.

* #### Lists
    They are very similar to arrays. They can contain any type of variable, and they can contain as many variables as you wish. List index starts from 0. Lists can also be iterated over in a very simple manner. To define, add items and iterate through a list, use the following code as reference. It's type is known as a `list`.
    ```python
    >>> my_list = [0] # initializes a list with one element
    >>> # now let's add more elements to the list using the append function
    >>> # The append function adds elements at the end of the list
    >>> type(my_list)
    <class `list`>
    >>> my_list.append(1)
    >>> my_list.append("2")
    >>> my_list.append(3.4)
    >>> print(my_list)
    [0, 1, '2', 3.4]
    >>> # accessing the list, indexing starts at 0
    >>> print(my_list[0])
    1
    >>> print(my_list[1])
    2
    >>> print(my_list[2])
    3.4
    >>> for x in my_list:
    ...     print(x)
    0
    1
    2
    3.4
    ```

* #### Tuples
    A tuple is a collection of objects which is ordered and immutable. Tuples are sequences, just like lists. The differences between tuples and lists are, the tuples cannot be changed unlike lists and tuples use parentheses, whereas lists use square brackets. To define a tuple and perform operations on it use the following code as reference. It's type is known as a `tuple`.
    ```python
    >>> # you can create a tuple in the following way
    >>> tuple_1 = ('a', 'b', 1, 2)
    >>> type(tuple_1)
    <class `tuple`>
    >>> tuple_2 = (1,2,3,4,5)
    >>> tuple_3 = "a","b","c"
    >>> empty_tuple = ()
    >>> tuple_with_one_value = (10,)
    >>> # access value in tuple
    >>> tuple_2[2]
    3
    >>> # changing the values of a tuple is not allowed
    >>> tuple_2[2] = 4
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    TypeError: 'tuple' object does not support item assignment
    >>> tuple_4 = tuple_1 + tuple_2 + tuple_3
    >>> tuple_4
    ('a', 'b', 1, 2, 1, 2, 3, 4, 5, 'a', 'b', 'c')
    >>> # removing an element in tuple is not allowed
    >>> del tuple_4
    >>> # this deletes the tuple itself
    >>> tuple_4
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    NameError: name 'tuple_4' is not defined
    >>> # tuples support length, concatenation, repitition, membership and iteration
    >>> # built-in tuple functions
    >>> # 1. len - gives total length of tuple
    >>> len(tuple_2)
    5
    >>> # 2. max - returns item in the tuple with max value
    >>> max(tuple_2)
    5
    >>> max(tuple_1)
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    TypeError: '>' not supported between instances of 'int' and 'str'
    >>> # 3. min - returns item in the tuple with min value
    >>> min(tuple_2)
    1
    >>> min(tuple_3)
    'a'
    >>> # 4. tuple - converts a list to a tuple
    >>> tuple([1,2,3])
    (1, 2, 3)
    ```

* #### Dictionaries
    They are similar to list and a tuple, but it works with key value pairs instead of indexes. Every value is associated with a key. A dictionary is not ordered and the keys are not associated with an index, hence indexing does not work in dictionaries. A Key or Value can be of any type. Let us consider a dictionary with students names as values and their roll numbers as keys.
    ```python
    >>> #empty dictionary
    >>> rollbook = {}
    >>> type(rollbook)
    <class `dict`>
    >>> #dictionary initialized with somevalues
    >>> rollbook = {1: "sweety", 2: "preeti"}
    >>> #now we'll add more values in to the dictionary
    >>> rollbook[3] = "arjun"
    >>> rollbook[4] = "lakshman"
    >>> #iterate through the dictionaries
    >>> for roll, name in rollbook.items():
    ...     print("roll number %d is %s" % (roll, name))
    roll number 1 is sweety
    roll number 2 is preeti
    roll number 3 is arjun
    roll number 4 is lakshman
    >>> #remove a value
    >>> del rollbook[1]
    >>> print(rollbook)
    {2: 'preeti', 3: 'arjun', 4: 'lakshman'}
    >>> #can also use pop method to remove a value
    >>> rollbook.pop(2)
    'preeti'
    >>> print(rollbook)
    {3: 'arjun', 4: 'lakshman'}
    ```

* #### Sets
    They are just lists with no duplicate entries. The `set` function is used to create a set from a list. Use the following for reference for creating a set
    ```python
    >>> numbers = [1,2,3,4,5,1,2,3,4,5]
    >>> set(numbers)
    {1, 2, 3, 4, 5}
    >>> nums = set(numbers)
    >>> type(nums)
    <class `set`>
    >>> # accessing a set is not as the same as list
    >>> nums[0]
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    TypeError: 'set' object is not subscriptable
    >>> for i in nums:
    ...     print(i)
    1
    2
    3
    4
    5
    ```
    Sets are a powerful tool in Python since they have the ability to calculate differences and intersections between other sets, let's look at an example
    ```python
    >>> set_a = set([1,2,3])
    >>> set_b = set([2,4])
    >>> # intersection : members attended both events
    >>> print(set_a.intersection(set_b))
    {2}
    >>> print(set_b.intersection(set_a))
    {2}
    >>> # symmetric_difference : members attended only one of the events
    >>> print(set_a.symmetric_difference(set_b))
    {1, 3, 4}
    >>> print(set_b.symmetric_difference(set_a))
    {1, 3, 4}
    >>> # difference : members attended only one event and not the other
    >>> print(set_a.difference(set_b))
    {1, 3}
    >>> print(set_b.difference(set_a))
    {4}
    >>> # union : receive a list of all participants
    >>> print(set_a.union(set_b))
    {1, 2, 3, 4}
    >>> print(set_b.union(set_a))
    {1, 2, 3, 4}
    ```

### Assignments
Assignments can be done on more than one variable "simultaneously" on the same line like this
```python
>>> a, b = 3, 4
>>> print(a,b)
```

### Arithmetic Operators
* Just as any other programming languages, the addition, subtraction, multiplication, and division operators can be used with numbers
    ```python
    >>> number = 1 + 2 * 3 / 4.0 - 0.5
    >>> print(number)
    2.0
    ```
* Another operator available is the __modulo operator %__, which returns the integer remainder of the division `dividend % divisor = remainder`
    ```python
    >>> remainder = 11 % 3
    >>> print(remainder)
    2
    ```
* Use two consecutive multiplication symbols for a power operator
    ```python
    >>> square = 4 ** 2
    >>> cube = 4 ** 3
    >>> print(square)
    16
    >>> print(cube)
    64
    ```
* Python follows order of operations :- When more than one operator appears in an expression, the order of evaluation depends on the rules of precedence. Python follows the same precedence rules for its mathematical operators that mathematics does.
    * Parentheses have the highest precedence and can be used to force an expression to evaluate in the order you want. Since expressions in parentheses are evaluated first, `2 * (3-1) is 4`, and `(1+1)**(5-2) is 8`. You can also use parentheses to make an expression easier to read, as in `(minute * 100) / 60`, even though it doesn’t change the result.
    * Exponentiation has the next highest precedence, so `2**1+1 is 3 and not 4`, and `3*1**3 is 3 and not 27`.
    * Multiplication and both division operators have the same precedence, which is higher than addition and subtraction, which also have the same precedence. So `2*3-1 yields 5 rather than 4`, and `5-2*2 is 1, not 6`.
    * Operators with the same precedence (except for **) are evaluated from left-to-right. In algebra we say they are left-associative. So in the expression `6-3+2`, the subtraction happens first, yielding 3. We then add 2 to get the result 5. If the operations had been evaluated from right to left, the result would have been `6-(3+2)`, which is 1.
* Simple operators can be executed on numbers and strings. Use addition to concatenate strings
    ```python
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
* Python also supports multiplying strings to form a string with a repeating sequence
    ```python
    >>> ones = "1" * 10
    >>> print(ones)
    1111111111
    ```
* Just as in strings, Python supports forming new lists with a repeating sequence using the multiplication operator
    ```python
    >>> print([1,2,3] * 3)
    [1, 2, 3, 1, 2, 3, 1, 2, 3]
    ```
* Lists can be joined with the addition operators
    ```python
    >>> first_list = [1,2,3,4,5]
    >>> second_list = [6,7,8,9,10]
    >>> full_list = first_list + second_list
    >>> print(full_list)
    [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    ```
* Mixing operators between numbers and strings is not supported - it throws a `TypeError`
    ```python
    >>> a, b, c = 1, 2, "hello"
    >>> print(a + b + c)
    Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
    TypeError: unsupported operand type(s) for +: 'int' and 'str'
    ```

### Basic String Operations
* To get the length of a string - use the `len` function
    ```python
    >>> my_string = "Hello world!"
    >>> print(len(my_string))
    12
    ```
* To print the index of the first occurence of a character - use the `index` method
    ```python
    >>> my_string = "Hello world!"
    >>> print(my_string.index('l'))
    2
    ```
* To count the number of times a character is repeated in a string - use the `count` method
    ```python
    >>> my_string = "Hello world!"
    >>> print(my_string.count('o'))
    2
    ```
* To print a slice of the string, the syntax is of the form `string[i:j]` - this prints the slice starting from index `i` and ends at index `j-1` or ends at `j`th character. __Note :__ If you just have one number in the brackets, it will give you the single character at that index. If you leave out the first number but keep the colon, it will give you a slice from the start to the number you left in.If you leave out the second number, it will give you a slice from the first number to the end. You can even put negative numbers inside the brackets. They are an easy way of starting at the end of the string instead of the beginning. This way, -4 means "4th character from the end"
    ```python
    >>> my_string = "Hello world!"
    >>> print(my_string[2:9])
    llo wor
    ```
* Extended slice syntax `[start:stop:step]`, the step skips characters. Let's skip one character and see
    ```python
    >>> my_string = "Hello world!"
    >>> print(my_string[2:11:2])
    lowrd
    ```
* To reverse a string
    ```python
    >>> my_string = "Hello world"
    >>> print(my_string[::-1])
    dlrow olleH
    ```
* To convert a string to all uppercase characters and lowercase characters, use the `upper` and `lower` methods
    ```python
    >>> my_string = "Hello world"
    >>> print(my_string.upper())
    HELLO WORLD
    >>> print(my_string.lower())
    hello world
    ```
* To determine whether a string starts with something and ends with something, use the `startswith` and `endswith` methods
    ```python
    >>> my_string = "Hello world!"
    >>> print(my_string.startswith("Hello"))
    True
    >>> print(my_string.endswith("Hello"))
    False
    ```
* To split a string into a bunch of strings based on a particular character or string, use the `split` method. It thows a `ValueError` if an empty seperator is provided
    ```python
    >>> my_string = "Hello world!"
    >>> print(my_string.split(" "))
    ['Hello', 'world!']
    >>> print(my_string.split(""))
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    ValueError: empty separator
    ```

### String Formatting
Python uses C-style string formatting to create new, formatted strings. The "%" operator is used to format a set of variables, together with a format string, which contains normal text together with "argument specifiers", special symbols like "%s" and "%d".

* __*repr*__ method __*TODO*__

* __*str*__ method __*TODO*__

* https://stackoverflow.com/questions/1984162/purpose-of-pythons-repr
* https://www.tutorialspoint.com/What-does-the-repr-function-do-in-Python-Object-Oriented-Programming
* https://www.journaldev.com/22460/python-str-repr-functions

* Let's print
    ```python
    >>> world = "World"
    >>> print("Hello, %s" % world)
    Hello, World
    >>> name, age = "Bob", 23
    >>> print("%s is %d years old." % (name, age))
    Bob is 23 years old.
    ```

* Any object which is not a string can be formatted using the %s operator as well. The string which returns from the *repr* method of that object is formatted as the string
    ```python
    >>> mylist = [1,2,3]
    >>> print("A list: %s" % mylist)
    A list: [1, 2, 3]
    ```

* Here are some basic argument specifiers to know about
    ```
    %s - String (or any object with a string representation, like numbers)
    %d - Integers
    %f - Floating point numbers
    %.<number of digits>f - Floating point numbers with a fixed amount of digits to the right of the dot.
    %x/%X - Integers in hex representation (lowercase/uppercase)
    ```

### Conditions
Python uses boolean variables to evaluate conditions. The boolean values True and False are returned when an expression is compared or evaluated.
* For example
    ```python
    >>> a = 5
    >>> print(a == 5)
    True
    >>> print(a == 1)
    False
    >>> print(a > 3)
    True
    ```
* The `if` statement follows the below syntax. `elif` is the keyword used in python instead of `else if`
    ```python
    >>> a, b = False, True
    >>> if a == True:
    ...     # do something
    ...     pass
    ... elif b == False:
    ...     # do something else
    ...     pass
    ... else:
    ...     # do another thing
    ...     pass
    ```
* __Note :__ Variable assignment is done using a single equals operator `"="`, whereas comparison between two variables is done using the double equals operator `"=="`. The `"not equals"` operator is marked as `"!="`.

### Boolean Operators
* The boolean operators in python are `and` and `or`. You can use them in the following way
    ```python
    >>> a, b = 1, 2
    >>> if a == 1 and a < 2:
    ...     print("value of a is 1 and value of a is less than 2")
    value of a is 1 and value of a is less than 2
    >>> if a == 1 or a == 2:
    ...     print("value of a is either 1 or 2")
    value of a is either 1 or 2
    ```
* The `in` operator is used to check if a specific object exists within an iterable object container, such as a list, tuple, dictionary.
    ```python
    >>> a = 5
    >>> if a in [1,3,5,7,9]:
    ...     print("the value of a is in present in the list")
    the value of a is in present in the list
    >>> a in (1,3,5,7,9)
    True
    >>> a in {1:'a', 3:'c', 5:'e', 7:'g', 9:'i'}
    True
    ```
* The `is` operator matches the instances themselves, unlike the equals operator `==` which matches the values of the variables.
    ```python
    >>> x = []
    >>> y = x
    >>> x is y
    True
    >>> x == y
    True
    >>> 
    >>> x, y = [], []
    >>> x is y
    False
    >>> x == y
    True
    ```
* The `not` operator, inverts a boolean expression.
    ```python
    >>> print(not False)
    True
    >>> a = True
    >>> not a
    False
    ```

### Loops
Two types of loops in Python, `for` and `while`.

* __*for loop*__ is used to iterate over a given sequence of objects. It's syntax is very simple, here's an example
    ```python
    >>> # iterating a list
    >>> for x in [1,2,3,4,5]:
    ...     print(x)
    1
    2
    3
    4
    5
    >>> # iterating a tuple
    >>> for x in (1,2,3,4,5):
    ...     print(x)
    1
    2
    3
    4
    5
    >>> # iterating keys in a dictionary
    >>> for x in {1:'a',2:'b',3:'c',4:'d',5:'e'}:
    ...     print(x)
    1
    2
    3
    4
    5
    >>> # iterating and accessing a dictionary
    >>> my_dictionary = {1:'a',2:'b',3:'c',4:'d',5:'e'}
    >>> for key in my_dictionary:
    ...     print(key, my_dictionary[key])
    1 a
    2 b
    3 c
    4 d
    5 e
    ```

* __*while loop*__ repeats as long as the boolean condition is true. Here's an example
    ```python
    >>> count = 0
    >>> while count < 2:
    ...     print(count)
    ...     count = count + 1
    0
    1
    ```

* *__range__ and __xrange__ functions :* range function returns a new list with numbers of that specified range, whereas xrange returns an iterator, which is more efficient. (Python 3 uses the range function, which acts like xrange. In Python 3 xrange does not exist). Note that the range function is zero based.
    ```python
    >>> range(2)
    range(0, 2)
    >>> for i in range(2):
    ...     print(i)
    0
    1
    ```

* __*break and continue statement :*__  break is used to exit a for loop or a while loop, whereas continue is used to skip the current block, and return to the "for" or "while" statement.
    ```python
    >>> count = 0
    >>> while count < 5:
    ...     print(count)
    ...     count = count + 1
    ...     if count == 2:
    ...             break
    ... 
    0
    1
    >>> for x in range(21):
    ...     if x % 4 != 0:
    ...             continue
    ...     print(x)
    0
    4
    8
    12
    16
    20
    ```

### Functions in Python
* Functions allow us to divide the code into useful blocks, make it more readable, reuse it. Functions in python are defined using the keyword `def`, followed with the function's name.
* To call the above function, just call the function name followed by curved brackets, you may also pass arguments if required.
    ```python
    >>> def my_func():
    ...     print("this is my first function")
    >>> my_func()
    this is my first function
    ```
* The above function does not return anything, the value of the function is `None`. To return a value to the caller - use the `return` keyword.
    ```python
    >>> def add(a, b):
    ...     print(a + b)
    >>> add(3,5)
    8
    >>> print(add(3,5))
    8
    None
    >>> def add2(a, b):
    ...     return a + b
    >>> add2(3,5)
    8
    >>> print(add2(3,5))
    8
    ```
* As `def` keyword defines the beginning of a block, it is known as a block keyword. Other block keywords we have seen till now are `if`, `for` and `while`.

__*Multiple Function Arguments*__
```python
>>> # Every function in Python receives a predefined number of arguments, if
>>> # declared normally, like this
>>> def my_func(a, b):
...     # do something with the two variables
...     print(a, "-", b)
>>> my_func(1,2)
1 - 2
>>> # It is possible to declare functions which receive a variable number of
>>> # arguments, using the following syntax
>>> def my_var_func(a, b, *remaining):
...     print(a, "-", b)
...     print(remaining)
>>> my_var_func(1,2)
1 - 2
()
>>> my_var_func(1, 2, 3)
1 - 2
(3,)
>>> my_var_func(1, 2, 3, 4, 5, 6)
1 - 2
(3, 4, 5, 6)
>>> # It is also possible to send functions arguments by keyword, so that the
>>> #  order of the argument does not matter, using the following syntax.
>>> def my_var_key_func(a, b, **options):
...     print(a, "-", b)
...     print(options)
>>> my_var_key_func(1,2)
1 - 2
{}
>>> my_var_key_func(1,2, c = 3, d = 4)
1 - 2
{'c': 3, 'd': 4}
>>> 
>>> def my_var_key_func_2(a, b, **options):
...     print(a, "-", b)
...     print(options)
...     # accessing the keywords
...     print('c :', options.get('c'))
...     print('d :', options.get('d'))
>>> my_var_key_func_2(1,2, c = 3, d = 4, e = 5)
1 - 2
{'c': 3, 'd': 4, 'e': 5}
c : 3
d : 4
```

__*Closures*__
A Closure is a function object that remembers values in enclosing scopes even if they are not present in memory. Firstly, a Nested Function is a function defined inside another function. It's very important to note that the nested functions can access the variables of the enclosing scope. However, at least in python, they are only readonly. However, one can use the "nonlocal" keyword explicitly with these variables in order to modify them.
```python
>>> # nested functions
>>> def outerFunction(text):
...     t = text
...     def innerFunction():
...             print(t)
...     innerFunction()
... 
>>> outerFunction('Hey!')
Hey!
```
A Closure is a function object that remembers values in enclosing scopes even if they are not present in memory.
* It is a record that stores a function together with an environment: a mapping associating each free variable of the function (variables that are used locally, but defined in an enclosing scope) with the value or reference to which the name was bound when the closure was created.
* A closure—unlike a plain function—allows the function to access those captured variables through the closure’s copies of their values or references, even when the function is invoked outside their scope.

```python
>>> # closures
>>> def outerFunction(text):
...     t = text
...     def innerFunction():
...             print(t)
...     return innerFunction
... 
>>> myFunction = outerFunction("Hey!")
>>> myFunction()
Hey!
```

__*Partial Functions*__
You can create partial functions in python by using the partial function from the functools library. Partial functions allow one to derive a function with x parameters to a function with fewer parameters and fixed values set for the more limited function.
```python
>>> from functools import partial
>>> def func(a,b,c):
...     return a + b**2 + c**3
... 
>>> func(1,2,3)
32
>>> partial_func = partial(func, 1)
>>> partial_func(2,3)
32
>>> partial_func(3,4)
74
>>> partial_func_2 = partial(func, 1, 2)
>>> 
>>> partial_func_2(3)
32
>>> partial_func_2(3, 4)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: func() takes 3 positional arguments but 4 were given
```

### JSON Serialization
Python provides built-in JSON libraries to encode and decode JSON. There are two basic formats for JSON data. Either in a string or the object datastructure. The object datastructure, in Python, consists of lists and dictionaries nested inside each other. The object datastructure allows one to use python methods (for lists and dictionaries) to add, list, search and remove elements from the datastructure. The String format is mainly used to pass the data into another program or load into a datastructure. To load JSON back to a data structure, use the `loads` method. This method takes a string and turns it back into the json object datastructure. To encode a data structure to JSON, use the `dumps` method. This method takes an object and returns a String.
```python
>>> import json
>>> json_string = '{"John":"Doe_1", "Jane":"Doe_2"}'
>>> json_object = json.loads(json_string)
>>> print(json_object)
{'John': 'Doe_1', 'Jane': 'Doe_2'}
>>> type(json_object)
<class 'dict'>
>>> json_string = ''
>>> json_string = json.dumps(json_object)
>>> print(json_string)
'{"John": "Doe_1", "Jane": "Doe_2"}'
```

### Modules and Packages
For information on building modules and packages follow the link - [Learn Python Moduling](https://www.learnpython.org/en/Modules_and_Packages)

<br>

[Go Next to Python Advanced](/python/advanced)