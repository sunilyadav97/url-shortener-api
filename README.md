# **URL Shortener API**  

This is a **URL Shortener** service built with **Go (Gin framework)** and **MongoDB**. The API allows users to create short URLs, redirect to original URLs, and update stored URLs. The project is fully **Dockerized** for easy installation and deployment.

## **Table of Contents**
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Environment Variables (.env)](#environment-variables-env)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Project Structure](#project-structure)
- [License](#license)

---

## **Features**
✅ Generate short URLs from long URLs  
✅ Redirect users from short URLs to original URLs  
✅ Update existing short URLs with new destinations  
✅ MongoDB integration for persistent data storage  
✅ Dockerized for easy deployment  
✅ Fast and efficient API using the **Gin framework**  

---

## **Prerequisites**  
Ensure you have the following installed before proceeding:

1. **Docker** - [Download & Install](https://www.docker.com/get-started)  
2. **Docker Compose** - [Install Docker Compose](https://docs.docker.com/compose/install/)  
3. **(Optional) Go 1.22+** (if you want to run the project without Docker) - [Download Go](https://golang.org/dl/)  

---

## **Installation**  

### **Step 1: Clone the Repository**  
```bash
git clone https://github.com/sunilyadav97/url-shortener-api.git
cd url-shortener-api
```

If you haven't initialized a Git repository yet, you can create a new folder and add all the files manually.

---

### **Step 2: Create a `.env` file**  
The `.env` file contains environment variables that configure the application. Create a `.env` file in the project's root directory:

```bash
touch .env
```

Now, open the `.env` file and add the following configuration:

```ini
# Server Configuration
PORT=8080  # Port where the API will run

# MongoDB Configuration
MONGO_URI=mongodb://mongo:27017  # MongoDB connection URI
MONGO_DB_NAME=url_shortener_db   # Database name
MONGO_COLLECTION_NAME=urls       # Collection name
```

---

### **Step 3: Install Dependencies (Only for Local Development)**
If you want to run the project without Docker, you need to install dependencies manually:

```bash
go mod tidy
```

---

## **Running the Project**

### **Option 1: Run with Docker (Recommended)**
Use **Docker Compose** to start the project:

```bash
docker-compose up --build
```

This will:
- Build the Go application as a Docker container  
- Start the **Go API** on `localhost:8080`  
- Start **MongoDB** as a separate container  

To stop the containers, press **`CTRL + C`** or run:

```bash
docker-compose down
```

---

### **Option 2: Run Locally (Without Docker)**
If you prefer running the project without Docker:

#### **Step 1: Start MongoDB Locally**  
Make sure you have MongoDB installed and running. If not, start it using:

```bash
mongod --dbpath <your-db-path>
```

#### **Step 2: Run the Go Application**
Execute the following command:

```bash
go run main.go
```

The server should now be running at:  
👉 **http://localhost:8080**

---

## **API Endpoints**

### **1. Create Short URL**
- **Endpoint:** `POST /api/v1/createShortUrl`
- **Request Body:**
  ```json
  {
    "url": "https://example.com/very/long/url"
  }
  ```
- **Response:**
  ```json
  {
    "shortUrl": "abcdefg"
  }
  ```

---

### **2. Redirect to Original URL**
- **Endpoint:** `GET /api/v1/{shortUrl}`
- **Response:**  
  A `301` redirect to the original URL.

---

### **3. Update Short URL**
- **Endpoint:** `PUT /api/v1/updateShortUrl/{shortUrl}`
- **Request Body:**
  ```json
  {
    "url": "https://example.com/new-destination"
  }
  ```
- **Response:**
  ```json
  {
    "status": "Successful",
    "message": "Short URL updated successfully"
  }
  ```

---

## **Testing with Postman**
1. **Create a Short URL:**  
   - **Method:** `POST`  
   - **URL:** `http://localhost:8080/api/v1/createShortUrl`
   - **Body (raw, JSON):**
     ```json
     {
       "url": "https://example.com/very/long/url"
     }
     ```
   - **Expected Response:**
     ```json
     {
       "shortUrl": "abcdefg"
     }
     ```

2. **Redirect to Original URL:**  
   - **Method:** `GET`  
   - **URL:** `http://localhost:8080/api/v1/abcdefg`  
   - **Expected Behavior:** Redirects to `https://example.com/very/long/url`.

3. **Update a Short URL:**  
   - **Method:** `PUT`  
   - **URL:** `http://localhost:8080/api/v1/updateShortUrl/abcdefg`
   - **Body (raw, JSON):**
     ```json
     {
       "url": "https://example.com/new-destination"
     }
     ```
   - **Expected Response:**
     ```json
     {
       "status": "Successful",
       "message": "Short URL updated successfully"
     }
     ```

---

## **Project Structure**

```
url-shortener-api/
├── .env                   # Environment variables
├── .gitignore             # Git ignore file
├── Dockerfile             # Docker configuration
├── docker-compose.yml     # Docker Compose for Go & MongoDB
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies checksum
├── main.go                # Entry point (Gin API)
├── README.md              # Project documentation
├── database/              # MongoDB connection
│   └── database.go
├── models/                # Data models
│   └── models.go
├── routes/                # API routes and handlers
│   └── routes.go
└── utils/                 # Utility functions
    └── shortener.go
```

---

## **License**
This project is licensed under the **MIT License**.

---

## **Summary**
✅ **`.env`** file contains important configurations like **MongoDB URI** and **server port**.  
✅ Use **`docker-compose up --build`** to run the project with **Docker**.  
✅ If running **locally**, install dependencies with `go mod tidy` and run `go run main.go`.  
✅ Test API with **Postman** or `curl` commands.
