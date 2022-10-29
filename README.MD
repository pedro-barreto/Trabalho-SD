 - Na pasta backend fica o banco de dados, para poder usar ele basta ter a database "escola".
 - Não precisa fazer tabela ou campo, somente o banco, pois o código fara a auto migration.
 - Após isso basta executar o main.go para rodar o banco.

 * Banco: mysql
 * User: root
 * Porta: 3306(padão)

 - OBS: use alguma API para testes, como "insomnia" ou "postman".
 - Os links para o CRUD:

 * GET Tudo    ->  localhost:8080/aluno/tudo
 * GET por ID  ->  localhost:8080/aluno/{id}
 * POST        ->  localhost:8080/aluno/
 * PUT         ->  localhost:8080/aluno/{id}
 * DELETE      ->  localhost:8080/aluno/{id}

 -> Observamos o banco pelo phpMyAdmin, baixando o xampp.
 -> Observamos o CRUD pelo "postman".
 -> O REST foi feito usando as bibliotecas "net/http" e "gorm"

 - No frontend nos fizemos uma simples chamada do banco, para testar a interação entre os dois containers.
 - Basta executar o "main.go" para exibir no terminal todos os cadastros do banco.

 *** OBS: Nos iremos detalhar mais os dados para a próxima etapa, nosso foco nessa etapa foi mais a interação ***