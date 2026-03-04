# Assignment Backend API

REST API สำหรับระบบจัดการข้อมูลผู้ป่วย พัฒนาด้วย **Go (Gin)** + **PostgreSQL** + **Nginx** รองรับ HTTPS และใช้ JWT Authentication

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.21+ |
| Framework | Gin |
| Database | PostgreSQL 15 |
| ORM | GORM |
| Auth | JWT (HS256) |
| Reverse Proxy | Nginx (HTTPS) |
| Container | Docker / Docker Compose |

---

## Project Structure

```
.
├── cmd/api/            # main.go — entry point
├── internal/
│   ├── domain/         # Entities, Interfaces (Clean Architecture)
│   ├── repository/     # Database layer (GORM)
│   ├── usecase/        # Business logic
│   ├── delivery/http/  # Gin Handlers
│   ├── middleware/     # JWT Auth Middleware
│   └── auth/           # JWT Service
├── config/             # DB Connection
├── nginx/              # Nginx config + SSL certs
├── test/               # Unit tests + Mocks
├── docs/               # Swagger docs
├── docker-compose.yml
└── Dockerfile
```

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) & Docker Compose

### 1. Clone repository

```bash
git clone <repo-url>
cd assignment-backend
```

### 2. ตั้งค่า Environment Variables

สร้างไฟล์ `.env` จาก template ด้านล่าง:

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=postgres
```

> **หมายเหตุ**: ในการรันผ่าน Docker Compose `DB_HOST` จะถูก override เป็น `postgres` อัตโนมัติ

### 3. รัน Application

```bash
docker compose up --build
```

Services ที่จะ start up:

| Service | Port |
|---|---|
| Nginx (HTTP → redirect HTTPS) | 80 |
| Nginx (HTTPS) | 443 |
| Go API | 8080 (internal only) |
| PostgreSQL | 5433 |

---

## API Endpoints

Base URL: `https://localhost`

### Staff

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/staff/create` | ❌ | สร้าง Staff ใหม่ |
| `POST` | `/staff/login` | ❌ | Login, รับ JWT Token |

#### POST `/staff/create`

```json
{
  "username": "nurse01",
  "password": "secret123",
  "hospital_name": "โรงพยาบาลกรุงเทพ"
}
```

#### POST `/staff/login`

```json
{
  "username": "nurse01",
  "password": "secret123",
  "hospital_name": "โรงพยาบาลกรุงเทพ"
}
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### Patient (ต้องใช้ JWT Token)

Header: `Authorization: Bearer <token>`

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/patient/create` | ✅ | สร้างผู้ป่วยใหม่ |
| `POST` | `/patient/search` | ✅ | ค้นหาผู้ป่วย (filter หลายเงื่อนไข) |
| `GET` | `/patient/search/:id` | ✅ | ค้นหาผู้ป่วยด้วย National ID |

#### POST `/patient/create`

```json
{
  "first_name_th": "สมชาย",
  "last_name_th": "ใจดี",
  "first_name_en": "Somchai",
  "last_name_en": "Jaidee",
  "patient_hn": "HN001",
  "national_id": "1234567890123",
  "date_of_birth": "1990-01-01",
  "gender": "m",
  "phone_number": "0812345678",
  "email": "somchai@email.com"
}
```

> **Gender values**: `m` = ชาย, `f` = หญิง, `o` = อื่นๆ

#### POST `/patient/search`

ส่ง field ที่ต้องการค้นหา (ทุก field เป็น optional):

```json
{
  "first_name_th": "สมชาย",
  "national_id": "1234567890123"
}
```

---

### Other

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/ping` | ✅ | Health check |
| `GET` | `/swagger/*any` | ❌ | Swagger UI |

---

## Authentication

API ใช้ **JWT Bearer Token**

1. เรียก `POST /staff/login` เพื่อรับ token
2. ใส่ token ใน header ทุก request ที่ต้องการ Auth:

```
Authorization: Bearer <your-token>
```

---

## Postman Setup

เนื่องจากใช้ Self-signed SSL Certificate:

1. ไปที่ **Settings → General**
2. ปิด **SSL certificate verification**
3. ใช้ URL เป็น `https://localhost/<endpoint>`

---

## Running Tests

```bash
go test ./test/... -v
```

ครอบคลุม:
- Auth (JWT generate/validate)
- Middleware (token validation)
- Staff Usecase + Handler
- Patient Usecase + Handler

---

## Swagger UI

เปิด browser ไปที่:

```
https://localhost/swagger/index.html
```
