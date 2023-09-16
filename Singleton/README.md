# Singleton Design Pattern

The Singleton design pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance. It's used to restrict the instantiation of a class to a single object and ensures that there's a single point of control for that class within an application.

**Why Use Singleton**:
- **Resource Management**: Singleton is useful when you want to manage a shared resource, such as a database connection, a configuration manager, or a thread pool, in a centralized manner to prevent unnecessary duplication.
- **Global Access**: It provides a global and consistent way to access a single instance of a class, making it convenient for different parts of an application to interact with that instance.

**Real Examples of Singleton**:

1. **Database Connection Pooling**:
    - In a web application, you often need a connection pool to interact with a database efficiently.
    - Using the Singleton pattern, you can ensure that there's a single, centralized connection pool instance that manages database connections.
    - This prevents multiple instances of the pool from being created, which would be wasteful and less efficient.

2. **Logging Service**:
    - In a large-scale application, you might want to centralize your logging functionality to record and monitor application events.
    - A Singleton Logger class can handle all logging operations and maintain a single log file or send logs to a centralized logging service.
    - This ensures that log entries from different parts of the application are consolidated and managed in a consistent manner.

## Implementation

In this package, the Singleton design pattern is implemented using Go's sync.Once to ensure that the Singleton instance is created only once. Here's a brief review of your code:

- `DpSingleton`: The Singleton struct that holds data.
- `GetInstance(newData string) *DpSingleton`: A function that returns the Singleton instance. It accepts a parameter `newData` to initialize the instance's data.

The Singleton pattern is employed to ensure that there's a single, centralized instance of `DpSingleton` within your application.

## Usage

To use the Singleton pattern, you can call the `GetInstance` function to obtain the Singleton instance. The instance will be created lazily the first time it's requested, and subsequent calls to `GetInstance` will return the same instance.

The `GetInstance` function returns the same instance with the provided data on each call, ensuring that there's only one instance of `DpSingleton` throughout your application.

The package includes test cases to ensure the correctness and thread safety of the Singleton pattern implementation.


## Concurrency
The Singleton pattern ensures that even in a highly concurrent environment, there will be only one instance of the Singleton, making it suitable for scenarios where shared resources need to be controlled.