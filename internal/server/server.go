package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/fajrirahmat/interview-aji/model"
	repository "github.com/fajrirahmat/interview-aji/repository"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Runner interface {
	io.Closer
	Run() error
}

type server struct {
	repo repository.DB
	grpc *grpc.Server
	UnimplementedCheckInOutServiceServer
}

func New() (Runner, error) {
	s := &server{}

	repo, err := repository.NewSQLLiteConnection("db/test.db")
	if err != nil {
		log.Fatalf(err.Error())
	}
	s.repo = repo
	grpcServer := grpc.NewServer()
	RegisterCheckInOutServiceServer(grpcServer, s)

	s.grpc = grpcServer
	return s, nil
}

func (s *server) Run() error {
	lis, _ := net.Listen("tcp", ":8080")
	mux := runtime.NewServeMux()
	dialopt := []grpc.DialOption{grpc.WithInsecure()}

	RegisterCheckInOutServiceHandlerFromEndpoint(context.Background(), mux, "127.0.0.1:8080", dialopt)

	m := cmux.New(lis)
	httpl := m.Match(cmux.HTTP1Fast())
	grpcl := m.Match(cmux.HTTP2(), cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
	https := &http.Server{
		Handler: mux,
	}
	go s.grpc.Serve(grpcl)
	go https.Serve(httpl)
	log.Println("Server running on :8080")
	return m.Serve()
}

//API get location
func (s *server) GetLocation(ctx context.Context, _ *emptypb.Empty) (*model.ListLocation, error) {
	locs, _ := s.repo.ListLocation(ctx)
	return &model.ListLocation{
		Locations: locs,
	}, nil
}

//API check IN/OUT
func (s *server) CheckInOut(ctx context.Context, request *model.CheckInOutRequest) (*model.CheckInOutResponse, error) {
	var data *model.CheckInOutResponse
	_, err := s.repo.AddCheckin(ctx, request)
	if err != nil {
		return data, err
	} else {
		return data, errors.New("success input data !")
	}
}

func (s *server) Close() error {
	s.grpc.GracefulStop()
	return nil
}
