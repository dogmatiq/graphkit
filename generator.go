package graphkit

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/enginekit/config"
	"github.com/dogmatiq/enginekit/message"
	"github.com/emicklei/dot"
)

// generator generates message-flow graphs for one or more Dogma applications.
type generator struct {
	id        int
	root      *dot.Graph
	app       *dot.Graph
	roles     map[message.Type]message.Role
	producers map[message.Type][]dot.Node
	consumers map[message.Type][]dot.Node
}

// nextID returns a unique ID to use for a sub-graph or node.
func (g *generator) nextID() string {
	g.id++
	return strconv.Itoa(g.id)
}

// generate builds and returns a graph of the given applications.
func (g *generator) generate(apps []dogma.Application) (*dot.Graph, error) {
	g.root = dot.NewGraph(dot.Directed)
	g.root.Attr("rankdir", "LR")
	g.root.Attr("splines", "spline")
	g.root.Attr("overlap", "false")
	g.root.Attr("outputMode", "nodesfirst")

	g.roles = map[message.Type]message.Role{}
	g.producers = map[message.Type][]dot.Node{}
	g.consumers = map[message.Type][]dot.Node{}

	for _, app := range apps {
		cfg, err := config.NewApplicationConfig(app)
		if err != nil {
			return nil, err
		}

		g.addApp(cfg)
	}

	g.addExternal()

	return g.root, nil
}

// addApp adds an application to the graph as a sub-graph.
func (g *generator) addApp(cfg *config.ApplicationConfig) {
	g.app = g.root.Subgraph(
		g.nextID(),
		dot.ClusterOption{},
	)
	g.app.Attr("label", makeLabel(cfg.Name(), "application"))

	for _, h := range sortHandlers(cfg.Handlers) {
		g.addHandler(h)
	}
}

// addHandler adds a handler to the graph as a node.
func (g *generator) addHandler(cfg config.HandlerConfig) {
	n := g.app.Node(g.nextID())
	n.Attr("label", makeLabel(cfg.Name(), cfg.HandlerType().String()))
	n.Attr("style", "filled")
	n.Attr("fillcolor", handlerColors[cfg.HandlerType()])
	n.Attr("shape", handlerShapes[cfg.HandlerType()])
	n.Attr("fontsize", nodeLabelFontSize)

	g.addEdges(cfg, n)
}

// addEdges adds edges describing the messages that are produced and consumed by
// a specific handler.
func (g *generator) addEdges(cfg config.HandlerConfig, n dot.Node) {
	for t, r := range cfg.ConsumedMessageTypes() {
		g.roles[t] = r
		g.consumers[t] = append(g.consumers[t], n)

		for _, p := range g.producers[t] {
			g.addEdge(p, n, t, r)
		}
	}

	for t, r := range cfg.ProducedMessageTypes() {
		g.roles[t] = r
		g.producers[t] = append(g.producers[t], n)

		for _, c := range g.consumers[t] {
			g.addEdge(n, c, t, r)
		}
	}
}

// addEdge adds an edge between two handler nodes.
// If the edge already exists, its label is expanded to include this message type.
func (g *generator) addEdge(src, dst dot.Node, t message.Type, r message.Role) {
	label := t.String() + r.Marker()

	// find an existing edge and add to its label.
	for _, e := range g.root.FindEdges(src, dst) {
		labels := strings.Split(
			e.Value("label").(string),
			"\n",
		)

		labels = append(labels, label)
		sort.Strings(labels)

		e.Label(strings.Join(labels, "\n"))
		return
	}

	e := g.root.Edge(
		src,
		dst,
		label,
	)

	e.Attr("fontsize", edgeLabelFontSize)
	e.Attr("color", roleColors[r])
	e.Attr("fontcolor", roleColors[r])
}

// addExternal adds nodes that represent producers and consumers that are
// external to any of the applications in the graph.
func (g *generator) addExternal() {
	for t, nodes := range g.producers {
		if _, ok := g.consumers[t]; !ok {
			for _, n := range nodes {
				g.addEdge(
					n,
					g.externalConsumer(),
					t,
					g.roles[t],
				)
			}
		}
	}

	for t, nodes := range g.consumers {
		if _, ok := g.producers[t]; !ok {
			for _, n := range nodes {
				g.addEdge(
					g.externalProducer(),
					n,
					t,
					g.roles[t],
				)
			}
		}
	}
}

// externalConsumer adds and returns a node representing an external consumer.
func (g *generator) externalConsumer() dot.Node {
	sg := g.root.Subgraph("external", dot.ClusterOption{})
	n := sg.Node("consumer")
	n.Attr("style", "filled")
	n.Attr("fillcolor", externalNodeColor)
	n.Attr("shape", externalNodeShape)
	n.Attr("fontsize", nodeLabelFontSize)
	return n
}

// externalConsumer adds and returns a node representing an external producer.
func (g *generator) externalProducer() dot.Node {
	sg := g.root.Subgraph("external", dot.ClusterOption{})
	n := sg.Node("producer")
	n.Attr("style", "filled")
	n.Attr("fillcolor", externalNodeColor)
	n.Attr("shape", externalNodeShape)
	n.Attr("fontsize", nodeLabelFontSize)
	return n
}

// makeLabel makes a node label from a name and type.
func makeLabel(n, t string) string {
	return fmt.Sprintf("%s\n(%s)", n, t)
}

// sortHandlers returns a set of handlers in an app, sorted by the number of
// message types they are aware of.
func sortHandlers(handlers map[string]config.HandlerConfig) []config.HandlerConfig {
	var sorted []config.HandlerConfig

	for _, h := range handlers {
		sorted = append(sorted, h)
	}

	sort.Slice(
		sorted,
		func(i, j int) bool {
			li := len(sorted[i].ConsumedMessageTypes()) + len(sorted[i].ProducedMessageTypes())
			lj := len(sorted[j].ConsumedMessageTypes()) + len(sorted[j].ProducedMessageTypes())

			return li < lj
		},
	)

	return sorted
}
