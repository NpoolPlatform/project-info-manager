package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/projectinfomgr"

	constant "github.com/NpoolPlatform/project-info-manager/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.ProjectInfoManagerClient) (cruder.Any, error)) (cruder.Any, error) {
	const CtxTimeoutSecond = 10
	_ctx, cancel := context.WithTimeout(ctx, CtxTimeoutSecond*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get project info manager connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewProjectInfoManagerClient(conn)

	return fn(_ctx, cli)
}

func GetProjectInfoManagerInfoOnly(ctx context.Context, conds cruder.FilterConds) (*npool.CoinDescription, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.ProjectInfoManagerClient) (cruder.Any, error) {
		// DO RPC CALL HERE WITH conds PARAMETER
		return &npool.CoinDescription{}, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get project info manager : %v", err)
	}
	return info.(*npool.CoinDescription), nil
}
