#!/bin/bash
SECONDS=0

cd $HOME/banking-app

msg () {
  echo -e "$1\n--------------------\n"
}

msg "Pulling from GitHub"
git fetch --all
git reset --hard origin/main
chmod +x prod_deploy.sh

msg "Building Docker image"
sudo docker build --tag banking-app .

# msg "Stopping Docker container"
# sudo docker stop banking-app
# sudo docker rm banking-app

# msg "Starting Docker container"
# sudo docker run \
# -d \
# --name banking-app \
# --expose 443 \
# -p 443:443 \
# -v /etc/letsencrypt:/etc/letsencrypt \
# -e SERVER_ENV=PROD \
# banking-app

# msg "Starting Postgres container"
# sudo docker run -d \
# --name db-postgres \
# -p 5432:5432 \
# --mount type=volume,src=app-db,target=/var/lib/postgresql/data \
# -e POSTGRES_PASSWORD=mysecretpassword \
# postgres:15.1-alpine

msg "Stopping containers"
sudo docker compose down

msg "Starting containers"
sudo docker compose up -d

msg "Pruning stale Docker images"
sudo docker image prune -f

duration=$SECONDS

echo
msg "Deploy finished in $(($duration % 60)) seconds."
msg "Press enter to exit"
read