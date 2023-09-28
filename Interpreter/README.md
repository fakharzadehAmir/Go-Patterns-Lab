# Interpreter Design Pattern

The Interpreter design pattern is a behavioral pattern that focuses on interpreting and evaluating domain-specific languages or expressions. It provides a way to define a language grammar and interpret sentences in that language. This pattern is particularly useful when you need to process and evaluate textual or symbolic expressions.

## Why Use Interpreter?

- **Custom Language Support**: Interpreter allows you to create and support custom languages or expression syntax within your application. This is valuable when you have specific rules or commands to be executed.

- **Extensibility**: It makes it easier to extend or modify the language by adding new grammar rules or expressions, without changing the core codebase.

- **Domain-Specific Language (DSL)**: Interpreter is well-suited for developing domain-specific languages tailored to particular problem domains, such as mathematical expressions, regular expressions, or configuration files.

- **Code Reusability**: By implementing an interpreter, you can reuse and share the parsing and evaluation logic across different parts of your application.

## When to Use Interpreter?

- Use Interpreter when you have a language or syntax to interpret, and you want to define its grammar and evaluate expressions or statements written in that language.

- Use it when you want to separate the parsing and interpretation concerns from the rest of the code, making it easier to manage and maintain.

- It's beneficial when you need to represent complex, hierarchical structures like abstract syntax trees (ASTs) for expressions.

## Real-World Examples

1. **Regular Expression Engines**: Many programming languages and tools use the Interpreter pattern to implement regular expression engines. The regular expressions are interpreted and matched against strings to perform pattern matching.

2. **SQL Query Execution**: Database systems use the Interpreter pattern to parse and execute SQL queries. The SQL query is interpreted, optimized, and executed to retrieve data from the database.

3. **Scripting Languages**: Scripting languages like JavaScript, Python, and Ruby use interpreters to execute code written in these languages. The interpreter parses and executes the script's statements.

4. **Mathematical Expression Evaluators**: Applications that evaluate mathematical expressions, such as calculators or mathematical modeling tools, use interpreters to parse and evaluate expressions like "2 + 3 * (4 - 1)."

## Code Description (Interpreter)

In the provided Go code example:

- `Expression` is defined as the interface that all expressions must implement. It requires the `interpret()` method to evaluate expressions and return an integer result.

- `Number` represents terminal expressions, holding numerical values.

- `Operation` represents non-terminal expressions that combine two sub-expressions using an operator. It evaluates expressions based on the operator.

- The code demonstrates a simple arithmetic expression interpreter capable of parsing and evaluating expressions like "2 + 3 - 1."

This code showcases the fundamental structure of an interpreter for basic arithmetic expressions, emphasizing the separation of terminal and non-terminal expressions and the evaluation process.
