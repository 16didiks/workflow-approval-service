# Workflow Approval Service

Sebuah REST API service sederhana untuk menangani sistem workflow approval (approval dokumen atau purchase request), dibangun menggunakan Golang, GORM, dan PostgreSQL.

How to run the application?

Clone repository
git clone <repository_url>
cd workflow-approval-service

---

Install dependencies
go mod tidy

---

Environment

APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=workflow_db
DB_DRIVER=postgres

---

Migrasi Database
go run cmd/server/main.go

---

Database configuration
Database PostgreSQL.
GORM digunakan sebagai ORM.

Tabel :
workflows → menyimpan alur approval
workflow_steps → menyimpan tahapan approval
requests → menyimpan data yang melewati workflow

Migrasi sudah otomatis dilakukan saat server dijalankan.

---

Endpoint API

Workflow
| Method | Endpoint | Deskripsi |
| ------ | ----------------- | ------------------------ |
| POST | `/workflows` | Membuat workflow baru |
| GET | `/workflows` | List semua workflow |
| GET | `/workflows/{id}` | Detail workflow tertentu |

Workflow Step
| Method | Endpoint | Deskripsi |
| ------ | ----------------------- | ------------------------------ |
| POST | `/workflows/{id}/steps` | Tambah step ke workflow |
| GET | `/workflows/{id}/steps` | List semua step dalam workflow |

Request
| Method | Endpoint | Deskripsi |
| ------ | ------------------------ | ---------------------------------- |
| POST | `/requests` | Membuat request baru |
| GET | `/requests/{id}` | Mendapatkan request berdasarkan ID |
| POST | `/requests/{id}/approve` | Approve request |
| POST | `/requests/{id}/reject` | Reject request |

---

Design Decisions

1. Clean Architecture
   - handler → HTTP layer
   - service → business logic / usecase
   - repository → database access
   - model → domain entity

2. Concurrency Safety
   - Endpoint approve menggunakan transaction (FindByIDForUpdate) untuk mencegah double approval / race condition.

3. Validation
   - Amount harus > 0
   - Request hanya bisa di-approve/reject jika status masih PENDING
   - CurrentStep selalu mulai dari 1

4. ORM
   - GORM dipakai untuk mempermudah akses database dan migrasi otomatis.

5. Error Handling
   - Custom error untuk ErrRequestNotFound dan ErrRequestAlreadyProcessed
   - Propagasi error melalui service dan handler

   ***

Asumsi / Trade-offs

- Validasi workflow name dan step level uniqueness hanya dilakukan di level service/handler minimal.
- Transaction management menggunakan GORM transaction.
- Tidak ada caching untuk requests/workflows karena fokus pada konsistensi data dan simplicity.
