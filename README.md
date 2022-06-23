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

## Working with struct methods
You can define methods to the output structs and those methods can be used to calculate value of any field.

There are a few conditions of defining methods

1. Methods must be value based not pointer based
2. Methods must be public
3. Add a method tag to the struct field having value equal to function name
4. Make sure the return type is same to the field you want to cast

### Working Example

```go
type Book struct {
    Id int64
    Title string
    Discount float64
}

// suppose your discount was in float 0-1
// and now you want to return discount by multiplying it with 100
// it can be done in your presentation layer like below

type BookEntity struct {
    Id int
    BookTitle string `entity:"Title" json:"Title"`
    Discount float64 `entity:"Discount" json:"-"`

    FinalDiscount int `method:"CreateDiscount" json:"Discount"`
}

func (b BookEntity) CreateDiscount() {
    b.FinalDiscount = int(100 * b.Discount)
    return b.FinalDiscount
}

response := Book{
    Id: 1,
    Title: "Hooked",
    Discount: 0.5,
}

c.Json(http.StatusOk, goentities.Present(response, BookEntity{}))
```

Just see the magic. With this it's easy to create derived fields in the presentation layer instead of writing them in bussiness logic.

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