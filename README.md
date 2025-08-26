# Card Service

A small **Go** service with **PostgreSQL** and **gRPC** definitions (see ```proto/```). The repository is set up for learning and practicing automated testing.

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
3. **Generate Protocol Buffers code**

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