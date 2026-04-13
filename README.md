# 🔗 Go URL Shortening Service

A high-performance and scalable URL shortening service built with Go, Gin, and MongoDB. This project provides RESTful APIs to create, manage, and track short URLs with analytics support.

---

## 🚀 Features

* 🔗 Create short URLs from long links
* 🔍 Retrieve original URL using short code
* ✏️ Update existing short URLs
* 🗑️ Delete short URLs
* 📊 Track click statistics
* 🔁 Fast redirection system
* 🧠 Structured error handling
* 🗃️ MongoDB integration

---

## 🛠️ Tech Stack

* Backend: Go (Golang)
* Framework: Gin
* Database: MongoDB

---

## 📁 Project Structure

- `config/` — Database configuration
- `controllers/` — API handlers
- `models/` — Data models
- `repository/` — Database logic
- `routes/` — API routes
- `utils/` — Helper functions
- `main.go` — Entry point
- `.env.example` — Environment variables template
- `README.md`

---

## ⚙️ Setup & Installation

### 1. Clone the Repository

git clone https://github.com/Rajkumar-coderm/go-url-shortening-service.git
cd go-url-shortening-service

---

### 2. Install Dependencies

go mod tidy

---

### 3. Setup Environment Variables

Create a `.env` file:

MONGO_LOCAL_URI=mongodb://localhost:27017
MONGO_DB_NAME=go-url-short
PORT=8080

⚠️ Do NOT commit `.env` file

---

### 4. Run the Server

go run main.go

Server will start at:
http://localhost:8080

---

## 📡 API Endpoints

- `POST /api/shorten` → Create short URL
- `GET /api/url/:code` → Get URL details
- `PUT /api/url/:code` → Update URL
- `DELETE /api/url/:code` → Delete URL
- `GET /api/url/:code/stats` → Get URL stats
- `GET /r/:code` → Redirect to original URL

---

## 📦 Example Responses

### ✅ Success

``` json
{
    "status": "success",
    "message": "Short URL created successfully",
    "data": {
        "code": "1Eo4n9",
        "created_at": 1776092904,
        "original_url": "https://web.whatsapp.com",
        "short_url": "http://localhost:8080/r/1Eo4n9",
        "updated_at": 1776092904
    }
}
```

---

### ❌ Error

``` json
{
    "status": "error",
    "message": "The field 'URL' is mandatory and cannot be empty. ",
    "data": null
}
```

---

## 📊 Stats Response

```json
{
    "status": "success",
    "message": "Stats retrieved successfully",
    "data": {
        "id": "69dd06e846bbe8823395c4a7",
        "code": "1Eo4n9",
        "original_url": "https://web.whatsapp.com",
        "clicks": 0,
        "created_at": 1776092904,
        "updated_at": 1776092904
    }
}
```

---

## 🔐 Security Best Practices

* Never commit `.env` files
* Use environment variables
* Rotate secrets if exposed
* Validate all inputs

---

## 🚀 Future Enhancements

* 🔐 User authentication
* 📊 Advanced analytics (device, location)
* ⚡ Redis caching
* 🚫 Rate limiting
* 🎯 Custom short URLs
* ⏳ Expiry links

---

## 💡 Use Cases

* Social media sharing
* Marketing campaigns
* Link tracking
* Developer tools

---

## 👨‍💻 Author

Rajkumar Gahane

---

## ⭐ Support

If you like this project:

* ⭐ Star the repo
* 🍴 Fork it
* 🛠️ Contribute
