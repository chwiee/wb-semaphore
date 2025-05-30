# Semaphore UI - Ansible and Terraform

## Struct

A estrutura do repositório se consiste em criar um fluxo de pastas separados por projetos no Semaphore-ui, a idéia é termos um centralizador segregado por projeto e assets para que quanto for feito um merge na main, seja criado o projeto caso não exista, crie as chaves necessárias (pre definidas), iventários e tasks conforme valores nas pastas


### Exemplo

xpto/                       -> nome do projeto 
/keys/                      -> pasta contendo as chaves 
/keys/xpto.pem              -> nome da chave 
/ansible/                   -> pasta com as automações para o ansible
/ansible/{x}.yml            -> onde {x} é o nome da Task que será criada
/terraform/                 -> pasta com as automações de terraform
/terraform/{x}/main.tf      -> onde {x} é o nome da Task a ser gerada

Algumas validações precisam ser feitas para que tenhamos um ótimo fluxo:

1. Pegar as mundanças da branch e considerar APENAS uma mudança (ou seja cada branch deverá ser padrão ajustar apenas um projeto)
2. Caso o projeto já existe 
3. Validar se as chaves já existe
4. Validar se o repositório já existe
5. Validar se as Tasks já existem 
6. Enviar status para o PR ** de preferencia triggar do github para o azure devops e também via actions

### Estrutura da Infra

-> Semaphore-UI           : Executando em um Cluster Kubernetes
-> Pipeline Agent         : Criar a imagem container com os pre-reqs se subir a imagem com todos os recursos necessários