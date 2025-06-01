# API Gateway com Microserviços: Autenticação e Produtos

Este projeto implementa uma arquitetura de **microserviços** em Go utilizando o framework **Fiber** e autenticação baseada em **JWT**.

- **Auth Service**: Serviço de autenticação (registro e login), gera tokens JWT.
- **Product Service**: Serviço de catálogo de produtos, protegido via autenticação.
- **API Gateway**: Intermedia as requisições, valida JWT e roteia para o serviço correto.

---

## ✅ Arquitetura

**Cliente**  
&nbsp;&nbsp;↓  
**API Gateway** (:3000)  
&nbsp;&nbsp;&nbsp;&nbsp;├─ `/auth/*` → **Auth Service ** (:3001)  
&nbsp;&nbsp;&nbsp;&nbsp;└─ `/api/*`  → **Product Service ** (:3002)

## ✅ Serviços e Rotas

### 📌 1. Auth Service — Porta `3001`

| Método  | Rota             | Descrição                        |
|----------|----------------|--------------------------------|
| GET      | /swagger      | Interface Swagger.      |
| POST      | /login |Realiza o login e retorna um **JWT** para utilizar na **API** .       |
| POST     | /register      | Registra um novo usuário.         |

---

### 📌 2. Product Service — Porta `3002`
| Método  | Rota             | Descrição                        |
|----------|----------------|--------------------------------|
| GET      | /products      | Lista todos os produtos       |
| GET      | /products/{id} | Obtém um produto por ID       |
| POST     | /products      | Cria um novo produto          |
| PUT      | /products/{id} | Atualiza um produto existente |
| DELETE   | /products/{id} | Remove um produto            |
| GET      | /swagger/*     | Acessa a documentação Swagger |
---

### 📌 3. API Gateway — Porta `3000`

- `/auth/*` → Proxy para Auth Service.
- `/api/*` → Proxy para Product Service, com validação de **JWT**.

---

## ✅ Como rodar o projeto
1. No gateway:
```bash
go mod tidy 
go run main.go
```
3. Em cada serviço (**auth** e **product**):

```bash
go mod tidy 
go run .
```
