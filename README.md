## Dev-Client-Go

**dev-client-go** is a client library for the Forem (dev.to) [developer api](https://developers.forem.com/api) written in Go. It provides fully typed methods for every operation you can carry out with the current api (beta)(0.9.7)

### Installation
> Go version >= 1.13
```sh
$ go get github.com/Mayowa-Ojo/dev-client-go
```

### Usage
Import the package and initialize a new client with your auth token(api-key).
To get a token, see the authentication [docs](https://developers.forem.com/api#section/Authentication)
```go
package main

import (
   dev "github.com/Mayowa-Ojo/dev-client-go"
)

func main() {
   token := <your-api-key>
   client, err := dev.NewClient(token)
   if err != nil {
      // handle err
   }
}
```

<hr style="border:1px solid gray"> </hr>

### Documentation
Examples on basic usage for some of the operations you can carry out.

#### Articles [[API doc](https://developers.forem.com/api#tag/articles)]
Articles are all the posts that users create on DEV that typically show up in the feed.

Example:

**Get published articles**

query parameters gives you options to filter the results 
```go
// ...
// fetch 10 published articles
articles, err := client.GetPublishedArticles(
   dev.ArticleQueryParams{
      PerPage: 10
   }
)

if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Articles: \n%+v", articles)
// ...
```

**Create an article**

you can pass the article content as string by setting the `Article.Body` field, or as a markdown file by passing the `filepath` as a second parameter
```go
// ...
payload := dev.ArticleBodySchema{}
payload.Article.Title = "The crust of structs in Go"
payload.Article.Published = false
payload.Article.Tags = []string{"golang"}

article, err := client.CreateArticle(payload, "article_sample.md")
if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Article: \n%+v", article)
// ...
```

#### Organizations [[API doc](https://developers.forem.com/api#tag/organizations)]
Example:

**Get an organization**
```go
// ...
organization, err := client.GetOrganization(orgname)
if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Organization: \n%+v", organization)
// ...
```

**Get users in an organization**
```go
// ...
users, err := client.GetOrganizationUsers(
   orgname,
   dev.OrganizationQueryParams{
      Page:    1,
      PerPage: 5,
   },
)

if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Users: \n%+v", users)
// ...
```

#### Comments [[API doc](https://developers.forem.com/api#tag/comments)]
Example:

**Get comments for article/podcast**
```go
// ...
comments, err := client.GetComments(
   dev.CommentQueryParams{
      ArticleID: articleID,
   },
)

if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Comments: \n%+v", comments)
// ...
```

**Get a single comment**
```go
// ...
comment, err := client.GetComment(commentID)
if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Comment: \n%+v", comment)
// ...
```

#### Listings [[API doc](https://developers.forem.com/api#tag/listings)]
Example:

**Get published listings**
```go
// ...
listings, err := client.GetPublishedListings(
   dev.ListingQueryParams{
      PerPage:  5,
      Category: "cfp",
   },
)

if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Listings: \n%+v", listings)
// ...
```

**Get a single listing**
```go
// ...
listing, err := client.GetListingByID(listingID)
if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Listings: \n%+v", listings)
// ...
```

<hr style="border:1px solid gray"> </hr>

### API methods
Here's a list of all methods matching every operation currently supported. Clicking them will also take you to the location in the test file to see usage examples.

[**Articles**]
[x] [GetPublishedArticles](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L11)

[x] [CreateArticle](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L75)

[x] [GetPublishedArticlesSorted](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L101)

[x] [GetPublishedArticleByID](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L131)

[x] [UpdateArticle](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L156)

[x] [GetPublishedArticleByPath](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L187)

[x] [GetUserArticles](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L207)

[x] [GetUserPublishedArticles](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L235)

[x] [GetUserUnPublishedArticles](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L265)

[x] [GetArticlesWithVideo](https://github.com/Mayowa-Ojo/dev-client-go/blob/main/articles_test.go#L295)

