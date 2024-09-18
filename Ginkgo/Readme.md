Ginkgo
Ginkgo is a BDD(behaviour based developemnt)-style testing framework that provides a descriptive and expressive syntax for writing tests. It organizes tests hierarchically and provides hooks to set up common states.

Key Concepts of Ginkgo:
Describe: Groups a set of related tests, similar to a test suite.
Context: Adds more specificity within a Describe block.
It: Represents a specific test case.
BeforeEach and AfterEach: Hooks that run before and after each It block.
JustBeforeEach: Runs after BeforeEach and right before It blocks.
Measure: Allows you to measure performance over multiple runs.

commands:
ginkgo bootstrap -> which will helps us to create the suite test file


Gomega
Gomega is a matcher library that works seamlessly with Ginkgo, providing expressive matchers for assertions in tests.

Key Concepts of Gomega:
Matchers: Functions that assert conditions in tests (e.g., Equal, HaveOccurred).
Ω (GomegaObject): Gomega uses this to wrap values for assertions (read as "Omega").
Expect: A more readable alternative to Ω for assertions.


Common Gomega Matchers:
Equal: Checks equality.
BeTrue / BeFalse: Checks if a value is true or false.
HaveOccurred: Checks if an error has occurred.
BeNil / NotTo(BeNil()): Checks if a value is nil.
ContainSubstring: Checks if a string contains a substring.
Succeed: Checks if a function returns nil.
