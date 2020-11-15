package dependencyinjection

const (
	TransientLifestyle = iota
	ScopedLifestyle
	SingletonLifestyle
)

var ServiceLifestyle int

type ServiceBuilderInterface interface {
	AddTransient(serviceKey string, newFunction func() *interface{}, deferFunction func())
	AddScoped(serviceKey string, newFunction func() *interface{}, deferFunction func())
	AddSingleton(serviceKey string, newFunction func() *interface{}, deferFunction func())
	BuildServiceProvider() ServiceProviderInterface
}

type ServiceResolverInterface interface {
	GetService(serviceKey string) (*interface{}, error)
	GetServices(serviceKey string) ([]*interface{}, error)
}

// Resolve singleton, transiant and scope services
type ServiceScopeInterface interface {
	ServiceResolverInterface
}

// Resolve singleton and transiant services, to resolve scopes use createScope
type ServiceProviderInterface interface {
	ServiceResolverInterface
	CreateScope() (*ServiceScopeInterface, error)
}
