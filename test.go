package main

type UpdatePasswordRequest struct {
	Password    string  `json:"password" xml:"password"`
	OldPassword *string `json:"oldPassword" xml:"oldPassword"`
}
type Pagination struct {
	Index string `json:"index" xml:"index"`
	Limit string `json:"limit" xml:"limit"`
	Total string `json:"total" xml:"total"`
}
type Merchant struct {
	CreatedAt   string  `json:"createdAt" xml:"createdAt"`
	UpdatedAt   string  `json:"updatedAt" xml:"updatedAt"`
	Id          string  `json:"id" xml:"id"`
	NodeId      string  `json:"nodeId" xml:"nodeId"`
	Name        string  `json:"name" xml:"name"`
	Description *string `json:"description" xml:"description"`
}
type Node struct {
	Id          string  `json:"id" xml:"id"`
	Name        string  `json:"name" xml:"name"`
	SecurityKey *string `json:"securityKey" xml:"securityKey"`
	Description *string `json:"description" xml:"description"`
	CreatedAt   string  `json:"createdAt" xml:"createdAt"`
	UpdatedAt   string  `json:"updatedAt" xml:"updatedAt"`
}
type DBConfig struct {
	Password string  `json:"password" xml:"password"`
	Database string  `json:"database" xml:"database"`
	Type     *string `json:"type" xml:"type"`
	Host     string  `json:"host" xml:"host"`
	Port     string  `json:"port" xml:"port"`
	Username string  `json:"username" xml:"username"`
}
type PhoneNumber struct {
	Number      string `json:"number" xml:"number"`
	CountryCode string `json:"countryCode" xml:"countryCode"`
}
type CreateSessionRequest struct {
	Account  *string `json:"account" xml:"account"`
	Password *string `json:"password" xml:"password"`
}
type Role string

const ROOT Role = "ROOT"
const ADMIN Role = "ADMIN"

type Session struct {
	Token string `json:"token" xml:"token"`
}
type CreateMerchantRequest struct {
	Account  string `json:"account" xml:"account"`
	Password string `json:"password" xml:"password"`
}
type TestObject struct {
	Name string `json:"name" xml:"name"`
	Id   string `json:"id" xml:"id"`
}
type NodeConfig struct {
	Api      *string `json:"api" xml:"api"`
	Server   *string `json:"server" xml:"server"`
	Database *string `json:"database" xml:"database"`
	Redis    *string `json:"redis" xml:"redis"`
}
type ServerConfig struct {
	Port      *string `json:"port" xml:"port"`
	JwtSecret *string `json:"jwtSecret" xml:"jwtSecret"`
}
type AdminList struct {
	Items      string `json:"items" xml:"items"`
	Pagination string `json:"pagination" xml:"pagination"`
}
type CreateAdminRequest struct {
	Account  string  `json:"account" xml:"account"`
	Password string  `json:"password" xml:"password"`
	Role     *string `json:"role" xml:"role"`
}
type CreateNodeRequest struct {
	Name        *string `json:"name" xml:"name"`
	Description *string `json:"description" xml:"description"`
}
type RedisConfig struct {
	Host     *string `json:"host" xml:"host"`
	Port     *string `json:"port" xml:"port"`
	Password *string `json:"password" xml:"password"`
}
type ApiConfig struct {
	MerchantUrl *string `json:"merchantUrl" xml:"merchantUrl"`
}
type Admin struct {
	Account     string  `json:"account" xml:"account"`
	PhoneNumber *string `json:"phoneNumber" xml:"phoneNumber"`
	Role        *string `json:"role" xml:"role"`
	CreatedAt   *string `json:"createdAt" xml:"createdAt"`
	Id          string  `json:"id" xml:"id"`
}
type MerchantList struct {
	Items      string `json:"items" xml:"items"`
	Pagination string `json:"pagination" xml:"pagination"`
}
type NodeList struct {
	Items      string `json:"items" xml:"items"`
	Pagination string `json:"pagination" xml:"pagination"`
}
type TencentCloudConfig struct {
	SecretId  *string `json:"secretId" xml:"secretId"`
	SecretKey *string `json:"secretKey" xml:"secretKey"`
}
type DatabaseType string

const MYSQL DatabaseType = "MYSQL"
const POSTGRES DatabaseType = "POSTGRES"

type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"
