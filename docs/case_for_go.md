# The Case for Go-Lang Usage for Backends

Building performant, scalable, and cost-effective backends is crucial for modern software applications. With the rise of serverless architectures and cloud computing, the choice of programming language can have a significant impact on the performance and cost of an application. In this document, we will evaluate the performance, scalability, and developer experience of three popular programming languages - TypeScript, Rust, and Go - in the context of building distributed, serverless applications. By understanding the strengths and weaknesses of each language, we aim to identify the best choice for building backends that meet specific requirements and constraints. The goal is to help developers strike a balance between performance, scalability, and cost-effectiveness while ensuring a positive developer experience.

Table of Contents

- [Pros/Cons of TypeScript, Rust, and Go](#proscons-of-typescript-rust-and-go)
  - [TypeScript](#typescript)
  - [Rust](#rust)
  - [Go](#go)
- [Performance](#performance)
- [Scalability](#scalability)
- [Developer Experience](#developer-experience)
- [Conclusion](#conclusion)
- [Addendum](#addendum)
  - [Paradigm Discussion - Functional versus Object-Oriented Programming](#paradigm-discussion---functional-versus-object-oriented-programming)
  - [Other languages not considered - Rust and Elixir](#other-languages-not-considered---ruby-and-elixir)

## Pros/Cons of TypeScript, Rust, and Go

### TypeScript

Pros:

- Easy to learn and use for developers familiar with JavaScript.
- Large developer community and ecosystem of libraries and tools.
- Supports modern language features such as async/await, decorators, and interfaces.
- Offers good performance for small to medium-sized applications.

Cons:

- Can have a performance overhead due to type-checking and compilation.
- Not ideal for low-level system programming.
- May not be ideal for large-scale, high-traffic applications.

### Rust

Pros:

- Fast and efficient, with low memory usage.
- Provides a high degree of safety and memory safety.
- Excellent for low-level system programming and performance-critical applications.
- Strongly typed with zero-cost abstractions.
- Good concurrency and parallelism support.

Cons:

- Steep learning curve for developers unfamiliar with systems programming.
- Smaller developer community and ecosystem of libraries and tools.
- More verbose than other languages, which can make it slower to write code.

### Go

Pros:

- Fast, efficient, and low memory usage.
- Provides good support for concurrency and parallelism.
- Simple and easy to learn, with a clean syntax and standard library.
- Offers good performance for medium to large-scale applications.
- Offers a good balance between performance and ease of use.

Cons:

- Less powerful than some other languages, with fewer language features.
- Garbage collection can impact performance in certain scenarios.
- Smaller ecosystem of libraries and tools compared to some other languages.

## Performance

**TypeScript** is a superset of JavaScript that adds optional static typing to the language. TypeScript has gained popularity among developers due to its ease of use and familiarity with JavaScript. However, when it comes to performance, TypeScript may not be the best choice. Since TypeScript is compiled into JavaScript, it has some performance overhead due to type-checking and compilation. However, the performance impact of TypeScript is usually negligible, especially for small to medium-sized applications.

In this code snippet, we're using TypeScript to calculate the 30th Fibonacci number. The Fibonacci sequence is a series of numbers where each number is the sum of the two numbers before it. The first two numbers in the sequence are 0 and 1, and the next number is the sum of the previous two numbers. The 30th Fibonacci number is 832040.  The code snippet below uses recursion to calculate the 30th Fibonacci number.  Give or take your machine's performance, the code snippet below should may take up to 17 seconds to run.

```TypeScript
function fibonacci(n: number): number {
  if (n < 2) {
    return n;
  }
  return fibonacci(n - 1) + fibonacci(n - 2);
}

const result = fibonacci(30);
console.log(result);
```

**Rust** is a systems programming language that focuses on speed and safety. Rust's performance is excellent due to its zero-cost abstractions and high-level control over system resources. Rust's ability to handle low-level system programming, combined with its performance, makes it an excellent choice for serverless functions that require high performance.

Using Rust, we can calculate the 30th Fibonacci number in less than a second.

```rust
fn fibonacci(n: u32) -> u32 {
    if n < 2 {
        return n;
    }
    fibonacci(n - 1) + fibonacci(n - 2)
}

fn main() {
    let result = fibonacci(30);
    println!("{}", result);
}
```

**Go** is a compiled language developed by Google that focuses on simplicity, performance, and concurrency. Go's performance is excellent due to its garbage collection, fast compilation, and optimized concurrency. Go's simplicity also makes it easy to write and maintain code. Go is a popular choice for building high-performance serverless applications due to its low memory footprint and fast execution speed.

Using Go, we can calculate the 30th Fibonacci number in approximately 1.2 seconds.

```go
package main

import "fmt"

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    result := fibonacci(30)
    fmt.Println(result)
}
```

When it comes to performance, Rust and Go are excellent choices for serverless functions.  Regards to performance in general, there is [an entertaining but informative video on YouTube that discusses a deeper performance testing between TypeScript, Rust, and Go](https://youtu.be/Z0GX2mTUtfo). Rust's focus on speed and safety makes it an excellent choice for low-level system programming, while **Go's simplicity, speed, and optimized concurrency make it an excellent choice for building high-performance serverless applications**. TypeScript, on the other hand, may not be the best choice for high-performance applications, but its ease of use and familiarity with JavaScript make it an excellent choice for smaller applications.

TypeScript is great, but it is definitely at a disadvantage here.  Let's continue to the next topic.

## Scalability

**TypeScript** works well with AWS Lambda and is a popular choice for building serverless applications.  It's easy to find examples of TypeScript serverless functions online.

```TypeScript
import { APIGatewayProxyHandler } from 'aws-lambda';

export const handler: APIGatewayProxyHandler = async (event, context) => {
  console.log('Handling request:', event);

  const response = {
    statusCode: 200,
    body: JSON.stringify({ message: 'Hello, World!' }),
  };

  console.log('Returning response:', response);
  return response;
};
```

**Rust** also works well with AWS Lambda.  Rust's fast execution speed, low memory usage, and strong memory safety features make it well-suited for building high-performance and secure serverless applications. Additionally, Rust has a growing ecosystem of libraries and tools specifically designed for building serverless functions, including the lambda_runtime and lambda_http crates, which make it easy to create and deploy serverless functions in Rust.

```rust
use lambda_http::{handler, lambda, IntoResponse, Request, Response};
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    lambda!(my_handler);
    Ok(())
}

fn my_handler(_: Request, _: lambda_runtime::Context) -> Result<impl IntoResponse, lambda_http::Error> {
    Ok(Response::builder()
        .status(200)
        .header("Content-Type", "application/json")
        .body(r#"{"message": "Hello, World!"}"#.into())
        .expect("Failed to render response body"))
}
```

**Go** is the biggest winner here as it offers some advantages compared to TypeScript and Rust for serverless applications:

- Fast startup time: Go has a very fast startup time, which is important for serverless applications because the function needs to start up quickly in response to an event.
- Low memory usage: Go uses less memory compared to TypeScript and Rust, which can be important for reducing the cost of running serverless functions in a cloud environment.
- Concurrent execution: Go's lightweight thread model and support for concurrency make it easy to write code that can handle multiple requests in parallel, which is important for scaling serverless applications.

```golang
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Handling request: %+v\n", request)

	response := Response{
		Message: "Hello, World!",
	}

	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Printf("Returning response: %+v\n", response)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
```

While TypeScript and Rust are both excellent languages with their own strengths, Go's combination of fast startup time, low memory usage, and support for concurrency makes it a great choice for building serverless applications.

## Developer Experience

TypeScript, Rust, and Go are also known for providing a great developer experience, with features that can make it easier and more enjoyable for developers to write and maintain code. Here are some ways in which these languages provide a good developer experience:

### /w TypeScript

- Type safety that catches errors at compile time, rather than at runtime.
- Great tooling support, including powerful code editors, linters, and debugging tools.
- A large and growing ecosystem of libraries and frameworks, with strong support for web development and server-side development.
- A strong focus on developer productivity, with features like optional chaining and nullish coalescing that can make code more concise and readable.

However...

- Type annotations can add complexity and verbosity to the code, especially for simple or small-scale projects.  For example: `const add = (a: number, b: number): number => a + b;` versus `const add = (a, b) => a + b;`
- The learning curve for TypeScript can be steeper than for other languages, especially for developers who are not already familiar with JavaScript.
- Code editors and other development tools may require additional configuration to work properly with TypeScript, which can be time-consuming.

### /w Rust

- A powerful type system that ensures memory safety and eliminates data races.
- Great tooling support, including a powerful package manager and code editors with strong syntax highlighting and code completion.
- A focus on performance and efficiency, with low-level control over system resources and efficient memory management.
- A growing ecosystem of libraries and frameworks, with support for web development, systems programming, and game development.

However...

- Rust's strong type system and [borrow checker](https://doc.rust-lang.org/beta/rust-by-example/scope/borrow.html) can be challenging for developers who are not used to working with memory-safe languages, especially when dealing with complex data structures or concurrency.
- The syntax and language features in Rust can be more complex and verbose than in other languages, which can make it harder for some developers to learn and write code.
- The [Rust compiler can be slower than some other languages](https://stackoverflow.com/questions/37362640/why-does-rust-compile-a-simple-program-5-10-times-slower-than-gcc-clang), especially when dealing with large codebases or complex projects.

### /w Go

- A simple and intuitive syntax that is easy to learn and read.
- Great tooling support, including a powerful command-line interface and a growing ecosystem of code editors and IDEs.
- A focus on performance and concurrency, with lightweight threads and efficient memory management.
- A growing ecosystem of libraries and frameworks, with support for web development, distributed systems, and cloud computing.

However...

- Go's focus on simplicity and minimalism can make it less expressive than some other languages, which can make some code less elegant and harder to read.
- Go's type system can be less flexible than other languages, which can be limiting for some projects.
- Go's tooling and package management system can be less user-friendly than other languages, especially for developers who are used to more sophisticated tools and frameworks.

TypeScript, Rust, and Go all provide a great developer experience, with features that can make it easier and more enjoyable to write and maintain code. However, there are also potential challenges and drawbacks that developers may face when working with these languages. These include issues with verbosity, complexity, and learning curve, as well as challenges with tooling, type systems, and performance.

## Cost Efficiency

### Cost of Development

**Sticking with TypeScript** while not utilizing the benefits of switching to Rust or Go can still be a viable choice for some projects. TypeScript is a mature language with a large and growing ecosystem, strong tooling support, and a focus on developer productivity and safety. While Rust and Go offer some unique benefits, such as performance, low-level control, and efficient memory management, they may not be necessary for all projects. If a project doesn't require these features, or if the development team doesn't have the expertise or experience to use these languages effectively, then sticking with TypeScript can be a good choice.  Additionally, staying with TypeScript can be a good choice if the project already has a significant investment in TypeScript code, and the cost of migrating to a new language would be prohibitive. It's often easier and more efficient to continue using the same language and build on the existing codebase, rather than starting from scratch with a new language. TypeScript is an excellent choice for web development.

**Switching to Rust from TypeScript** can offer several benefits, especially for projects that require low-level control, performance, and memory safety. Rust is a systems programming language that provides developers with low-level control over memory allocation and management, without sacrificing safety or security.  However, switching to Rust can also have some challenges, especially for developers who are not used to working with memory-safe languages or who are not familiar with Rust's syntax and features. Rust can have a steeper learning curve than some other languages, and its strong type system and borrow checker can be challenging for developers to work with at first. Additionally, the tooling and ecosystem for Rust may not be as mature or robust as those for more established languages like TypeScript.  The return on investment (ROI) of sticking with TypeScript and losing out on the benefits of Go or Rust will depend on the specific needs of our shop and the goals of the project(s). TypeScript is a mature language with a large and growing ecosystem, strong tooling support, and a focus on developer productivity and safety.  Rust is an excellent choice for building network servers.

**Switching to Go from TypeScript** can offer several benefits, especially for projects that require high concurrency, low memory footprint, and fast compilation times. Go is designed to be a simple, efficient, and scalable language that makes it easy to write highly concurrent programs.  However, switching to Go can also have some challenges, especially for developers who are not used to working with statically-typed languages or who are not familiar with Go's syntax and features. Go's type system and error handling can be challenging for developers to work with at first, especially if they are used to more dynamic languages like TypeScript. Additionally, the tooling and ecosystem for Go may not be as mature or robust as those for more established languages like TypeScript.  The return on investment (ROI) of switching from TypeScript to Rust will depend on the specific needs of our shop and the goals of the project.  The learning curve and migration costs associated require a significant intestment in time and resources.

The ease of learning and use of Go can certainly make it easier for a development team to transition from JavaScript to Go. Go is a simple and intuitive language that is relatively easy to learn, especially for developers who are already familiar with C-like syntax. Additionally, Go's focus on simplicity and minimalism can make it easier to write and read code, which can make it faster and more efficient for developers to work with.  This ease of learning and use can translate into a measurable return on investment for a development shop that is converting from JavaScript to Go. For example, developers who are able to learn and use Go quickly may be able to ramp up on new projects more quickly, which can save time and reduce development costs. Additionally, the simpler and more efficient nature of Go code can reduce the amount of code that needs to be written and maintained, which can save time and reduce the risk of errors and bugs.

Go is excellent choice for building highly concurrent and distributed systems, such as network services, cloud-based microservices, and serverless applications.

### Cost of Operation - Serverless

Serverless applications can be a cost-effective way to build and run applications, but they can also be expensive if they are not designed and implemented properly. Here are some ways in which TypeScript, Rust, and Go can help to keep serverless applications cost-efficient:

- **Efficient resource utilization:** All three languages, TypeScript, Rust, and Go, have a strong focus on efficient resource utilization, which can help to reduce the overall cost of running serverless applications. This can be achieved through features such as efficient memory management, low memory overhead, and optimized code execution.
    a. **Rust** is often considered to be the most efficient language in terms of resource utilization because of its ownership model, which helps to ensure that memory is managed correctly and efficiently, reducing the risk of memory leaks and other performance issues. Additionally, Rust's built-in optimizer produces highly optimized code that can run efficiently on both CPUs and GPUs.
    b. However, both TypeScript and Go are also known for their efficient resource utilization. TypeScript benefits from the strengths of the underlying JavaScript engine, which has been optimized for years, and Go has built-in support for concurrency, lightweight threads, and a garbage collector, which can all help to ensure efficient memory management and optimized code execution.
- **Rapid startup times:** Serverless applications are designed to scale up and down quickly based on demand. This means that fast startup times are essential to ensure that the application can respond to changes in demand in a timely manner. TypeScript, Rust, and Go are all designed to provide fast startup times, which can help to keep serverless applications cost-effective.
    a. **Go** is often considered to be the fastest language for startup times because of its simplicity and minimalism. Go's lightweight threads, called "goroutines," can be created and destroyed quickly, and its simple syntax can help to reduce the time required for compilation and startup.
    b. TypeScript and Rust can also provide fast startup times. TypeScript's compilation process is generally fast, especially for small and medium-sized applications, and its support for incremental compilation can help to reduce the time required for recompilation. Rust also provides fast startup times because of its optimized code execution and its use of LLVM, a compiler infrastructure that can produce fast and efficient code.

- **Concurrency support:** Serverless applications often need to handle multiple requests simultaneously. Concurrency support is essential to ensure that the application can handle these requests efficiently, without incurring additional costs. TypeScript, Rust, and Go all have strong support for concurrency, which can help to keep serverless applications cost-effective.
    a. **Go** is often considered to be the strongest language for concurrency support because of its built-in support for lightweight threads called "goroutines," as well as its channels, which can be used to communicate between goroutines. This makes it easy to write highly concurrent and distributed systems that can handle multiple requests efficiently.
    b. TypeScript and Rust also have strong support for concurrency. TypeScript supports asynchronous programming through its support for Promises and async/await syntax, which can help to write highly concurrent applications that handle multiple requests efficiently. Rust's ownership model helps to ensure that code is thread-safe and free of data races, which can make it easier to write highly concurrent applications that handle multiple requests efficiently.

- **Smaller code size:** Smaller code size can help to reduce the overall cost of running serverless applications, as it reduces the amount of code that needs to be deployed and executed. TypeScript, Rust, and Go all have features that can help to reduce code size, such as built-in optimization and tree-shaking.
    a. **Rust** is often considered to be the best language for smaller code size because of its built-in support for static linking, which can help to reduce the size of the final executable. Additionally, Rust's ownership model helps to ensure that code is free of unused or unnecessary components, which can help to reduce the code size further.
    b. TypeScript and Go also have features that can help to reduce code size. TypeScript's incremental compilation can help to avoid recompiling unchanged code, which can help to reduce the size of the final code bundle. Go's built-in optimization and tree-shaking can help to reduce the size of the final executable by removing unused components.

- **Smaller container size:** Smaller container size can also help to reduce the overall cost of running serverless applications, as it reduces the amount of resources that are required to deploy and run the application. TypeScript, Rust, and Go all have features that can help to reduce container size, such as built-in support for static linking and compiled binaries.
    a. **Rust** is often considered to be the best language for smaller container size because of its built-in support for static linking and compiled binaries. This makes it possible to create fully self-contained binaries that can be deployed without requiring additional dependencies or runtime libraries. Additionally, Rust's ownership model helps to ensure that only the necessary components are included in the final executable, further reducing the container size.
    b. TypeScript and Go also have features that can help to reduce container size. TypeScript's incremental compilation can help to avoid recompiling unchanged code, which can help to reduce the size of the final container. Go's built-in support for static linking and compiled binaries can also help to reduce the size of the final container.

Overall, TypeScript, Rust, and Go can all help to keep serverless applications cost-effective by providing efficient resource utilization, rapid startup times, strong concurrency support, smaller code and container size, and other cost-saving features. However, the specific cost savings will depend on the individual application and the way it is designed and implemented.

## Conclusion

The choice of programming language is a critical factor in building performant, scalable, and cost-effective backends. In this document, we evaluated the performance, scalability, and developer experience of three popular programming languages - TypeScript, Rust, and Go - in the context of building distributed, serverless applications. While all three languages offer unique strengths and advantages, we found that Go stands out as the best choice for some backend projects.

Go's simplicity, efficiency, and ease of use make it an excellent choice for building scalable and performant serverless applications. Its support for concurrency and lightweight threads makes it easy to handle multiple requests simultaneously, while its efficient memory management and optimized code execution help to reduce costs.

- Go's support for lightweight threads can be beneficial is in the development of web servers that need to handle a large number of concurrent requests. By using goroutines, Go's lightweight threads, developers can easily handle thousands of connections simultaneously, allowing web servers to efficiently scale and handle high traffic loads without incurring additional costs. This approach can be much more efficient than using traditional operating system threads, which are heavier and more resource-intensive.
- In a serverless environment, efficient memory management and optimized code execution can help to reduce costs. Go's garbage collector automatically frees up memory that is no longer needed, while its compiled code executes quickly and efficiently, reducing the amount of compute resources needed to run the application. For example, a serverless function processing large amounts of data can run faster and more efficiently in Go, resulting in lower costs.

While TypeScript and Rust also offer many advantages for building serverless applications, Go's unique combination of simplicity and performance make it an excellent choice for specific use cases. We encourage developers to consider Go as a potential option when building performant, scalable, and cost-effective backends."

Of course, the decision to switch to using Go for some backend projects will depend on the specific requirements and constraints of each project, as well as the expertise and preferences of the development team.

## Addendum

### Paradigm Discussion - Functional versus Object-Oriented Programming

There will be a method to my madness, I promise.  Let's start with a quick overview of functional and object-oriented programming.

Functional programming emphasizes immutability, pure functions, and higher-order functions, which can lead to code that is easier to reason about, test, and parallelize. Functional programming is also well-suited for applications that deal with large amounts of data, such as data analysis, machine learning, and scientific computing. However, functional programming can have a steeper learning curve, and some developers may find it more difficult to write and maintain functional code compared to object-oriented code.  There are 3 key tenets of functional programming:

- **Immutability**: Immutable data cannot be changed after it is created. This makes it easier to reason about the state of a program, and it can also lead to better performance because the compiler can optimize away unnecessary copies of data.
- **Pure functions**: A pure function is a function that has no side effects and always returns the same output given the same input. Pure functions are easier to reason about and test, and they can be easily parallelized.
- **Higher-order functions**: A higher-order function is a function that takes a function as an argument or returns a function as a result. Higher-order functions can be used to abstract common patterns, which can lead to more modular and reusable code.

It's easy to see how these tenets can lead to code that is easier to reason about, test, and parallelize.  For example, consider the following code:

```javascript
const numbers = [1, 2, 3, 4, 5];
const doubled = numbers.map((number) => number * 2);
console.log(doubled);
```

This code uses the `map` function to double each number in the array. The `map` function takes a function as an argument, and it returns a new array containing the results of calling the function on each element in the original array. In this case, the function passed to `map` is a pure function that doubles its argument. This code is easy to reason about because it's clear that the `doubled` array contains the same number of elements as the `numbers` array, and it's also clear that the elements in the `doubled` array are twice as large as the elements in the `numbers` array.

Object-oriented programming emphasizes encapsulation, inheritance, and polymorphism, which can lead to code that is more modular, reusable, and extensible. Object-oriented programming is also well-suited for applications that deal with complex state and behavior, such as user interfaces, games, and simulations. However, object-oriented programming can lead to code that is more tightly coupled, and it can be more difficult to reason about the behavior of an object-oriented system as a whole compared to a functional system.  So, the 3 key tenets of object-oriented programming are:

- **Encapsulation**: Encapsulation is the process of hiding the implementation details of a class from the rest of the program. This makes it easier to change the implementation of a class without breaking other parts of the program that depend on that class.
- **Inheritance**: Inheritance is the process of creating a new class from an existing class. The new class inherits the behavior of the existing class, and it can also add new behavior. Inheritance can be used to reduce duplication in a program by sharing behavior across multiple classes.
- **Polymorphism**: Polymorphism is the process of reusing the same interface for multiple types. This can lead to more modular and reusable code because a function can operate on different types of data as long as they all implement the same interface.

For example, consider the following code:

```javascript
class Animal {
  constructor(name) {
    this.name = name;
  }

  speak() {
    console.log(`${this.name} makes a noise.`);
  }
}
```

This code defines an `Animal` class that has a `name` property and a `speak` method. The `speak` method prints a message to the console, but it doesn't actually say what noise the animal makes. This code is easy to reason about because it's clear that all animals have a `name` property and a `speak` method, and it's also clear that the `speak` method doesn't actually say what noise the animal makes.  Now, let's create a `Dog` class that inherits from the `Animal` class:

```javascript
class Dog extends Animal {
  speak() {
    console.log(`${this.name} barks.`);
  }
}
```

This code defines a `Dog` class that inherits from the `Animal` class. The `Dog` class overrides the `speak` method so that it prints a message to the console that says what noise a dog makes. This code is also easy to reason about because it's clear that all dogs have a `name` property and a `speak` method, and it's also clear that the `speak` method prints a message to the console that says what noise a dog makes.

Why do I bring this up in the context of TypeScript, Rust, and Go?  Some functional programming concepts are better suited for writing highly concurrent and distributed systems, which are common in many modern backends. While TypeScript, Rust, and Go are not strictly functional languages, they do offer some functional programming features and can be used to build highly concurrent and distributed systems. These languages are well-suited to building backend systems that require high availability and fault tolerance.  OOP concepts can cause problems in highly concurrent and distributed systems because of its reliance on mutable state and shared data. In a concurrent system, multiple threads or processes may need to access and modify the same data simultaneously, which can lead to race conditions, deadlocks, and other types of synchronization issues.  This is not to say that OOP is inherently bad or unsuitable, but focusing on immutability and pure functions will mitigate the aforementioned issues.

**TypeScript** provides many features that make it easier to write large-scale applications, including classes, interfaces, and modules. TypeScript also support functional programming concepts like lambdas and higher-order functions. Additionally, TypeScript is often used with Node.js to build server-side applications, which can be highly concurrent and distributed.

**Rust** provides many features that make it easier to write highly concurrent and distributed systems, including lightweight threads and an ownership model that ensures memory safety and eliminates data races. Rust also supports functional programming concepts like closures and iterators, which can make code more concise and easier to reason about.

**Go** is purposefully designed to make it easy to write highly concurrent and distributed systems. Go provides lightweight threads, called goroutines, and a simple syntax that makes it easy to write concurrent code. Go also includes features like channels, which can be used to communicate between goroutines, and a garbage collector, which can simplify memory management. While Go is not a functional language, it does support some functional programming concepts, like first-class functions and closures.

Languages such as C#, Java, and Python were not considered for this distributed application discussion because, while they are also widely used for building backend systems, they may not be as well-suited for highly concurrent and distributed systems as TypeScript, Rust, and Go. C# and Java are primarily object-oriented languages, although they do offer some support for functional programming concepts. Python, on the other hand, is a general-purpose language that is often used for data science and web development, but it is not specifically designed for building highly concurrent and distributed systems.

In distributed, serverless applications, where the system may be composed of multiple processes or functions running on different machines, the potential for concurrency issues is even greater. This is because each process or function may be running independently and concurrently, with the potential for shared data or resources. In this context, functional programming can be particularly useful for reducing the complexity of concurrent and distributed systems, by emphasizing immutability and pure functions that avoid shared state and side effects.

Also, C#, Java, and Python are not specifically designed for building serverless applications, although they can still be used for this purpose. Serverless applications typically require a different architectural approach, with a focus on event-driven and stateless computing. Languages like TypeScript, Rust, and Go are well-suited for building serverless applications because they provide support for concurrency, fault tolerance, and memory safety, which are important considerations for serverless computing. Additionally, these languages provide efficient runtime performance, which is critical for serverless applications that are often charged based on resource consumption.

### Other languages not considered - Ruby and Elixir

While Ruby and Elixir are both powerful languages that are well-suited for building highly concurrent and distributed systems, they do have some potential drawbacks.

One potential disadvantage of Ruby is its runtime performance. Ruby is an interpreted language, which can make it slower than compiled languages like C++ or Rust. Additionally, Ruby's garbage collector can sometimes cause performance issues, especially in large-scale systems. While Ruby has a large ecosystem of libraries and frameworks, it may not have the same level of support for building highly concurrent and distributed systems as other languages like Elixir.  A very timely article titled ["Whatever Happened to Ruby?"](https://www.infoworld.com/article/3687219/whatever-happened-to-ruby.html) was published when I started writing this document, and it provides a good overview of the current state of Ruby.

Elixir, on the other hand, is a relatively new language, and it may not have the same level of community support as more established languages. Additionally, while Elixir provides excellent support for concurrency, it may not be as well-suited for building systems that require a lot of mutable state. Finally, while Elixir runs on the Erlang virtual machine, which provides excellent support for fault tolerance and distributed computing, it may not have the same level of support for low-level systems programming as other languages like Rust.  Elixir has a different syntax and programming paradigm than many other popular languages, which can make it challenging for developers who are not familiar with functional programming to learn.  Still, if you want to learn more about the pros and cons of Elixir, I recommend reading this article titled ["The Pros & Cons of Elixir Programming Language"](https://www.rswebsols.com/tutorials/programming/pros-cons-elixir-programming-language).
