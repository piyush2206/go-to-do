# go-to-do
A simple 'To Do' lists application.
The App server exposes REST APIs to perform CRUD for the To-Do lists.

## Tech Stack
App server is in GoLang which uses [echo](https://echo.labstack.com/) as web framework.
MongoDB is used as persistant database to store lists and tasks.

## Build and Run
`docker-compose up --build`

## APIs
### Create List
**Method**: POST

**Route**: /lists

**Content type**: application/json

**Request Params**

Key | Type | Value (example)
--- | --- | ---
*name* | `string` | office


### Get All Lists
**Method**: GET

**Route**: /lists


### Get List
**Method**: GET

**Route**: /lists/:id

**Note**: `id` to be taken from GET /lists API response


### Delete List
**Method**: DELETE

**Route**: /lists/:id

**Note**: `id` to be taken from GET /lists API response


### Create Task
**Method**: POST

**Route**: /tasks

**Content type**: application/json

**Request Params**

Key | Type | Value (example)
--- | --- | ---
*listId* | `string (ObjectId)` | 5349b4ddd2781d08c09890f3
*message* | `string` | Move tasks on jira board

**Note**: `listId` to be taken from GET /lists API response


### Update Task
**Method**: PUT

**Route**: /tasks

**Request Params**

Key | Type | Value (example)
--- | --- | ---
*taskId* | `string (ObjectId)` | 5349b4ddd2781d08c09890f3
*message* | `string` | Review tasks on jira board

**Note**: `taskId` to be taken from GET /lists/:id API response


### Delete List
**Method**: DELETE

**Route**: /tasks/:id

**Note**: `taskId` to be taken from GET /lists/:id API response


