# Experiment on mkcert
Use mkcert to create SSL certificate for localhost.

## Environment
- Webserver (nginx)
- App server (golang)

## Setup
- Self-signed certificate must be created on the host machine to access web
  server from the host machine.  
  Run the below command to download `mkcert` and create certificate
  ```
  ./web/download-mkcert.sh
  ./web/create-cert.sh
  ```
