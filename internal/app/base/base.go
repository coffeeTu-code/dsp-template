package base

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"

	"dsp-template/api/base"
)

func main() {

	// 应用程序层传输安全性（ALTS）是Google开发的双向身份验证和传输加密系统。
	// 它用于保护Google基础架构内的RPC通信。ALTS与双向TLS类似，但经过设计和优化可满足Google生产环境的需求。
	{
		altsTC := alts.NewClientCreds(alts.DefaultClientOptions())
		conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(altsTC))
		base.NewBaseServiceClient(conn)
	}

	// ---
	{
		altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
		server := grpc.NewServer(grpc.Creds(altsTC))
		base.RegisterBaseServiceServer(server, &BI{})
	}

}

type BI struct {
}

func (b *BI) FindBase(ctx context.Context, in *base.BaseRequest) (*base.BaseResponse, error) {
	return nil, nil
}
