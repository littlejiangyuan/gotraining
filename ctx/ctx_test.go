package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T)  {
	var channel = make(chan struct {})
	go func() {
		time.Sleep(time.Second)
		//close(channel)
	}()

	for {
		select {
		case <-channel:
			fmt.Println("job end...")
			channel<- struct{}{}
		default:
			time.Sleep(time.Second )
			fmt.Println("working...")
		}
	}
	//time.Sleep(time.Second * 50)
}

//	简单说一下这段代码的功能，创建一些Context ca, cb, cc, cd， 它们都是前一个的子Context, cc设置了超时时间，3秒。主协程运行5秒后执行取消操作
func TestCtx(t *testing.T)  {
	//基于context.Background()创建一个可以取消的context. 返回cancel是一个func, 调用它就相当于取消这个任务
	ca, cancel := context.WithCancel(context.Background())
	//WithValue生成一个新的ctx, 携带一个键值对的数据，键为key1,值为val1
	cb := context.WithValue(ca, "key1", "val1")
	fmt.Println(cb.Value("key1")) //这儿输出为val1
	//基于cb生成新的ctx cc。cc是一个带有超时时间的Context,一旦3秒到了以后，任务就应该结束
	cc, _ := context.WithTimeout(cb, time.Second * 3)
	//基于cc创建一个子context cd.
	cd := context.WithValue(cc, "key2", "val2")

	go work(ca, "ca job")  //创建一个协程，监听ca的状态
	go work(cb, "cb job")
	go work(cc, "cc job")
	go jobWithValue(cd, "cd job")

	//主协程休眠5秒
	time.Sleep(5*time.Second)
	cancel() //ca取消操作
	time.Sleep(time.Second)


	/* 输出
	val1
	cd job  key1的值为： val1  --在这儿我们可以知道，Context中的键值对数据是可以带到子Context的
	ca job  执行中...
	cc job  执行中...
	cb job  执行中...
	cb job  执行中...
	cc job  执行中...
	ca job  执行中...
	cd job  key1的值为： val1
	cb job  执行中...
	cd job  key1的值为： val1
	ca job  执行中...
	cc job  执行中...
	cd job  key1的值为： val1
	cc job  执行中...
	ca job  执行中...
	cb job  执行中...
	cd job  工作结束    --cd结束了
	cb job  执行中...
	ca job  执行中...
	cc job  工作结束    --在这儿可以看到到了3秒的超时时间后，cc结束工作了， 它的子Context cd也已经结束工作了
	ca job  工作结束
	cb job  工作结束    --ca执行取消操作后，ca, cb都结束工作了
	*/
}

//一个死循环，当context结束时（包括超时，正常结束或者任务取消）输出工作结束，否则输出执行中...
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
			value:=ctx.Value("key1").(string)
			println(name, " key1的值为：", value)
			time.Sleep(time.Second)
		}
	}
}



