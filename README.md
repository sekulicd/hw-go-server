# Trafik Reverse Proxy with Let's Encrypt and Docker Services

This repository provides a simple example on how to use Trafik as a reverse proxy for Docker services, securing your services with SSL certificates from Let's Encrypt. In this setup, we run two instances of a simple "Hello, World!" GoLang web server, each bound to a different subdomain.

## Prerequisites

Ensure that you have Docker and Docker Compose installed on your machine. You also need a registered domain name, as we'll be issuing SSL certificates for specific subdomains.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/sekulicd/hw-go-server.git
    ```

2. Modify the `docker-compose.yml` file:
    - Replace `dusan.sekulic.mne@gmail.com` with your email address
    - Replace `hw1.dusansekulic.me` and `hw2.dusansekulic.me` with your desired subdomains

3. Run Docker Compose:
    ```bash
    docker-compose up -d
    ```

## Usage

After the Docker Compose setup has finished, you can access your services through your specified subdomains, they should be served over HTTPS, thanks to the Let's Encrypt certificates issued by Trafik.

## Architecture

The `traefik` service is our entrypoint and routes incoming requests to the appropriate service based on the request's Host header.

It's configured to automatically generate Let's Encrypt SSL certificates for our specified subdomains. The generated certificates are stored in the `./letsencrypt/acme.json` file.

The `hw1` and `hw2` services are simple "Hello, World!" GoLang web servers. They're each built from the Dockerfile included in the repository and are assigned a unique subdomain through Docker Compose labels.

## Notes

Remember, the Let's Encrypt CA server used here is the staging server. For production, you should use the production server `https://acme-v02.api.letsencrypt.org/directory`.

Also, note that the external network 'web' is required to be pre-created before running the `docker-compose up -d` command. If it doesn't exist, create it with `docker network create web`.
