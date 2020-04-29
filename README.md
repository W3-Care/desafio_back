![W3Care](https://w3.care/assets/images/w3care-logo-4.svg)

# Desafio Backend
Com grandes poderes, vem grandes chocolates suíços... ou alguma coisa assim...


Neste desafio, o propósito é avaliar o raciocínio empregado na resolução de problemas e as tecnologias empregadas na solução.

## Problema
De acordo com nosso lema: **Um segundo faz toda diferença!**, o desafio consiste em desenvolver um client (app ou web) que se comunique em tempo real com um servidor e outro client de forma autenticada...


Beleza, n etendi nada...


### Contextualizando
Sua solução deve conter uma **API REST** que autentica um usuário (login e senha) e retorna um token, id e categoria de forma que estes dados sejam utilizados para se comunicar (autenticadamente) em tempo real com outro usuário.


#### Usuários
Os usuários serão divididos entre **Paciente** e **Médico**. 

#### Comunicação entre os usuários
Toda vez que um **Paciente** for autenticado, ele deve ir para uma sala de espera, onde irá esperar até um **Médico** o chamar para conversa (texto plano).

###### Chat
Apenas um **Médico** e um **Paciente** podem estar na mesma sala de chat, ou seja, o **Médico** incia a conversa com um **Paciente** o removendo da fila de espera.

###### Finalizando
O **Médico** finalizará a sala de chat desconectando o **Paciente** atual e chamando o próximo da fila. Ao finalizar, deve-se enviar: dataInicio, horaInicio,dataFim,horaFim, idPaciente, idMedico para a **API REST** e persistir o log.


## REQUISITOS
1. A fila deve ser persistida, ou seja, deve ser gerenciada por um banco de dados
2. Deverá existir um CRUD de usuários
3. Utilize Docker
4. Beba água


### Legal ter
* Interface gráfica
* Documentação
* Testes automatizados

# Como submeter
Faça um fork deste repositório e quando finalizar, faça um **pull request** com o seu nome.