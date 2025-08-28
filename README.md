# Card Service (Go + gRPC)

Card Service is an **educational project** written in Go that provides a **gRPC API** for practicing test automation.  
It simulates three main entities:  

- **Cards**  
- **Transactions**  
- **Users**  

The service is intended as a sandbox for QA engineers, automation testers, and developers to explore gRPC, write tests, and practice real-world scenarios.

---

## 📋Prerequisites

Before you begin, ensure your system meets the following requirements:

- [Go](https://golang.org/) (version 1.18 or higher recommended)
- [PostgreSQL](https://www.postgresql.org/)
- Protocol Buffers compiler (`protoc`)
- Go plugins for gRPC:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
- (Optional) Allure CLI (for test reports)

---

## ⚒️Installation & Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/Vasto-GoQa/card_service.git
   cd card_service
2. **Install Go dependencies**

   ```bash
   go mod download
3. (If you haven't pb.go files) **Generate Protocol Buffers code**

Make sure you have `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` installed.  
Then run:

   ```bash
   mkdir -p generated
   rotoc --go_out=generated --go_opt=paths=source_relative --go-grpc_out=generated --go-grpc_opt=paths=source_relative proto/card_service.proto
   ```

4. **Initialize PostgreSQL**

Create a new PostgreSQL database.
Run the SQL script located at ```init/db/init.sql``` to set up tables.

5. **Configure database connection**

Open ```config.go```.
Update the database connection details:

- host
- port
- username
- password
- database name

6. **(Optional) Clean and tidy dependencies**

7. **Running the Service**

Start the gRPC server with:

```bash
go run cmd/server/main.go cmd/server/config.go
```

8. **Running Tests**

Navigate to the test folder:

```bash
cd test/tests
```

Run tests:

```bash
go test
```

9. **Generating Test Reports with Allure**

After running tests, go to the Allure results folder:

```bash
cd allure-results
```

Start the Allure server to view the report:

```bash
allure serve
```

---

## 🧪 Automated Tests

The repository already includes **automated tests** for:

- **Cards** → `test/tests/card_test.go`  
- **Transactions** → `test/tests/transaction_test.go`  

👉 The **Users** entity is intentionally left without tests.

**This is your practice assignment:**

- Explore the `user_service.go` API, user model and operations with users which we have in `card_service.proto`.  
- Write gRPC tests (in Go, Python, or another language).  
- Cover both **positive** and **negative** cases

---

## 🏗️Project Structure (high level)

- ```cmd/server/``` – entry point for the service (main, config, wiring)
- ```proto/``` – Protocol Buffers definitions
- ```generated/``` – generated gRPC Go code (created by protoc)
- ```internal/``` – internal packages (business logic, data access, etc.)
- ```init/db/``` – database initialization scripts
- ```test/``` – tests and reporting artifacts

---

## 🅰️🅿️1️⃣API Definitions

The gRPC API is defined in ```proto/card_service.proto```. Regenerate the stubs using the Generate gRPC code step above whenever the proto changes.

---

## 📝Notes

- This is a learning project; implementation details may evolve.

- If ```protoc``` or plugins aren’t found, ensure ```GOPATH/bin``` (or the install location) is on your ```PATH```.

- If the server cannot connect to PostgreSQL, double-check credentials and that the ```init.sql``` has been executed.

---

## 🎯 Who is this project for?

- **Beginners** → learn gRPC API testing.  
- **Automation engineers** → practice building robust test suites.  
- **Mentors** → use as a workshop/demo project.  

---

## 💬 Feedback

Project repo: [Card Service on GitHub](https://github.com/Vasto-GoQa/card_service)  

Please leave suggestions and feedback in [GitHub Issues](https://github.com/Vasto-GoQa/card_service/issues).
