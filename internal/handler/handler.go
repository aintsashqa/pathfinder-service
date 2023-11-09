package handler

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/aintsashqa/pathfinder-service/internal/types"
)

var (
	ErrInvalidPoint = errors.New("invalid point")
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
		explored  = make([]*types.Node, 0, 256)
		reachable = []types.Node{
			{
				Position: args.Entry,
				Weight:   0,
				Distance: args.Entry.Distance(args.Final),
				Previous: nil,
			},
		}

		fields = []types.Field{
			{Position: types.Point{-args.Step, 0}, Weight: fieldWeight},
			{Position: types.Point{0, -args.Step}, Weight: fieldWeight},
			{Position: types.Point{args.Step, 0}, Weight: fieldWeight},
			{Position: types.Point{0, args.Step}, Weight: fieldWeight},
		}
	)

	// TODO: cancel all calculations if context is done or canceled

	if args.UseExtraFields {
		fields = append(fields,
			types.Field{Position: types.Point{-args.Step, -args.Step}, Weight: extraFieldWeight},
			types.Field{Position: types.Point{args.Step, -args.Step}, Weight: extraFieldWeight},
			types.Field{Position: types.Point{args.Step, args.Step}, Weight: extraFieldWeight},
			types.Field{Position: types.Point{-args.Step, args.Step}, Weight: extraFieldWeight},
		)
	}

	if h.validatePoint(args.Entry, fields, args.Objects) != nil {
		return nil, fmt.Errorf("%w: entry %s", ErrInvalidPoint, args.Entry.String())
	}

	if h.validatePoint(args.Final, fields, args.Objects) != nil {
		return nil, fmt.Errorf("%w: final %s", ErrInvalidPoint, args.Final.String())
	}

	for len(reachable) > 0 {
		current := reachable[0]
		for i := range reachable {
			if reachable[i].Cost() < current.Cost() {
				current = reachable[i]
			}
		}

		if current.Position.Eq(args.Final) {
			for !current.Position.Eq(args.Entry) {
				out = append(out, current.Position)
				current = *current.Previous
			}

			slices.Reverse(out)
			break
		}

		reachable = slices.DeleteFunc(reachable, func(n types.Node) bool {
			return n == current
		})

		for i := range fields {
			searchable := current.Position.Add(fields[i].Position)

			if slices.ContainsFunc(explored, func(n *types.Node) bool {
				return n.Position.Eq(searchable)
			}) {
				continue
			}

			next := types.Node{
				Position: searchable,
				Weight:   fields[i].Weight + current.Weight,
				Distance: searchable.Distance(args.Final),
				Previous: &current,
			}

			unavailable := false
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

			explored = append(explored, &current)
			reachable = append(reachable, next)
		}
	}

	return
}

func (h *Handler) validatePoint(p types.Point, fields []types.Field, objs []*types.Object) error {
	counter := 0

	for i := range fields {
		current := p.Add(fields[i].Position)

		if slices.ContainsFunc(objs, func(o *types.Object) bool { return o.Position.Eq(current) && o.Unavailable }) {
			counter++
		}
	}

	if counter == len(fields) {
		return ErrInvalidPoint
	}

	return nil
}
