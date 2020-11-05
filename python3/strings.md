# String Operations/Formatting

## Contents
* Basic String Operations
* String Formatting

## Basic String Operations
* To get the length of a string
    ```
    >>> mystring = "Hello world!"
    >>> print(len(mystring))
    12
    ```

* To print the index of the first occurence of a character
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.index('l'))
    2
    ```

* To count the number of times a character is repeated in a string
    ```
    >>> mystring = "Hello world!"
    >>> print(mystring.count('o'))
    2
    ```

* 

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