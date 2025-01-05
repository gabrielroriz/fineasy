# Criar o container:
docker run --name fineasy-postgresql -e \\ 
POSTGRES_USER=gabrielroriz -e \\
POSTGRES_PASSWORD=12345678 -e \\
POSTGRES_DB=fineasydb \\
-p 5432:5432 \\
-d postgres:latest

# Se já exisitr o container, então:
docker start fineasy-postgresql


# docker exec -it fineasy-postgresql bash

# psql -U gabrielroriz -d fineasydb

# SELECT ARRAY_AGG(ano) AS anos
# FROM (
#     SELECT DISTINCT EXTRACT(YEAR FROM created_at) AS ano
#     FROM flows
# ) subquery;