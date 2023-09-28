# Iterator Design Pattern

The Iterator design pattern is a behavioral pattern that provides a way to access elements of an aggregate object (a collection) sequentially without exposing its underlying representation. It offers a standardized way to iterate through various data structures.

## Why Use Iterator?

- **Encapsulation**: Iterator pattern encapsulates the traversal of a collection, abstracting away the complexity of iterating over different data structures.

- **Separation of Concerns**: It separates the client code (the code using the iterator) from the collection's implementation. Clients don't need to know about the inner workings of the collection.

- **Flexibility**: Iterator allows different types of iterators to be used with the same collection, promoting flexibility and code reusability.

- **Clean Code**: It results in cleaner, more readable code by removing the need for repetitive and error-prone iteration logic in client code.

## When to Use Iterator?

- Use Iterator when you have a collection of objects, and you want to provide a way to traverse the elements sequentially without exposing the internal structure of the collection.

- It's useful when you want to support multiple types of traversal (e.g., forward, backward) of a collection without modifying its code.

- Use it when you want to decouple the client code from the specifics of how elements are stored or organized in the collection.

## Real-World Examples

1. **Java Collections Framework**: The Java Collections Framework uses iterators extensively to traverse collections like lists, sets, and maps.

2. **C# LINQ**: In C#, Language Integrated Query (LINQ) allows querying collections using a fluent syntax. It relies on iterators to process data in collections.

3. **Python's Iterators**: Python provides built-in support for iterators using the `iter()` and `next()` functions. It's used for iteration over lists, dictionaries, and other iterable objects.

4. **File System Iteration**: When scanning a file system for files and directories, an iterator can be used to traverse the directory tree without exposing the underlying file structure.

## Code Description (Iterator)

In the provided Go code example:

- `Iterator` is an interface that defines methods for checking if there are more items (`hasNext()`) and getting the next item (`next()`) in the collection.

- `MenuIterator` is a concrete iterator that implements the `Iterator` interface. It iterates through a collection of menu items (`MenuItem`) in a menu.

- `Menu` is a concrete aggregate that holds a collection of menu items. It provides a method to create an iterator for the menu items.

- `MenuItem` represents an individual menu item with attributes like name, description, and price.

The code demonstrates how to implement the Iterator pattern to traverse a collection of menu items, ensuring that the underlying structure of the collection is encapsulated, and client code can iterate through the items without knowing the details of the menu's internal representation.  
