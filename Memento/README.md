# Memento Design Pattern

The Memento design pattern is a behavioral pattern that allows an object to capture its internal state and externalize it, so the object can be restored to that state later. It is particularly useful when you need to provide an undo mechanism or save and restore the state of an object without exposing its internal structure. The pattern achieves this by defining three key roles: the Originator, the Memento, and the Caretaker.

## Why Use Memento?

- **State Snapshot**: Memento pattern enables the capturing of an object's internal state at a specific point in time, providing a snapshot of its data.

- **Encapsulation**: It encapsulates the object's state, ensuring that it's not accessible directly from outside, maintaining information hiding.

- **Undo and Redo**: Memento can be used to implement undo and redo functionality, allowing users to revert changes.

- **Persistence**: It's useful in scenarios where you need to save and restore an object's state, such as in text editors or games.

## When to Use Memento?

- Use Memento when you need to capture and restore an object's state, such as implementing an undo feature.

- Use Memento when you want to encapsulate the object's state, preventing direct access from outside.

- Use Memento when you want to implement version control, history, or snapshots of an object's state.

## Real-World Examples

1. **Text Editors**: Text editors often use the Memento pattern to enable undo and redo functionality, allowing users to revert to previous states of a document.

2. **Game State Management**: In video games, the Memento pattern can be used to save and restore the game's state, including player positions, scores, and progress.

3. **Version Control Systems**: Version control systems like Git use a form of the Memento pattern to capture and manage changes to source code and other files over time.

4. **Database Rollbacks**: In database management systems, Memento can be employed to roll back changes to a previous database state in case of errors or data corruption.

## Implementation

In the provided Go code example, we implement the Memento pattern to demonstrate how to capture and restore the state of an object. The Originator (`Originator`), which represents the object with a state to be saved, creates and restores the Mementos (`Memento`). The Caretaker (`Caretaker`) manages the collection of saved Mementos, allowing the Originator to undo changes by restoring the state from a Memento.

## Code Description

- The `Originator` class holds the internal state and provides methods to set, capture a snapshot (`takeSnapshot`), and restore (`restore`) the state from a Memento.

- The `Memento` class stores the state captured by the Originator.

- The `Caretaker` class manages the collection of Mementos and provides methods to save (`write`) the state and undo (`undo`) changes by restoring the previous state.

This code showcases the Memento design pattern, enabling the capture and restoration of an object's state, making it suitable for various scenarios requiring state management and history tracking.
