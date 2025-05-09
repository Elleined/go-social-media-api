# go-notes
Notes for Go

# Naming conventions
- Package names: lowercase and single word
- File names: lowercase all
- Variables and Constants:
    - Local: camelCase
    - Exported: PascalCase
- Functions
    - Local: camelCase
    - Exported: PascalCase
- Structs and Interface: PascalCase

# Data types
- Any type: denote with `any` and `interface{}`
- Numeric types
  Both positive and negative numbers
    - int: default
    - int8: -128 to 127
    - int16: -32,768 to 32,767
    - int32: -2,147,483,648 to 2,147,483,647
    - int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
  ```go
  var age int = 18
  or
  age := 18
  ```
  Positive numbers
    - unint: default
    - uint8: 0 to 255
    - uint16: 0 to 65,535
    - uint32: 0 to 4,294,967,295
    - uint64: 0 to 18,446,744,073,709,551,615
  ```go
  var age uint = 18
  or
  age := 18
  ```
  Decimal numbers
    - float32: Approx. 6-7 decimal digits
    - float64: Approx. 15 decimal digits
  ```go
  var age float64  = 18.1
  or
  age := 18.1
  ```
    - complex64, complex 128: Real or Imaginarytnumbers

- Text types
    - string: String literally with double-qoutes ""
  ```go
  var name string = "Juan"
  or
  name := "Juan"
  ```
    - rune: Character with single-qoutes ``
  ```go
  var initial rune = 'J'
  or
  initial := 'J'
  ```
    - byte: byte

- bool: True or false
  ```go
  var isPresent bool = true
  or
  isPresent := true
  ```
- pointer: Stores the memory address of a variable (Will be elaborated more later)

- Data Structure Types
    - slice: Dynamic size array just like ArrayList in java
  ```go
  names := []string {
    "Juan",
    "Pedro"
  }
  ```
    - map: key-value pair
  ```go
  persons := map[string] int {
    "Juan": 25,
    "Pedro": 20
  }
  ```
    - struct: same concept with java classes
  ```go
  type Person struct {
    Name string
    Age int
  }

  person := Person {
    Name: "Juan",
    Age: 25
  }
  ```
    - interface: same concept as intefaces in java
  ```go
  type Human interface {
    Walk()
  }
  ```
    - channel: used for go routines communication.
        - chan T: can send and receive data of type T
        - chan <- T: Can only send data of type T
        - <- chan T: Can only receive data of type T
  ```

## Most use datatypes
- int (defaults 0)
- unint
- float64 (defaults 0.0)
- string (defaults "")
- rune (defaults '')
- bool (defaults false)
- struct, map, slice, and interface (defaults nil)
- map, slice, and channel should declared with make() function

## When to use specific numerical dataypes
- Unles you have a good reason. It is recommended to use the default numerical datatype `int` and `uint`
- When the performance optimization are concerned use the specific numerical datatypes

# Variables
- Please note that unused variable are not allowed in go it will cause an compilation error.
- Long hand syntax
```go
var name string = "Juan";
```

- Short hand syntax
  In short hand syntax the type is inferred automatically. With the use of `:=`. Which means it automatically detects the datatype based on given value, In this example it will be inferred as string
```
name := "Juan"
```

# Constants
- constant as the name itself it cannot be changed
- Also worth noting that constant should be defined at compile time
- So constant can be only created or derived from other constants.
- Its good to declare constant in all caps.
```go
const NAME string = "Red"
```

# Format Specifiers
| Specifier  | Description                        | Example Output                        |
|------------|------------------------------------|--------------------------------------|
| `%v`       | Default format (any value)         | `42`, `true`, `{Alice 30}`           |
| `%+v`      | Struct with field names            | `{Name:Alice Age:30}`                |
| `%#v`      | Go syntax representation           | `main.Person{Name:"Alice", Age:30}`  |
| `%T`       | Type of the value                  | `int`, `string`, `main.Person`       |
| `%d`       | Integer (base 10)                  | `123`                                |
| `%b`       | Binary representation              | `1111011`                            |
| `%o`       | Octal representation               | `173`                                |
| `%x` / `%X`| Hexadecimal (lowercase / uppercase)| `7b` / `7B`                          |
| `%f`       | Floating point (decimal)           | `3.141593`                           |
| `%e` / `%E`| Scientific notation                | `3.141593e+00` / `3.141593E+00`      |
| `%s`       | String                             | `Hello`                              |
| `%q`       | Double-quoted string               | `"Hello"`                            |
| `%p`       | Pointer address                    | `0xc000012080`                       |
| `%c`       | Character (ASCII value)            | `'A'` (for `65`)                     |
| `%t`       | Boolean                            | `true`, `false`                      |


# IF with short statement
```go
Syntax:
if statement; expression {
   // body
}

if score := 10; score > 10 {
   fmt.Printf("Score is greater than %d", score)
}
```
- Here the statement is only scoped within the if statement and cannot be access outside. This is useful for the variable that are only used once for if statement one liner.

# Function
## Function Signature
```go
func name(parameters...) (returns...) {

}

// Without any parameters and returns
func foo() {

}

// With parameters and returns
// Multiple data type declaration
func foo(bar, foo int) (int, int) {
    return bar, foo
}
```

## Named Returns
- Is basically a return with a name and this was used for more code readability and named return is initialize with default value for primitives and nil with objects.
```
func bar(param1, param2 int) (ret1, ret2 int) (
    
    return ret1, ret2
}
```

## Receiver function
- Usually used inside struct to mimic other programming languange methods inside classes, since go doesn't have class go receiver function is the go work around for the struct to have a method inside.
- Basically saying here's a function `Name() {}` attached this to `func (u *User)` with final result of.
```go
func (u *User) Name() {

}
```
- So now its like this user has now this method. also can also be denoted as this user has received this method thats why its called receiver function.

### Example of receiver function
```go
type User struct {
  id   int
  name string
}

func (u *User) Name() {
      return name
}

func main() {
   user := Person { 1, "foo" }

   // To use the receiver function is:
   user.Name() 
}
```
(u *User): is the method receiver. and just like saying `this`.

Name(): is the method.

### Conclusion for receiver function
- Used to mimic method inside class in other programming languange.

## Anonymous function
a function without a name, often assignes as variable or called immediately. Just like in js.
```go
greet := func(name string) {
    fmt.Println("Hello", name)
}
greet("Alice")
```

## Closures function
a returned function that takes variable to its outer scope
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```
## Higher order function
a function that takes a function as parameter or return them
usecase is just like stream api methodd in java. map, filter, and etc...
```go
func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

add := func(x, y int) int { return x + y } // Anonymous function
result := operate(3, 4, add)
```


# Pointer
```go
/*
1. x holds the value of 10.
2. And has the memory address of 10101.
*/
x := 10
// outputs 10

/*
1. &x just like saying get the memory address of x (10101)
2. ptr holds the memory address of x (10101)
*/
ptr := &x
// outputs 10101

/*
1. *ptr just like saying get the value of memory address that ptr pointer pointing to.
2. And ptr pointer points to memory address of (10101)
3. And memory address of (10101) is x and has a value of 10.
4. Which outputs to 10.
*/
*ptr
// outputs 10
```
`*` is basically saying that get the value of that memory address. Basically saying that "What is the value of memory address that ptr pointer pointing to".

`*` And when you see `*variable` think of it just like interacting to the real object itself.

`&` is saying that get the memory address of the variable. Basically saying that saying that "Where is the x"

Another example:
```go
func main() {
  x := 10
  fmt.Println("Before modifying: ", x)
  // Pass the memory address of x
  modify(&x)
  fmt.Println("After modifying: ", x)
}

func modify(num *int) {
    *num = 15
}

/*
outputs
10
15
*/
```

Another example:
Check the swap.go it covers the referencing and dereferencing of a pointer.

## Benefits of using pointers
- Avoids copying of large data sets and only passing its pointer so that any modification happens inside a function it will also affect the original copy. Instead of passing its value in function it means that you just created 2 whole same large dataset and when youre done inside the function you return the modified copy and reassign that to original variable so much work here.
- Basically you just passing the address (which is just a random alphanumeric values) instead of a full copy.

## Best use case of pointer
- When working with maps, struct, slices, channels, and array. Because its avoids passing the whole large data sets inside a function instead only working wiith its memory address and modify it accordingly and efficiently.

# Defer, Panic, and Recover
# defer keyword
1. For code cleanup.
2. Executed after the function call.
3. Acts as final block in try catch in any other language. Meaning its like saying "no matter happens in this function, make sure to run this code".
4. defer is executed in LIFO manner, meaning the last deferred funtion is executed first. Like a down to up
5. And when the defer keyword is executed it will be computed right there.
6. Direct replacement for try catch and only as finally block in the essense
   Example:
```go
defer fmt.Println("This will be printed third")
defer fmt.Println("This will be printed second")

fmt.Println("This will be printed first")

// output
This will be printed first
This will be printed second
This will be printed third
```

# panic keyword
1. When panic is executed the rest of the code will not be executed and all the deferred keyword met along the way will be executed.
2. Only use for fatals error where the program cannot be recoverable.
3. It's like saying "something has gone so wrong here that normal execution cannot continue safely."

## difference of error and panic
1. panic is like error in java where program cannot be recoverable after an error.
2. error is like exception in java where program are recoverable after an error.

Example:
```go
defer fmt.Println("This will be executed even theres a panic")

fmt.Println("A panic is about to happen")
panic("Somethin went wrong")

// output

A panic is about to happen
This will be executed even theres a panic
panic: Something went wrong
```

# recover keyword
- Must be inside a defer function to use.
- So basicallly recover is used to execute the code even after a panic occurs
- The return value of recover() will the value the passed on the panic(x any) function
- Also used for monitoring, logging, and debugging.
- When theres no recover() after a panic() the program will be terminated
- Acts as safety net after encountering a panic
  Example:
```go
func recoverSample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic: %v", r)
		}
	}()

	fmt.Println("Starting the function...")

	panic("Something went wrong!")
	
	fmt.Println("This will never be executed!")
}

func main() {
	fmt.Println("Before having a panic")
	recoverSample()
	fmt.Println("After having a panic")
}

// output
Before having a panic
Starting the function
Recover from panic: Something went wrong
After having a panic
```

## Conclusion about defer, panic, and recover
1. defer is for code cleanup and logging
2. panic is for signaling unrecoverable error or the program is invalid state and should be terminated.
3. recover is for handling a panic to gracefully handle and prevent system crashes and also used for logging.

# exit keyword
1. Exits the program immediately. Thus ignoring all the defer, panic, and recover keyword.

# init keyword
1. No parameter and return values
2. Executed before the main function
3. Only be used once per package
```go
func init() {
	fmt.Println("Initializing this package")
}

func main() {
	fmt.Println("From main funtion")
}

// outputs
Initializing this package
From main function
```

## Conclusion
1. Don't use panic and recover all over the place. use panic for unrecoverable error else use the error and return a error in recoverable situation
2. Use panic and recover by pair is a good practice

# Go database connectivity methods
## sql package method
1. Open(): returns sql.DB, and used to open a database connection
   // username:password@tcp(host:port)/database
   sql.Open("mysql", "root:root@tcp(localhost:3306)/myDB")

## sql.DB methods
1. Ping(): checks the database connection
2. Close(): close database connection `defer db.Close()`
3. Exec(): returns sql.Result, and commonly used for INSERT, UPDATE, and DELETE
4. Query(): returns sql.Rows, and commonly used for SELECT

## sql.Rows methods
1. Next(): moves the cursor to the next row.
2. Scan(): copies the values of the current row into the provided variables (usually pointers). and commonly used after the .Next() method inside a while loop
3. Close(): close the rows connection `defer rows.Close()`

## sql.Result methods
1. LastInsertedId() (int, error): Usually used after an INSERT
2. RowsAffected() (int, error): Usually used after an UPDATE and DELETE

# What is goroutines
1. Is a lightweight thread managed by Go runtime
2. Enables execution of functions allowing you to perform multiple tasks in a single go program.
3. Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value
4. Goroutines are non blocking
5. Goroutines runs asynchronously basically their execution are not defines as expected it depends whoever goroutine finished first the will be executed.

## goroutine syntax
- created with `go` keyword
```go
go myFunction(arguments...)
```

## Important note!
1. When main thread is terminated the goroutines will be terminated immediately. Use channels instead.

## FAQS about Goroutine
1. Goroutine cannot have return value, you must use a channel to receive and send data through goroutines safely.
2. You can use anonymous function for goroutines

## Goroutine common pitfalls
1. Exiting the main thread before the goroutines finishes. In sample use the time.Sleep(). Basically give time for the go routines to finish before the main thread.
2. Multiple goroutines accessing the same variable. Use channel, sync.Mutex, or sync/atomic
3. Using variables outside of scope. The fix is pass the variable as parameter in goroutine function instead always.

# What is channel
1.

## When to close the channel
1. Only close the sending channel not a receiving channel

## Why use channels
1. Prevents race conditions
2. Avoid shared memory locks
3. Enforce message-passing concurrently in a safer manner.

# Conclusion about goroutines and channels
1. Goroutines run tasks concurrently and channel are used for the goroutines to synchonize and communicate.
2. Goroutines don't return any values and channel provide a way to pass values safely.
3. Can run in thousand asynchonously and channel helps coordinate the goroutines interactions

# Analogy of goroutines and channel working together
- Its like goroutine is the worker and the walkie-talkie are their channel. They cannot shout to each other. They will use the walkie-talkie(channel) to communicate and coordinate effectively while working independently and concurrently.

# WaitGroup
1. Is just use to block the main thread for our goroutines to fiinsh executing before closing the main thread

# File and Folder
## packages
1. path/filepath: For pathing
    - Abs(): returns the absolute path of the file
    - Ext(): returns only the file extension name
    - Clean(): removes unecessary characters
    - Join(): join the paths with os specific operator
    - Base(): returns only the last element
2. os: For CRUD of files and folder