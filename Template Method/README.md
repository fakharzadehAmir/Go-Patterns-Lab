# Template Method Design Pattern

The Template Method design pattern is a behavioral pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure. It promotes code reusability by allowing different implementations of specific parts of an algorithm while keeping the overall algorithm intact.

## Why Use Template Method?

- **Code Reusability**: Template Method encourages reusing common code in the superclass while allowing variations in subclasses. This reduces code duplication.

- **Consistency**: It enforces a consistent algorithm structure across multiple implementations, ensuring that key steps are executed in a predefined order.

- **Extensibility**: Subclasses can extend and override specific parts of the algorithm to add new functionality or modify behavior as needed.

- **Maintenance**: Changes to the algorithm can be made in a single place (the superclass), making it easier to maintain.

## When to Use Template Method?

- Use Template Method when you have an algorithm with a fixed structure but multiple variations in its steps.

- Use Template Method when you want to define a common interface for a family of algorithms while allowing individual algorithms to vary.

- Use Template Method when you want to avoid code duplication by moving shared code to a superclass.

## Real-World Examples

1. **Beverage Preparation**: A coffee or tea preparation system can use the Template Method. The overall process (boiling water, adding ingredients, etc.) remains the same, but the specific ingredients and their proportions can vary for different beverages.

2. **Software Build Systems**: Build systems like Make or Ant use the Template Method to define the build process. The general steps (compilation, linking, etc.) remain consistent, but the specific commands and dependencies may change.

3. **Game Development**: Game development frameworks often use the Template Method to define the game loop. The loop structure is fixed (update, render, etc.), but the game logic within those steps can vary.

4. **Document Conversion**: In a document conversion system, the conversion process can follow a template. The steps for reading, processing, and writing documents are consistent, but the handling of different document formats can vary.

## Implementation

In the provided Go code example, we implement a one-time password (OTP) generation and notification system using the Template Method design pattern. The `OneTimePass` struct defines the template method `generateAndSendPassword`, while concrete OTP classes (`SMS`, `Email`, `PhoneCall`) implement specific steps (`getRandomPass`, `cacheOTP`, `getMessage`, `Notify`) for different delivery methods.

This code demonstrates how to use the Template Method pattern to maintain a consistent OTP generation and notification process while allowing flexibility in how OTPs are delivered to users via SMS, email, or phone call.

## Code Description

- The `IOneTimePassword` interface defines the methods required for OTP generation, caching, message formatting, and notification.

- The `OneTimePass` struct contains the template method `generateAndSendPassword`, which orchestrates the OTP generation and notification process.

- Concrete OTP classes (`SMS`, `Email`, `PhoneCall`) extend the `OneTimePass` struct and implement the specific steps for their respective delivery methods.

By following this pattern, we achieve a consistent OTP generation and notification process while allowing for extensibility in adding new delivery methods or modifying existing ones.
