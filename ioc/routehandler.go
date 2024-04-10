package ioc

import (
	"reflect"
)

type RouteHandlerFunc func() error

func (c *Container) RegisterRouteHandler(
  handler interface{},
  route string,
) {
  c.register(route, handler, Transient)
}

func (scope *ContainerScope) ResolveRouteHandler(route string) (RouteHandlerFunc, bool) {
  constructor, ok := scope.mainContainer.mappings[route]
       
  if !ok { return nil, false }

  result, ok := scope.createRouteHandler(constructor)

  if !ok { return nil, false }

  return result, true
}

func (scope *ContainerScope) createRouteHandler(constructor Constructor) (RouteHandlerFunc, bool) {
  numParams := len(constructor.dependencyTypes);
  dependencies := make([]reflect.Value, numParams);
  
  for i := 0; i < numParams; i++ {
    dependencyType := constructor.dependencyTypes[i];
    dependency, ok := resolveDependecy(scope, dependencyType);
    
    if !ok { return nil, false }

    dependencies[i] = dependency;
  }

  routeHandlerFunc := func() error {
    result := reflect.ValueOf(constructor.constructor).Call(dependencies)

    if result[0].IsNil() {
      return nil
    } else {
      return result[0].Interface().(error)
    }
  }

  return routeHandlerFunc, true
}
