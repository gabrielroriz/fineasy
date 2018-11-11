
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

  

## Commands

### List  Flows

By default, this command will list all flows from current month and year.

    fin -lf -m=<month> -y=<year>
    
### List  Categories
    fin -lc
    
  ### List  Sources
    fin -lc
    
  ### List  Wallets
    fin -lw
    
## Database scheme

![alt text](https://github.com/gabrielroriz/cli-fineasy/blob/master/schema/db-schema.png)