<div align="center">
  <img alt="arbif" src="https://github.com/thudf/arbif-api/assets/54460874/536dd845-4741-4169-878a-d3bb9cb85c9a" width="auto" heigth="auto"/>
</div>

</br>

<div align="center">

  <p align="center" >
    <a href="#-sobre"> Sobre </a> |
    <a href="#-tecnologias"> Tecnologias </a> |
    <a href="#-iniciar"> Iniciar </a> |
    <a href="#-executar"> Executar </a> |
    <a href="#-licença"> Licença </a>
  </p>

</div>

## 📚 Sobre

Projeto desenvolvido para a prática de conceitos de backend de forma prática.</br>
O projeto tem por objetivo criar uma aplicação fictícia que se assemelha as entregas feitas no dia a dia da RV.

## 🚀 Tecnologias

As princiais tecnologias utilizadas na construção da API:

- [Go](https://go.dev/)
- [MySQL](https://www.mysql.com/)
- [uuid](https://pkg.go.dev/github.com/google/uuid@v1.3.1)
- [Mux](https://pkg.go.dev/github.com/gorilla/mux@v1.8.0)

## 💻 Iniciar

Para importar a documentação das rotas no Postman, clique no botão abaixo:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10025318-83870870-1392-4c84-b746-ee3473a19be2?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D10025318-83870870-1392-4c84-b746-ee3473a19be2%26entityType%3Dcollection%26workspaceId%3Dc2caa706-33b6-4923-a730-6a702d3d04aa#?env%5BENV_ARBIF_LOCAL%5D=W3sia2V5IjoidXJsIiwidmFsdWUiOiJodHRwOi8vbG9jYWxob3N0OjUwMDAiLCJlbmFibGVkIjp0cnVlLCJ0eXBlIjoiZGVmYXVsdCIsInNlc3Npb25WYWx1ZSI6Imh0dHA6Ly9sb2NhbGhvc3Q6NTAwMCIsInNlc3Npb25JbmRleCI6MH1d)

### Requisitos

- [Go](https://go.dev/)
- [Colima](https://github.com/abiosoft/colima)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Iniciar DB
```sh
# Verificar o status do colima
$ colima status

# Iniciar o colima, caso não esteja rodando
$ colima start

# Verificar se a instância do mysql está rodando
$ docker ps

# Iniciar instância caso não esteja rodando
$ cd .devenv
$ docker-compose up -d
```

## :zap: Executar
#### Clone o projeto e acesse a pasta
```sh
$ git clone https://github.com/thudf/arbif-api.git && cd arbif-api
```
#### Configure as variáveis ambiente
```sh
$ cp .env.example .env
```
#### Iniciando a aplicação
```sh
# Instale as dependências
$ go mod tidy && go mod vendor

# Inicie a aplicação
$ go run main.go
```

## 🖊 Licença

Este projeto é desenvolvido sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para saber mais detalhes.

</br>
</br>
</br>

<div align="center">
  <img src="https://media.giphy.com/media/9IZjYtYKV1103Pnu56/giphy.gif" width="200px">
  <p><sup>By <a href="https://www.linkedin.com/in/arthur-d-afonseca-885757183/">Arthur D'Afonseca</a><sup></p>
</div>
