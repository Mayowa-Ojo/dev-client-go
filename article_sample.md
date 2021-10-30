### Introduction

Go doesn't support inheritance in the classical sense; instead, in encourages composition as a way to extend the functionality of types. This is not a notion peculiar to Go. **Composition over inheritance** is a known principle of OOP and is featured in the very first chapter of the Design Patterns book.

Embedding is an important Go feature making composition more convenient and useful. While Go strives to be simple, embedding is one place where the essential complexity of the problem leaks somewhat. In this series of short posts, I want to cover the different kinds of embedding Go supports, and provide examples from real code (mostly the Go standard library).

There are three kinds of embedding in Go:

- Structs in structs (this part)
- Interfaces in interfaces (part 2)
- Interfaces in structs (part 3)

#### Embedding structs