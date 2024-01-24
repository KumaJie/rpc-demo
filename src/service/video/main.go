package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/etcd"
	"rpc-douyin/src/util/log"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.Server.Video.Port))
	if err != nil {
		panic(err)
	}
	// 默认的接收数据大小约为4MB
	s := grpc.NewServer(grpc.MaxRecvMsgSize(config.Cfg.File.Max))
	video.RegisterVideoServiceServer(s, &VideoServiceImpl{})

	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}
	e := etcd.Endpoint{
		Name: config.Cfg.Server.Video.Name,
		Addr: config.Cfg.Server.Video.Host,
		Port: config.Cfg.Server.Video.Port,
	}
	err = service.Register(&e)
	if err != nil {
		panic(err)
	}
	log.Info("VideoService: service start", zap.Int("port", config.Cfg.Server.Video.Port))

	//consumer, _ := mq.NewConsumer()
	//topic := "publish"
	//partitions, _ := consumer.Partitions(topic)
	//for _, partition := range partitions {
	//	cp, _ := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	//	go CoverHandler(cp.Messages())
	//}

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

//func CoverHandler(messages <-chan *sarama.ConsumerMessage) {
//	for msg := range messages {
//		title := string(msg.Value)
//		log.Info("VideoService: start generate cover", zap.String("title", title))
//		coverPath := path.Join(config.Cfg.File.Dir, title) + ".jpg"
//		cmd := exec.Command("ffmpeg",
//			"-i", path.Join(config.Cfg.File.Dir, title)+".mp4",
//			"-vframes", "1",
//			"-update", "true",
//			"-y",
//			"-f", "image2",
//			coverPath)
//		if err := cmd.Run(); err != nil {
//			log.Error("VideoService: generate cover failed", zap.String("cover", coverPath), zap.Error(err))
//			return
//		}
//		log.Info("VideoService: generate cover succeed", zap.String("cover", coverPath))
//	}
//}
