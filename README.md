
![W3Care](https://w3.care/assets/images/w3care-logo-4.svg)

  

# Desafio Backend - Solução
 

## Rodando a solução
Após clonar, na raiz do projeto:

    docker-compose up
### Pgadmin
* O pgMyAdmin roda no endereço: http://localhost:16543
* As informações de login e senha estão no docker-compose.yml
  ## Explicando a Solução
Foram criados 3 projetos: 
* Serviço de usuários e filas, em Java
* Servidor de mensagens, em nodejs
* frontend, em Angular 9
### Serviço de usuários
* Rodando no http://localhost:8080
* Uma aplicação feita em SpringBoot fornece o CRUD de usuários, a autenticação e o gerenciamento de filas
* O banco de dados utilizado foi o postgres
* A fila é persistida na tabela chat_queue, que também funciona como log, com data inicio, data fim e status
#### Documentaçao
* Foi criado um swagger com a documentaçao do serviço em [http://localhost:8080/swagger-ui.html](http://localhost:8080/swagger-ui.html)
#### Autenticação
* A autenticação é realizada através de Spring Security
* A autorização é feita através de um token JWT, sendo assim, stateless
#### Limitações
* Não foi feito teste unitário (o tempo foi um fator importante nesse quesito)
* Alguns pequenos fluxos estão descritos na camada de Controller. Entendo que não é o ideal.
### Servidor de mensagens
* Rodando em http://localhost:3000/
* Uma aplicação simples escrita em nodejs + socket.io
* Funciona simplesmente como um broker entre quem manda mensagem e quem recebe a mensagem
#### Limitações
* Não possui testes unitários
* Não foi implementada autenticação. Iria usar o mesmo token jwt para autorizar (ou nao) o envio de mensagens. O tempo me limitou nessa implementação
### Frontend
* Rodando em http://localhost:8081
* Feito em Angular + Bootstrap
* Uma chat plenamente funcional em tempo real entre médicos e pacientes
#### Fluxo
* Quando um paciente loga, este entra numa fila (FIFO)
* Quando um médico loga ele tem a possibilidade de iniciar um novo atendimento
* Esse novo atendimento carrega o primeiro usuário que esteja esperando na fila e abre uma sala de chat
* Ao final do atendimento o médico tem a opçao de Encerrar esse atendimento. A sala é fechada e o médico pode solicitar outro paciente que esteja na fila. Caso não haja mais pacientes esperando, uma mensagem é exibida.
## Pontos de atenção
* As portas de configuração de backend estão **FIXAS** no frontend, portanto não devem ser alteradas no docker-compose.yml
* Entendo que este não é o ideal, mas devido ao tempo, foi a melhor solução que cheguei pra esse momento.