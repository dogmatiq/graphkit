package graphkit

import (
	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/configkit/message"
)

const (
	foreignNodeShape = "box"
	foreignNodeColor = "#aaaaaa"
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
