// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: video.proto

package video

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	user "rpc-douyin/src/proto/user"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`   //视频数据
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"` // 视频标题
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PublishRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PublishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户ID
}

func (x *PublishListRequest) Reset() {
	*x = PublishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListRequest) ProtoMessage() {}

func (x *PublishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListRequest.ProtoReflect.Descriptor instead.
func (*PublishListRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *PublishListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *PublishListResponse) Reset() {
	*x = PublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListResponse) ProtoMessage() {}

func (x *PublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListResponse.ProtoReflect.Descriptor instead.
func (*PublishListResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *PublishListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type PublishIdListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId []int64 `protobuf:"varint,1,rep,packed,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *PublishIdListResponse) Reset() {
	*x = PublishIdListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishIdListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishIdListResponse) ProtoMessage() {}

func (x *PublishIdListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishIdListResponse.ProtoReflect.Descriptor instead.
func (*PublishIdListResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *PublishIdListResponse) GetVideoId() []int64 {
	if x != nil {
		return x.VideoId
	}
	return nil
}

type PublishCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户ID
}

func (x *PublishCountRequest) Reset() {
	*x = PublishCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCountRequest) ProtoMessage() {}

func (x *PublishCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCountRequest.ProtoReflect.Descriptor instead.
func (*PublishCountRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

func (x *PublishCountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PublishCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PublishCountResponse) Reset() {
	*x = PublishCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCountResponse) ProtoMessage() {}

func (x *PublishCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCountResponse.ProtoReflect.Descriptor instead.
func (*PublishCountResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

func (x *PublishCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type FeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime *int64 `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3,oneof" json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	UserId     *int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
}

func (x *FeedRequest) Reset() {
	*x = FeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRequest) ProtoMessage() {}

func (x *FeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRequest.ProtoReflect.Descriptor instead.
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{6}
}

func (x *FeedRequest) GetLatestTime() int64 {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return 0
}

func (x *FeedRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

type FeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`     // 视频列表
	NextTime  *int64   `protobuf:"varint,2,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time,omitempty"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (x *FeedResponse) Reset() {
	*x = FeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResponse) ProtoMessage() {}

func (x *FeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResponse.ProtoReflect.Descriptor instead.
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{7}
}

func (x *FeedResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *FeedResponse) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

type GetVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetVideoRequest) Reset() {
	*x = GetVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRequest) ProtoMessage() {}

func (x *GetVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{8}
}

func (x *GetVideoRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *GetVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video *Video `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *GetVideoResponse) Reset() {
	*x = GetVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoResponse) ProtoMessage() {}

func (x *GetVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoResponse.ProtoReflect.Descriptor instead.
func (*GetVideoResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{9}
}

func (x *GetVideoResponse) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频唯一标识
	Author        *user.User `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                                     // 视频作者信息
	PlayUrl       string     `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string     `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64      `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64      `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool       `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string     `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{10}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *user.User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x32, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x6d, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x22, 0x71, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x05, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0xd8, 0x03, 0x0a, 0x0c, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_video_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),        // 0: proto.video.PublishRequest
	(*PublishListRequest)(nil),    // 1: proto.video.PublishListRequest
	(*PublishListResponse)(nil),   // 2: proto.video.PublishListResponse
	(*PublishIdListResponse)(nil), // 3: proto.video.PublishIdListResponse
	(*PublishCountRequest)(nil),   // 4: proto.video.PublishCountRequest
	(*PublishCountResponse)(nil),  // 5: proto.video.PublishCountResponse
	(*FeedRequest)(nil),           // 6: proto.video.FeedRequest
	(*FeedResponse)(nil),          // 7: proto.video.FeedResponse
	(*GetVideoRequest)(nil),       // 8: proto.video.GetVideoRequest
	(*GetVideoResponse)(nil),      // 9: proto.video.GetVideoResponse
	(*Video)(nil),                 // 10: proto.video.Video
	(*user.User)(nil),             // 11: proto.user.User
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_video_proto_depIdxs = []int32{
	10, // 0: proto.video.PublishListResponse.video_list:type_name -> proto.video.Video
	10, // 1: proto.video.FeedResponse.video_list:type_name -> proto.video.Video
	10, // 2: proto.video.GetVideoResponse.video:type_name -> proto.video.Video
	11, // 3: proto.video.Video.author:type_name -> proto.user.User
	0,  // 4: proto.video.VideoService.VideoPublish:input_type -> proto.video.PublishRequest
	1,  // 5: proto.video.VideoService.GetPublishList:input_type -> proto.video.PublishListRequest
	1,  // 6: proto.video.VideoService.GetPublishId:input_type -> proto.video.PublishListRequest
	4,  // 7: proto.video.VideoService.PublishCount:input_type -> proto.video.PublishCountRequest
	6,  // 8: proto.video.VideoService.Feed:input_type -> proto.video.FeedRequest
	8,  // 9: proto.video.VideoService.GetVideo:input_type -> proto.video.GetVideoRequest
	12, // 10: proto.video.VideoService.VideoPublish:output_type -> google.protobuf.Empty
	2,  // 11: proto.video.VideoService.GetPublishList:output_type -> proto.video.PublishListResponse
	3,  // 12: proto.video.VideoService.GetPublishId:output_type -> proto.video.PublishIdListResponse
	5,  // 13: proto.video.VideoService.PublishCount:output_type -> proto.video.PublishCountResponse
	7,  // 14: proto.video.VideoService.Feed:output_type -> proto.video.FeedResponse
	9,  // 15: proto.video.VideoService.GetVideo:output_type -> proto.video.GetVideoResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishIdListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishCountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishCountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_video_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_video_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
