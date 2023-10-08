# Introduction

# Stack

# How to Run the App Locally on a Windows PC

1. Ensure you have at least Go version 1.20.6 installed
2. Have Docker Desktop installed
3. Build the postgres database image running the following command

```
docker run -d `
--name db-postgres `
-p 5432:5432 `
--mount type=volume,src=app-db,target=/var/lib/postgresql/data `
-e POSTGRES_PASSWORD=mysecretpassword `
postgres:15.1-alpine
```

4. You will have to manually access the database container to create a users table and populate it with records. A database migration and seeding script is on my to do list.

5. Run the following command to build the Docker image and to upload it to Docker Hub

`docker build --tag banking-app .`

6. Run `docker compose up -d` to run both images on a connected network
7. Access http://localhost:8080 to check the home page is being served
