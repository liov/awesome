```go

var producer *nsq.Producer
producer.Publish(topic, []byte(message))

// 消费者
type Consumer struct{}

// 主函数

// 处理消息
func (*Consumer) HandleMessage(msg *nsq.Message) error {
log.Info("receive", msg.NSQDAddress, "message:", string(msg.Body))
return nil
}

```