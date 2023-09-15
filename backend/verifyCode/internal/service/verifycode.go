package service

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	pb "verifyCode/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: GetCode(int(req.Length), req.Type),
	}, nil
}

func GetCode(l int, t pb.TYPE) string {
	var chars string
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		chars = RandCode("0123456789", 4, l)
	case pb.TYPE_LETTER:
		chars = RandCode("abcdefghijklmnopqrstuvwxyz", 5, l)
	case pb.TYPE_MIXED:
		chars = RandCode("0123456789abcdefghijklmnopqrstuvwxyz", 6, l)
	default:
		fmt.Println("验证码类型暂不支持")
	}
	return chars
}

func RandCode(chars string, idxBits, l int) string {
	// 形成掩码
	idxMask := 1<<idxBits - 1
	// 63位可以使用的最大次数
	idxMax := 63 / idxBits

	// 利用string builder 构建结果缓冲
	sb := strings.Builder{}
	sb.Grow(l)

	// 循环生成随机数
	// i 索引
	// cache 随机数缓存
	// remain 随机数还能用多少次
	for i, cache, remain := l-1, rand.Int63(), idxMax; i >= 0; {
		// 随机缓存不足，重新生成
		if remain == 0 {
			cache, remain = rand.Int63(), idxMax
		}
		// 利用掩码生成随机索引，有效索引为小于字符集合长度
		if idx := int(cache & int64(idxMask)); idx < len(chars) {
			sb.WriteByte(chars[idx])
			i--
		}
		// 利用下一组随机数位
		cache >>= idxBits
		remain--
	}
	return sb.String()
}
