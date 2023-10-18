package interceptor

import (
	"context"
	"strconv"
	"strings"
	"time"

	"chative-server-go/utils/common/constant"
	"chative-server-go/utils/common/utils"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type ClientInterceptor struct {
	AppId string
	Key   string
}

func (cli *ClientInterceptor) Interceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	appid := cli.AppId
	key := cli.Key
	uuidStr := uuid.New().String()
	nonce := strings.ReplaceAll(uuidStr, "-", "")
	timestamp := time.Now().UnixMilli()

	reqByte, err := proto.Marshal(req.(proto.Message)) //prototext.MarshalOptions{Multiline: true}.Marshal(req.(proto.Message))
	if err != nil {
		log.WithFields(log.Fields{"api": "interceptor"}).Error(err)
		return err
	}
	dataToSign := utils.PreSignedData(appid, timestamp, nonce, reqByte)
	sig := utils.Sign([]byte(key), dataToSign)
	log.WithFields(log.Fields{"api": "interceptor", "data size:": len(dataToSign), "signature": sig}).Info("sign")
	ctx = metadata.AppendToOutgoingContext(ctx, constant.MetaDataKeyAppid, appid, constant.MetaDataKeyTimestamp, strconv.FormatInt(timestamp, 10),
		constant.MetaDataKeyNonce, nonce, constant.MetaDataKeyVersion, "2", constant.MetaDataKeySignature, sig)
	// Calls the invoker to execute RPC
	err = invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		log.WithFields(log.Fields{"api": "interceptor"}).Error(err)
	}
	return err
}
