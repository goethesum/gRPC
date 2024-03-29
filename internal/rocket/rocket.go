//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/goethesum/gRPC/internal/rocket Store

package rocket

import "context"

// Rocket - should contain the definition of our
// rocker
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - definces the interface we expect
// updating the rocker inventory
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service, responsible for
// updating the rocket inventory
type Service struct {
	Store Store
}

// New - returns a new instance of our rocker service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - retrieves a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket from our inventory
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
