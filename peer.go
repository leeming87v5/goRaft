package raft

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/moxiaomomo/goRaft/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Peer stores others nodes' info
type Peer struct {
	server           *server
	Name             string
	Host             string
	Client           string
	voteRequestState int

	lastActivity      time.Time
	heartbeatInterval time.Duration

	mutex sync.RWMutex
}

// RequestVoteRespChan response of a vote request
type RequestVoteRespChan struct {
	Failed   bool
	Resp     *pb.VoteResponse
	PeerHost string
}

// AppendLogRespChan response of an append-entry request
type AppendLogRespChan struct {
	Failed   bool
	Resp     *pb.AppendEntriesResponse
	PeerHost string
}

// NewPeer creates a new peer instance
func NewPeer(server *server, name, host string, heartbeatInterval time.Duration) *Peer {
	return &Peer{
		server:            server,
		Name:              name,
		Host:              host,
		voteRequestState:  NotYetVote,
		heartbeatInterval: heartbeatInterval,
	}
}

// SetVoteRequestState set current state about votion
func (p *Peer) SetVoteRequestState(state int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.voteRequestState = state
}

// VoteRequestState returns the current state about votion
func (p *Peer) VoteRequestState() int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.voteRequestState
}

// RequestVoteMe send request for leader votion
func (p *Peer) RequestVoteMe(lastLogIndex, lastTerm uint64) {
	conn, err := grpc.Dial(p.Host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("dail rpc failed, err: %s\n", err)
		if conn != nil {
			conn.Close()
		}
		return
	}
	defer conn.Close()

	client := pb.NewRequestVoteClient(conn)
	pb := &pb.VoteRequest{
		Term:          p.server.currentTerm,
		LastLogIndex:  lastLogIndex,
		LastLogTerm:   lastTerm,
		CandidateName: p.server.conf.Name,
		Host:          p.server.conf.Host,
	}
	res, err := client.RequestVoteMe(context.Background(), pb)

	resp := &RequestVoteRespChan{
		Resp:     res,
		PeerHost: p.Host,
	}

	if err != nil {
		resp.Failed = true
		p.server.ch <- resp
	} else {
		resp.Failed = false
		p.server.ch <- resp
	}
}

// RequestAppendEntries send a request to append entries
func (p *Peer) RequestAppendEntries(entries []*pb.LogEntry, sindex, lindex, lterm uint64) {
	if p.server.State() != Leader {
		fmt.Println("only leader can request append entries.")
		return
	}

	conn, err := grpc.Dial(p.Host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("dail rpc failed, err: %s\n", err)
		if conn != nil {
			conn.Close()
		}
		return
	}
	defer conn.Close()

	client := pb.NewAppendEntriesClient(conn)

	req := &pb.AppendEntriesReuqest{
		Term:              p.server.currentTerm,
		FirstLogIndex:     sindex,
		PreLogIndex:       lindex,
		PreLogTerm:        lterm,
		CommitIndex:       p.server.log.CommitIndex(),
		LeaderName:        p.server.conf.Host,
		LeaderHost:        p.server.conf.Host,
		LeaderExHost:      p.server.conf.Client,
		HeartbeatInterval: p.server.heartbeatInterval,
		Entries:           entries,
	}

	res, err := client.AppendEntries(context.Background(), req)

	resp := &AppendLogRespChan{
		Failed:   false,
		Resp:     res,
		PeerHost: p.Host,
	}

	if err != nil {
		resp.Failed = true
		//		fmt.Printf("leader reqeust AppendEntries failed, err:%s\n", err)
		p.server.ch <- resp
		return
	}

	if res.Success {
		p.server.ch <- resp
	} else {
		el := []*pb.LogEntry{}
		for _, e := range p.server.log.entries {
			if e.Entry.GetIndex() <= res.Index {
				continue
			}
			el = append(el, e.Entry)
		}
		req := &pb.AppendEntriesReuqest{
			Term:              p.server.currentTerm,
			FirstLogIndex:     sindex,
			PreLogIndex:       res.Index,
			PreLogTerm:        res.Term,
			CommitIndex:       p.server.log.CommitIndex(),
			LeaderName:        p.server.conf.Host,
			LeaderHost:        p.server.conf.Host,
			LeaderExHost:      p.server.conf.Client,
			HeartbeatInterval: p.server.heartbeatInterval,
			Entries:           el,
		}

		res, err = client.AppendEntries(context.Background(), req)

		resp := &AppendLogRespChan{
			Failed:   false,
			Resp:     res,
			PeerHost: p.Host,
		}

		if err != nil {
			//			fmt.Printf("leader reqeust AppendEntries failed, err:%s\n", err)
			resp.Failed = true
			p.server.ch <- resp
		} else {
			fmt.Printf("Appendentries succeeded: %s %+v\n", p.Host, res)
			p.server.ch <- resp
		}
	}

	//TODO
}
