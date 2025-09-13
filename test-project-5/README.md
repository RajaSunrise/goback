# TEST-PROJECT-5

test-project-5 backend API

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- POSTGRESQL Database Server
- Docker (optional, for containerization)
- Make (optional, for automation)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/test-project-5.git
    cd test-project-5
    ```

2.  **Set up environment variables:**
    ```bash
    cp .env.example .env
    ```
    *Update the `.env` file with your configuration, especially database credentials.*

3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Run database migrations:**
    *Ensure your database server is running and the credentials in `.env` are correct.*
    ```bash
    make migrate
    ```

5.  **Run the application:**
    ```bash
    make dev
    ```
    The server will start on `http://localhost:8080` (or the port specified in `.env`).

## 📁 Project Structure

A brief overview of the project structure based on the **SIMPLE ARCHITECTURE Architecture**:
