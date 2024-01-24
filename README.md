# 🦦 | RegistrationConsultationOfBooks

### Sobre o projeto:

O projeto teve início em uma feira de ciências simples da minha escola, durante o segundo ano do ensino médio. Inicialmente, o projeto consistia apenas em persistir dados no lado do cliente (ClientSide) via localStorage, utilizando as tecnologias básicas da web (HTML, CSS, TypeScript, Webpack). Com uma proposta da direção para a construção de um sistema mais robusto, foi concebido este projeto.

Este servidor tem como objetivo solucionar problemas internos relacionados à consulta de livros na biblioteca pública da instituição de ensino Professora Maria Angelina Gomes.

- Por enquanto o projeto é apenas uma app monolítica, futuramente um arquitetura de microservices.
Estamos estudando formas de implementar um micro serviço de emails.

### 🤖 | Tecnologias:

> ⚠ Obs: Todas as tecnologias foram escolhidas levando em consideração a qualidade de software, custo de produção, tempo e afinidade;

[![GoLang Badge](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![GoLang Badge](https://img.shields.io/badge/Echo-00ADD8?style=for-the-badge&logo&logoColor=white)](https://go.dev/)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
[![JWT Badge](https://img.shields.io/badge/jwt-181818?style=for-the-badge&logo=json-web-tokens&logoColor=yellow)](https://jwt.io/)

## 👨‍💻 | Modelagem de software

> ⚠ Obs: Este modulo é destinado a recrutadores e demais colegas desenvolvedores que desejam se inspirar no desenvolvimento desta aplicação. Vale lembrar que esse é um dos meus primeiros projetos após 2 anos de estudos intensos e apenas poucas semanas de golang. Não leve muito em considerações os erros de arquitetura de repositórios ou boas práticas de programação.

### Entities

- User

`
    UserEntity: {
        id: cuid,
        username: string,
        adm: boolean,
        email: email,
        password_hash: hash
    } `

- Book

`
    BookEntity: {
        id: uuid,
        name: string,
        author: string,
        quantity: int8,
        available: int8, 
    } `

### RFs (Requisitos Funcionais)

- [ ] Deve ser possível que a instituição possa criar registrar um livro

- [ ] Deve ser possível que a instituição posso criar usuários para registro de dados.

- [ ] Deve ser possível que a instituição possa listar os livros.

- [ ] Deve ser possivel fazer a consulta de livros

- [ ] Deve ser possivel cadrastar um email pra aluno

### RNs (Regras de negócio)

- [ ] Só deve ser possivel criar um usuário se o usuário for admim

- [ ] Só deve ser possível que a cusulta seja feita se o livro estiver disponivel

- [ ] Só deve ser possível criar um usuário se o email estiver disponivel

- [ ] O aluno só poderá consultar o livro se cadastrar um email junto ao livro

- [ ] O aluno só poder ficar com o livro por 15 dias.

### RNFs (Requisitos NÃO-Funcionais)

- [ ] A senha do usuário precisa estar criptografada(hash);
- [X] Os dados da aplicação precisam estar persistidos em MongoDb;


version: 0.2.3