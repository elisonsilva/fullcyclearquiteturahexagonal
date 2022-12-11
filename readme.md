# Módulo de Arquitetura Hexagonal Curso Full Cycle

```bash

# Start Docker
docker-compose up -d

# Execute Docker
docker exec -it -u 0 appproduct bash

# Go Tests
go test ./...

# 3. Go Server
go run main.go http
# ... open Postman


# Commands by CobraCli

# View commands
go run main.go cli --help

# Create
go run main.go cli -a=create -n="Product Cli2" -p=25.0

# Get by ID
go run main.go cli -a=get -i={ID_STRING}

# Update Enable or Disable
go run main.go cli -a=enable -i={ID_STRING}
go run main.go cli -a=disable -i={ID_STRING}

```

# Problems solution


```bash
# Permission denied
# apt-get install sqlite3
# E: Could not open lock file /var/lib/dpkg/lock-frontend - open (13: Permission denied)
# E: Unable to acquire the dpkg frontend lock (/var/lib/dpkg/lock-frontend), are you root?
docker exec -it -u 0 appproduct bash
```


[[1](https://forum.code.education/forum/topico/cobra-init-nao-funciona-1416/)]
```bash
# Cobra init não funciona

# Só precisa executar agora
cobra-cli init

# Para adicionar novo comando:
cobra-cli add cli

# Para executar o novo comando:
go run main.go cli
```



# Testes Postman

```bash
# Postman
# http://localhost:9000/product

#Send method POST and body raw type Json

{
    "name" : "Testando do postman",
    "price" : 25.99
}

# Response
{
    "ID": "....",
    "Name": "Testando do postman",
    "Price": 25.99,
    "Status": "disabled"
}
```