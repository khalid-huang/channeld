module channeld.clewcat.com/channeld/examples/channeld-ue-chat

go 1.16

require (
	channeld.clewcat.com/channeld v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	go.uber.org/zap v1.19.1
	google.golang.org/protobuf v1.27.1
)

replace channeld.clewcat.com/channeld => ../..
