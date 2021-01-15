package dataServer

import (
	"context"
	"encoding/json"
	"etcGasDataCenter/db"
	"etcGasDataCenter/types"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sync"
)

const (
	kafkaConn1 = "localhost:9092" //本机
	kafkaConn2 = "127.0.0.1:9093"
	kafkaConn3 = "127.0.0.1:9094"
	topic      = "topic1"
)

//代理
//var brokerAddrs = []string{kafkaConn1, kafkaConn2, kafkaConn3}

//生产数据
func Producer(msgdata []byte, id string) {
	config := sarama.NewConfig()

	// 等待服务器 所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	//producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	producer, err := sarama.NewSyncProducer([]string{types.KafkaIp}, config)

	if err != nil {
		log.Println("sarama.NewSyncProducer errpr:", err)
		return
	}

	defer func() {
		_ = producer.Close()
	}()

	data := new(types.KafKaBillHourMsg)
	err = json.Unmarshal(msgdata, &data)
	if err != nil {
		log.Println("dd +++++++++++++++++json.Unmarshal error:", err)
		return
	}

	d := new(types.KafKaReply)

	d.Head.Topic = data.Head.Topic                              //消息类型 如exitdata
	d.Head.Index = data.Head.Index                              //消息序号,自增
	d.Head.Topicreply = data.Head.Topicreply                    // 消息回执的主题
	d.Head.Id = data.Head.Id                                    //数据ID
	d.Head.Topictime = time.Now().Format("2006-01-02 15:04:05") // 入kafka的时间
	d.Head.Lane_id = data.Head.Lane_id                          //车道id
	d.Head.Parking_id = data.Head.Parking_id                    //停车场id
	d.Head.Company_id = data.Head.Company_id                    //公司id
	d.Head.Source_type = data.Head.Source_type                  // ddd  vs  zdz

	d.Data.Id = id

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}

	log.Println("回调通知 msgdataTopic = ", data.Head.Topicreply, ",value =d: ", d.Head)
	msg.Topic = data.Head.Topicreply
	//将字符串转换为字节数组
	hdvalue, _ := json.Marshal(d)
	//log.Println("dvalue", string(hdvalue))
	msg.Value = sarama.ByteEncoder(hdvalue)

	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Println("Send message Fail", err)
	}
	log.Printf("Partition = %d, offset=%d\n", partition, offset)
}

var (
	wg sync.WaitGroup
)

//本机测试消费者
func Consumer() {
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	//conf := &sarama.Config{}
	ConsumerGroup, grouperr := sarama.NewConsumerGroup([]string{"localhost:9092"}, "12324", nil)
	if grouperr != nil {
		panic(grouperr)
	}
	log.Println(ConsumerGroup)
	//ConsumerGroup.Consume()
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				log.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

type Kafka struct {
	brokers []string
	topics  []string
	//OffsetNewest int64 = -1
	//OffsetOldest int64 = -2
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
}

func NewKafka() *Kafka {
	return &Kafka{
		brokers:           brokers,
		topics:            topics,
		group:             group,
		channelBufferSize: 2,
		ready:             make(chan bool),
		version:           "1.1.1",
	}
}

var brokers = []string{"172.18.70.21:9092"}
var topics = []string{types.DdkafkaTopic, types.DdkafkaHourTopic} //不参与编译
var group = "39"

func (p *Kafka) Init() func() {
	log.Println("++++++++++++++++++++++++++++++++++kafka init...")

	version, err := sarama.ParseKafkaVersion(p.version)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++Error parsing Kafka version:  ", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = -2                                   // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = p.channelBufferSize                         // channel长度

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(p.brokers, p.group, config)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++++++++Error creating consumer group client: ", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			//util.HandlePanic("client.Consume panic", log.StandardLogger())
		}()
		for {
			if err := client.Consume(ctx, p.topics, p); err != nil {
				log.Println("++++++++++++++++++++++++++++++++Error from consumer: ", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println("+++++++++++++++++++++++++ctx.Err():", ctx.Err())
				return
			}
			p.ready = make(chan bool)
		}
	}()
	<-p.ready
	log.Infoln("+++++++++++++++++++++++++++++Sarama consumer up and running!...")
	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		log.Println("+++++++++++++++++++++++++++++kafka close")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Println("++++++++++++++++++++++++++++++++++Error closing client: ", err)
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (p *Kafka) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(p.ready)
	return nil
}

//清理
// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (p *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

//消费主张
// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (p *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// 具体消费消息
	for message := range claim.Messages() {
		msg := string(message.Value)
		log.Println("+++++++++++++++++++++++++++++++msg:", msg)
		time.Sleep(time.Second)
		//run.Run(msg)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

func UseKafka() {
	k := NewKafka()
	f := k.Init()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	f()
}

type consumerGroupHandler struct {
	name string //groupname
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

//消费主张
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("++++++++++++++%s group Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		//消息处理

		//小时统计
		if msg.Topic == types.DdkafkaHourTopic {
			log.Println("小时统计消息已接受，正处理中+++++++++++++++消息已接受，正处理中+++++++++++")

			err, id := ProcessMessage(msg.Topic, msg.Value)
			if err != nil {
				log.Println("执行ProcessMessage  error:", err)
			}

			// 手动确认消息
			sess.MarkMessage(msg, "")
			//发送回调
			Producer(msg.Value, id)
		}

		//kafka单点出口原始消息入库
		if msg.Topic == types.DdkafkaTopic {
			log.Println("kafka单点出口原始消息入库的消息已接受，正处理中+++++++++++++++消息已接受，正处理中+++++++++++")

			err, id := ProcessMessage(msg.Topic, msg.Value)
			if err != nil {
				log.Println("执行ProcessMessage  error:", err)
			}

			// 手动确认消息
			sess.MarkMessage(msg, "")
			//发送回调
			Producer(msg.Value, id)

			//kafka单点出口原始消息入库成功
			if err == nil {
				//更新回调时间
				uperr := db.ChedckyssjDataUpdate(id)
				if uperr != nil {
					log.Println("db.ChedckyssjDataUpdate error :", uperr)
				}
			}

		}
		log.Println("消息处理完成+++++++消息处理完成+++++++++++")
	}
	return nil
}

//处理消息  msg 消息数据
func ProcessMessage(topic string, msg []byte) (error, string) {
	log.Println("正执行处理消息:ProcessMessage【topic,msg[0:10]的值】 :", topic, string(msg[0:10]))
	switch topic {
	//处理单点流水
	case types.DdkafkaTopic:
		data := new(types.KafKaMsg)
		err := json.Unmarshal(msg, &data)
		if err != nil {
			log.Println("执行  types.DdkafkaTopic  json.Unmarshal error:", err)
			return err, ""
		}

		log.Println("数据中心接收的数据data:", data)
		//数据入库
		log.Println("执行 db.DataStorage(data)，进行数据入库 ")
		inerr := db.DataStorage(data)
		if inerr != nil {
			log.Println("单点车道出口数据入库失败,inerr:", inerr)
			return inerr, data.Data.Bill_id
		}
		//回复回调通知
		return nil, data.Data.Bill_id

		//处理单点ETC加油小时统计
	case types.DdkafkaHourTopic:
		//log.Println("++++++ topic:获取的值：", string(msg))
		data := new(types.KafKaBillHourMsg)
		err := json.Unmarshal(msg, &data)
		if err != nil {
			log.Println("执行小时统计  types.KafKaBillHourMsg  json.Unmarshal error:", err)
			return err, ""
		}

		log.Println("数据中心接收的小时统计数据data:", data)
		//小时统计数据入库
		log.Println("执行 db.HourDataStorage(data)，进行数据入库 ")

		inerr := db.HourDataStorage(data)
		if inerr != nil {
			log.Println("单点车道出口数据中心接收的小时统计数据入库失败:", inerr)
		}
		//回复回调通知
		return nil, data.Data.Bill_id

	case "topic1":
		log.Println(string(msg))
		return nil, ""
	}

	return nil, ""
}

func handleErrors(group *sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Done()
	for err := range (*group).Errors() {
		log.Println("ERROR", err)
	}
}

//消费  group name == c1
func consume(group *sarama.ConsumerGroup, wg *sync.WaitGroup, name string) error {
	log.Println(name+" group "+"start ok ++++++消费 kafka consume ++++++ name:", name)
	wg.Done()
	ctx := context.Background()
	for {
		//c1 ：group  topics := []string{"zdzBillExitDataCollectTopic", "topic1","sun", "billDataCollectTopic"}
		//ddtopic := types.DdkafkaTopic
		//ddHourtopic := types.DdkafkaHourTopic
		//	topics := []string{"topic1"}
		var topics []string
		topics = append(topics, types.DdkafkaTopic)
		topics = append(topics, types.DdkafkaHourTopic)
		log.Println("+++++++++++++++++++消费的topics:", topics)

		//name c1
		handler := consumerGroupHandler{name: name}

		err := (*group).Consume(ctx, topics, handler)
		if err != nil {
			log.Println("++++++++++++++++++【(*group).Consume  error】  +++++++++++++++++++++", err)
			return err
		}
		log.Println("+++++++++++++++++++++++[kafka ok]+++++++++++++++++++++++++")
	}
}

//main 调用 消费kafka
func ConsumerGroup() error {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0
	client, err := sarama.NewClient([]string{types.KafkaIpa, types.KafkaIpb, types.KafkaIpc, types.KafkaIp}, config)
	defer func() {
		_ = client.Close()
	}()

	if err != nil {
		log.Println("++++++++++++++++++++++++++sarama.NewClient 执行出错: ", err)
		return err
	}
	//c1 组
	group1, err := sarama.NewConsumerGroupFromClient("c1", client)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++sarama.NewConsumerGroupFromClient 执行出错: ", err)
		return err

	}
	//group2, err := sarama.NewConsumerGroupFromClient("c2", client)
	//if err != nil {
	//	panic(err)
	//}
	//group3, err := sarama.NewConsumerGroupFromClient("c3", client)
	//if err != nil {
	//	panic(err)
	//}
	defer func() {
		_ = group1.Close()
	}()

	//defer group2.Close()
	//defer group3.Close()

	wg.Add(1)

	//处理kafka数据
	cerr := consume(&group1, &wg, "c1")
	if cerr != nil {
		return cerr
	}
	//go consume(&group2,&wg,"c2")
	//go consume(&group3,&wg,"c3")
	wg.Wait()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
	return nil
}
