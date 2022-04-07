package server

import (
	"WikiLinkParser/config"
	"WikiLinkParser/proto"
	"WikiLinkParser/queue_info"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedWikiLinkParserServer
	qInfo      queue_info.QueueInfo
	results    map[uint64]chan *queue_info.ParseResponse
	mutex      sync.Mutex
	nextTaskId uint64
}

func (s *server) PublishTask(_ context.Context, req *proto.ParseRequest) (*proto.ParseResponse, error) {
	s.mutex.Lock()
	curTaskId := s.nextTaskId
	s.nextTaskId++
	s.mutex.Unlock()

	s.results[curTaskId] = make(chan *queue_info.ParseResponse)

	s.qInfo.PublishTask(queue_info.ParseRequest{
		StartPage:  req.InitPage,
		TargetPage: req.TargetPage,
		TaskId:     curTaskId,
	})

	res := <-s.results[curTaskId]
	return &proto.ParseResponse{
		TraceLen: uint32(res.TraceLen),
		Trace:    res.Trace,
	}, nil
}

func (s *server) EstablishConnection(_ context.Context, _ *proto.EmptyMsg) (*proto.ConnectionStatus, error) {
	return &proto.ConnectionStatus{Successful: true}, nil
}

func (s *server) AbortConnection(_ context.Context, _ *proto.EmptyMsg) (*proto.ConnectionStatus, error) {
	return &proto.ConnectionStatus{Successful: true}, nil
}

func Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.SERVER_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	servImpl := server{
		qInfo:      queue_info.CreateQueue(),
		nextTaskId: 0,
		results:    make(map[uint64]chan *queue_info.ParseResponse),
		mutex:      sync.Mutex{},
	}

	servImpl.qInfo.InitQueueConnection(config.QUEUE_HOST, config.QUEUE_PORT, SEND_QUEUE, RECEIVE_QUEUE)
	defer servImpl.qInfo.AbortConnection()

	msgHandler := func(msg queue_info.ParseMsg) {
		switch t := msg.(type) {
		case queue_info.ParseResponse:
			fmt.Println(t.TraceLen)
			fmt.Println(t.Trace)
			servImpl.mutex.Lock()
			servImpl.results[t.TaskId] <- &t
			delete(servImpl.results, t.TaskId)
			servImpl.mutex.Unlock()
		}
	}

	go servImpl.qInfo.GetResults(&msgHandler, queue_info.ParseResponse{})

	s := grpc.NewServer()
	proto.RegisterWikiLinkParserServer(s, &servImpl)
	log.Printf("SERVER listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
