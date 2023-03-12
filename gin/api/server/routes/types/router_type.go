package routetypes
type RouterType struct {
    prefix string
}
type RouterInterface interface {
    New() RouterInterface
    InitRoutes()
}
