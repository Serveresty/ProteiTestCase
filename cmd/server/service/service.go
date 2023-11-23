package service

import (
	"context"
	"fmt"
	getdata "proteitestcase/internal/server_data/get_data"
	"proteitestcase/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyDEMServer struct {
	api.UnimplementedDEMServer
}

func (s *MyDEMServer) GetInfoAboutUser(ctx context.Context, req *api.GetInfoRequest) (*api.GetInfoResponse, error) {
	_, err := CheckAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("Not authenticated: %v", err)
	}
	if (req.UsersData.Id == nil) && (req.UsersData.Name == "") && (req.UsersData.Email == "") && (req.UsersData.WorkPhone == "") {
		usersData, err1 := getdata.GetAllUsers()
		if err1 != nil {
			return &api.GetInfoResponse{
				Status:    status.New(codes.NotFound, "").String(),
				UsersData: []*api.OutputUsersData{},
			}, err1
		}
		return &api.GetInfoResponse{
			Status:    status.New(codes.OK, "").String(),
			UsersData: usersData,
		}, nil
	}

	usersData, err := getdata.GetUsersByFilter(req.UsersData)
	if err != nil {
		return &api.GetInfoResponse{
			Status:    status.New(codes.NotFound, "").String(),
			UsersData: []*api.OutputUsersData{},
		}, err
	}

	return &api.GetInfoResponse{
		Status:    status.New(codes.OK, "").String(),
		UsersData: usersData,
	}, nil
}

func (s *MyDEMServer) CheckAbsenceStatus(ctx context.Context, req *api.AbsenceStatusRequest) (*api.AbsenceStatusResponse, error) {
	_, err := CheckAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("Not authenticated: %v", err)
	}
	if (req.InputAbsenceData.DateFrom == nil) && (req.InputAbsenceData.PersonIds == nil) && (req.InputAbsenceData.DateTo == nil) {
		usersData, err1 := getdata.GetAllAbsence()
		if err1 != nil {
			return &api.AbsenceStatusResponse{
				Status:      status.New(codes.NotFound, "").String(),
				AbsenceData: []*api.OutputAbsenceData{},
			}, err1
		}
		return &api.AbsenceStatusResponse{
			Status:      status.New(codes.OK, "").String(),
			AbsenceData: usersData,
		}, nil
	}

	usersData, err := getdata.GetAbsenceByFilter(req.InputAbsenceData)
	if err != nil {
		return &api.AbsenceStatusResponse{
			Status:      status.New(codes.NotFound, "").String(),
			AbsenceData: []*api.OutputAbsenceData{},
		}, err
	}

	return &api.AbsenceStatusResponse{
		Status:      status.New(codes.OK, "").String(),
		AbsenceData: usersData,
	}, nil
}
