package signal

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

//https://www.jianshu.com/p/ae72ad58ecb6

//POSIX.1-1990
//信号	    值	动作	    说明
//SIGHUP	1	Term	终端控制进程结束(终端连接断开)
//SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
//SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
//SIGILL	4	Core	非法指令(程序错误、试图执行数据段、栈溢出等)
//SIGABRT	6	Core	调用abort函数触发
//SIGFPE	8	Core	算术运行错误(浮点运算错误、除数为零等)
//SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)
//SIGSEGV	11	Core	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
//SIGPIPE	13	Term	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
//SIGALRM	14	Term	时钟定时信号
//SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
//SIGUSR1	30,10,16	Term	用户保留
//SIGUSR2	31,12,17	Term	用户保留
//SIGCHLD	20,17,18	Ign	子进程结束(由父进程接收)
//SIGCONT	19,18,25	Cont	继续执行已经停止的进程(不能被阻塞)
//SIGSTOP	17,19,23	Stop	停止进程(不能被捕获、阻塞或忽略)
//SIGTSTP	18,20,24	Stop	停止进程(可以被捕获、阻塞或忽略)
//SIGTTIN	21,21,26	Stop	后台程序从终端中读取数据时触发
//SIGTTOU	22,22,27	Stop	后台程序向终端中写数据时触发

// 在SUSv2和POSIX.1-2001标准中的信号列表:
//信号	    值			动作		说明
//SIGTRAP	5			Core	Trap指令触发(如断点，在调试器中使用)
//SIGBUS	0,7,10		Core	非法地址(内存地址对齐错误)
//SIGPOLL				Term	Pollable event (Sys V). Synonym for SIGIO
//SIGPROF	27,27,29	Term	性能时钟信号(包含系统调用时间和进程占用CPU的时间)
//SIGSYS	12,31,12	Core	无效的系统调用(SVr4)
//SIGURG	16,23,21	Ign		有紧急数据到达Socket(4.2BSD)
//SIGVTALRM	26,26,28	Term	虚拟时钟信号(进程占用CPU的时间)(4.2BSD)
//SIGXCPU	24,24,30	Core	超过CPU时间资源限制(4.2BSD)
//SIGXFSZ	25,25,31	Core	超过文件大小资源限制(4.2BSD)

//kill pid的作用是向进程号为pid的进程发送SIGTERM（这是kill默认发送的信号），该信号是一个结束进程的信号且可以被应用程序捕获。
// 若应用程序没有捕获并响应该信号的逻辑代码，则该信号的默认动作是kill掉进程。这是终止指定进程的推荐做法。
//kill -9 pid则是向进程号为pid的进程发送SIGKILL（该信号的编号为9），从本文上面的说明可知，SIGKILL既不能被应用程序捕获，也不能被阻塞或忽略，其动作是立即结束指定进程。
// 通俗地说，应用程序根本无法“感知”SIGKILL信号，它在完全无准备的情况下，就被收到SIGKILL信号的操作系统给干掉了，显然，在这种“暴力”情况下，应用程序完全没有释放当前占用资源的机会。
// 事实上，SIGKILL信号是直接发给init进程的，它收到该信号后，负责终止pid指定的进程。在某些情况下（如进程已经hang死，无法响应正常信号），就可以使用kill -9来结束进程。
//若通过kill结束的进程是一个创建过子进程的父进程，则其子进程就会成为孤儿进程（Orphan Process），
// 这种情况下，子进程的退出状态就不能再被应用进程捕获（因为作为父进程的应用程序已经不存在了），不过应该不会对整个linux系统产生什么不利影响。

//Linux Server端的应用程序经常会长时间运行，在运行过程中，可能申请了很多系统资源，也可能保存了很多状态，
// 在这些场景下，我们希望进程在退出前，可以释放资源或将当前状态dump到磁盘上或打印一些重要的日志，也就是希望进程优雅退出（exit gracefully）。
//从上面的介绍不难看出，优雅退出可以通过捕获SIGTERM来实现。具体来讲，通常只需要两步动作：
//1）注册SIGTERM信号的处理函数并在处理函数中做一些进程退出的准备。信号处理函数的注册可以通过signal()或sigaction()来实现，
// 其中，推荐使用后者来实现信号响应函数的设置。信号处理函数的逻辑越简单越好，通常的做法是在该函数中设置一个bool型的flag变量以表明进程收到了SIGTERM信号，准备退出。
//2）在主进程的main()中，通过类似于while(!bQuit)的逻辑来检测那个flag变量，一旦bQuit在signal handler function中被置为true，
// 则主进程退出while()循环，接下来就是一些释放资源或dump进程当前状态或记录日志的动作，完成这些后，主进程退出。

func ExampleNotifyAll() {
	c := make(chan os.Signal)
	signal.Notify(c)

	for cs := range c {
		switch cs {
		case os.Interrupt:
			os.Exit(0)
		default:
			fmt.Println("default")
		}
	}

	// output:
}

func ExampleNotify() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	for cs := range c {
		switch cs {
		case os.Interrupt:
			return
		default:
			fmt.Println("default")
		}
	}
	bufio.NewReader(os.Stdin).ReadLine()
	// output:
}
