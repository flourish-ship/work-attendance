package socket

import (
	"fmt"
	"github.com/flourish-ship/skeleton/debuger"
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/kataras/iris"
)

var (
	iris_instance *iris.Framework
	chartRoom1    = "room1"
)

type (
	clientPage struct {
		Title string
		Host  string
	}

	ChartRoomApi struct {
	}
)

func (chart *ChartRoomApi) Chart(frame *iris.Framework) {
	iris_instance = frame
	chartApi := frame.Party(consts.URL_PRFIX + "/chart")
	{
		chartApi.Get("", ChartConnect)
	}
}

func ChartConnect(c *iris.Context) {

	iris_instance.Config.Websocket.Endpoint = "/chart/room"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		debuger.Display("您已经进入房间", chartRoom1)
		c.Join(chartRoom1)

		c.On("chat", func(message string) {
			c.To(chartRoom1).Emit("chat", "From"+c.ID()+": "+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID :%s has been disconnected!", c.ID())
		})
	})
}
