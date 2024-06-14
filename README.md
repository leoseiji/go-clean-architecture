# go-clean-architecture

Comando docker-compose: docker-compose up -d

Comando para rodar a aplicação: go run main.go wire_gen.go

Portas:
Web Server: 8000
gRPC Server: 50051
GraphQL: 8080

Arquivos requests API:
create_order.http
list_orders.http

GraphQL (localhost:8080):
mutation:mutation createOrder {
  createOrder(input: {id: "7643", Price: 12, Tax: 2}) {
    id
    Price
    Tax
  }
}

Query:
{
  orders {
    id,
    Price,
    Tax,
  }
}

gRPC:
evans -r repl
package pb
service OrderService
call CreateOrder
call ListOrder