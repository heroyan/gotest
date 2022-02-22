package ctx

import (
	"context"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

type If1 interface {
	Done() bool
}

type G1 struct {
	address string
}

func (g *G1) Done() bool {
	fmt.Println("hi, i am g1, my address is" + g.address)
	return false
}

type S1 struct {
	If1
	name string
	age int
}

func (s1 *S1) Done() bool {
	fmt.Println("hi, i am s1, my name is" + s1.name)
	return true
}

func TestContext(t *testing.T)  {
	g1 := &G1{
		"guang dong shen zhen",
	}
	g1.Done()
	s1 := &S1{g1, "fafa", 18}
	s1.Done()
	fmt.Println(s1.If1.Done())

	convey.Convey("test with value", t, func() {
		wv := context.WithValue(context.Background(), "hahakey", "hahavalue")
		wv2 := context.WithValue(wv, "hahakey2", "hahavalue1")
		convey.So(wv2.Value("hahakey"), convey.ShouldEqual, "hahavalue")
		convey.So(wv2.Value("hahakey2"), convey.ShouldEqual, "hahavalue1")
		convey.So(wv.Value("hahakey"), convey.ShouldEqual, "hahavalue")
		cancelCtx, cancelFunc := context.WithCancel(wv)
		t.Log(cancelCtx)
		t.Log(wv2.(fmt.Stringer))
		deadlineCtx, cancelFunc2 := context.WithTimeout(cancelCtx, time.Second)
		go func() {
			for {
				select {
				case <-deadlineCtx.Done():
					t.Log("deadline reached")
					return
				default:

				}
			}
		}()
		cancelFunc2()
		time.Sleep(time.Second)
		if 1 == 2 {
			cancelFunc2()
			cancelFunc()
		}
		convey.So(cancelCtx.Value("jiujiujiu"), convey.ShouldEqual, nil)
		convey.So(deadlineCtx.Value("hahakey"), convey.ShouldEqual, "hahavalue")
	})
}

func TestEmpty(t *testing.T) {

}
