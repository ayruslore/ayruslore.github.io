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
CamelCase ------------------ __*TODO*__

### Variables and Types
As python is object oriented and not statically typed, you do not need to declare variables, or declare their type before using them. *Every variable in Python is an object.* Now let us cover the basic types of variables in python.

* __*Numbers*__ Python supports three types of numbers - integers, floating point numbers, complex numbers.

    * To define an integer, use the following syntax. It's type is known as an `int`
        ```python
        >>> myint = 5
        >>> print(myint)
        5
        ```
    
    * To define a floating point number, use the following syntax. It's type is known as a `float`
        ```python
        >>> myfloat = 5.0
        >>> print(myfloat)
        5.0
        >>> myfloat = float(5)
        >>> print(myfloat)
        5.0
        ```
    
    * To define a complex number, use the following syntax. It's type is know as `complex`
        ```python
        >>> c1 = 3 + 6j
        >>> c2 = complex(4,7)
        >>> type(c1)
        <class complex>
        >>> c1
        (3+6j)
        >>> c2
        (4+7j)
        >>> # Each complex has two parts real and imaginary.
        >>> # Hence the type complex offers two attributes - real, imag
        >>> # It also offers a method - conjugate() -- returns the conjugate value
        >>> c1.real
        3.0
        >>> c1.imag
        6.0
        >>> c2.real
        4.0
        >>> c2.imag
        7.0
        >>> c1.conjugate()
        (3-6j)
        >>> c2.conjugate()
        (4-7j)
        ```
        `Note :` Normal addition, subtraction, multiplication and division of complex numbers is supported, but they do not support comparison operators like `<`, `>`, `<=`, `>=` and it will throw a *TypeError*.

* __*Strings*__ are defined either with a single quote or a double quotes. Use the following syntax to define a string. It's type is known as `str`
    ```python
    >>> mystring = 'hello'
    >>> print(mystring)
    hello
    >>> mystring = "hello"
    >>> print(mystring)
    hello
    ```
    The difference between using the quotes is that using double quotes makes it easy to includes apostrophes.
    ```python
    >>> mystring = "Don't worry about apostrophes"
    >>> print(mystring)
    Don't worry about apostrophes
    ```
    There are additional variations on defining strings that make it easier to include things such as carriage returns, backslashes and Unicode characters.

* __*Lists*__ are very similar to arrays. They can contain any type of variable, and they can contain as many variables as you wish. List index starts from 0. Lists can also be iterated over in a very simple manner. To define, add items and iterate through a list, use the following code as reference. It's type is known as a `list`.
    ```python
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

* __*Tuples*__ A tuple is a collection of objects which ordered and immutable. Tuples are sequences, just like lists. The differences between tuples and lists are, the tuples cannot be changed unlike lists and tuples use parentheses, whereas lists use square brackets. It's type is known as a `tuple`.
    ```python
    >>> # you can create a tuple in the following way
    >>> tuple1 = ('a', 'b', 1, 2)
    >>> tuple2 = (1,2,3,4,5)
    >>> tuple3 = "a","b","c"
    >>> empty_tuple = ()
    >>> tuple_with_one_value = (10,)
    >>> # access value in tuple
    >>> tuple2[2]
    3
    >>> tuple2[2] = 4
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    TypeError: 'tuple' object does not support item assignment
    >>> tuple4 = tuple1 + tuple2 + tuple3
    >>> tuple4
    ('a', 'b', 1, 2, 1, 2, 3, 4, 5, 'a', 'b', 'c')
    >>> # removing an element in tuple is not allowed
    >>> del tuple4
    >>> # this deletes the tuple itself
    >>> tuple4
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    NameError: name 'tuple4' is not defined
    >>> # tuples support length, concatenation, repitition, membership and iteration
    >>> # built-in tuple functions
    >>> # 1. len - gives total length of tuple
    >>> len(tuple2)
    5
    >>> # 2. max - returns item in the tuple with max value
    >>> max(tuple2)
    5
    >>> max(tuple1)
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    TypeError: '>' not supported between instances of 'int' and 'str'
    >>> # 3. min - returns item in the tuple with min value
    >>> min(tuple2)
    1
    >>> min(tuple3)
    'a'
    >>> # 4. tuple - converts a list to a tuple
    >>> tuple([1,2,3])
    (1, 2, 3)
    ```

* __*Dictionaries*__ are similar to list and a tuple, but it works with key value pairs instead of indexes. Every value is associated with a key. A dictionary is not ordered and the keys are not associated with an index, hence indexing does not work in dictionaries. A Key or Value can be of any type. Let us consider a dictionary with students names as values and their roll numbers as keys.
    ```python
    >>> #empty dictionary
    >>> rollbook = {}
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

* Python follows order of operations :- When more than one operator appears in an expression, the order of evaluation depends on the rules of precedence. Python follows the same precedence rules for its mathematical operators that mathematics does.
    * Parentheses have the highest precedence and can be used to force an expression to evaluate in the order you want. Since expressions in parentheses are evaluated first, `2 * (3-1) is 4`, and `(1+1)**(5-2) is 8`. You can also use parentheses to make an expression easier to read, as in `(minute * 100) / 60`, even though it doesn’t change the result.
    * Exponentiation has the next highest precedence, so `2**1+1 is 3 and not 4`, and `3*1**3 is 3 and not 27`. Can you explain why?
    * Multiplication and both division operators have the same precedence, which is higher than addition and subtraction, which also have the same precedence. So `2*3-1 yields 5 rather than 4`, and `5-2*2 is 1, not 6`.
    * Operators with the same precedence (except for **) are evaluated from left-to-right. In algebra we say they are left-associative. So in the expression `6-3+2`, the subtraction happens first, yielding 3. We then add 2 to get the result 5. If the operations had been evaluated from right to left, the result would have been `6-(3+2)`, which is 1.

* Use two consecutive multiplication symbols for a power operator
    ```python
    >>> square = 4 ** 2
    >>> cube = 4 ** 3
    >>> print(square)
    16
    >>> print(cube)
    64
    ```

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
    >>> mystring = "Hello world!"
    >>> print(len(mystring))
    12
    ```

* To print the index of the first occurence of a character - use the `index` method
    ```python
    >>> mystring = "Hello world!"
    >>> print(mystring.index('l'))
    2
    ```

* To count the number of times a character is repeated in a string - use the `count` method
    ```python
    >>> mystring = "Hello world!"
    >>> print(mystring.count('o'))
    2
    ```

* To print a slice of the string, the syntax is of the form `string[i:j]` - this prints the slice starting from index `i` and ends at index `j-1` or ends at `j`th character.
    
    __Note :__ If you just have one number in the brackets, it will give you the single character at that index. If you leave out the first number but keep the colon, it will give you a slice from the start to the number you left in. If you leave out the second number, it will give you a slice from the first number to the end. You can even put negative numbers inside the brackets. They are an easy way of starting at the end of the string instead of the beginning. This way, -4 means "4th character from the end"
    ```python
    >>> mystring = "Hello world!"
    >>> print(mystring[2:9])
    llo wor
    ```

* Extended slice syntax `[start:stop:step]`, the step skips characters. Let's skip one character and see
    ```python
    >>> mystring = "Hello world!"
    >>> print(mystring[2:11:2])
    lowrd
    ```

* To reverse a string
    ```python
    >>> mystring = "Hello world"
    >>> print(mystring[::-1])
    dlrow olleH
    ```

* To convert a string to all uppercase characters and lowercase characters, use the `upper` and `lower` methods
    ```python
    >>> mystring = "Hello world"
    >>> print(mystring.upper())
    HELLO WORLD
    >>> print(mystring.lower())
    hello world
    ```

* To determine whether a string starts with something and ends with something, use the `startswith` and `endswith` methods
    ```python
    >>> mystring = "Hello world!"
    >>> print(mystring.startswith("Hello"))
    True
    >>> print(mystring.endswith("Hello"))
    False
    ```

* To split a string into a bunch of strings based on a particular character or string, use the `split` method. It thows a `ValueError` if an empty seperator is provided
    ```python
    >>> mystring = "Hello world!"
    >>> print(mystring.split(" "))
    ['Hello', 'world!']
    >>> print(mystring.split(""))
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
    ```

* The `is` operator matches the instances themselves, unlike the equals operator `==` which matches the values of the variables.
    ```python
    >>> x, y = [], []
    >>> print(x == y)
    True
    >>> print(x is y)
    False
    ```

* The `not` operator, inverts a boolean expression.
    ```python
    >>> print(not False)
    True
    ```

### Loops
Two types of loops in Python, `for` and `while`.

* __*for* loop__ for loop is used to iterate over a given sequence of objects. It's syntax is very simple, here's an example
    ```python
    >>> for x in [1,2,3,4,5]:
    ...     print(x)
    1
    2
    3
    4
    5
    ```

* __*while* loop__ while loop repeats as long as the boolean condition is true. Here's an example
    ```python
    >>> count = 0
    >>> while count < 2:
    ...     print(count)
    ...     count = count + 1
    0
    1
    ```

* __range__ and __xrange__ functions : range function returns a new list with numbers of that specified range, whereas xrange returns an iterator, which is more efficient. (Python 3 uses the range function, which acts like xrange. In Python 3 xrange does not exist). Note that the range function is zero based.
    ```python
    >>> range(2)
    range(0, 2)
    >>> for i in range(2):
    ...     print(i)
    0
    1
    ```

* __*break* and *continue* statement__ : break is used to exit a for loop or a while loop, whereas continue is used to skip the current block, and return to the "for" or "while" statement.
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
* As we all know that functions allow us to divide the code into useful blocks, make it more readable, reuse it. Functions in python are defined using the keyword `def`, followed with the function's name.

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

* As `def` keyword defines the beginning of a block, it is known as a block keyword. Other block keywords are `if`, `for` and `while`.

### Modules and Packages
For information on building modules and packages follow the link - [Learn Python Moduling](https://www.learnpython.org/en/Modules_and_Packages)

<br>

[Go Next to Python Advanced](/python/advanced)