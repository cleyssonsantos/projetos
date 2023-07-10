# **Sistema de Gerenciamento de Tarefas (To-Do List)**

Este é um sistema de gerenciamento de tarefas básico que permite criar, visualizar, atualizar e excluir tarefas. Ele foi desenvolvido utilizando HTML, CSS, JavaScript, jQuery, Ajax, PHP e MySQL.

## **Configuração**

Siga as instruções abaixo para configurar o projeto em seu ambiente local:

### **Pré-requisitos**

- Servidor web (por exemplo, Apache)
- PHP (versão 7 ou superior) instalado
- MySQL (versão 5.6 ou superior) instalado

### **Passo a passo**

1. Faça o download do código-fonte do projeto para o seu computador ou clone-o a partir do repositório.
2. Abra o arquivo **`database.sql`** em um editor de texto e execute as instruções SQL no seu servidor MySQL para criar o banco de dados e a tabela necessários para o sistema.
3. Abra o arquivo **`config.php`** e atualize as configurações de conexão com o banco de dados de acordo com as configurações do seu ambiente.

```
// Configurações do banco de dados
define('DB_HOST', 'localhost'); // endereço do servidor MySQL
define('DB_USERNAME', 'seu_usuario'); // nome de usuário do banco de dados
define('DB_PASSWORD', 'sua_senha'); // senha do banco de dados
define('DB_NAME', 'nome_do_banco'); // nome do banco de dados

```

1. Copie todos os arquivos do projeto para o diretório raiz do seu servidor web.
2. Acesse o sistema através do seu navegador, digitando a URL correspondente ao diretório onde os arquivos foram copiados.

## **Funcionalidades**

O sistema de gerenciamento de tarefas possui as seguintes funcionalidades:

### **1. Criar Tarefa**

- Na página inicial, preencha o campo "Nova Tarefa" com o nome da tarefa que deseja criar.
- Clique no botão "Adicionar" para criar a tarefa. A tarefa será exibida na lista de tarefas.

### **2. Visualizar Tarefas**

- Na página inicial, todas as tarefas cadastradas serão exibidas na lista de tarefas.
- Cada tarefa exibida possui um botão "Editar" e um botão "Excluir" associados a ela.

### **3. Atualizar Tarefa**

- Para atualizar uma tarefa, clique no botão "Editar" correspondente à tarefa desejada.
- Na página de edição, você pode modificar o nome da tarefa.
- Clique no botão "Salvar" para atualizar a tarefa.

### **4. Excluir Tarefa**

- Para excluir uma tarefa, clique no botão "Excluir" correspondente à tarefa desejada.
- Uma caixa de diálogo de confirmação será exibida. Clique em "OK" para confirmar a exclusão ou em "Cancelar" para voltar.

## **Considerações Finais**

Este é um sistema básico de gerenciamento de tarefas que demonstra as operações de CRUD utilizando as tecnologias HTML, CSS, JavaScript, jQuery, Ajax, PHP e MySQL. Ele pode ser usado como ponto de partida para desenvolver um sistema mais completo ou ser modificado conforme suas necessidades.