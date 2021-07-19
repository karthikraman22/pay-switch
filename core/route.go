package core

import "fmt"

type Route struct {
	Id              string
	ConnectedClient *Client
	Active          bool
}

type RouteManager struct {
	routes map[string]*Route
}

func NewRouteManager() *RouteManager {
	return &RouteManager{routes: make(map[string]*Route)}
}

func (rm *RouteManager) AddRoute(routeId string, client *Client) {
	route := &Route{Id: routeId, ConnectedClient: client, Active: false}
	rm.routes[routeId] = route
}

func (rm *RouteManager) ActiveRoute(routeId string) {
	route := rm.routes[routeId]
	if route != nil {
		route.Active = true
	}
}

func (rm *RouteManager) DeactiveRoute(routeId string) {
	route := rm.routes[routeId]
	if route != nil {
		route.Active = false
	}
}

func (rm *RouteManager) PrintRoutes() {
	fmt.Printf("rm.routes: %v\n", rm.routes)
}
