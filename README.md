#  🦦| RegistrationConsultationOfBooks

### Sobre o projeto: 

O projeto teve início em uma feira de ciências simples da minha escola, durante o segundo ano do ensino médio. Inicialmente, o projeto consistia apenas em persistir dados no lado do cliente (ClientSide) via localStorage, utilizando as tecnologias básicas da web (HTML, CSS, TypeScript, Webpack). Com uma proposta da direção para a construção de um sistema mais robusto, foi concebido este projeto.

Este servidor tem como objetivo solucionar problemas internos relacionados à consulta de livros na biblioteca pública da instituição de ensino Professora Maria Angelina Gomes.

### Tecnologias:

> ⚠ Obs: Todas as tecnologias foram escolhidas levando em consideração a qualidade de software, custo de produção, tempo e afinidade;

[![GoLang Badge](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![MySql Badge](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
[![JWT Badge](https://img.shields.io/badge/jwt-181818?style=for-the-badge&logo=json-web-tokens&logoColor=yellow)](https://jwt.io/) 
## 👨‍💻 | Modelagem de software

> ⚠ Obs: Este modulo é destinado a recrutadores e demais colegas desenvolvedores que desejam se inspirar no desenvolvimento desta aplicação. Vale lembrar que esse é um dos meus primeiros projetos após 2 anos de estudos intensos. Não leve muito em considerações os erros de arquitetura de repositórios ou boas práticas de programação. 

### RFs (Requisitos Funcionais)
- ✅ Deve ser possível que a instituição possa criar usuários;
- 🔴 Deve ser possível que o usuário pegue seu perfil;
- ✅ Deve ser possível cadastrar livros;
- 🔴 Deve ser possível fazer a consulta de um livro especifico;
- 🔴 Deve ser possível obter a relação de quantidade todos os livros;
- 🔴 Deve ser possível criar Alunos para a consulta de livros;
- 🔴 Deve ser possível o Aluno realizar o check-in dos livros;
### RNs (Regras de negócio)

- ✅ Só poderá haver um aluno com determinado Email;
- 🔴 O aluno só poderá consultar um livro por vez;
- 🔴 O aluno terá que entregar o livro em até 15 dias;

### RNFs (Requisitos NÃO-Funcionais)
- ✅ A senha do usuário precisa estar criptografada(hash);
- ✅ Os dados da aplicação precisam estar persistidos em um db MySQL;
- ✅ Caso não exista a tabela deverá ser criada;
