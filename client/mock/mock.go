//go:generate go run -mod=mod github.com/golang/mock/mockgen  -package=mock -source=../services.go -destination=./services.go
//go:generate go run -mod=mod github.com/golang/mock/mockgen  -package=mock -source=../backend.go -destination=./backend.go

package mock
