package mainrpc

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"chative-server-go/internal/config"
	"chative-server-go/proto"
	"chative-server-go/utils/common/interceptor"

	"google.golang.org/grpc"
)

var (
	client *rpcClient
)

type rpcClient struct {
	cliTeam proto.InternalTeamsServiceClient
	appID   string
	pid     string

	cliNotify proto.InternalNotifyServiceClient

	cliAccount proto.InternalAccountsServiceClient
}

func (c *rpcClient) sendNotify(content string, uids, gids []string, apn string) error {
	req := &proto.NotifySendRequest{
		Content: &content,
		Uids:    uids,
		Gids:    gids,
	}
	if apn != "" {
		req.DefaultNotification = &apn
	}
	res, err := c.cliNotify.SendNotify(context.Background(), req)
	if err != nil {
		return err
	}
	if *res.Status != 0 {
		return errors.New(strconv.Itoa(int(*res.Status)) + ":" + *res.Reason)
	}
	return nil
}

func (c *rpcClient) joinTeams(team string, uids []string) error {
	joins := make([]*proto.JoinLeaveRequest_JoinLeaveInfo, len(uids))
	for i := range uids {
		joins[i] = &proto.JoinLeaveRequest_JoinLeaveInfo{
			Team: &team,
			Uid:  &uids[i],
		}
	}
	parentID, id, Ancestors, Status, OrderNum, Domain, Remark := uint64(0), uint64(0), "", false, uint32(0), "chative", ""
	c.cliTeam.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateRequest{
		Name: &team, Id: &id, ParentId: &parentID, Ancestors: &Ancestors,
		Status: &Status, OrderNum: &OrderNum, Appid: &c.appID, Pid: &c.pid, Domain: &Domain, Remark: &Remark})
	res, err := c.cliTeam.Join(context.Background(), &proto.JoinLeaveRequest{Joinleaves: joins, Pid: &c.pid, Appid: &c.appID})
	if err != nil {
		return err
	}
	if *res.Status == 0 {
		return nil
	}
	c.cliTeam.Leave(context.Background(), &proto.JoinLeaveRequest{Joinleaves: joins, Pid: &c.pid, Appid: &c.appID})
	return errors.New(*res.Reason)

}

func (c *rpcClient) genLoginInfo(uid, ua string, supportTransfer bool) ([]byte, error) {
	res, err := c.cliAccount.GenLoginInfo(context.Background(), &proto.LoginInfoReq{
		Uid:             &uid,
		Ua:              &ua,
		SupportTransfer: &supportTransfer,
	})
	if err != nil {
		return nil, err
	}
	if *res.Status != 0 {
		err = errors.New(strconv.Itoa(int(*res.Status)) + ":" + *res.Reason)
		return nil, err
	}
	resData := new(proto.LoginInfoRes)
	if err = res.Data.UnmarshalTo(resData); err != nil {
		return nil, err
	}
	return json.Marshal(resData)
}

func (c *rpcClient) blockConversation(operator, conversation string, blockStatus int32) error {
	resp, err := c.cliAccount.BlockConversation(context.Background(),
		&proto.BlockConversationRequest{Operator: &operator, ConversationId: &conversation,
			Block: &blockStatus})
	if err != nil {
		return err
	}

	if *resp.Status != 0 {
		return errors.New(*resp.Reason)
	}
	return nil
}

func Init(c config.Config) {
	inter := interceptor.ClientInterceptor{AppId: c.MainGrpc.AppID, Key: c.MainGrpc.AppSecret}
	conn, err := grpc.Dial(c.MainGrpc.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(inter.Interceptor))
	if err != nil {
		panic(err)
	}

	cliTeam := proto.NewInternalTeamsServiceClient(conn)
	cliNotify := proto.NewInternalNotifyServiceClient(conn)
	cliAccount := proto.NewInternalAccountsServiceClient(conn)
	client = &rpcClient{cliTeam: cliTeam, cliNotify: cliNotify, cliAccount: cliAccount,
		appID: c.MainGrpc.AppID, pid: c.MainGrpc.PID}
}

func JoinTeams(team string, uids []string) error {
	return client.joinTeams(team, uids)
}

func SendNotify(content string, uids []string, apn string) error {
	return client.sendNotify(content, uids, nil, apn)
}

func SendGroupNotify(content string, gids []string, apn string) error {
	return client.sendNotify(content, nil, gids, apn)
}

func GenLoginInfo(uid, ua string, supportTransfer bool) ([]byte, error) {
	return client.genLoginInfo(uid, ua, supportTransfer)
}

func BlockConversation(operator, conversation string, blockStatus int32) error {
	return client.blockConversation(operator, conversation, blockStatus)
}
