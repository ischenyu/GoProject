package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

// Packet 代表数据包结构
type Packet struct {
	SeqNum uint32 // 序列号
	Data   []byte // 实际数据
}

// Server 处理 UDP 数据包并发送 ACK
func startServer(ctx context.Context, addr string) error {
	// 监听 UDP 地址
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345})
	if err != nil {
		return fmt.Errorf("listen failed: %v", err)
	}
	defer conn.Close()

	// 用于记录已接收的序列号
	receivedSeq := make(map[uint32]bool)
	buf := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// 设置读取超时
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			n, clientAddr, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}

			// 解析数据包
			if n < 4 {
				continue
			}
			seqNum := binary.BigEndian.Uint32(buf[:4])
			data := buf[4:n]

			// 避免重复处理
			if receivedSeq[seqNum] {
				continue
			}
			receivedSeq[seqNum] = true

			// 发送 ACK
			ack := make([]byte, 4)
			binary.BigEndian.PutUint32(ack, seqNum)
			conn.WriteToUDP(ack, clientAddr)

			// 模拟处理数据
			fmt.Printf("Server received packet %d: %s\n", seqNum, string(data))
		}
	}
}

// Client 发送数据包并等待 ACK
func startClient(ctx context.Context, addr string, messages []string) error {
	// 建立 UDP 连接
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345})
	if err != nil {
		return fmt.Errorf("dial failed: %v", err)
	}
	defer conn.Close()

	// 使用 sync.Pool 优化内存分配
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	// 用于跟踪已发送但未确认的包
	pending := make(map[uint32][]byte)
	var mu sync.Mutex
	var seqNum uint32

	// ACK 接收 goroutine
	ackChan := make(chan uint32, 100)
	go func() {
		buf := make([]byte, 4)
		for {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}
			if n == 4 {
				ackSeq := binary.BigEndian.Uint32(buf[:4])
				ackChan <- ackSeq
			}
		}
	}()

	// 发送数据包
	for _, msg := range messages {
		seqNum++
		buf := pool.Get().([]byte)
		binary.BigEndian.PutUint32(buf[:4], seqNum)
		copy(buf[4:], msg)

		mu.Lock()
		pending[seqNum] = buf[:4+len(msg)]
		mu.Unlock()

		// 重传逻辑
		for attempts := 0; attempts < 3; attempts++ {
			conn.Write(pending[seqNum])

			// 等待 ACK 或超时
			timer := time.NewTimer(500 * time.Millisecond)
			select {
			case ackSeq := <-ackChan:
				if ackSeq == seqNum {
					fmt.Printf("Client received ACK for packet %d\n", seqNum)
					mu.Lock()
					delete(pending, seqNum)
					pool.Put(buf)
					mu.Unlock()
					break
				}
			case <-timer.C:
				fmt.Printf("Timeout for packet %d, retrying...\n", seqNum)
				continue
			case <-ctx.Done():
				return ctx.Err()
			}
			timer.Stop()
		}
	}

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 启动服务器
	go startServer(ctx, "127.0.0.1:12345")
	time.Sleep(100 * time.Millisecond) // 确保服务器启动

	// 启动客户端
	messages := []string{"Hello", "World", "Reliable", "UDP"}
	err := startClient(ctx, "127.0.0.1:12345", messages)
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
	}
}
