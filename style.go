package graphkit

import (
	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/configkit/message"
	"github.com/emicklei/dot"
)

const fontName = "Helvetica"

var roleColors = map[message.Role]string{
	message.CommandRole: "#0066ff",
	message.EventRole:   "#ff6600",
	message.TimeoutRole: "#444444",
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
func styleApp(g *dot.Graph, cfg configkit.Application) {
	g.Attr("fontname", fontName)
	g.Attr("penwidth", "0")
	g.Attr("style", "filled")
	g.Attr("fillcolor", "#eeeeee")
}

// styleHandler applies style attributes to a graph node representing a Dogma
// message handler.
func styleHandler(n dot.Node, cfg configkit.Handler) {
	n.Attr("fontname", fontName)
	n.Attr("style", "filled")
	n.Attr("margin", "0.15")
	n.Attr("penwidth", "2")
	n.Attr("color", "#eeeeee")
	n.Attr("fillcolor", handlerColors[cfg.HandlerType()])
	n.Attr("shape", handlerShapes[cfg.HandlerType()])
}

// styleMessageEdge applies style attributes to an edge representing message
// flow between two handlers.
func styleMessageEdge(e dot.Edge, r message.Role) {
	e.Attr("fontname", fontName)
	e.Attr("penwidth", "2")
	e.Attr("arrowsize", "0.75")
	e.Attr("color", roleColors[r])
	e.Attr("fontcolor", roleColors[r])
}

// styleForeignNode applies style attributes to a graph node representing a
// foreign message producer or consumer.
func styleForeignNode(n dot.Node, r message.Role) {
	n.Attr("fontname", fontName)
	n.Attr("style", "filled")
	n.Attr("margin", "0.15")
	n.Attr("penwidth", "2")
	n.Attr("color", "#ffffff")
	n.Attr("fillcolor", roleColors[r])
	n.Attr("shape", "box")
}
