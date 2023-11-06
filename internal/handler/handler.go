package handler

import (
	"context"
	"slices"

	"github.com/aintsashqa/pathfinder-service/internal/types"
)

const (
	fieldWeight      float64 = 10
	extraFieldWeight float64 = 15
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

type Args struct {
	// TODO: add Size
	Entry          types.Point
	Final          types.Point
	Objects        []*types.Object
	Step           float64
	UseExtraFields bool
}

func (h *Handler) Handle(_ context.Context, args *Args) (out []types.Point, err error) {
	var (
		explored  []*types.Node
		reachable = []*types.Node{
			{
				Position: args.Entry,
				Weight:   0,
				Distance: args.Entry.Distance(args.Final),
				Previous: nil,
			},
		}

		fields = []*types.Field{
			{
				Position: types.Point{-args.Step, 0},
				Weight:   fieldWeight,
			},
			{
				Position: types.Point{0, -args.Step},
				Weight:   fieldWeight,
			},
			{
				Position: types.Point{args.Step, 0},
				Weight:   fieldWeight,
			},
			{
				Position: types.Point{0, args.Step},
				Weight:   fieldWeight,
			},
		}
	)

	// TODO: cancel all calculations if context is done or canceled

	if args.UseExtraFields {
		fields = append(fields,
			&types.Field{
				Position: types.Point{-args.Step, -args.Step},
				Weight:   extraFieldWeight,
			},
			&types.Field{
				Position: types.Point{args.Step, -args.Step},
				Weight:   extraFieldWeight,
			},
			&types.Field{
				Position: types.Point{args.Step, args.Step},
				Weight:   extraFieldWeight,
			},
			&types.Field{
				Position: types.Point{-args.Step, args.Step},
				Weight:   extraFieldWeight,
			},
		)
	}

	for len(reachable) > 0 {
		current := reachable[0]
		for _, node := range reachable {
			if node.Cost() < current.Cost() {
				current = node
			}
		}

		if current.Position.Eq(args.Final) {
			for current != nil && !current.Position.Eq(args.Entry) {
				out = append(out, current.Position)
				current = current.Previous
			}

			slices.Reverse(out)
			break
		}

		reachable = slices.DeleteFunc(reachable, func(n *types.Node) bool {
			return n == current
		})

		for _, field := range fields {
			var (
				searchable  = current.Position.Add(field.Position)
				unavailable = false

				exists = slices.ContainsFunc(explored, func(n *types.Node) bool {
					return n.Position.Eq(searchable)
				})

				next = types.Node{
					Position: searchable,
					Weight:   field.Weight + current.Weight,
					Distance: searchable.Distance(args.Final),
					Previous: current,
				}
			)

			if exists {
				continue
			}

			for _, obj := range args.Objects {
				if obj.Position.Eq(next.Position) {
					if obj.Unavailable {
						unavailable = true
						break
					}

					next.Weight += obj.Weight
					break
				}
			}

			if unavailable {
				continue
			}

			explored = append(explored, current)
			reachable = append(reachable, &next)
		}
	}

	return
}
