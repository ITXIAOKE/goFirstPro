package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"strings"
	"time"
	global2 "xkshop/v1/xkshop_srv/user_srv/global"
	model2 "xkshop/v1/xkshop_srv/user_srv/model"
	proto2 "xkshop/v1/xkshop_srv/user_srv/proto"
)

type UserServer struct {
}

func ModelToResponse(user model2.User) *proto2.UserInfoResponse {
	userInfoRsp := proto2.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil { //这里birthday有可能为空，所有这里要判断一下
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return &userInfoRsp
}

//分页功能

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//用户列表接口实现

func (s *UserServer) GetUserList(ctx context.Context, req *proto2.PageInfoRequest) (*proto2.UserListResponse, error) {
	var users []model2.User
	result := global2.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	resp := &proto2.UserListResponse{}
	resp.Total = int32(result.RowsAffected)

	global2.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		resp.Data = append(resp.Data, userInfoRsp)
	}

	return resp, nil
}

//根据手机号查询用户接口

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto2.MobileRequest) (*proto2.UserInfoResponse, error) {
	var user model2.User
	result := global2.DB.Where(&model2.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return userInfoRsp, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *proto2.IdRequest) (*proto2.UserInfoResponse, error) {
	var user model2.User
	result := global2.DB.First(&user, req.Id) //通过id查询用户，id是主键
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return userInfoRsp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto2.CreateUserInfo) (*proto2.UserInfoResponse, error) {
	//新建用户
	var user model2.User
	result := global2.DB.Where(&model2.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = req.Mobile
	user.NickName = req.NickName

	//密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd) //生产用的password

	//用户存数据库
	result = global2.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	userInfoRsp := ModelToResponse(user)
	return userInfoRsp, nil

}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto2.UpdateUserInfo) (*emptypb.Empty, error) {
	//更新个人用户信息
	var user model2.User
	result := global2.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	birthDay := time.Unix(int64(req.Birthday), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = global2.DB.Save(&user) //更新数据库用户信息
	if result != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) CheckPassword(ctx context.Context, req *proto2.PasswordCheckInfo) (*proto2.CheckResponse, error) {
	//检查密码
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	pwdInfo := strings.Split(req.EncryptedPassword, "$")
	//fmt.Println(pwdInfo) //使用debug方法查找切割字符串后，数组中是第一个元素
	check := password.Verify(req.PassWord, pwdInfo[2], pwdInfo[3], options)
	return &proto2.CheckResponse{Success: check}, nil
}
