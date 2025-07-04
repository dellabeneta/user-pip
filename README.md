## User-PIP

![Pipeline](https://github.com/dellabeneta/user-pip/actions/workflows/main.yaml/badge.svg)
![GHCR version](https://img.shields.io/badge/ContainerVersion-v44-blue)

Uma aplicação web simples em Go para exibir o endereço IP do usuário, a hora da requisição e o nome do pod/servidor (útil para ambientes Kubernetes).

### Aviso Importante

**Este projeto é destinado a fins de teste, demonstração e aprendizado.** Não utilize para coleta de dados sensíveis em produção sem as devidas adequações de segurança.

### Funcionalidades

- Exibe o IP real do visitante (considerando proxies)
- Mostra a data e hora da requisição
- Exibe o nome do pod/servidor
- Interface web responsiva e minimalista
- Pronto para Docker e Kubernetes

### Começando

**Pré-requisitos**
- Go 1.24+  
- Docker (opcional)  
- Kubernetes (opcional)

**Instalação Local**
```bash
git clone https://github.com/seu-usuario/user-pip.git
cd user-pip
go run main.go
```
A aplicação estará disponível em `http://localhost:8080`

**Usando Docker**
```bash
docker build -t user-pip .
docker run -p 8080:8080 user-pip
```

**Deploy no Kubernetes**
```bash
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
Acesse via NodePort em `http://<NODE_IP>:30080`

### Tecnologias

- **Backend**: Go 1.24+
- **Frontend**: HTML5, CSS3 (inline)
- **Container**: Docker Alpine
- **Orquestração**: Kubernetes

### Como Funciona

O serviço exibe o IP do visitante, a data/hora da requisição e o nome do pod/servidor, útil para troubleshooting e testes em ambientes distribuídos.

### Configuração

O serviço pode ser configurado através das seguintes variáveis de ambiente:

| Variável | Descrição         | Padrão |
|----------|-------------------|--------|
| PORT     | Porta do servidor | 8080   |

### Estrutura do Projeto

```
della@ubuntu:~/projetos/user-pip$ tree
.
├── Dockerfile
├── go.mod
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── LICENSE
├── main.go
├── nuke.sh
├── pip.png
├── README.md
└── templates
    ├── favicon.ico
    └── index.html

3 directories, 12 files
della@ubuntu:~/projetos/user-pip$
```

### Scripts Úteis

**nuke.sh**: Script para limpeza completa do Docker (containers, imagens, volumes e redes)

```bash
chmod +x nuke.sh
./nuke.sh
```

### Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais