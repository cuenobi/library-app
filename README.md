# Library App

Library App is a book borrowing and returning management system designed for libraries. This application provides an easy-to-use interface for users to borrow and return books, and for administrators to manage the entire book lending process.

## Features

- **User System**: Allow users to sign up, log in, and view their borrowing history.
- **Book Management**: Admin can add books to the library's inventory.
- **Borrow/Return Books**: Users can borrow and return books with simple actions.
- **Book Availability**: Check if a book is available for borrowing.
- **Search Books**: Search books by title, or category.
- **Librarian Dashboard**: The Librarian can view the status of all borrowed books and manage user activities.

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
   ```

2. Create .env file:
   Copy the .env.example file (if available) and rename it to .env. Update the environment variables as necessary.

3. Build and start the application with Docker Compose:

   ```bash
   docker-compose up --build
   ```

4. Wait for the containers to initialize. The backend will be available at http://localhost:8080, and the frontend will be available at http://localhost:3000.

   Stopping the Application

   ```bash
   docker-compose down
   ```

### Accessing Swagger API Documentation

This application comes with Swagger UI for easy access to the API documentation. To view the API documentation: 1. Make sure the backend is running (at http://localhost:8080). 2. Open your browser and go to the following URL:

```bash
http://localhost:8080/swagger/
```

This will open the Swagger UI where you can explore and interact with the API endpoints.

### Username/Password for access Library App

- Librarian:
  • Username: john_doe
  • Password: password123
- Member: *Unsupported for feature
  • Username: jane_doe
  • Password: password123

### Notes
   - For accessing features, please primarily use the Librarian account.
   - This application was developed within a limited timeframe (7 days) as part of a job application test.
   - This application is still under development. Features may be incomplete or unstable.
   Below are some features I plan to add or improve in the future:
	•	Borrowing Check: Implement a system that prevents a member from borrowing a book if they already have an active loan.
	•	Book Status Update: Change the book’s status to Not Available when borrowed and update it back when returned.
	•	Member Registration: Allow users to sign up and manage their accounts properly.
   - If you need additional information or have any questions, feel free to ask!
   - Some features have not been added yet. The code may not be fully completed or follow best practices in the frontend part as I am still learning it. I apologize for any inconvenience.
