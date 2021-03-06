# Api pagar.me

## OBS: Quando clonado em maquina com sistema operacional Windows o formato do arquivo run.sh está sendo modificado para CLF ele deve ser convertido novamente para LF

pagar-me-test é uma api rest desenvolvida com base nas tecnologias Golang e banco de dados Neo4j.

DOCUMENTAÇÃO DA API: https://leonardo-vieira.stoplight.io/docs/pagarme-test/7fgxmcwkxtz2m-pagarme

## Tech
#### pagar-me-test usa:

* [Golang]
* [Neo4j]
* [Docker]
* [Docker Compose]

Código publicado no GitHub

## Endpoints:
  - POST /person/create : Cria um nó do tipo pessoa
  - DELETE /person/delete/{person_id} : remove um nó do tipo pessoa
  - POST /person/kinship : busca qual o grau de parentesco entre  os nós
  - POST /person/relashion : Cria um relacionamento (Pai/Filho) entre dois nós
  - POST /person/bacon_number : Busca o bacon number entre dois nós
  - GET /person/genealogy/{person_id} : Busca a arvore Genealogica do nó informado

  As features de não permitir relacionamento incestuoso e de busca em nivel de descendentes foram implementadas.
 
### Instalação e Build

pagar-me-test precisa de [Docker](https://docs.docker.com/desktop/windows/install/) para rodar.
pagar-me-test precisa de [Docker compose](https://docs.docker.com/desktop/windows/install/)
pagar-me-test precisa de [Golang](https://docs.docker.com/compose/install/) na versão go1.18.3 .

### Rodando o projeto
```sh
$ docker compose up 
```
Após a execução desse processo a API estara rodando no endereço: http://localhost:80/

### Acessando o Banco Neo4j

O Neo4J pode ser acessado pelo endereço: http://localhost:7474/browser/
Nesse endereço é possivel ver os nós e seus relacionamentos como na imagem a seguir:

![neo4j](https://i.ibb.co/MDYb4Fh/neo4j.png)

## OBS: a senha de acesso ao neo4j e o username podem ser encontradas dentro do arquivo .env
