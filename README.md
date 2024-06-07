# GO API

This is a simple **CRUD** API for maintaining a list of posts. It is built using **Golang** and **PostgreSql**.

## Installation

1. Clone the repository and navigate to the directory

```bash
git clone https://github.com/ahmedYasserM/goapi.git
cd goapi
```

2. build the docker image of the api

```bash
docker compose build
```

3. Run the docker containers

```bash
docker compose up -d
```

4. The API will be available at `http://localhost:7000`

## API Endpoints

1. **GET** `/posts` - Get all posts
2. **GET** `/posts/{id}` - Get a single post
3. **POST** `/posts` - Create a new post
4. **PATCH** `/posts/{id}` - Update a post
5. **DELETE** `/posts/{id}` - Delete a post

## API Payload

1. **POST** `/posts` - Create a new post

```json
{
  "author": "Post Title",
  "content": "Post Content"
}
```

2. **PATCH** `/posts/{id}` - Update a post

```json
{
  "author": "Post Title",
  "content": "Post Content"
}
```

or

```json
{
  "author": "Post Title"
}
```

or

```json
{
  "content": "Post Content"
}
```
