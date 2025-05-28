# User-PIP

Este projeto é uma aplicação web simples em Go que exibe o endereço IP do usuário, a hora da requisição e o nome do pod (quando executado em Kubernetes).

## Funcionalidades

- Mostra o IP real do visitante (considerando proxies).
- Exibe a data e hora da requisição.
- Mostra o nome do pod (útil para testes em ambientes Kubernetes).
- Interface simples em HTML.

## Como executar localmente

1. **Pré-requisitos:**  
   - [Go](https://golang.org/dl/) 1.24+
   - [Docker](https://www.docker.com/) (opcional)

2. **Executando diretamente com Go:**
   ```sh
   go run main.go
   ```
   Acesse [http://localhost:8080](http://localhost:8080) no navegador.

3. **Executando com Docker:**
   ```sh
   docker build -t user-pip .
   docker run -p 8080:8080 user-pip
   ```

## Deploy no Kubernetes

1. Crie o namespace:
   ```sh
   kubectl apply -f k8s/namespace.yaml
   ```

2. Faça o deploy:
   ```sh
   kubectl apply -f k8s/deployment.yaml
   kubectl apply -f k8s/service.yaml
   ```

3. Acesse via NodePort em `http://<NODE_IP>:30080`

## Estrutura do Projeto

```
.
├── main.go
├── go.mod
├── Dockerfile
├── nuke.sh
├── templates/
│   └── index.html
└── k8s/
    ├── deployment.yaml
    ├── namespace.yaml
    └── service.yaml
```

## Limpeza de recursos Docker

Use o script `nuke.sh` para remover todos os containers, imagens, volumes e redes Docker (use com cuidado):

```sh
./nuke.sh
```

## Licença

MIT