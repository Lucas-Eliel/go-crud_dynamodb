Feature: Conferindo Cenários da API

  Scenario: Lista todos os usuários
    Given os usuarios salvos no banco de dados
    When quando eu chamar a api na rota /users
    Then deve retornar os usuarios salvos no banco de dados