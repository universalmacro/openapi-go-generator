# openapi go generator
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
Then will generate code below.
```go
type CreateSessionRequest struct {
	Password        *string `json:"password" xml:"password"`
	ShortMerchantId *string `json:"shortMerchantId" xml:"shortMerchantId"`
	Account         *string `json:"account" xml:"account"`
}
```
About APIs definition.
```yml
paths:
  /sessions:
    post:
      tags:
        - Session
      summary: "Create session"
      operationId: CreateSession
      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/CreateSessionRequest"
      responses:
        "201":
          description: "Success"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Session"
```
Then the api interface and Binding function is generated. The prefix is depending on the first tags.
```go
type SessionApi interface {
	CreateSession(ctx *gin.Context)
}
func SessionApiBinding(router *gin.Engine, api SessionApi) {
	router.POST("/sessions", api.CreateSession)
}
```
