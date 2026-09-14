# Sany Language Specification
**Version:** 0.1 (Draft)  
**Status:** Approved Specification  
**Reference Runtime Implementation:** Go (Pure Go, Zero CGO)  

---

## 1. Philosophy
Sany is a concise, pragmatic, general-purpose programming language designed for backend development and native agentic software engineering.

Rather than treating AI agents as external third-party libraries or bolted-on frameworks, Sany's computational model elevates agents, execution context, tasks, and tool orchestration to first-class citizens within the language semantics and runtime.

### Core Tenets
* **Pragmatism & Minimalism**: Sany prioritizes predictability over syntactic novelty. There is one idiomatic way to achieve a task, with minimal language keywords and zero hidden magic.
* **Low Boilerplate, High Clarity**: Code is clean, direct, and readable. The developer expresses intent without drowning in glue-code, wrappers, or ceremonial typing.
* **Native Intelligence Primitives**: Autonomous decision loops, context propagation, and tool dispatching are native semantics of the runtime, operating with the same natural ease as functions and loops.
* **Backend Foundation**: Sany is a robust general-purpose backend language capable of handling I/O, networking, data structures, and services with elegance and high reliability.

---

## 2. Design Goals
* **Unified Agentic Paradigm**: Eradicate the friction between standard business logic and agentic workflows.
* **Predictable Runtime Behavior**: Explicit behavior over implicit assumptions; deterministic control over non-deterministic AI outputs.
* **Gentle Learning Curve**: A developer proficient in modern languages (Go, Python, TypeScript) can write idiomatic Sany within hours.
* **Zero-Framework Tooling**: Declarations in Sany naturally expose metadata, schemas, and contracts without requiring external schema libraries (such as Pydantic or Zod).

---

## 3. Non-Goals (v0.1)
* **Low-level Systems Programming**: Sany does not provide raw memory manipulation, manual allocation/deallocation, or pointer arithmetic. It is designed for high-level application logic, I/O, and network orchestration.
* **Direct Browser UI Rendering**: Sany is a backend and orchestration language; it does not compile to client-side DOM code.
* **Complex Metaprogramming**: Sany avoids runtime AST mutation, macros, or hidden reflection proxies to preserve determinism and readability.

---

## 4. Execution Model

### 4.1. Implementation Architecture (v0.1)
* **AST Tree-Walking Interpreter**: For v0.1, Sany is executed via an Abstract Syntax Tree (AST) tree-walking interpreter implemented in pure Go (`Eval(node, scope)`).
* **Evolution Path**: The AST-to-Runtime boundary is cleanly decoupled to allow a direct transition to a Bytecode Compiler and Stack-based Virtual Machine (VM) in v0.2+ without breaking language semantics.
* **Memory Management**: Handled automatically by the host Go runtime's concurrent Garbage Collector.

### 4.2. Concurrency Model
* **Structured Concurrency (No Function Coloring)**: Functions and tasks are synchronous and direct by default. There is no `async/await` syntax coloring and infecting call hierarchies.
* **Lightweight Concurrent Tasks**: Concurrency is explicit at the call site or agent level, mapped directly to Go's lightweight Goroutines without exposing low-level threading primitives.

---

## 5. Lexical Structure

### 5.1. Character Set & Encoding
Sany source files are encoded in UTF-8. Identifiers are case-sensitive and must start with an ASCII letter or underscore (`_`), followed by letters, digits, or underscores.

### 5.2. Comments
* **Single-line comments**: Start with `//` and extend to the end of the line.
* **Multi-line comments**: Enclosed in `/* ... */` and may span multiple lines.

### 5.3. Whitespace & Semicolons
* Whitespace (spaces and tabs) is used to separate tokens. Indentation is free-form and not semantically significant.
* Semicolons (`;`) are completely optional. Newlines terminate statements. Semicolons are only used to separate multiple statements on the same line.

---

## 6. Syntax

### 6.1. Identifiers & Keywords
Reserved keywords in Sany v0.1:
```
var       const     func      class     agent     task      tool
if        else      for       in        break     continue  return
and       or        not       true      false     null      self
import    export    error
```

---

## 7. Variables & Constants

### 7.1. Mutable Variables (`var`)
Variables are declared using `var` and can be reassigned:
```sany
var count = 0
count = count + 1
```

### 7.2. Immutable Constants (`const`)
Constants are declared using `const` and cannot be reassigned once bound:
```sany
const maxRetries = 3
const apiBaseUrl = "https://api.sany.org"
```

### 7.3. Scoping & Shadowing
Sany enforces strict **Block Scoping**. Variables declared within a block (enclosed by `{ }`, such as in `if` statements or `for` loops) are only accessible within that block.
**Shadowing** is permitted: a variable declared in an inner block can have the same name as a variable in an outer block, shadowing the outer variable for the duration of the inner block.

---

## 8. Type System

### 8.1. Pragmatic Gradual Typing
Sany is dynamically typed by default with optional type annotations.
```sany
var score = 100            // dynamic type (inferred)
var score: int = 100       // annotated type with contract validation
```

When type annotations are supplied, they fulfill two roles:
1. **Runtime boundary contract validation**: Guards parameters and structured fields against invalid data (especially outputs returned by LLMs).
2. **Zero-boilerplate schema generation**: Annotated functions and tools automatically emit standardized JSON Schemas for AI model tool-calling without third-party libraries.

### 8.2. Primitive Types
* `int`: 64-bit signed integer.
* `float`: 64-bit IEEE 754 floating point.
* `bool`: Boolean value (`true` or `false`).
* `string`: Immutable UTF-8 encoded text.
* `null`: Represents intentional absence of value.
* `error`: Represents a captured failure state or execution error.
* `any`: The wildcard dynamic type (implicit default when annotations are omitted).

### 8.3. Type Coercion & Equality
* Equality (`==`, `!=`) checks value equality.
* Sany avoids JavaScript-style loose type coercions: `1 == "1"` evaluates to `false`.
* Explicit conversions are provided via standard primitives: `int("42")`, `string(100)`.

---

## 9. Strings & Literals

### 9.1. Interpolated Strings
Double-quoted strings support native expression interpolation using `{expr}` without requiring special prefixes:
```sany
var name = "Raphael"
var age = 21
print("User {name} is {age} years old.")
```

### 9.2. Raw & Multi-line Strings
Enclosed in backticks (`` `...` ``), raw strings preserve formatting and do not evaluate escape characters. Ideal for prompts and JSON payloads:
```sany
const systemPrompt = `
You are a precise data extractor.
Output strictly JSON matching the required schema.
`
```

---

## 10. Control Flow

### 10.1. Conditionals (`if`, `else if`, `else`)
Parentheses around conditions are optional. Explicit curly braces are mandatory:
```sany
if score >= 90 {
    print("Grade A")
} else if score >= 75 {
    print("Grade B")
} else {
    print("Grade C")
}
```

### 10.2. Logical Operators
Logical operations use readable English keywords: `and`, `or`, `not`.
```sany
if active and not blocked {
    login()
}
```

### 10.3. Unified Loop (`for`)
Sany does not have a `while` keyword. The `for` statement covers all loop constructs:
```sany
// Condition-only loop (while-style)
for count < 10 {
    count += 1
}

// Collection iteration
for item in items {
    process(item)
}

// Infinite loop with break
for {
    if shouldExit() {
        break
    }
}
```

---

## 11. Collections

### 11.1. Arrays
Zero-indexed, ordered, dynamic lists:
```sany
var fruits = ["apple", "banana", "orange"]
print(fruits[0])        // "apple"
fruits.append("grape")
print(fruits.length())  // 4
```

### 11.2. Maps
Key-value associative dictionaries:
```sany
var user = {
    "name": "Raphael",
    "role": "admin"
}

// Accessible via indexing or property dot notation
print(user["name"])     // "Raphael"
print(user.role)        // "admin"
```

---

## 12. Functions & Classes

### 12.1. Functions
Declared using the `func` keyword. First-class citizens, supporting lexical scoping and closures:
```sany
func calculate_total(subtotal: float, taxRate: float = 0.1): float {
    return subtotal + (subtotal * taxRate)
}
```

### 12.2. Classes
Classes define blueprints for objects with state and behavior.
* Instantiation does not require the `new` keyword.
* Methods use `self` to reference the instance, which is implicitly available inside method bodies (no manual `self` argument in parameter definitions).
* Constructors use `func init(...)`:

```sany
class User {
    var name: string
    var email: string

    func init(name: string, email: string) {
        self.name = name
        self.email = email
    }

    func describe(): string {
        return "{self.name} <{self.email}>"
    }
}

var user = User("Raphael", "raphael@sany.org")
print(user.describe())
```

---

## 13. Errors as Values

### 13.1. Philosophy
Errors are values, not hidden runtime control jumps. Sany has a built-in `error` primitive type. When a function fails, it explicitly returns an `error` value instead of throwing an exception.

### 13.2. Ergonomic Fallback & Propagation via `or`
The `or` keyword serves as a concise short-circuiting fallback operator. If the expression on the left evaluates to an `error` or `null`, the `or` operator immediately evaluates and yields the expression on the right:
```sany
// Fallback to a default value
var port = env.get("PORT") or "8080"

// Early return on error
var config = loadConfig("config.json") or return "Failed to load config"
```

---

## 14. Modules & Imports

### 14.1. Export & Visibility
By default, all declarations (`func`, `var`, `const`, `class`, `agent`, `task`, `tool`) are private to the file in which they are defined. To expose them for import by other files, use the `export` keyword:
```sany
export func process_data() { ... }
export const timeout = 30
```

### 14.2. Module Loading (Decentralized)
Sany uses a decentralized package management model, importing modules directly via their Git repository URLs, allowing a zero-infrastructure ecosystem (similar to Go and Deno). Standard libraries use short names:
```sany
// Standard library imports
import "http"
import { get, post } from "http"

// Third-party decentralized imports
import "github.com/euraphhh/sany/x/math"
```

---

## 15. Tools

### 15.1. Declaration
Tools are functions marked with the `tool` keyword. They include an explicit semantic docstring explaining what the tool does to an AI model:
```sany
tool search_web(query: string, limit: int = 5): array {
    "Searches the web for recent technical news and articles matching the query"
    var response = http.get("https://api.search.com?q={query}&limit={limit}")
    return response.json()
}
```

### 15.2. Auto-Schema Introspection
The Sany runtime automatically inspects `query: string`, `limit: int = 5`, the return type, and the docstring, exposing a standard JSON Schema for AI providers without external configuration.

---

## 16. Agents

### 16.1. Declaration
An `agent` is an autonomous reasoning actor declared as a first-class language structure:
```sany
agent Researcher {
    model = "gpt-4o"
    prompt = "You are a senior tech research analyst. Produce deep, objective analyses."
    tools = [search_web]
}
```

### 16.2. Model Resolution
The `model` property accepts:
* Shorthand string identifier resolved automatically by runtime environment variables: `"gpt-4o"`, `"claude-3-5-sonnet"`, `"ollama/llama3"`.
* Configured model instance from the `model` standard library: `Model.anthropic("claude-3-5-sonnet", temperature = 0.2)`.

---

## 17. Tasks

### 17.1. Declaring Tasks within Agents
A `task` represents a discrete cognitive capability of an agent, with typed parameters, an objective instruction, and a typed return contract:
```sany
class AnalysisReport {
    var topic: string
    var summary: string
    var key_takeaways: array
}

agent Researcher {
    model = "gpt-4o"
    prompt = "You are an expert research analyst."
    tools = [search_web]

    task analyze(topic: string): AnalysisReport {
        "Conduct thorough research on the given topic and formulate an AnalysisReport."
    }
}
```

### 17.2. Invoking Tasks
Invoking a task is syntactically direct:
```sany
var bot = Researcher()
var report = bot.analyze("Advancements in Solid State Batteries")
print(report.summary)
```

### 17.3. The Execution Cycle
When a task is invoked:
1. Runtime constructs the execution context (system prompt + task prompt + input parameters + tool schemas).
2. Runtime calls the configured model driver.
3. If the model emits tool calls, the runtime executes the corresponding Sany `tool` functions, injects responses, and continues the cognitive loop.
4. Once the model outputs the final answer, the runtime validates and coerces the result into the expected return type (e.g. `AnalysisReport`).

---

## 18. Multi-Agent Collaboration

### 18.1. Explicit Orchestration (Deterministic Code)
Developers can orchestrate multiple agents using standard Sany code:
```sany
var researcher = Researcher()
var writer = TechnicalWriter()

var data = researcher.analyze("Quantum Computing in 2026")
var article = writer.draft(data)
```

### 18.2. Autonomous Delegation (Task as Tool)
Because an agent's `task` has a typed signature and docstring, **it can be passed directly as a tool to another agent**:
```sany
var researcher = Researcher()

agent Coordinator {
    model = "gpt-4o"
    prompt = "You coordinate research and publishing workflows."
    tools = [researcher.analyze]

    task publish(subject: string): string {
        "Produce a complete publication on the subject by delegating research as needed."
    }
}
```

---

## 19. Context & Memory

### 19.1. Stateless by Default
Tasks are stateless by default. Each invocation starts with a clean conversation context to prevent token leakage and state contamination across backend requests.

### 19.2. Session Memory
Stateful conversational memory can be explicitly enabled when needed:
```sany
var bot = Researcher(memory = true)
// Or via explicit session
var session = bot.session()
session.chat("First question")
session.chat("Follow-up question") // Remembers first question
```

### 19.3. Execution Context & Timeouts
The standard `Context` carries timeouts, cancellation, and tracing metadata:
```sany
import "context"

var ctx = Context(timeout = 30) // 30 seconds deadline
var report = bot.analyze("Topic", with = ctx)
```

---

## 20. Standard Library & Go Runtime Integration

### 20.1. CLI Toolchain (Pure Go)
Distributed as a single, zero-dependency binary `sany`:
* `sany run <file.sany>`: Executes a Sany program.
* `sany check <file.sany>`: Performs type checking and syntax validation.
* `sany fmt <file.sany>`: Formats code according to canonical style rules.

### 20.2. Core Standard Libraries (v0.1)
* `http`: High-performance HTTP client and server based on Go's `net/http`.
* `json`: JSON serialization and deserialization.
* `io`: File system reading and writing.
* `env`: Operating system environment variable access.
* `time`: Durations, timestamps, and sleeping.
* `model`: Model drivers for OpenAI, Anthropic, Ollama, and generic OpenAI-compatible endpoints.

---

## 21. Future Considerations (v0.2+)
* **Bytecode Compiler & VM**: Replacement of the AST tree-walking interpreter with a dedicated Go-based bytecode VM for raw CPU speed optimization.
* **Hot-Reloading for Agents**: Dynamic updating of agent prompts and tools without restarting backend servers.
* **Persistent Memory Backends**: Built-in adapter interfaces for vector databases and semantic memory stores.
