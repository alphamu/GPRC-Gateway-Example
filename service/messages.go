package service

import (
	"github.com/alphamu/goecho/proto"
	"golang.org/x/net/context"
		"github.com/alphamu/goecho/persist"
)

type Server struct{
}

func (s *Server) PostMessages(ctx context.Context, in *proto.MessagesRequest) (*proto.MessagesResponse, error) {
	if len(in.Messages) > 0 {
		_, err := persist.WriteMessageToDb(in.Messages, "0", "1")
		if err != nil {
			return nil, err
		}
	}
	return &proto.MessagesResponse{Success: true}, nil
}

func (s *Server) GetMessages(ctx context.Context, in *proto.MessagesRequest) (*proto.MessagesResponse, error) {
	messages, err := persist.ReadMessagesForUser("1")
	return &proto.MessagesResponse{Messages: messages}, err
}

func (s *Server) Echo(ctx context.Context, in *proto.EchoRequest) (*proto.EchoResponse, error) {
	return &proto.EchoResponse{Message:"OK"}, nil
}