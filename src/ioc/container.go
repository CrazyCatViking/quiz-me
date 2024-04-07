package ioc

import (
	"fmt"
	"reflect"
	"strings"
)

type RegistrationScope int64

const (
  Singleton RegistrationScope = iota
  Scoped RegistrationScope = iota
  Transient RegistrationScope = iota
)

type Constructor struct {
  constructor interface{}
  dependencyTypes []reflect.Type
  scope RegistrationScope
}

type Container struct {
  mappings map[string]Constructor
  singletons map[string]reflect.Value
}

type ContainerScope struct {
  mainContainer *Container
  instances map[string]reflect.Value
}

func (c *Container) register(key string, constructor interface{}, scope RegistrationScope) {
  dependencies := getDependencies(constructor)

  c.mappings[key] = Constructor {
    constructor: constructor,
    dependencyTypes: dependencies,
    scope: scope,
  }
}

func (c *Container) OpenScope() *ContainerScope {
  return &ContainerScope {
    mainContainer: c,
    instances: make(map[string]reflect.Value),
  }  
}

func getDependencies(constructor interface{}) []reflect.Type {
  constructorType := reflect.TypeOf(constructor);
  numParams := constructorType.NumIn();
  
  dependencies := make([]reflect.Type, numParams);
  
  for i := 0; i < numParams; i++ {
    dependencyType := constructorType.In(i);
    dependencies[i] = dependencyType;
  }
  
  return dependencies;
}

func NewContainer() *Container {
  return &Container {
    mappings: make(map[string]Constructor),
    singletons: make(map[string]reflect.Value),
  }
}

func RegisterScoped[t interface{}](c *Container, constructor interface{}) {
  typeName := getTypeName[t]()
  c.register(typeName, constructor, Scoped)
}

func RegisterTransient[t interface{}](c *Container, constructor interface{}) {
  typeName := getTypeName[t]()
  c.register(typeName, constructor, Transient)
}

func RegisterSingleton[t interface{}](c *Container, constructor interface{}) {
  typeName := getTypeName[t]()
  c.register(typeName, constructor, Singleton)
}

func RegisterRouteHandler[t interface{}](
  c *Container,
  constructor interface{},
  route string,
) {
  typeName := getTypeName[t]()
  key := strings.Join([]string{typeName, route}, "-")

  c.register(key, constructor, Transient)
}

func ResolveRequestHandler[T interface{}](scope *ContainerScope, route string) (*T, bool) {
  typeName := getTypeName[T]()
  key := strings.Join([]string{typeName, route}, "-")

  constructor, ok := scope.mainContainer.mappings[key]
  
  if !ok { return nil, false }

  result, ok := createInstance(scope, constructor)

  if !ok { return nil, false }
  
  instance := result.Interface().(T)

  return &instance, true
}

func Resolve[T interface{}](scope *ContainerScope) *T {
  typeName := getTypeName[T]()

  constructor := scope.mainContainer.mappings[typeName]

  if constructor.scope == Singleton {
    result, ok := scope.mainContainer.singletons[typeName]
    if ok {
      return result.Interface().(*T)
    }
  } else if constructor.scope == Scoped {
    result, ok := scope.instances[typeName]
    if ok {
      return result.Interface().(*T)
    }
  }

  instance, _  := createInstance(scope, constructor)

  if constructor.scope == Singleton {
    scope.mainContainer.singletons[typeName] = instance  
  } else if constructor.scope == Scoped {
    scope.instances[typeName] = instance 
  }
  
  return instance.Interface().(*T)
}

func createInstance(scope *ContainerScope, constructor Constructor) (reflect.Value, bool) {
  numParams := len(constructor.dependencyTypes);
  dependencies := make([]reflect.Value, numParams);
  
  for i := 0; i < numParams; i++ {
    dependencyType := constructor.dependencyTypes[i];
    dependency, ok := resolveDependecy(scope, dependencyType);
    
    if !ok { return reflect.Value{}, false }

    dependencies[i] = dependency;
  }
  
  result := reflect.ValueOf(constructor.constructor).Call(dependencies)
  instance := result[0]

  return instance, true
}

func resolveDependecy(scope *ContainerScope, t reflect.Type) (reflect.Value, bool) {
  var typeName string

  if t.Kind() == reflect.Ptr {
    typeName = t.Elem().String()
  } else {
    typeName = t.String()
  }

  constructor, ok := scope.mainContainer.mappings[typeName]
  
  if !ok { return reflect.Value{}, false }

  if constructor.scope == Singleton {
    result, ok := scope.mainContainer.singletons[typeName]
    if ok {
      return result, true
    }
  } else if constructor.scope == Scoped {
    result, ok := scope.instances[typeName]
    if ok {
      return result, true
    }
  }

  instance, ok := createDependencyInstance(scope, constructor, t)
  
  if constructor.scope == Singleton {
    scope.mainContainer.singletons[typeName] = instance
  } else if constructor.scope == Scoped {
    scope.instances[typeName] = instance
  }

  return instance, true
}

func createDependencyInstance(
  scope *ContainerScope,
  constructor Constructor,
  t reflect.Type,
) (reflect.Value, bool) {
  numParams := len(constructor.dependencyTypes);
  dependencies := make([]reflect.Value, numParams);
  
  for i := 0; i < numParams; i++ {
    dependencyType := constructor.dependencyTypes[i];
    dependency, ok := resolveDependecy(scope, dependencyType);

    if !ok { return reflect.Value{}, false }

    dependencies[i] = dependency;
  }
  
  result := reflect.ValueOf(constructor.constructor).Call(dependencies)
  instance := result[0].Convert(t)

  return instance, true
}

func getTypeName[T any]() string {
  return reflect.TypeOf((*T)(nil)).Elem().String();
}
