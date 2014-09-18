    起先是把几个结构完全照搬C的邻接表版本。
之后发觉完全没必要自己实现链表，go有container/list 可以直接用。

再后来查找list的用法资料的时候，看到有人在SO问，为什么LIST使用并不太频繁，然后有人说一般你要用list的地方都可以通过方便强大的slice来实现。

于是vnode的 arlist改成了用slice，后面实际要用的时候才发觉初始化时个大麻烦。因为用之前要用make初始化，make([]int,0)的话，后面index为1的时候就报错说index out of range了，囧，我以为slice是很灵活自动扩张的。。。

不知道还有没有vector，C++的vector就挺适合来替代邻接表里的arc（边）链表的。因为vector可以push数据到末尾。

看来slice跟vector还是有区别的。
slice如果是make([]int,10) 那么默认初始化为0，如果再append的话，是形成11个元素的slice，而不是像vector那样只在实际自己塞入的元素后面。语义上有区别。话说vector的一个实现方式是不是就是list，所以才能表现出list的特性。。


golang的一大问题就是调试不方便，对于稍微复杂点的逻辑（算法）来说，没IDE环境调试确实不太方便。VS2013 EXPRESS 能自由下断点查看非常方便。

GO版本的更多时间是花在查找go的语法和包的用法，比如类型后置，指针中的*前置，比如数组是[]int,[]*int， 初始化方式，struct的用法，slice的初始化，map的初始化，range的使用，

最后又改回用list.New() list.List实现。
list的member function 中Front返回的是list这个包里定义的Element的指针，里面有个Value interface{} 的域来存放实际自己给进去的值。但是取出来要按照原来的类型使用还得cast回去
e.(int) 这样。



吃完饭继续写，9点左右（ 不到）终于调对了。


最终合并了源点的计算和其他
// 无需特殊对待源点直接相连的点，源点的距离被设置为0，则在挑选最短距离的x
// 时自然首先挑选的就是源点，进而计算并更新其他距离


================Bellman Ford===================
之前是看RIP看到bellman ford的，当时还没看dijkstra，所以糊里糊涂的。现在实现过dijkstra再看bellman ford觉得相当清楚，bellman ford通过每次都全扫描所有的边来更新值，共v-1次全扫，最后第v次是用于检测是否有负权值循环圈的存在，因为V点的无环图，两点之间最多只有v-1条边。另外，go的test就是爽哇，上次dijkstra写好了几个测试用例，这次就直接写测试函数就可以验证在无负边的情况下是否正确了，太方便了~
