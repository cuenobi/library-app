# Library App

Library App is a book borrowing and returning management system designed for libraries. This application provides an easy-to-use interface for users to borrow and return books, and for administrators to manage the entire book lending process.

## Features

- **User System**: Allow users to sign up, log in, and view their borrowing history.
- **Book Management**: Admin can add, edit, and remove books in the library's inventory.
- **Borrow/Return Books**: Users can borrow and return books with simple actions.
- **Book Availability**: Check if a book is available for borrowing.
- **Search Books**: Search books by title, author, or category.
- **Admin Dashboard**: Admins can view the status of all borrowed books and manage user activities.

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: Next.js (TypeScript)
- **Database**: PostgreSQL

## Installation

To run the project using Docker and Docker Compose, follow these steps:

### Prerequisites

Make sure you have the following installed:
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)

### Steps to Run the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/library-app.git
   cd library-app

2.	Create .env file:
Copy the .env.example file (if available) and rename it to .env. Update the environment variables as necessary.

3.	Build and start the application with Docker Compose:
    ```bash
    docker-compose up --build

4.  Wait for the containers to initialize. The backend will be available at http://localhost:8080, and the frontend will be available at http://localhost:3000.

Stopping the Application
    ```bash
    docker-compose down