# Assan Elnur 20B030463

# golang_project: Online book store
 A website for selling books online. It allows you to get acquainted with books in an interesting category, get information about their authors and make online sales.

## REST API:
```
http://localhost:8081/api/v1

Category:
GET    /category
POST   /category
GET    /category/:id
PUT    /category/:id
DELETE /category/:id
GET    /category/:id/books

Book:
GET    /book
POST   /book
GET    /book/:id
PUT    /book/:id
DELETE /book/:id

Auth:
POST    /users
PUT     /users/activated
POST    /users/login

```
## DB Structure:
```
TABLE categories (
    id          bigserial [PRIMARY KEY],
    created_at  timestamp,
    updated_at  timestamp,
    title       text,
    description text
);

TABLE authors (
    id              bigserial [PRIMARY KEY],
    created_at      timestamp,
    updated_at      timestamp,
    name           text,
    info           text,
    age            integer
);

TABLE books (
    id              bigserial [PRIMARY KEY],
    created_at      timestamp,
    updated_at      timestamp,
    title           text,
    description     text,
    price           integer,
    category        bigserial,
    author          bigserial,
    FOREIGN KEY (category)
        REFERENCES categories(id),
    FOREIGN KEY (author)
        REFERENCES authors(id)
); 

REF: books.author < author.id
REF: books.category < category.id
```
