# ===============================
# Projeto: site-produto
# Go idiomático + Fedora Atomic
# ===============================

APP_NAME := site-produto
CMD_DIR  := ./cmd/web

BIN_DIR  := bin
BIN_PATH := $(BIN_DIR)/$(APP_NAME)

# Diretório real no Fedora Atomic (/opt -> /var/opt)
INSTALL_DIR := /var/opt/site-produto
INSTALL_BIN := $(INSTALL_DIR)/bin
INSTALL_TPL := $(INSTALL_DIR)/templates
INSTALL_STATIC := $(INSTALL_DIR)/static

.PHONY: help run build clean install restart status logs deploy

## help: lista os comandos disponíveis
help:
	@echo "Comandos:"
	@echo "  make run       - roda em modo desenvolvimento"
	@echo "  make build     - gera o binário em ./bin"
	@echo "  make clean     - remove artefatos de build"
	@echo "  make install   - copia binário, templates e static para /var/opt"
	@echo "  make restart   - reinicia o serviço via systemd (HOST)"
	@echo "  make status    - mostra status do serviço (HOST)"
	@echo "  make logs      - acompanha logs do serviço (HOST)"
	@echo "  make deploy    - build + install + restart"

## run: desenvolvimento local
run:
	go run $(CMD_DIR)

## build: gera binário
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH) $(CMD_DIR)

## clean: limpa build
clean:
	rm -rf $(BIN_DIR)

## install: instala no host (requer sudo)
install:
	sudo mkdir -p $(INSTALL_BIN)
	sudo cp $(BIN_PATH) $(INSTALL_BIN)/
	sudo cp -r templates $(INSTALL_DIR)/
	sudo cp -r static $(INSTALL_DIR)/
	sudo chown -R site-produto:site-produto $(INSTALL_DIR)
	sudo restorecon -Rv $(INSTALL_DIR)

## restart: reinicia serviço (HOST)
restart:
	sudo systemctl restart $(APP_NAME)

## status: status do serviço (HOST)
status:
	sudo systemctl status $(APP_NAME)

## logs: logs do serviço (HOST)
logs:
	sudo journalctl -u $(APP_NAME) -f

## deploy: build + install + restart
deploy: build install restart
	@echo "Deploy concluído com sucesso."
