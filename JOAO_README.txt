Não consegui gerar o Docker, a maquina que estou utilizando é da empresa em que trabalho e a 
virtualização está bloqueada, estou 350km do meu notebook pessoal.

Para rodar:
1 - executar "go run main.go" com um mongodb sendo executado no localhost, caso queira 
escolher outro host para o BD basta alterar o arquivo config.toml.

2 - Após a súbida faça um GET pelo browser mesmo em http://localhost:3000/api/v1/users/auto/config
    A aplicação irá criar 4 usuários, são eles "patient1", "patient2", "doctor1", "doctor2"
    Todos com a mesma senha "secret"

3 - Para acessar a tela de chat acesse http://localhost:3000/simplechat/

4 - Estou disponibilizando junto ao projeto uma collection do postman com todos os metodos disponíveis.

Como funciona:

1-Ao logar com um "Patient" será criada uma fila de chamada
2-Ao logar com um "Doctor" será disponibiliad uma opção para iniciar o atendimento do primeiro paciente da fila
3-A ordem dos fatores não altera o produto, deslogar/logar mantem o chat ativo
4-O chat só é encerrado por comando do "Doctor"
5-A conversa não é armazenada, porém todos as demais informações são persistentes

Features:
- Desenvolvimento em GoLang
- Autenticação por JWT
- WebSockets
- Persistencia em MongoDB
- CRUD completo de Usuários, Filas e Salas de Chat
- Interface em VUE

Como informei no batepapo, sou especialista em Java com conhecimentos avançados em outras linguagens
mas nunca desenvolvi em Golang ou Flutter, conforme conversamos.

Com o intuito de demonstrar um pouco da minha capacidade de aprendizado, resolvi fazer em GoLang o desafio.
Trabalhei um total de 13hs no projeto, saindo do ZERO de conhecimento em Golang.
Tentei gerar a interface em Flutter, mas depois de 2h eu desisti e lancei uma interface simples em VUE.
Espero que considere o tempo que tive para aprender a linguagem e desenvolver o desafio, 
não tive tempo habil para fazer bem feito, apenas atender o máximo da necessidade no tempo disponível
e na linguagem solicitada.

Não tive tempo de trabalhar em testes automatizados ou Documentação.

Caso tenha interesse estou disponível para demonstrar outros projetos melhor elaborados, testados e documentados.

Aguardo suas considerações.