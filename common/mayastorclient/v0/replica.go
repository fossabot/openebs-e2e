package v0

import (
	"context"
	"fmt"
	"time"

	mayastorGrpc "github.com/openebs/openebs-e2e/common/mayastorclient/v0/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type v0MayastorReplica struct {
	Uuid  string                            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                                       // uuid of the replica
	Pool  string                            `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`                                       // name of the pool
	Thin  bool                              `protobuf:"varint,3,opt,name=thin,proto3" json:"thin,omitempty"`                                      // thin provisioning
	Size  uint64                            `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`                                      // size of the replica in bytes
	Share mayastorGrpc.ShareProtocolReplica `protobuf:"varint,5,opt,name=share,proto3,enum=mayastor.ShareProtocolReplica" json:"share,omitempty"` // protocol used for exposing the replica
	Uri   string                            `protobuf:"bytes,6,opt,name=uri,proto3" json:"uri,omitempty"`                                         // uri usable by nexus to access it
}

func (msr v0MayastorReplica) Description() string {
	return fmt.Sprintf("Uuid=%s; Pool=%s; Thin=%v; Size=%d; Share=%s; Uri=%s;",
		msr.Uuid, msr.Pool, msr.Thin, msr.Size, msr.Share, msr.Uri)
}

func (msr v0MayastorReplica) GetUuid() string {
	return msr.Uuid
}

func (msr v0MayastorReplica) GetPool() string {
	return msr.Pool
}

func (msr v0MayastorReplica) GetThin() bool {
	return msr.Thin
}

func (msr v0MayastorReplica) GetSize() uint64 {
	return msr.Size
}
func (msr v0MayastorReplica) GetShareString() string {
	return msr.Share.String()
}

func (msr v0MayastorReplica) GetUri() string {
	return msr.Uri
}

func listReplica(address string) ([]v0MayastorReplica, error) {
	var replicaInfos []v0MayastorReplica
	var err error
	addrPort := fmt.Sprintf("%s:%d", address, mayastorPort)
	conn, err := grpc.Dial(addrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logf.Log.Info("listReplica", "error", err)
		return replicaInfos, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logf.Log.Info("listReplicas", "error on close", err)
		}
	}(conn)

	c := mayastorGrpc.NewMayastorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	var response *mayastorGrpc.ListReplicasReply
	retryBackoff(func() error {
		response, err = c.ListReplicas(ctx, &null)
		return err
	})

	if err == nil {
		if response != nil {
			for _, replica := range response.Replicas {
				ri := v0MayastorReplica{
					Uuid:  replica.Uuid,
					Pool:  replica.Pool,
					Thin:  replica.Thin,
					Size:  replica.Size,
					Share: replica.Share,
					Uri:   replica.Uri,
				}
				replicaInfos = append(replicaInfos, ri)
			}
		} else {
			err = fmt.Errorf("nil response for ListReplicas on %s", address)
			logf.Log.Info("listReplicas", "error", err)
		}
	} else {
		err = niceError(err)
		logf.Log.Info("listReplicas", "error", err)
	}
	return replicaInfos, err
}

// RmReplica remove a replica identified by node and uuid
func RmReplica(address string, uuid string) error {
	logf.Log.Info("RmReplica", "address", address, "UUID", uuid)
	addrPort := fmt.Sprintf("%s:%d", address, mayastorPort)
	conn, err := grpc.Dial(addrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logf.Log.Info("rmReplicas", "error", err)
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logf.Log.Info("RmReplicas", "error on close", err)
		}
	}(conn)
	c := mayastorGrpc.NewMayastorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req := mayastorGrpc.DestroyReplicaRequest{Uuid: uuid}
	retryBackoff(func() error {
		_, err = c.DestroyReplica(ctx, &req)
		return err
	})

	return niceError(err)
}

// CreateReplicaExt create a replica on a mayastor node
func CreateReplicaExt(address string, uuid string, size uint64, pool string, thin bool) error {
	shareProto := mayastorGrpc.ShareProtocolReplica_REPLICA_NVMF
	logf.Log.Info("CreateReplica", "address", address, "UUID", uuid, "size", size, "pool", pool, "Thin", thin, "Share", shareProto)
	addrPort := fmt.Sprintf("%s:%d", address, mayastorPort)
	var err error

	conn, err := grpc.Dial(addrPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logf.Log.Info("createReplica", "error", err)
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logf.Log.Info("CreateReplicaExt", "error on close", err)
		}
	}(conn)
	c := mayastorGrpc.NewMayastorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req := mayastorGrpc.CreateReplicaRequest{
		Uuid:  uuid,
		Size:  size,
		Thin:  thin,
		Pool:  pool,
		Share: shareProto,
	}

	retryBackoff(func() error {
		_, err = c.CreateReplica(ctx, &req)
		return err
	})

	return niceError(err)
}

// CreateReplica create a replica on a mayastor node, with parameters
//
//	 thin fixed to false and share fixed to NVMF.
//	Other parameters must be specified
func CreateReplica(address string, uuid string, size uint64, pool string) error {
	return CreateReplicaExt(address, uuid, size, pool, false)
}

// ListReplicas given a list of node ip addresses, enumerate the set of replicas on mayastor using gRPC on each of those nodes
// returns accumulated errors if gRPC communication failed.
func ListReplicas(addrs []string) ([]v0MayastorReplica, error) {
	var accErr error
	var replicaInfos []v0MayastorReplica
	for _, address := range addrs {
		replicaInfo, err := listReplica(address)
		if err == nil {
			replicaInfos = append(replicaInfos, replicaInfo...)
		} else {
			if accErr != nil {
				accErr = fmt.Errorf("%v;%v", accErr, err)
			} else {
				accErr = err
			}
		}
	}
	return replicaInfos, accErr
}

// RmNodeReplicas given a list of node ip addresses, delete the set of replicas on mayastor using gRPC on each of those nodes
// returns errors if gRPC communication failed.
func RmNodeReplicas(addrs []string) error {
	var accErr error
	for _, address := range addrs {
		replicaInfos, err := listReplica(address)
		if err == nil {
			for _, replicaInfo := range replicaInfos {
				err = RmReplica(address, replicaInfo.GetUuid())
			}
		}
		if err != nil {
			if accErr != nil {
				accErr = fmt.Errorf("%v;%v", accErr, err)
			} else {
				accErr = err
			}
		}
	}
	return accErr
}

// FindReplicas given a list of node ip addresses, enumerate the set of replicas on mayastor using gRPC on each of those nodes
// returns accumulated errors if gRPC communication failed.
func FindReplicas(uuid string, addrs []string) ([]v0MayastorReplica, error) {
	var accErr error
	var replicaInfos []v0MayastorReplica
	for _, address := range addrs {
		replicaInfo, err := listReplica(address)
		if err == nil {
			for _, repl := range replicaInfo {
				if repl.GetUuid() == uuid {
					replicaInfos = append(replicaInfos, repl)
				}
			}
		} else {
			if accErr != nil {
				accErr = fmt.Errorf("%v;%v", accErr, err)
			} else {
				accErr = err
			}
		}
	}
	return replicaInfos, accErr
}
