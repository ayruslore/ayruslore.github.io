# Python Advanced

## Contents
* Classes and Objects
* *else* clause in loops

## Classes and Objects
* Objects are an encapsulation of variables and functions into a single entity. Objects get their variables and functions from classes. Classes are essentially a template to create your objects.

* `self` keyword is passed as a parameter to the methods of an class. This is passed in the declaration so as to access the current attibutes of the instance created.

* Here's an example of a class
    ```python
    >>> class Circle:
    ...     radius = 5
    ...     def area(self):
    ...             print(3.14 * self.radius * self.radius)
    ...     def circumference(self):
    ...             print(2 * 3.14 * self.radius)
    >>> c1 = Circle()
    >>> c2 = Circle()
    >>> c1.area()
    78.5
    >>> c1.circumference()
    31.400000000000002
    >>> c2.radius = 4
    >>> c2.area()
    50.24
    >>> c2.circumference()
    25.12
    >>> print(c1.radius, " ", c2.radius)
    5   4
    ```
    In the above example we declared a class named `Circle`. We have created two objects of the `Circle` class called `c1` and `c2`. To call the methods you can use `c1.area()`. To access an attribute of the object, you can do that using `c2.radius`.

## *else* clause in loops
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