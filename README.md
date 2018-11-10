# CLI-fineasy

CLI-fineasy is a command-line-interface software to manage cash flow using **Go** with **PostgreSQL**.

## Setup Environment

 1. Execute this Queries in your PostgreSQL:

    CREATE DATABASE fineasy;
    CREATE TYPE fluxType as enum('expense', 'income');
 
 2. Define environment variables with your DB configs on this way:

**FIN_DB_HOST**="host"
**FIN_DB_PORT**="port"
**FIN_DB_USER**="user"
**FIN_DB_PASSWORD**="password"

## Database scheme



