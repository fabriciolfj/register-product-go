# register-product-go
- instalando o wire
````
go install github.com/google/wire/cmd/wire@latest
go get github.com/google/wire
````
- crie um arquivo wire.go indicando as dependencias. Não esqueça das diretivas abaixo:
````
//go:build wireinject
// +build wireinject
````