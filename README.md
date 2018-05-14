# go-to-do
A simple 'To Do' lists application.
The App server exposes REST APIs to perform CRUD for the To-Do lists.

### Tech Stack
App server is in GoLang which uses [echo](https://echo.labstack.com/) as web framework.
MongoDB is used as persistant database to store lists and tasks.

### Build and Run
`docker-compose up --build`

### APIs
##### Create List
**Method**: POST

**Route**: /lists

**Content type**: application/json

**Request Params**

Key | Type | Value (example)
--- | --- | ---
*name* | `string` | office


##### Get All Lists
**Method**: GET

**Route**: /lists


##### Get List
**Method**: GET

**Route**: /lists/:id

##### Delete List
**Method**: DELETE

**Route**: /lists/:id


##### Create Task
**Method**: POST

**Route**: /tasks

**Content type**: application/json

**Request Params**

Key | Type | Value (example)
--- | --- | ---
*listId* | `string` | ObjectId
*message* | `string` | move tasks on jira board

