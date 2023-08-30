## Propósito
A ideia desse projeto é facilitar a minha organização pessoal das tarefas que eu fiz e quanto tempo levei para encerrar.

## Como usar
Para usar esse projeto, basta clonar o repositório e subir os caontainers usando o docker-compose.

```bash
git clone
cd todo-api
docker-compose up -d
```

Após isso, você pode executar o arquivo `main.go` usando os containers que acabou de subir.

```bash
docker compose exec -it server sh

cd ../app
go run main.go
```
