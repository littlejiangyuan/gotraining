package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtx(t *testing.T)  {
	ca, cancel := context.WithCancel(context.Background())
	cm := context.WithValue(ca, "key1", "val1") //WithValue生成一个新的ctx, 携带一个键值对的数据，键为key1,值为val1
	fmt.Println(cm.Value("key1")) //这儿输出为val1

	go work(ca, "job1")

	cb, _ := context.WithTimeout(cm, time.Second * 3)
	go work(cb, "job2")

	cc := context.WithValue(cb, "key2", "val2")
	go jobWithValue(cc, "job3")


	time.Sleep(5*time.Second)
	cancel()
	time.Sleep(time.Second)
}


func work(ctx context.Context, name string){
	for{
		select{
		case <-ctx.Done():
			println(name," 工作结束")
			return
		default:
			println(name," 执行中...")
			time.Sleep(time.Second)
		}
	}
}


func jobWithValue(ctx context.Context, name string){
	for{
		select {
		case <-ctx.Done():
			println(name," 工作结束")
			return
		default:
			value:=ctx.Value("key2").(string)
			println(name, " key2的值为：", value)
			time.Sleep(time.Second)
		}
	}
}

