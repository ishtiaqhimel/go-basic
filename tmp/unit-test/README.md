## Unit Testing

Unit testing ensures software quality to develop reliable software systems 
and maintain them long term. It is an important software development practice 
that can be used for ensuring it should not destroy the functionality
of other parts when we modify some parts of an application.

Unit testing is a kind of automated testing process in which the smallest pieces of testable software in
the application, called units, are individually and independently tested to determine whether they behave
exactly as designed.

---

### Test-Driven Development (TDD)

Test-driven development (TDD) is a software development process that follows test-first development
in which unit tests are written before the production code.

The TDD approach is highly recommended by those who are using agile development methodologies
for their software delivery. 

Here are the steps involved in TDD:
1. Add a unit test to define a new functional requirement.
2. Run all tests and see whether the new unit test gets a fail.
3. Write some code to pass the tests.
4. Run the tests.
5. Refactor the code.

---

### Unit Testing with Go

The following conventions are used to write a new test suite:
- Create a source file with a name ending in `_test.go`.
- Within the test suite (the source file ends with `_test.go`), write functions with
signature `func TestXxx(*testing.T)`.

To get help with `go test`, run `go help test`.

To get help with the various flags used by the `go test`, run `go help testflag`.

---

### Benchmark Unit Tests

Benchmark our code allows us to analyze the performance of a unit of work.
Here is the convention for writing benchmark tests `func BenchmarkXxx(*testing.B)`.
We write benchmark functions inside the `_test.go` files. The benchmark tests are executed by the `go
test` command when its benchmark (`-bench`) flag is provided.

---

### Verifying Example Code

Example code provides support for running and verifying for packages,
functions, types, and methods.
Here are the naming conventions used to declare examples for the package, a function F, a type T, and a
method M on type T:
```
func Example() // Example test for package
func ExampleF() // Example test for function F
func ExampleT() // Example test for type T
func ExampleT_M() // Example test for M on type T
```
Within the example test functions is a concluding line comment that begins with `"Output:"` and is
compared with the standard output of the function when the tests are run.

In addition to verifying the example
code, example tests are available as examples for package documentation. When documentation is
generated with the `godoc` tool, the example code in the example test functions is available as an example in
the documentation.

---

### Skipping Test Cases

We can skip some of the test cases by leveraging the `Skip` function provided by the
testing package and providing the short (`-short`) flag to the `go test` command. This is useful in some
specific scenarios. If we want to skip some time-consuming test cases when running the tests, we can
leverage the capability of skipping test cases.

In another scenario, some test cases may require a dependency to resources such as a configuration
file or an environment variable that should be provided for running those tests. If these resources are not
available during the execution of those tests, we can simply skip those tests instead of letting them fail. The
testing package provides a `Skip` method of `testing.T` type that allows us to skip test cases.

---

### Running Tests Cases in Parallel

Although test cases are run sequentially, we can run test cases in parallel if we want to speed up the test
execution. When we run a large set of sequential test cases, we can leverage the capability of running
tests in parallel to speed up the execution. To run a test case in parallel, call the `Parallel` method of the
`testing.T` type as the first statement in the test case.

With the parallel flag, we specify running how many test cases at a time in parallel. If we don’t specify the
parallel flag, it defaults to `runtime.GOMAXPROCS(0)`, which is `1`, so the parallel tests will be run one at a time.

---

### BDD Testing in Go

`Ginkgo` is a behavior-driven development (BDD) – based testing framework that lets you write expressive tests in Go to specify
application behaviors. `Ginkgo` is a great choice of package for practicing BDD for our software development process. `Gomega` is a matcher 
library that is best paired with the `Ginkgo` package. Although `Gomega`
is a preferred matching library for `Ginkgo`, it is designed to be matcher-agnostic.

---

