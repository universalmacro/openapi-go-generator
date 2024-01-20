package main

import "github.com/gin-gonic/gin"

type SessionApi interface {
	CreateSession(gin.Context)
}
type NodeApi interface {
	GetNodeDatabaseConfig(gin.Context)
}
type MerchantApi interface {
	ListNodeMerchants(gin.Context)
}
type AdminApi interface {
	CreateAdmin(gin.Context)
}
type ServerConfig struct {
	Port      *string `json:"port" xml:"port"`
	JwtSecret *string `json:"jwtSecret" xml:"jwtSecret"`
}
type TencentCloudConfig struct {
	SecretId  *string `json:"secretId" xml:"secretId"`
	SecretKey *string `json:"secretKey" xml:"secretKey"`
}
type PhoneNumber struct {
	CountryCode string `json:"countryCode" xml:"countryCode"`
	Number      string `json:"number" xml:"number"`
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
type CreateMerchantRequest struct {
	Password string `json:"password" xml:"password"`
	Account  string `json:"account" xml:"account"`
}
type NodeConfig struct {
	Api      *string `json:"api" xml:"api"`
	Server   *string `json:"server" xml:"server"`
	Database *string `json:"database" xml:"database"`
	Redis    *string `json:"redis" xml:"redis"`
}
type RedisConfig struct {
	Host     *string `json:"host" xml:"host"`
	Port     *string `json:"port" xml:"port"`
	Password *string `json:"password" xml:"password"`
}
type DBConfig struct {
	Password string  `json:"password" xml:"password"`
	Database string  `json:"database" xml:"database"`
	Type     *string `json:"type" xml:"type"`
	Host     string  `json:"host" xml:"host"`
	Port     string  `json:"port" xml:"port"`
	Username string  `json:"username" xml:"username"`
}
type Admin struct {
	CreatedAt   *string `json:"createdAt" xml:"createdAt"`
	Id          string  `json:"id" xml:"id"`
	Account     string  `json:"account" xml:"account"`
	PhoneNumber *string `json:"phoneNumber" xml:"phoneNumber"`
	Role        *string `json:"role" xml:"role"`
}
type CreateSessionRequest struct {
	Account  *string `json:"account" xml:"account"`
	Password *string `json:"password" xml:"password"`
}
type Session struct {
	Token string `json:"token" xml:"token"`
}
type Pagination struct {
	Index string `json:"index" xml:"index"`
	Limit string `json:"limit" xml:"limit"`
	Total string `json:"total" xml:"total"`
}
type Merchant struct {
	Name        string  `json:"name" xml:"name"`
	Description *string `json:"description" xml:"description"`
	CreatedAt   string  `json:"createdAt" xml:"createdAt"`
	UpdatedAt   string  `json:"updatedAt" xml:"updatedAt"`
	Id          string  `json:"id" xml:"id"`
	NodeId      string  `json:"nodeId" xml:"nodeId"`
}
type NodeList struct {
	Items      string `json:"items" xml:"items"`
	Pagination string `json:"pagination" xml:"pagination"`
}
type ApiConfig struct {
	MerchantUrl *string `json:"merchantUrl" xml:"merchantUrl"`
}
type MerchantList struct {
	Items      string `json:"items" xml:"items"`
	Pagination string `json:"pagination" xml:"pagination"`
}
type Ordering string

const ASCENDING Ordering = "ASCENDING"
const DESCENDING Ordering = "DESCENDING"

type Node struct {
	UpdatedAt   string  `json:"updatedAt" xml:"updatedAt"`
	Id          string  `json:"id" xml:"id"`
	Name        string  `json:"name" xml:"name"`
	SecurityKey *string `json:"securityKey" xml:"securityKey"`
	Description *string `json:"description" xml:"description"`
	CreatedAt   string  `json:"createdAt" xml:"createdAt"`
}
type DatabaseType string

const MYSQL DatabaseType = "MYSQL"
const POSTGRES DatabaseType = "POSTGRES"

type UpdatePasswordRequest struct {
	Password    string  `json:"password" xml:"password"`
	OldPassword *string `json:"oldPassword" xml:"oldPassword"`
}
type TestObject struct {
	Id   string `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}
type CreateNodeRequest struct {
	Name        *string `json:"name" xml:"name"`
	Description *string `json:"description" xml:"description"`
}
type Role string

const ROOT Role = "ROOT"
const ADMIN Role = "ADMIN"
