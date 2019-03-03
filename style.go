package graphkit

import (
	"github.com/dogmatiq/enginekit/handler"
	"github.com/dogmatiq/enginekit/message"
)

const (
	nodeLabelFontSize = "14"
	edgeLabelFontSize = "12"

	externalNodeShape = "oval"
	externalNodeColor = "#aaffaa"
)

var roleColors = map[message.Role]string{
	message.CommandRole: "#0066ff",
	message.EventRole:   "#ff6600",
}

var handlerShapes = map[handler.Type]string{
	handler.AggregateType:   "oval",
	handler.ProcessType:     "diamond",
	handler.IntegrationType: "box",
	handler.ProjectionType:  "cylinder",
}

var handlerColors = map[handler.Type]string{
	handler.AggregateType:   "#fff89d",
	handler.ProcessType:     "#c086c5",
	handler.IntegrationType: "#ffde00",
	handler.ProjectionType:  "#a3d350",
}
