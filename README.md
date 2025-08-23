# Go Interpreter Series

This repository is my learning journey of building an interpreter from scratch in GO

The goal of this repo is to document each step as I gradually build a tiny interpreter in Go, starting from a basic arithmetic evaluator to something more advanced.

---

## 📖 About
Interpreters are programs that execute code directly without compiling it into machine code.  
This series walks through the foundations of interpreters:

- Lexical analysis (tokenizing input text)  
- Parsing (understanding expressions)  
- Evaluating (producing results)  

## 📂 Versions
 Each version progressively adds new features to the interpreter:
  
- **Version 1** → Supports addition and subtraction `1 + 1`, `2 - 1`
- **Version 2** → Supports chained addition & substraction  `30+40-20+5`  
- **Version 3** → Supports multiplication and division `2 * 3`, `6 / 2`
- **Version 4** → Supports Chained arithmetic expressions (`2 + 3 - 1 * 4 / 2`)  
- **Version 5** → Supports full arithmetic with nested parentheses  `7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8))`
-  **Version 6** →  Has a basic Parser & AST with all the  previous feature intact

## 🚀 Running the Project

Make sure you have Go installed ([download here](https://go.dev/dl/)).

Clone the repo:
```bash
git clone https://github.com/TanvirTian/go-interpreter-series.git
cd go-interpreter-series
```