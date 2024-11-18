
# Invoice Management System

## Project Structure

```plaintext
invoice-system/
├── main.go                  # Entry point for the application
├── models/                  # Contains database models
│   └── models.go
├── controllers/             # Contains controllers for API endpoints
│   └── invoice_controller.go
├── routes/                  # Route definitions
│   └── routes.go
├── database/                # Database connection
│   └── database.go
├── .env                     # Environment variables
├── go.mod                   # Go module configuration
├── go.sum                   # Dependency lock file
└── README.md              # Documentation
```

---

## Prerequisites

1. **Go (v1.19 or above)** installed on your system.
2. **MySQL (v5.7 or above)** installed and running.
3. **Postman** or any REST client to test the APIs.

---

## Setup Instructions

### 1. Clone the Repository

```bash
cd invoice-management
```

---

### 2. Install Dependencies

Make sure you have Go modules enabled. Then run:

```bash
go mod tidy
```

---

### 3. Configure Environment Variables

Create a `.env` file in the root of the project and add the following variables:

```plaintext
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=password
DB_NAME=invoice_system
DB_PORT=3306
```

Replace the values with your MySQL credentials.

---

### 4. Set Up the Database

#### Step 1: Create the Database
Log in to MySQL and create a database:
```sql
CREATE DATABASE invoice_system;
```

#### Step 2: Run Migrations
Use the provided SQL migration file in `migrations/001_create_tables.sql` to create the necessary tables:
```bash
cat migrations/001_create_tables.sql | mysql -u root -p invoice_system
```

---

### 5. Start the Application

Run the application:
```bash
go run main.go
```

The server will start at **http://localhost:8080**.

---

## API Documentation

### Base URL
`http://localhost:8080/api`

---

### Endpoints

#### 1. Get All Invoices
**Request**: `GET /api/invoices?page=1&limit=50`  
**Query Parameters**:  
- `page` (optional, default = 1): The page number.  
- `limit` (optional, default = 50): Number of records per page.  

**Response**:
```json
[
    {
        "id": 1,
        "invoice_id": "INV-001",
        "issue_date": "2024-01-01",
        "due_date": "2024-01-31",
        "subject": "Spring Marketing Campaign",
        "from_company": "Discovery Designs",
        "to_company": "Barrington Publishers",
        "subtotal": 10000.00,
        "tax": 1000.00,
        "total": 11000.00,
        "status": "PAID",
        "items": [
            {
                "id": 1,
                "item_type": "Service",
                "description": "Design",
                "quantity": 41.00,
                "unit_price": 230.00,
                "amount": 9430.00
            }
        ]
    }
]
```

---

#### 2. Create an Invoice
**Request**: `POST /api/invoices`  
**Request Body**:
```json
{
    "invoice_id": "INV-001",
    "issue_date": "2024-01-01",
    "due_date": "2024-01-31",
    "subject": "Spring Marketing Campaign",
    "from_company": "Discovery Designs",
    "from_address": "41 St Vincent Place, Glasgow, Scotland",
    "to_company": "Barrington Publishers",
    "to_address": "17 Great Suffolk Street, London, UK",
    "subtotal": 10000.00,
    "tax": 1000.00,
    "total": 11000.00,
    "status": "PAID",
    "items": [
        {
            "item_type": "Service",
            "description": "Design",
            "quantity": 41.00,
            "unit_price": 230.00,
            "amount": 9430.00
        }
    ]
}
```

**Response**:
```json
{
    "id": 1,
    "invoice_id": "INV-001",
    "issue_date": "2024-01-01",
    "due_date": "2024-01-31",
    "subject": "Spring Marketing Campaign",
    "from_company": "Discovery Designs",
    "to_company": "Barrington Publishers",
    "subtotal": 10000.00,
    "tax": 1000.00,
    "total": 11000.00,
    "status": "PAID"
}
```

---

#### 3. Delete an Invoice
**Request**: `DELETE /api/invoices/:id`  
**Path Parameter**:  
- `id` (required): The ID of the invoice to delete.  

**Response**:
```json
{
    "message": "Invoice deleted"
}
```

---

## Testing the Application

1. **Postman**:
   - Import the APIs into Postman and test the endpoints.
   - Ensure the pagination (`?page=1&limit=50`) works for large datasets.

2. **CLI**:
   Use `curl` to test endpoints. Example:
   ```bash
   curl -X GET "http://localhost:8080/api/invoices?page=1&limit=10"
   ```

---

## Performance and Scalability

1. **Big Data Handling**:
   - The `GET /api/invoices` endpoint supports **pagination** (`?page=1&limit=50`).
   - Indexing is applied to the `id` and `invoice_id` fields for faster lookups.

2. **Dependency Injection**:
   - Services are decoupled from controllers using Go interfaces, ensuring modularity.

---

## Future Enhancements

1. Add authentication and authorization (e.g., JWT).
2. Add unit tests for services and controllers.
3. Dockerize the application for deployment.

---

## License

This project is licensed under the MIT License.
