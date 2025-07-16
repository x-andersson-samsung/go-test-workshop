# Go test workshop

## Plan

### **Part 1: Fundamentals of Testing in Go (1 hour)**
- Introduction to Go’s standard testing library
- Introduction to Test-Driven Development (TDD)
- Test naming, output, and failure diagnostics
---
### **Part 2: Test Structure (1 hour)**
- Table-Driven Tests
- Writing DRY and readable test suites
- Using `t.Run` to structure subtests
- Hands-on: Table-driven tests for various scenarios
---
### **Part 3: Test Organization and Coverage (1 hour)**
- Organizing tests in packages
- Code coverage: generating and interpreting with `go test -cover`
- Using `go tool cover` for visual coverage
- Hands-on: Refactor tests and check coverage
---
### **Break (30 minutes)**
---
### **Part 4: Mocks and Dependency Isolation (1 hour)**
- Interfaces for testability
- Manual mocking vs. mocking frameworks
- Mocking time (`clockwork`)
- Hands-on: Writing and using mocks in service layer tests
---
### **Part 5: Advanced Topics – Benchmarking, Concurrency, and Integration (1 hour)**
- Writing performance benchmarks with `BenchmarkXxx`
- Testing concurrent code (data races, `-race` flag, `testing/synctest`)
- Creating mock containers for integration tests.
- Hands-on: Writing a benchmark and a concurrent test
---
### **Part 6: Test Frameworks and Best Practices (1 hour)**
- Using `stretchr/testify` for more expressive assertions
- Setup / teardown strategies ()
- Golden files and snapshot testing
-  Hands-on: Convert basic tests to use `testify`
---
### **Takeaways:**
By the end of the training, participants will:
- Understand how to write and structure tests in Go
- Be able to test real-world codebases with unit, integration, and concurrent tests
- Know how to use tools like `testify`, `go test`, and code coverage utilities effectively
- Be familiar with Go testing patterns and performance testing (edited) 
