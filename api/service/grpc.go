package service

import (
	"context"
	"io"
	"log"
	"strings"
	"sync"

	pb "github.com/RohanDoshi21/messaging-platform/proto"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	pb.GrpcServerServiceServer
	Clients map[string]pb.GrpcServerService_SendMessageServer
	mu      sync.Mutex
}

type Payload struct {
	Username string
}

func (server *GrpcServer) SendMessage(stream pb.GrpcServerService_SendMessageServer) error {
	payload, ok := stream.Context().Value(payloadHeaderKey(payloadHeader)).(string)
	// fmt.Println(payload)
	if !ok {
		return status.Errorf(codes.Internal, "missing required token")
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			// The client has closed the connection.
			break
		}

		if server.Clients[payload] == nil {
			server.mu.Lock()
			server.Clients[payload] = stream
			server.mu.Unlock()
		}

		if err != nil {
			return status.Errorf(codes.Internal, "Error receiving message: %v", err)
		}

		// Find the receiver by username.
		server.mu.Lock()
		receiver, ok := server.Clients[message.Reciever]
		if !ok {
			// If the receiver or sender is not found, send an error message back to the sender.
			// Avoid Deadlock of the server
			server.mu.Unlock()
			continue
		}

		sender, ok := server.Clients[payload]
		server.mu.Unlock()

		if !ok {
			// If the receiver or sender is not found, send an error message back to the sender.
			continue
		}

		messageUUID := uuid.New().String()

		// Forward the message to the receiver.
		err = receiver.Send(&pb.Message{
			Sender:   payload,
			Receiver: message.Reciever,
			Message:  message.Message,
			Id:       messageUUID,
		})
		if err != nil {
			log.Printf("Error sending message to %s: %v", message.Reciever, err)
			continue
		}

		// Send the same message back to the sender as a confirmation.
		err = sender.Send(&pb.Message{
			Sender:   payload,
			Receiver: message.Reciever,
			Message:  message.Message,
			Id:       messageUUID,
		})
		if err != nil {
			log.Printf("Error sending confirmation message to %s: %v", payload, err)
			continue
		}
	}

	// Remove the sender from the Clients map when the client disconnects.
	server.mu.Lock()
	delete(server.Clients, payload)
	server.mu.Unlock()
	return nil
}

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
	payloadHeader       = "payload"
)

func (server *GrpcServer) UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := server.AuthInterceptor(info.FullMethod, ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

type payloadHeaderKey string

func (server *GrpcServer) AuthInterceptor(method string, ctx context.Context) (context.Context, error) {
	// Extract the metadata from the context.
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "metadata not found")
	}

	// Get the authorization token from metadata if present.
	authTokens := md[authorizationHeader]
	if len(authTokens) == 0 {
		// No token found, but it's optional, so return the unmodified context.
		return ctx, nil
	}

	authHeader := authTokens[0] // Assuming a single token is sent in the header.
	fields := strings.Fields(authHeader)

	if len(fields) < 2 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth header format: %v", fields)
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization type: %v", authType)
	}
	accessToken := fields[1]

	claims, err := casdoorsdk.ParseJwtToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token %v invalid", authType)
	}

	ctx = context.WithValue(ctx, payloadHeaderKey(payloadHeader), claims.User.Id)
	return ctx, nil
}

type customServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (css *customServerStream) Context() context.Context {
	return css.ctx
}
func (server *GrpcServer) StreamAuthInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()
	ctx, err := server.AuthInterceptor(info.FullMethod, ctx)
	if err != nil {
		return err
	}
	newStream := &customServerStream{
		ServerStream: stream,
		ctx:          ctx,
	}
	return handler(srv, newStream)
}
