package web

import (
	"context"
	"encoding/json"
	"net"

	"log"
	"net/http"

	authService "github.com/antoha2/auth/service"
	authEndpoints "github.com/antoha2/auth/transport/endpoints"
	pb "github.com/antoha2/auth/transport/protoAPI"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	config "github.com/antoha2/auth/config"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func (wImpl *webImpl) StartHTTP() error {

	AuthOptions := []httptransport.ServerOption{
		//httptransport.ServerBefore(wImpl.UserIdentify),
	}

	signInHandler := httptransport.NewServer(
		authEndpoints.MakeSignInEndpoint(&wImpl.authService),
		decodeMakeSignInRequest,
		encodeResponse,
		AuthOptions...,
	)

	signUpAdminHandler := httptransport.NewServer(
		authEndpoints.MakeSignUpAdminEndpoint(&wImpl.authService),
		decodeMakeSignUpAdminRequest,
		encodeResponse,
		AuthOptions...,
	)

	signUpUserHandler := httptransport.NewServer(
		authEndpoints.MakeSignUpUserEndpoint(&wImpl.authService),
		decodeMakeSignUpUserRequest,
		encodeResponse,
		AuthOptions...,
	)

	deleteUserHandler := httptransport.NewServer(
		authEndpoints.MakeDeleteUserEndpoint(&wImpl.authService),
		decodeMakeDeleteUserRequest,
		encodeResponse,
		httptransport.ServerBefore(wImpl.UserIdentify),
	)

	updateUserHandler := httptransport.NewServer(
		authEndpoints.MakeUpdateUserEndpoint(&wImpl.authService),
		decodeMakeUpdateUserRequest,
		encodeResponse,
		httptransport.ServerBefore(wImpl.UserIdentify),
	)

	/* parseTokenHandler := httptransport.NewServer(
		authEndpoints.MakeParseTokenEndpoint(&wImpl.authService),
		decodeMakeParseTokenRequest,
		encodeResponse,
		AuthOptions..., //httptransport.ServerBefore(wImpl.UserIdentify),
	)

	getRolesHandler := httptransport.NewServer(
		authEndpoints.MakeGetRolesEndpoint(&wImpl.authService),
		decodeMakeGetRolesRequest,
		encodeResponse,
		AuthOptions..., //httptransport.ServerBefore(wImpl.UserIdentify),
	) */

	r := mux.NewRouter() //I'm using Gorilla Mux, but it could be any other library, or even the stdlib

	r.Methods("POST").Path("/auth/sign-in").Handler(signInHandler)
	r.Methods("POST").Path("/auth/sign-up/admin").Handler(signUpAdminHandler)
	r.Methods("POST").Path("/auth/sign-up/user").Handler(signUpUserHandler)
	r.Methods("POST").Path("/auth/deleteUser").Handler(deleteUserHandler)
	r.Methods("POST").Path("/auth/updateUser").Handler(updateUserHandler)

	//r.Methods("POST").Path("/auth/parseToken").Handler(parseTokenHandler)
	//r.Methods("POST").Path("/auth/getRoles").Handler(getRolesHandler)

	wImpl.server = &http.Server{Addr: config.HTTPAddr}
	log.Printf("(auth) Запуск HTTP-сервера на http://127.0.0.1%s\n", wImpl.server.Addr) //:8180

	if err := http.ListenAndServe(wImpl.server.Addr, r); err != nil {
		log.Println(err)
	}

	return nil

}

func (wImpl *webImpl) StartGRPC(a *authService.AuthService) error {

	listener, err := net.Listen("tcp", config.GRPCAddr) //:8183

	if err != nil {
		grpclog.Fatalf("failed to listen v", err)
		return err
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	grpcImpl := NewGRPC(a)

	// wImpl.authService.ParseToken()
	// wImpl.authService.GetRoles()

	pb.RegisterTaskServiceServer(grpcServer, grpcImpl) //&authservice.AuthService{}

	log.Printf("(auth) Запуск GRPC-сервера на http://127.0.0.1%s\n", config.GRPCAddr) //:8183
	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("err 333333333333333333: %v", err)
		return err
	}

	return nil
}

func (s *GRPCImpl) ParseToken(c context.Context, request *pb.ParseTokenRequest) (response *pb.ParseTokenResponse, err error) {

	id, err := s.authService.ParseToken(request.Token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	response = &pb.ParseTokenResponse{
		Id: int32(id),
	}
	return response, nil
}

func (s *GRPCImpl) GetRoles(c context.Context, request *pb.GetRolesRequest) (response *pb.GetRolesResponse, err error) {

	roles := s.authService.GetRoles(int(request.Id))
	response = &pb.GetRolesResponse{
		Roles: roles,
	}
	return response, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func decodeMakeSignInRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeSignUpAdminRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.SignUpAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeSignUpUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.SignUpUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeParseTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.ParseTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeGetRolesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.GetRolesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
