# Criar o container:
docker run --name fineasy-postgresql -e \\ 
POSTGRES_USER=gabrielroriz -e \\
POSTGRES_PASSWORD=12345678 -e \\
POSTGRES_DB=fineasydb \\
-p 5432:5432 \\
-d postgres:latest

# Se já exisitr o container, então:
docker start fineasy-postgresql

