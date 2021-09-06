# Golang API REST
> Projeto criado para simulação de um banco de dados de uma livraria.

## Linguagem do projeto
Projeto feito com a linguagem golang e o framework fiber.


## Inicialição do projeto
-Com a ferramenta Go já instalada em seu computador clone o projeto e abra o terminal na pasta src e digite `ro run main.go`.



## Rotas
Projeto criado usando API REST com 4 verbos HTTP `GET, POST, PUT e DELETE` fazendo requisições para um banco de dados NoSQL.


### Método GET
- Para consultas poderá ser utilizados 3 parâmetros : "id", "title", "author". Utilize no seu testador de requisições de sua preferência uma rota por vez:
`http://localhost:{PORT}/api/books/:id?`,
`http://localhost:{PORT}/api/books/title/:title?` ou
`http://localhost:{PORT}/api/books/author/:author?`.
OBS: Caso não tenha parâmetros a consulta ainda sim poderá ser utilizada mas no modo global.
o retorno será a lista de livros que deseja.

### Método POST 
- Para adicionar um novo livro no banco de dados com o testador de sua preferência coloque na rota POST : 
`http://localhost:{PORT}/api/books`
 e no body:
 
```json
{
	title: "Exemplo",
	author: "Exemplo2",
	isbn10: "Exemplo3",
}
```
Se for um sucesso a requisição deverá retornar:
```json
{
	"InsertedID":"IDEexemplo"
}
```

### Método PUT
- Para atualizar um **livro**, você terá que usar a rota `http://localhost:{PORT}/api/books/:id` no método `PUT`, substituindo o `:id` pelo identificador do livro,  com a mesma estrutura de dados do **método POST**.
Se for um sucesso a requisição deverá retornar:
```json
{
	"MatchedCount": N,
  "ModifiedCount": N,
  "UpsertedCount": N,
  "UpsertedID": null
}
```
### Método DELETE 
- Para deletar um **livro**, você terá que usar a rota `http://localhost:{PORT}/api/books/:id` no método `DELETE`, substituindo o `:id` pelo identificador do livro.

Se for um sucesso a requisição deverá retornar:
```json
{
	"DeletedCount": 1
}
```

## Ferramentas utilizadas
- golang
- fiber
- mongodb
- insomnia
