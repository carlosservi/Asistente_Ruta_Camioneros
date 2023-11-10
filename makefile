# Configuración del proyecto
APP_NAME = Asistente_Ruta_Camioneros
MAIN_FILE = route.go
TESTS_DIR = ./tests

# Comandos
GO = go
BUILD_FLAGS = -o $(BIN_DIR)/$(APP_NAME)

# Directorios
SRC_DIR = ./src
BIN_DIR = ./bin

# Objetivo predeterminado
.PHONY: all
all: clean build test

.PHONY: build
build:
	@echo "Compilando el proyecto..."
	$(GO) build $(BUILD_FLAGS) $(SRC_DIR)/$(MAIN_FILE)
	@echo "Listo."

.PHONY: clean
clean:
	@echo "Limpiando archivos binarios..."
	@rm -rf $(BIN_DIR)/*
	@echo "Listo."

.PHONY: run
run: build
	@echo "Ejecutando el proyecto..."
	$(BIN_DIR)/$(APP_NAME)

.PHONY: test
test:
	@echo "Ejecutando tests..."
	$(GO) test $(TESTS_DIR)/...
	@echo "Listo."

.PHONY: help
help:
	@echo "Uso del Makefile:"
	@echo "  make          - Compila, construye el proyecto y ejecuta tests"
	@echo "  make build    - Compila el proyecto"
	@echo "  make test     - Ejecuta los tests"
	@echo "  make clean    - Limpia los archivos binarios"
	@echo "  make run      - Compila y ejecuta el proyecto"
