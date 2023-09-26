# Command Design Pattern

The Command design pattern is a behavioral pattern that provides a way to encapsulate and decouple the sender of a request from the receiver, allowing you to parameterize objects with operations, delay or queue requests, and support undoable operations. This pattern is particularly useful when you want to encapsulate a request as an object, allowing for dynamic parameterization of clients and support for various operations without changing their code.

## Why Use Command?

- **Separation of Concerns**: Command pattern separates the sender and receiver of a request, preventing tight coupling between them. This leads to more maintainable and flexible code.

- **Undo/Redo Functionality**: It allows for the implementation of undo and redo operations easily by storing the state of executed commands.

- **Queueing and Scheduling**: Commands can be queued or scheduled for execution, providing control over when and how commands are processed.

- **Extensibility**: New commands can be added without modifying existing client code, making it easy to extend the functionality.

## When to Use Command?

- Use Command when you need to parameterize objects with operations in a flexible and decoupled way.

- Use Command when you want to support undoable operations or transactional behavior.

- Use Command when you want to implement a request-processing pipeline or scheduling system.

## Real-World Examples

1. **Remote Control**: A universal remote control can be designed using the Command pattern. Each button press represents a command that can be executed on various devices like TVs, DVD players, or sound systems.

2. **Text Editor**: In a text editor, commands like "Cut," "Copy," "Paste," and "Undo" can be implemented as Command objects. This allows users to execute these operations and undo/redo them.

3. **Job Scheduling Systems**: Command pattern can be used in job scheduling systems where different tasks or jobs can be scheduled for execution at specified times or in response to events.

4. **Game Input Handling**: In video games, player input commands (e.g., move, jump, attack) can be implemented using the Command pattern, allowing for easy customization of controls.

## Implementation

In the provided Go code example, we implement the Command design pattern for controlling a light. Three command classes (`TurnOnCommand`, `TurnOffCommand`, and `ChangeBrightCommand`) encapsulate the actions of turning the light on, turning it off, and changing its brightness. These commands are executed on a `Light` object.

## Code Description

- The `ILightCommand` interface defines a common `execute` method for all concrete command objects. This allows commands to be executed uniformly.

- The `Light` struct represents the receiver in the Command pattern, which is the object that the commands act upon. It has methods like `On`, `Off`, and `ChangeBright`.

- Concrete command classes (`TurnOnCommand`, `TurnOffCommand`, and `ChangeBrightCommand`) implement the `ILightCommand` interface and provide specific behavior for each command.

- Clients can create and execute these commands, decoupling the client code from the details of the `Light` object's implementation.

This code demonstrates how to implement the Command pattern to control the behavior of a `Light` object without directly coupling the client code to the `Light`'s methods. It promotes encapsulation, flexibility, and ease of extension.
