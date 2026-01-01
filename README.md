# Workflow Approval Service

A simple REST API service in Golang to manage workflow approvals (documents, purchase requests, internal requests).

---

## 🚀 Running the Application

1. Clone the repository:
git clone https://github.com/USERNAME/REPO_NAME.git
cd REPO_NAME

2. Environment:
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=workflow_db

3. Run migrations:
   go run config/migration.go
   
4. Start The Server:
   go run cmd/server/main.go
   Server will run at http://localhost:8080.

5. Database:
   PostgreSQL
   Managed via GORM
   Tables: workflows, workflow_steps, requests

6. API Endpoints:
   Workflows
   POST /workflows - Create workflow
   GET /workflows - List workflows
   GET /workflows/{id} - Get workflow by ID
   
   Workflow Steps
   POST /workflows/{id}/steps - Add step to workflow
   GET /workflows/{id}/steps - List steps of a workflow
   
   Requests
   POST /requests - Create request
   GET /requests/{id} - Get request
   POST /requests/{id}/approve - Approve request
   POST /requests/{id}/reject - Reject request

7. Design Decisions
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

8. Asumsi / Trade-offs
   - Validasi workflow name dan step level uniqueness hanya dilakukan di level service/handler minimal.
   - Transaction management menggunakan GORM transaction.
   - Tidak ada caching untuk requests/workflows karena fokus pada konsistensi data dan simplicity.

