[![CI](https://github.com/mjucimara/site-produto-go/actions/workflows/ci.yml/badge.svg)](https://github.com/mjucimara/site-produto-go/actions/workflows/ci.yml)

# Site Produto

Aplicação web em Go, minimalista e idiomática, voltada para divulgação de serviços, com templates HTML, assets estáticos e deploy via systemd em Fedora Atomic.

O foco do projeto é **simplicidade operacional**, **previsibilidade** e **boas práticas**, sem frameworks desnecessários.

---

## ✨ Características

* Go puro (`net/http`, `html/template`)
* Estrutura idiomática (`cmd/`, `internal/`, `pkg/`)
* Templates HTML com layout compartilhado
* Assets estáticos (CSS)
* Endpoint de healthcheck (`/health`)
* Makefile para padronizar build e deploy
* CI automática com GitHub Actions
* Deploy real via systemd (Fedora Atomic)

---

## 📁 Estrutura do Projeto

```
cmd/web            # entrypoint da aplicação
internal/          # domínio e infraestrutura
pkg/               # código reutilizável
templates/         # templates HTML
static/            # CSS e assets
bin/               # binário local (ignorado no Git)
```

---

## 🚀 Desenvolvimento

Rodar localmente:

```bash
make run
```

Acesse:

* [http://localhost:8080](http://localhost:8080)
* [http://localhost:8080/health](http://localhost:8080/health)

---

## 🏗️ Build

Gerar binário local:

```bash
make build
```

---

## 📦 Deploy (Fedora Atomic)

O projeto foi pensado para **build em toolbox** e **execução no host**.

Deploy completo:

```bash
make deploy
```

Verificar status:

```bash
make status
make logs
```

---

## ⚙️ Configuração

A aplicação lê configuração via variáveis de ambiente.

Exemplo (systemd):

```ini
Environment=ADDR=:8080
```

---

## 🤖 CI (GitHub Actions)

A pipeline executa automaticamente em cada push ou pull request:

* `go build`
* `go test`
* `go vet`

A versão do Go é explicitamente travada para garantir previsibilidade.

---

## 🎯 Objetivo do Projeto

Este repositório serve como:

* base para novos projetos em Go
* exemplo de aplicação web simples e bem estruturada
* referência de deploy em ambiente Linux imutável
* material de portfólio técnico

---

## 📄 Licença

MIT

---
