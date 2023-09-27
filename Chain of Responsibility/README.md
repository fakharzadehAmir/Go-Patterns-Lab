# Chain of Responsibility Design Pattern

The Chain of Responsibility design pattern is a behavioral pattern that allows an object to pass a request along a chain of handlers. Each handler decides either to process the request or to pass it to the next handler in the chain. This pattern is particularly useful when there are multiple handlers for a request, and the handler that can fulfill the request is not known in advance.

## Why Use Chain of Responsibility?

- **Decoupling**: Chain of Responsibility promotes loose coupling between senders and receivers of requests. The sender doesn't need to know the exact handler; it just passes the request to the first handler in the chain.

- **Flexibility**: It allows you to add or remove handlers dynamically without changing the client code. This makes the system more flexible and extensible.

- **Single Responsibility Principle**: Each handler has a single responsibility, making the code easier to understand and maintain.

- **Reduces Conditional Statements**: Instead of having nested conditional statements to determine the appropriate handler, the responsibility is delegated to separate handler objects.

## When to Use Chain of Responsibility?

- Use Chain of Responsibility when there are multiple objects that can handle a request, and the client is unsure which object will handle it.

- Use it when you want to decouple the sender and receiver of a request, allowing for more flexibility and scalability.

- It's useful when you want to avoid a "fat" class with too many conditional statements for handling different requests.

## Real-World Examples

1. **Middleware in Web Frameworks**: Many web frameworks use the Chain of Responsibility pattern to process HTTP requests. Each middleware component handles a specific aspect, such as authentication, logging, or compression, and passes the request to the next middleware.

2. **Event Handling in GUIs**: Graphical user interface (GUI) libraries often employ this pattern for handling user events like mouse clicks and keyboard input. The event is passed through a chain of event handlers until it's processed or ignored.

3. **Logging Systems**: As demonstrated in the code example, logging systems can use the Chain of Responsibility pattern to handle log messages based on their severity levels. Handlers can filter and route logs to different destinations like console, file, or a remote server.

4. **Approval Workflows**: In business applications, approval workflows for documents or transactions can be modeled using this pattern. Each step in the approval process is a handler, and the request flows through the chain until it's approved or rejected.

## Code Description (Chain of Responsibility)

In the provided Go code example:

- `LogLevel` represents severity levels for logging: `INFO`, `WARN`, and `ERR`.

- `ILogger` is an interface defining the methods `LogMessage` and `setNextLogger`, which logger classes must implement.

- `ConsoleLogger` and `FileLogger` are concrete logger classes that implement the `ILogger` interface. They handle log messages based on severity and pass them to the next logger in the chain.

The code demonstrates how log messages can be processed through a chain of responsibility based on their severity levels. It showcases the decoupling of loggers, making it easy to extend and configure the logging behavior.