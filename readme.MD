# Introduction

I am wanting to become more familiar with Golang, Docker and DevOps so I will be combining it with my interest in Fintech. I will be creating basic functionalities of a digital bank to get more hands on practice.

Things I hope to achieve:

- Build basic prototypes of banking functionalities with Golang
- Automate deployments
- User Docker to containerise functionalities and services
- Explore Kubernetes and how it can fit in to a multi-containered application
- Monitoring dashboards

# Stack

# How to Run the App Locally on a Windows PC

1. Ensure you have at least Go version 1.20.6 installed
2. Have Docker Desktop installed
3. You will have to manually access the database container to create a users table and populate it with records. A database migration and seeding script is on my to do list.
4. Run the `docker build --tag banking-app .` to build the Docker image locally
5. Run `docker compose up -d` to run both images on a connected network
6. Access http://localhost:8080 to check the home page is being served

## How to Access the database locally

1. Ensure you have the database container running. You can chekc by running `docker ps`
2. Run `docker exec -it banking-app-db-1 bash` to start an interactive terminal session with the running container
3. Connect to the psql server by running `psql -U postgres`
4. You can run SQL commands to check the tables
