package handler_test

import (
	"context"
	"testing"

	"github.com/aintsashqa/pathfinder-service/internal/handler"
	"github.com/aintsashqa/pathfinder-service/internal/types"
	"github.com/stretchr/testify/require"
)

func TestHandler_Handle(t *testing.T) {
	type want struct {
		path []types.Point
		err  error
	}

	var (
		hdlr  = handler.New()
		tests = []struct {
			name string
			args handler.Args
			want want
		}{
			{
				name: "success_with_result",
				args: handler.Args{
					Entry:          types.Point{0, 0},
					Final:          types.Point{3, 1},
					Objects:        nil,
					Step:           1,
					UseExtraFields: false,
				},
				want: want{
					path: []types.Point{
						{1, 0},
						{2, 0},
						{3, 0},
						{3, 1},
					},
					err: nil,
				},
			},
			{
				name: "success_with_result_using_extra_fields",
				args: handler.Args{
					Entry:          types.Point{0, 0},
					Final:          types.Point{3, 1},
					Objects:        nil,
					Step:           1,
					UseExtraFields: true,
				},
				want: want{
					path: []types.Point{
						{1, 0},
						{2, 1},
						{3, 1},
					},
					err: nil,
				},
			},
			{
				name: "success_path_around_objects",
				args: handler.Args{
					Entry: types.Point{0, 0},
					Final: types.Point{3, 0},
					Objects: []*types.Object{
						{
							Field: types.Field{
								Position: types.Point{1, 0},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{1, 1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{0, -1},
								Weight:   0,
							},
							Unavailable: true,
						},
					},
					Step:           1,
					UseExtraFields: false,
				},
				want: want{
					path: []types.Point{
						{0, 1},
						{0, 2},
						{1, 2},
						{2, 2},
						{2, 1},
						{2, 0},
						{3, 0},
					},
					err: nil,
				},
			},
			{
				name: "success_path_around_objects_with_high_weight",
				args: handler.Args{
					Entry: types.Point{0, 0},
					Final: types.Point{0, 5},
					Objects: []*types.Object{
						{
							Field: types.Field{
								Position: types.Point{0, 1},
								Weight:   10,
							},
							Unavailable: false,
						},
						{
							Field: types.Field{
								Position: types.Point{0, 2},
								Weight:   10,
							},
							Unavailable: false,
						},
						{
							Field: types.Field{
								Position: types.Point{0, 3},
								Weight:   10,
							},
							Unavailable: false,
						},
						{
							Field: types.Field{
								Position: types.Point{0, 4},
								Weight:   10,
							},
							Unavailable: false,
						},
					},
					Step:           1,
					UseExtraFields: false,
				},
				want: want{
					path: []types.Point{
						{-1, 0},
						{-1, 1},
						{-1, 2},
						{-1, 3},
						{-1, 4},
						{-1, 5},
						{0, 5},
					},
					err: nil,
				},
			},
			{
				name: "fail_invalid_entry_point",
				args: handler.Args{
					Entry: types.Point{0, 0},
					Final: types.Point{2, 0},
					Objects: []*types.Object{
						{
							Field: types.Field{
								Position: types.Point{1, 0},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{-1, 0},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{0, 1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{0, -1},
								Weight:   0,
							},
							Unavailable: true,
						},
					},
					Step:           1,
					UseExtraFields: false,
				},
				want: want{
					path: nil,
					err:  handler.ErrInvalidPoint,
				},
			},
			{
				name: "fail_invalid_final_point_using_extra_fields",
				args: handler.Args{
					Entry: types.Point{-2, 0},
					Final: types.Point{0, 0},
					Objects: []*types.Object{
						{
							Field: types.Field{
								Position: types.Point{1, 0},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{-1, 0},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{0, 1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{0, -1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{-1, -1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{1, -1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{-1, 1},
								Weight:   0,
							},
							Unavailable: true,
						},
						{
							Field: types.Field{
								Position: types.Point{1, 1},
								Weight:   0,
							},
							Unavailable: true,
						},
					},
					Step:           1,
					UseExtraFields: true,
				},
				want: want{
					path: nil,
					err:  handler.ErrInvalidPoint,
				},
			},
		}
	)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := hdlr.Handle(context.Background(), &tt.args)

			require.ErrorIs(t, err, tt.want.err)
			require.ElementsMatch(t, got, tt.want.path)
		})
	}
}
