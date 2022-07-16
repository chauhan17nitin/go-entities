# go-entities

![Version](https://img.shields.io/github/v/tag/chauhan17nitin/go-entities)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/chauhan17nitin/go-entities)](https://pkg.go.dev/github.com/chauhan17nitin/go-entities)

go-entities is a support to api responses structs. It sits on the the top of model structs to prevent exposing unwanted fields in api response.

## Introduction
Sometimes you can return some sensitive data via your api if you directly returns your model struct in the response.

So it's always adviced to cast you models to api specific responses. go-entities helps you in casting your model structs to api specific responses to prevent leaking sensitive data. 

## What are Entities?
Entities are meant to be reusable response formats / (structs for go). Entites can be used to cast fields, add extra fields during final presentation output, conditional casting, adding derived fields in response, nesting with other entities and much more.

### Basic Example
List of fields you want to return in api response

```go
type User struct {
	UserName string `json:"UserName"`
	Name     string `json:"Name"`
	Location string `json:"Location"`
}
```

## Using go-entities
As Entites are meant to be reusable so they should be present in a specific folder so that different controllers can use them.

.

    ├── entities                   # Entities (represents response types)
    ├── controllers                # controllers for api's
    ├── classes                    # Bussiness logic layer
    └── README.md

Under entities you can have all the entities which you wants to use in your service. (Surely you can do sub structuring under entities too if required).

Below is how a user enitity will look
entities/user.go
```go
type User struct {
	UserName string `json:"UserName" entity:"UserName"`
	Name     string `json:"Name" entity:"Name"`
	Location string `json:"Location" entity:"Location"`
}
```

While under classes User class will surely have more fields classes/user.go
```go
type User struct {
	ID       int
	UserName string
	Name     string
	Location string
	Phone    string
}
```

It's pretty sure that in fetch basic details you will not want to return Id, Phone of user, so in controller you can cast your class struct to entities struct like below.

```go
user := userHandle.GetUserBasicDetails(c)
c.Json(http.StatusOk, goentities.Present(user, entities.User{}))

```

It works using the value defined in entity tag in the User entity struct.

## Deriving fields using struct methods

You can define methods to the output structs and those methods can be used to calculate value of any field.

There are a few conditions of defining methods

1. Methods must be value based not pointer based
2. Methods must be public
3. Add a method tag to the struct field having value equal to function name
4. Make sure the return type is same to the field you want to cast

### Working Example

classes/book.go
```go
type Book struct {
    Id int64
    Title string
    Discount float64
}
```

entities/book.go
```go
// suppose your discount was in float 0-1
// and now you want to return discount by multiplying it with 100
// it can be done in your presentation layer like below

type BookEntity struct {
    Title string `entity:"Title" json:"Title"`
    Discount float64 `entity:"Discount" json:"-"`

    FinalDiscount int `method:"CreateDiscount" json:"Discount"`
}

func (b BookEntity) CreateDiscount() {
    b.FinalDiscount = int(100 * b.Discount)
    return b.FinalDiscount
}

```

```go
response := Book{
    Id: 1,
    Title: "Hooked",
    Discount: 0.5,
}

c.Json(http.StatusOk, goentities.Present(response, BookEntity{}))
```

Just see the magic. With this it's easy to create derived fields in the presentation layer instead of writing them in bussiness logic.

## Nesting structs

entities
```go
type struct User {
	UserName string `json:"UserName" entity:"UserName"`
	Name     string `json:"Name" entity:"Name"`
	Location string `json:"Location" entity:"Location"`
    Address  Address `json:"Address" entity:"Address"`
}

type struct Address {
    City string `json:"City" entity:"City"` 
    State string`json:"State" entity:"State"`
    PinCode int `json:"PinCode" entity:"Pin"`
}
```

classes 
```go
type User struct {
	ID       int
	UserName string
	Name     string
	Location string
	Phone    string
    Address Address
}

type struct Address {
    City  string
    State string
    Pin   int
}
```

## Working with Pointers
Currently we support pointers in input struct only not in entity struct.

Like Below example is invalid

```go
type Book struct {
    Id int64
    Title string
    Author string
}

type BookEntity struct {
    BookTitle *string `entity:"Title"`
    AuthorOfBook *string `entity:"Author"`
}
```

But something like this will work 

```go
type Book struct {
    Id int64
    Title *string
    Author *string
}

type BookEntity struct {
    BookTitle string `entity:"Title"`
    AuthorOfBook string `entity:"Author"`
}
```

```go
bookTitle := "Hooked"
bookAuthor := "Nir Eyal"
response := Book{
    Id: 1,
    Title: &bookTitle,
    Author: &bookAuthor,
}

// this will work
c.Json(http.StatusOk, goentities.Present(response, BookEntity{}))

// this will also work
c.Json(http.StatusOk, goentities.Present(&response, BookEntity{}))
```