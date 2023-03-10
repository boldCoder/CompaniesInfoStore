## Run Project:
---------------

### Run in Docker Container:
----------------------------
1. Install Docker in your machine.
2. Navigate to the repository directory, open terminal and run: `docker-compose up`.
3. In Postman app, use url as `localhost:9000/[followed by the endpoints mentioned below]`.
4. Run `docker-compose down` to remove the containers.
5. At the end, run `docker image prine -a`.

### Run in Local Machine:
1. Install golang in your system.
2. Copy git repo on your local machine. [Repo Link](https://github.com/boldCoder/CompaniesInfoStore)
3. Open terminal, navigate to the folder, where the repo is downloaded.
4. To install project dependencies, Run: `go mod tidy` 
5. Run: `go build -o main ./cmd/`
6. Run: `./main`


### Authentication:
-------------------
1. Open Postman app. 
2. In order to generate JWT token, first signup with email and password.
3. After the user is successfully created, you can login ang get an authentication token in the Cookie. 
4. After successful login, now we can use the generated token to `create`, `update`, `fetch` and `delete` operation.
5. In order to perform the `CRUD` operations: 
    1. In Postman app, go to `Authorization` tab, choose `API Key` as Type.
    2. Provide Key as `Authentication` and Value as token generated while looging in, for each request for requesting a resource.



### Endpoints:
--------------
    1. For SignUp:
    -----------
    Method: POST 
    Endpoint: host:port/user/signup 
    Body:
    ```
    {
    "email":"gaurav1@gmail.com",
    "password":"qwerty12"
    }
    ```


    2. For Login:
    -----------
    Method: POST 
    Endpoint: host:port/user/login 
    Body:
    ```
    {
    "email":"gaurav1@gmail.com",
    "password":"qwerty12"
    }
    ```


    3. Create Resource:
    ----------------
    Method: POST 
    Endpoint: host:port/company/create 
    Body:
    ```
    [
        {
        "name": "Demo",
        "description": "This is company A description",
        "employee_strength": 50,
        "registered": true,
        "type": "Corporations"
        }
    ]
    ```

    4. For Update a resource:
    ----------------------
    Method: PATCH 
    Endpoint: host:port/company/update
    Body:
    ```
    {
    "name": "Demo",
    "description": "Patch Call in description",
    "employee_strength": 253,
    "registered": true,
    "type": "Sole-Proprietorship"
    }
    ```


    5. For Get resource by ID:
    -----------------------
    GET /company/get/{id}
    Method: GET 
    Endpoint: host:port/company/get/{id} 


    6. For Delete resource by ID:
    --------------------------
    Method: DELETE 
    Endpoint: host:port/company/delete/{id}

   