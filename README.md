
# Bank App Server

This is a basic bank application server developed with Go, providing essential banking functionalities such as user management, account management, transaction history, money transfers, and scheduled transfers.

## Features

### User Management
- **Signup**: New users can register using a username and password.
- **Login**: Users can log in and receive a JWT token for authenticated requests.
- **Password Reset**: Users can reset their password after verifying the old one.
- **Token Refresh**: Tokens can be refreshed when they are close to expiration.

### Account Management
- **Multiple Accounts**: Users can create multiple accounts and manage them.
- **Account Balance**: Users can query their account balances.
- **Account List**: Users can list all of their bank accounts.

### Transactions
- **Transfer**: Users can transfer money between accounts.
- **Deposit**: Users can deposit money into their accounts.
- **Withdraw**: Users can withdraw money from their accounts.
- **Transaction History**: Users can view the history of all transactions made in their accounts.

### Scheduled Transfers
- **Scheduled Transfers**: Users can schedule transfers to be made on a future date.

## API Documentation

Please refer to the `API_DOCUMENTATION.md` file for a complete list of API endpoints, their inputs, and outputs.

## Project Structure

The project's directory structure is described in the `PROJECT_STRUCTURE.md` file.

## How to Run the Project

### Prerequisites
- Go (latest stable version)
- SQLite

### Steps to Run:
1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   ```

2. **Navigate to the project directory:**
   ```bash
   cd bank-app-server
   ```

3. **Install the dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   go run src/cmd/main.go
   ```

   The server will start at `http://localhost:8080`.

## Database
- The project uses SQLite for database management.
- The database file `bank.db` will be automatically generated upon running the server.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

