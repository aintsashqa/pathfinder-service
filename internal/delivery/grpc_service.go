package delivery

import (
	"context"
	"errors"

	"github.com/aintsashqa/pathfinder-service/api/proto"
	"github.com/aintsashqa/pathfinder-service/internal/handler"
	"github.com/aintsashqa/pathfinder-service/internal/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcService struct {
	proto.UnimplementedPathfinderServiceServer
	handler *handler.Handler
}

func NewGrpcService(handler *handler.Handler) *GrpcService {
	return &GrpcService{
		handler: handler,
	}
}

func (s *GrpcService) GetPath(ctx context.Context, r *proto.GetPathRequest) (*proto.GetPathResponse, error) {
	var (
		entry = s.fromProtoPoint(r.GetEntry())
		final = s.fromProtoPoint(r.GetFinal())
	)

	objs := make([]*types.Object, 0, len(r.GetObjects()))
	for _, pObj := range r.GetObjects() {
		obj := s.fromProtoObject(pObj)
		objs = append(objs, &obj)
	}

	args := handler.Args{
		Entry:          entry,
		Final:          final,
		Objects:        objs,
		Step:           r.GetStep(),
		UseExtraFields: r.GetUseExtraFields(),
	}

	list, err := s.handler.Handle(ctx, &args)
	if err != nil {
		if errors.Is(err, handler.ErrInvalidPoint) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	path := make([]*proto.Point, 0, len(list))
	for i := range list {
		point := s.toProtoPoint(list[i])
		path = append(path, &point)
	}

	return &proto.GetPathResponse{Path: path}, nil
}

func (s *GrpcService) fromProtoPoint(p *proto.Point) types.Point {
	return types.Point{p.GetX(), p.GetY()}
}

func (s *GrpcService) fromProtoObject(obj *proto.Object) types.Object {
	return types.Object{
		Field: types.Field{
			Position: s.fromProtoPoint(obj.GetPosition()),
			Weight:   obj.GetWeight(),
		},
		Unavailable: obj.GetUnavailable(),
	}
}

func (s *GrpcService) toProtoPoint(p types.Point) proto.Point {
	return proto.Point{X: p.X(), Y: p.Y()}
}
