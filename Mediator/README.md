# Mediator Design Pattern

The Mediator design pattern is a behavioral pattern that promotes loose coupling between components by centralizing their communication through a mediator object. It helps manage complex interactions and relationships among objects by reducing direct dependencies between them. This pattern is particularly useful when the interactions between objects become intricate and lead to a web of interconnected dependencies.

## Why Use Mediator?

- **Decoupling**: Mediator allows you to reduce direct dependencies between objects, making the system more modular and easier to maintain.

- **Centralized Control**: It provides a centralized point for controlling and coordinating interactions among objects, which can help in avoiding chaotic and error-prone communication.

- **Promotes Reusability**: Components can be reused more easily since they are decoupled from each other and rely on a common mediator.

- **Simplifies Maintenance**: Changes to the communication logic or the addition of new components can be done more seamlessly without impacting existing components.

## When to Use Mediator?

- Use Mediator when a system has many objects that interact in complex ways, and you want to reduce the dependencies between them.

- Use Mediator when you want to centralize control and logic related to object interactions, making the system easier to understand and maintain.

- Use Mediator when you want to promote reusability of components by isolating their communication.

## Real-World Examples

1. **Air Traffic Control System**: In air traffic control, various aircraft need to communicate and coordinate their movements. A mediator (the air traffic controller) manages the interactions between aircraft to ensure safe takeoffs, landings, and routing.

2. **Chat Applications**: In chat applications, multiple users communicate with each other through messages. A mediator manages the sending and receiving of messages between users, ensuring they are properly delivered.

3. **Stock Exchange Systems**: In stock exchange systems, different components, such as traders and order books, need to interact. A mediator orchestrates the communication and execution of orders.

4. **Smart Home Systems**: In smart home systems, various devices like lights, thermostats, and sensors need to work together. A mediator can coordinate their actions based on user preferences and environmental conditions.

## Implementation

In the provided Go code example, we implement an air traffic control (ATC) system using the Mediator design pattern. The ATC acts as the mediator that manages the arrivals and departures of private and commercial aircraft. The aircraft are the colleagues that communicate with the ATC for landing permissions and departures.

## Code Description

- The `ITrafficControlMediator` interface defines methods for the mediator to handle aircraft arrivals and departures.

- The `ATC` struct is a concrete mediator that manages the landing queue and runway readiness. It implements the mediator interface.

- The `IAircraft` interface represents aircraft and defines methods for arrival, departure, and landing permission.

- `PrivateAircraft` and `CommercialAircraft` are concrete colleagues representing private and commercial aircraft, respectively. They communicate with the mediator for landing permissions and departures.

This code demonstrates how the Mediator pattern centralizes control and coordination in a complex system, such as air traffic control, making it easier to manage interactions between objects.  