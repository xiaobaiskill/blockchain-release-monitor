package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleErr(err error) error {
	switch err {
	default:
		return status.Errorf(codes.Internal, "internal error")
	}
}
