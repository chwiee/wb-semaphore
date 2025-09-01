⚡ Semaphore API Wrapper (Go)

Uma API em Go para interagir com o Semaphore UI
, permitindo listar projetos, inventories, templates e executar tasks de forma simples e prática.

✨ Features

📂 Listagem de Projetos: obtenha todos os projetos disponíveis no Semaphore

🔑 Gerenciamento de Keys: visualize as chaves associadas a um projeto

🗄️ Inventários: filtre inventários por nome e ID

📦 Repositórios: acesse repositórios configurados no Semaphore

📝 Templates: encontre templates por nome

⚙️ Execução de Tasks: rode tasks diretamente pelo nome do projeto, inventory e template, sem precisar decorar IDs

🛠️ Tecnologias

🐹 Go 1.22+

🌐 Chi Router

🔄 errgroup

🗃️ JSON + HTTP Client nativo do Go

🚀 Como rodar
# Clone o repositório
git clone https://github.com/seu-user/seu-repo.git
cd seu-repo

# Configure variáveis de ambiente
export SEMAPHORE_URL=http://localhost:3000/api
export SEMAPHORE_TOKEN=seu_token_aqui

# Rodar a API
go run cmd/api/main.go


A API estará disponível em:

http://localhost:8000

🔍 Endpoints principais
📂 Listar projetos
GET /projects

📑 Detalhes de um projeto por ID
GET /project/{id}

⚡ Executar task por nomes
GET /task/{project}/{inventory}/{task}


Esse endpoint resolve os IDs automaticamente e dispara a execução no Semaphore.

🧪 Testes

Rodar testes unitários:

go test ./... -v

📌 Exemplo de execução de task via cURL
curl -X POST "http://localhost:8000/task/meu-projeto/meu-inventory/minha-task"

🤝 Contribuição

Faça um fork do projeto

Crie sua branch de feature (git checkout -b minha-feature)

Commit suas alterações (git commit -m 'Adiciona nova feature')

Push para a branch (git push origin minha-feature)

Abra um Pull Request 🚀