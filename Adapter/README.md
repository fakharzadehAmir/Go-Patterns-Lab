# Adapter Design Pattern

The Adapter design pattern is a structural pattern that allows objects with incompatible interfaces to work together. It acts as a bridge between two incompatible interfaces, making them compatible without altering their source code. This pattern is particularly useful when integrating new features or components into an existing system.

**Why Use Adapter Design Pattern**:
- **Integration**: It facilitates the integration of new functionalities or third-party components without modifying existing code.
- **Reusability**: It promotes code reusability by allowing you to use existing classes/interfaces with new systems.
- **Maintenance**: It simplifies maintenance by isolating changes or updates within the adapter.

**Real Examples of Adapter Design Pattern**:

1. **Legacy Systems Integration**:
    - In enterprise software, the Adapter pattern is used to integrate legacy systems or databases with modern applications.

2. **API Compatibility**:
    - When working with external APIs, adapters can be used to make the API's interface compatible with the application's requirements.

3. **Library Compatibility**:
    - In software libraries or frameworks, adapters can be implemented to provide compatibility with different versions or platforms.

4. **Hardware Abstraction**:
    - In embedded systems, adapters are used to abstract hardware interactions, making it easier to switch between different hardware components.

## Implementation

In the provided Go package, we have implemented the Adapter design pattern for integrating two different database systems, PostgreSQL and MongoDB. Let's break down the key components:

- **Client**: The `Client` struct represents a client that interacts with a database. It has a method `SaveNewData` that allows clients to save new data to the database.

- **RelationalDatabase**: The `RelationalDatabase` interface defines the method `CreateRecord` for creating a new record in the database. It acts as an adapter interface that both PostgreSQL and MongoDB adapters implement.

- **Postgres**: The `Postgres` struct represents a PostgreSQL database. It has a method `CreateRecord` to create a new record in PostgreSQL.

- **MongoDB**: The `MongoDB` struct represents a MongoDB database. It has methods `CreateDocument` and `GetDocument` to create and retrieve data from MongoDB.

- **MongoAdapter**: The `MongoAdapter` struct acts as an adapter for MongoDB to implement the `RelationalDatabase` interface. It wraps the MongoDB functionality and provides an adapter method for creating records.

## Usage

To use the Adapter pattern for integrating databases, follow these steps:

1. **Create Client**: Instantiate a `Client` object, specifying the new data you want to save.

2. **Choose Database**: Create an instance of either `Postgres` or `MongoDB`, depending on the database system you want to use.

3. **Save Data**: Call the `SaveNewData` method on the client, passing the chosen database instance as an argument. This will save the new data to the selected database.

4. **Check Data (Optional)**: If needed, you can retrieve the saved data using the appropriate database-specific methods.
