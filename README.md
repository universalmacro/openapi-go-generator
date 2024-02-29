# openapi go gin server generator
`openapi-go-generator` is a code generator. It can generate go struct model. And generator gin router binder. The binder automatically bind the router into gin.Router.
It also generate interface of APIs.
## Feature

## Sample use
```bash
go run main.go ./openapi.yml main ./api.go
```

Let say there are swagger.yml.
```yml
components:
  schemas:
    CreateSessionRequest:
      description: "Login request"
      type: object
      properties:
        shortMerchantId:
          type: string
        account:
          type: string
        password:
          type: string
```
Then will generate below code.
```go
type CreateSessionRequest struct {
	Password        *string `json:"password" xml:"password"`
	ShortMerchantId *string `json:"shortMerchantId" xml:"shortMerchantId"`
	Account         *string `json:"account" xml:"account"`
}
```
