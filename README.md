# go-entities

go-entities is a support to api responses structs. It sits on the the top of model structs to prevent exposing unwanted fields in api response.

## Examples

```go
type Book struct {
    Id int64
    Title string
    Author string
}

type BookEntity struct {
    BookTitle string `entity:"Title"`
    AuthorOfBook string `entity:"Author"`
}
```

Now you don't want to expose the Id field to the frontend, then you can just present your response with goentities

```go
// let's suppose response is the data which you are getting from service layer and now you want to present it using your book entity
response := Book{
    Id: 1,
    Title: "Hooked",
    Author: "Nir Eyal",
}

// Now you can present the response using go-entities like below
c.Json(http.StatusOk, goentities.Present(response, BookEntity{}))
```

Boom you are done, it will only output the exported fields from Book struct and will fill the BookEntity struct using that and return the data accordingly.

## Limitations
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
```

But something like these will work 

```go
type Book struct {
    Id int64
    Title *string
    Author *string
}

type BookEntity struct {
    BookTitle *string `entity:"Title"`
    AuthorOfBook *string `entity:"Author"`
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