# Full Stack Todo Application

## Introduction

Welcome to the Full Stack Todo Application! This project is a comprehensive solution for managing your todos, built with a modern tech stack. The frontend is powered by React, TypeScript, ChakraUI, and React Query, providing a fast and responsive user interface. The backend is developed using Go, Fibre, and Air, ensuring a robust and scalable server-side application. MongoDB is used as the database to store the todo items efficiently.

## Technologies Used

### Frontend

- **React (using Vite)**: A JavaScript library for building user interfaces. Vite is used for fast and optimized builds.
- **TypeScript**: A strongly typed programming language that builds on JavaScript, giving you better tooling at any scale.
- **ChakraUI**: A simple, modular, and accessible component library that gives you the building blocks to build your React applications.
- **React Query**: A powerful data-fetching library that simplifies data synchronization between your frontend and backend.

### Backend

- **Go**: An open-source programming language that makes it easy to build simple, reliable, and efficient software.
- **Fibre**: An Express-inspired web framework built in Go with performance and scalability in mind.
- **Air**: A live-reloading development tool for Go applications to improve productivity.

### Database

- **MongoDB**: A NoSQL database known for its high performance, high availability, and easy scalability.

## Steps to Run the Application

### Starting the Frontend

1. **Install Dependencies**:
   ```bash
   cd client
   npm install
   ```
2. **Start the Development Server**
   ```bash
   npm run dev
   ```

### Starting the Frontend

1. **Install Air package**:

   ```bash
   go install github.com/cosmtrek/air@latest

   ```

2. **Set Up the Environment**:

   ```env
   MONGODB_URI=mongodb://localhost:27017/todoapp
   PORT=<PORTNUMBER>
   ```

3. **Start the Server with Air**:
   ```bash
   air
   ```

### Thank you for checking out the project! This simple todo application serves as an excellent starting point for creating full stack applications. By exploring this project, you will gain insights into:

- **Creating APIs with Go**: Learn how to build efficient and scalable APIs using the Go programming language and the Fibre web framework.
- **Integrating APIs with React**: Discover how to seamlessly integrate your Go APIs with a React frontend using Tanstack Query (formerly React Query) for efficient data fetching and state management.

### This project lays a solid foundation for building more complex and feature-rich applications. Dive into the code, experiment with the stack, and see how you can extend its capabilities to suit your needs.

### Happy coding!
