# Python Basics
Welcome to the basics of Python Programming.

## Contents
* Hello, World!
* Indentation
* Variables and Types
* Assignments
* Arithmetic Operators
* Basic String Operations
* String Formatting
* Conditions
* Boolean Operators
* Loops

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

* To define an integer, use the following syntax. It's type is known as an `int`
    ```
    >>> myint = 5
    >>> print(myint)
    5
    ```

* To define a floating point number, use the following syntax. It's type is known as a `float`
    ```
    >>> myfloat = 5.0
    >>> print(myfloat)
    5.0
    >>> myfloat = float(5)
    >>> print(myfloat)
    5.0
    ```

* To define a complex number __*TODO*__

__Strings__ Strings are defined either with a single quote or a double quotes. Use the following syntax to define a string. It's type is known as `str`
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

__Lists__ Lists are very similar to arrays. They can contain any type of variable, and they can contain as many variables as you wish. List index starts from 0. Lists can also be iterated over in a very simple manner. To define, add items and iterate through a list, use the following code as reference. It's type is known as a `list`.
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

__Tuples__ A tuple is a fixed size list. __*TODO*__

## Assignments
* Assignments can be done on more than one variable "simultaneously" on the same line like this
    ```
    >>> a, b = 3, 4
    >>> print(a,b)
    ```

## Arithmetic Operators
* Just as any other programming languages, the addition, subtraction, multiplication, and division operators can be used with numbers
    ```
    >>> number = 1 + 2 * 3 / 4.0 - 0.5
    >>> print(number)
    2.0
    ```

* Another operator available is the __modulo operator %__, which returns the integer remainder of the division `dividend % divisor = remainder`
    ```
    >>> remainder = 11 % 3
    >>> print(remainder)
    2
    ```

* Python follows order of operators. __*TODO*__

* Use two consecutive multiplication symbols for a power operator
    ```
    >>> square = 4 ** 2
    >>> cube = 4 ** 3
    >>> print(square)
    16
    >>> print(cube)
    64
    ```

* Simple operators can be executed on numbers and strings. Use addition to concatenate strings
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

* Python also supports multiplying strings to form a string with a repeating sequence
    ```
    >>> ones = "1" * 10
    >>> print(ones)
    1111111111
    ```

* Just as in strings, Python supports forming new lists with a repeating sequence using the multiplication operator
    ```
    >>> print([1,2,3] * 3)
    [1, 2, 3, 1, 2, 3, 1, 2, 3]
    ```

* Lists can be joined with the addition operators
    ```
    >>> first_list = [1,2,3,4,5]
    >>> second_list = [6,7,8,9,10]
    >>> full_list = first_list + second_list
    >>> print(full_list)
    [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    ```

* Mixing operators between numbers and strings is not supported - it throws a `TypeError`
    ```
    >>> a, b, c = 1, 2, "hello"
    >>> print(a + b + c)
    Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
    TypeError: unsupported operand type(s) for +: 'int' and 'str'
    ```

## Basic String Operations
* To get the length of a string - use the `len` function
    ```
    >>> mystring = "Hello world!"
    >>> print(len(mystring))
    12
    ```

* To print the index of the first occurence of a character - use the `index` method
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.index('l'))
    2
    ```

* To count the number of times a character is repeated in a string - use the `count` method
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.count('o'))
    2
    ```

* To print a slice of the string, the syntax is of the form `string[i:j]` - this prints the slice starting from index `i` and ends at index `j-1` or ends at `j`th character. `Note : If you just have one number in the brackets, it will give you the single character at that index. If you leave out the first number but keep the colon, it will give you a slice from the start to the number you left in. If you leave out the second number, it will give you a slice from the first number to the end. You can even put negative numbers inside the brackets. They are an easy way of starting at the end of the string instead of the beginning. This way, -4 means "4th character from the end"`
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring[2:9])
    llo wor
    ```

* Extended slice syntax `[start:stop:step]`, the step skips characters. Let's skip one character and see
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring[2:11:2])
    lowrd
    ```

* To reverse a string
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring[::-1])
    !dlrow olleH
    ```

* To convert a string to all uppercase characters and lowercase characters, use the `upper` and `lower` methods
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.upper())
    HELLO WORLD!
    >>> print(mystring.lower())
    hello world!
    ```

* To determine whether a string starts with something and ends with something, use the `startswith` and `endswith` methods
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.startswith("Hello"))
    True
    >>> print(mystring.endswith("Hello"))
    False
    ```

* To split a string into a bunch of strings based on a particular character or string, use the `split` method. It thows a `ValueError` if an empty seperator is provided
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.split(" "))
    ['Hello', 'world!']
    >>> print(mystring.split(""))
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    ValueError: empty separator
    ```

## String Formatting
Python uses C-style string formatting to create new, formatted strings. The "%" operator is used to format a set of variables, together with a format string, which contains normal text together with "argument specifiers", special symbols like "%s" and "%d".

* __repr__ method

* __string__ method

* Let's print
    ```
    >>> world = "World!"
    >>> print("Hello, %s" % world)
    Hello, World!
    >>> name, age = "Bob", 23
    >>> print("%s is %d years old." % (name, age))
    Bob is 23 years old.
    ```

* Any object which is not a string can be formatted using the %s operator as well. The string which returns from the *repr* method of that object is formatted as the string
    ```
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

## Conditions
Python uses boolean variables to evaluate conditions. The boolean values True and False are returned when an expression is compared or evaluated.

* For example
    ```
    >>> a = 5
    >>> print(a == 5)
    True
    >>> print(a == 1)
    False
    >>> print(a > 3)
    True
    ```

* The `if` statement follows the below syntax. `elif` is the keyword used in python instead of `else if`
    ```
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

* `Note : Variable assignment is done using a single equals operator "=", whereas comparison between two variables is done using the double equals operator "==". The "not equals" operator is marked as "!=".`

## Boolean Operators
* The boolean operators in python are `and` and `or`. You can use them in the following way
    ```
    >>> a, b = 1, 2
    >>> if a == 1 and a < 2:
    ...     print("value of a is 1 and value of a is less than 2")
    value of a is 1 and value of a is less than 2
    >>> if a == 1 or a == 2:
    ...     print("value of a is either 1 or 2")
    value of a is either 1 or 2
    ```

* The `in` operator is used to check if a specific object exists within an iterable object container, such as a list, tuple, dictionary.
    ```
    >>> a = 5
    >>> if a in [1,3,5,7,9]:
    ...     print("the value of a is in present in the list"
    the value of a is in present in the list
    ```

* The `is` operator matches the instances themselves, unlike the equals operator `==` which matches the values of the variables.
    ```
    >>> x, y = [], []
    >>> print(x == y)
    True
    >>> print(x is y)
    False
    ```

* The `not` operator, inverts a boolean expression.
    ```
    >>> print(not False)
    True
    ```

## Loops
Two types of loops in Python, `for` and `while`

* __*for* loop__

* __*while* loop__

* __*break* statement__

* __*continue* statement__