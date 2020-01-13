# Estudos - Aprenda Go com Testes

Fontes do estudos realizados da linguagem Go.

Link do curso: https://github.com/quii/learn-go-with-tests  
Link do curso (pt-BR): https://larien.gitbook.io/aprenda-go-com-testes/

## Instalações
Abaixo seguem informações sobre as instalações necessárias para a configuração do ambiente de desenvolvimento.

**Compilador**  
Acesse o endereço [https://golang.org/dl/](https://golang.org/dl/) e baixe e instale o compilador da linguagem Go (next, next, finish);

**Documentação**  
A versão 1.13 (ou superior) não incui mais a godoc, para instalá-la execute o comando:


    go get golang.org/x/tools/cmd/godoc

Após a instalação, para iniciá-la execute o comando:


    godoc -http :8000

Feito isso, a documentação estará disponível através do endereço [http://localhost:8000](http://localhost:8000).

**Debugger**  
Para instalar o debugger, execute o comando:

    go get -u github.com/go-delve/delve/cmd/dlv

**Linter**  
Para instalar o linter, execute o comando:

    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

**Editor de textos / Extensão**  
Utilizando o VS Code ([https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)), instale a extensão para Go (ms-vscode.go);

13/01/2020 - Estudos na linguagem Go paralizados devido a priorização do estudo de PHP.