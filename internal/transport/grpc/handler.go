package grpc

import (
	"context"
	"log"
	"net"

	"github.com/goethesum/gRPC/internal/rocket"
	rkt "github.com/goethesum/gRPC/internal/rocket"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RocketService - define the interface that the concrete implementation
// has to adhere to
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
}

// New - returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("could not listen on port 50051")
		return err
	}
	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("failed to serve: %s/n", err)
		return err
	}
	return nil

}

// GetRocket - retrieves a rocket by id and returns the response.
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("Get Rocket gRPC Endpoint Hit")

	rocket, err := h.RocketService.GetRocketByID(ctx, req.Id)
	if err != nil {
		log.Println("Failed to retrieve rocket by ID")
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

// AddRocket - adds a rocket to the database
func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC endpoint hit")
	if _, err := uuid.Parse(req.Rocket.ID); err != nil {
		errorStatus := status.Error(codes.InvalidArgument, "uuid is not valid")
		log.Print("given uuid is not valid")
		return &rkt.AddRocketResponse{}, errorStatus
	}
	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})
	if err != nil {
		log.Println("failed to insert rocket into database")
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

// DeleteRocket - handlerf for deleting a rocket
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("Delete rocket gPRC endpoint hit")
	err := h.RocketService.DeleteRocket(ctx, req.Rocket.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.Del1eteRocketResponse{
		Status: "successfully delete rocket",
	}, nil
}
