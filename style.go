package graphkit

import (
	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/configkit/message"
	"github.com/emicklei/dot"
)

var roleColors = map[message.Role]string{
	message.CommandRole: "#0066ff",
	message.EventRole:   "#ff6600",
}

var handlerShapes = map[configkit.HandlerType]string{
	configkit.AggregateHandlerType:   "box",
	configkit.ProcessHandlerType:     "box",
	configkit.IntegrationHandlerType: "box",
	configkit.ProjectionHandlerType:  "cylinder",
}

var handlerColors = map[configkit.HandlerType]string{
	configkit.AggregateHandlerType:   "#fff89d",
	configkit.ProcessHandlerType:     "#c086c5",
	configkit.IntegrationHandlerType: "#ffde00",
	configkit.ProjectionHandlerType:  "#a3d350",
}

// styleApp applies style attributes to a sub-graph representing an application.
//
// If cfg is nil, the graph represents a foreign application.
func styleApp(g *dot.Graph, cfg configkit.Application) {
	g.Attr("penwidth", "0")
	g.Attr("style", "filled")
	g.Attr("fillcolor", "#eeeeee")
	g.Attr("fontname", "Helvetica")
}

// styleHandler applies style attributes to a graph node representing a Dogma
// message handler.
//
// If cfg is nil, the node represents a handler within a foreign application.
func styleHandler(n dot.Node, cfg configkit.Handler) {
	if cfg == nil {
		n.Attr("style", "filled")
		n.Attr("fillcolor", "#aaaaaa")
		n.Attr("shape", "box")
		n.Attr("color", "#eeeeee")
		return
	}

	n.Attr("fillcolor", handlerColors[cfg.HandlerType()])
	n.Attr("shape", handlerShapes[cfg.HandlerType()])
	n.Attr("style", "filled")
	n.Attr("fontname", "Helvetica")
	n.Attr("margin", "0.15")
	n.Attr("penwidth", "2")
	n.Attr("color", "#eeeeee")
}

// styleMessageEdge applies style attributes to an edge representing message
// flow between two handlers.
func styleMessageEdge(e dot.Edge, r message.Role) {
	e.Attr("penwidth", "2")
	e.Attr("arrowsize", "0.75")
	e.Attr("fontname", "Helvetica")
	e.Attr("color", roleColors[r])
	e.Attr("fontcolor", roleColors[r])
}
