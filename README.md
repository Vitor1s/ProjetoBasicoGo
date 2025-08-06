# Projeto Básico Go - Conversor de Temperatura

## 📋 Sobre o Projeto

Este projeto foi desenvolvido como parte do curso de **Fundamentos de Go** da Digital Innovation One (DIO), ministrado pela instrutora **Tenille Martins**. O objetivo é aplicar os conceitos fundamentais aprendidos sobre **variáveis, valores e tipos em Go** através de um conversor de temperatura prático.

## 🎯 Objetivos de Aprendizado

- Aplicar os fundamentos da linguagem Go
- Trabalhar com variáveis e tipos primitivos (`float64`)
- Implementar funções puras e reutilizáveis
- Seguir as convenções e boas práticas Go
- Estruturar um projeto Go de forma organizada

## 🚀 Funcionalidades

O projeto implementa um conversor que transforma temperatura de **Kelvin para Celsius**, demonstrando:

- Função `kelvinParaCelsius()` para conversão de temperatura
- Exemplo prático com o ponto de ebulição da água (373.15K → 100°C)
- Uso correto de tipos `float64` para precisão científica
- Saída formatada em português

## 📁 Estrutura do Projeto

```
ProjetoBasicoGo/
├── .gitattributes          # Configuração Git
├── cmd/
│   └── main.go            # Ponto de entrada da aplicação
└── README.md              # Documentação do projeto
```

### Convenções Aplicadas:
- **cmd/**: Diretório padrão Go para executáveis
- **Função pura**: Sem efeitos colaterais, facilita testes
- **Separação de responsabilidades**: Lógica de conversão separada da apresentação

## 🔧 Como Executar

```bash

# Executar o programa
go run cmd/main.go
```

### Saída esperada:
```
O ponto e ebulição em Celsius é: 100
```

## 💡 Conceitos Aplicados

### Variáveis e Tipos:
- **`float64`**: Tipo usado para precisão em cálculos científicos
- **Declaração de variáveis**: `K := 373.15`
- **Atribuição**: `celsius := kelvinParaCelsius(K)`

### Funções:
- **Função pura**: `kelvinParaCelsius(k float64) float64`
- **Parâmetros tipados**: Entrada e saída bem definidas
- **Reutilização**: Função pode ser chamada para qualquer valor Kelvin

### Pacotes:
- **`fmt`**: Para formatação e saída de dados
- **`main`**: Pacote principal do executável

## 📚 Fundamentos Demonstrados

1. **Declaração de variáveis** com inferência de tipo
2. **Funções** como blocos reutilizáveis de código
3. **Tipos primitivos** e sua aplicação prática
4. **Importação de pacotes** da biblioteca padrão
5. **Estrutura básica** de um programa Go

## 👩‍🏫 Créditos

Projeto desenvolvido seguindo as orientações da instrutora **Tenille Martins** no curso de **Fundamentos de Go** da [Digital Innovation One (DIO)](https://dio.me).

## 🛠️ Tecnologias

- **Go 1.x**
- **Package fmt** (biblioteca padrão)

---

> Este projeto representa a aplicação prática dos conceitos fundamentais de Go, servindo como base para desenvolvimento de aplicações mais complexas.