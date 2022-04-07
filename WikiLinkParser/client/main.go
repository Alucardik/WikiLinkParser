package client

import (
	"WikiLinkParser/config"
	"WikiLinkParser/proto"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
	"time"
)

type client struct {
	dialer      proto.WikiLinkParserClient
	conn        *grpc.ClientConn
	isConnected bool
}

var cl = client{isConnected: false}

func (c *client) checkState() bool {
	if !c.isConnected {
		fmt.Println("You are not connected to the server")
	}

	return c.isConnected
}

func (c *client) Connect(address string) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Couldn't connect to server: %v\n", err)
		return
	}

	c.conn = conn
	c.dialer = proto.NewWikiLinkParserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = cl.dialer.EstablishConnection(ctx, &proto.EmptyMsg{})
	if err != nil {
		if err := cl.conn.Close(); err != nil {
		}
		log.Fatalf("Couldn't connect to server: %v\n", err.Error())
	}
	c.isConnected = true
}

func (c *client) Publish(initPage, targetPage string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := c.dialer.PublishTask(ctx, &proto.ParseRequest{
		InitPage:   initPage,
		TargetPage: targetPage,
	})

	if err != nil {
		log.Printf("Couldn't get response from server: %v\n", err)
	}

	fmt.Println("------SERVER RESPONSE------\n", "TraceLen: ", resp.TraceLen, "\nTrace:\n", resp.Trace)
	fmt.Println("-------------------------")
}

func (c *client) Disconnect() {
	if !c.checkState() {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.dialer.AbortConnection(ctx, &proto.EmptyMsg{})
	if err != nil {
		log.Printf("Error while Disconnecting: %v\n", err)
	}

	if err := cl.conn.Close(); err != nil {
	}

	c.isConnected = false
}

func Run() {
	cl.Connect(fmt.Sprintf("localhost:%d", config.SERVER_PORT))
	defer cl.Disconnect()

	fmt.Println("----\tYou have launched WikiParser client\t----\nsee available commands in the README")
	for reader := bufio.NewReader(os.Stdin); ; {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		switch strings.TrimSpace(cmd) {
		case "publish":
			fmt.Println("Enter first wiki page address:")
			initPage, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("ERROR", err)
				continue
			}

			fmt.Println("Enter target wiki page address:")
			targetPage, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("ERROR", err)
				continue
			}

			cl.Publish(strings.TrimSpace(initPage), strings.TrimSpace(targetPage))
		case "exit":
			cl.Disconnect()
			os.Exit(0)
		default:
			fmt.Println("Unknown command")

		}
	}
}
