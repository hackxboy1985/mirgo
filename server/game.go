package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/yenkeia/mirgo/proc/mirtcp"
)

type Game struct {
	Conf Config
	DB   *Database
	Env  *Environ
}

func NewGame(conf Config) *Game {
	g := new(Game)
	g.Conf = conf
	g.DB = NewDB(conf.MirDB, conf.AccountDB)
	g.Env = g.NewEnv()
	return g
}

func (g *Game) ServerStart() {

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("tcp.Acceptor", "server", g.Conf.Addr, queue)

	proc.BindProcessorHandler(p, "mir.server.tcp", g.EventHandler)

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()

}
