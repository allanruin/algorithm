
###楔子
&emsp;&emsp;起先是把几个结构完全照搬C的邻接表版本。
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


###Bellman Ford
&emsp;&emsp;之前是看RIP看到bellman ford的，当时还没看dijkstra，所以糊里糊涂的。现在实现过dijkstra再看bellman ford觉得相当清楚，bellman ford通过每次都全扫描所有的边来更新值，共v-1次全扫，最后第v次是用于检测是否有负权值循环圈的存在，因为V点的无环图，两点之间最多只有v-1条边。另外，go的test就是爽哇，上次dijkstra写好了几个测试用例，这次就直接写测试函数就可以验证在无负边的情况下是否正确了，太方便了~


###DFS深度优先搜索
&emsp;&emsp;实现的时候出过几个问题。
1.忘了由于是递归实现，在函数内部定义的visited无法被各个递归变量共享。后来改成传递给visited给DFS函数。
2.开始实现代码的时候并没有很强烈的从递归的参数要求入手。递归的时候要接收一个参数s，指明从哪个编号的点（vertex）开始继续DFS，但一开始写到递归的地方前没有想到这个，于是就写了一个先遍历vnode链（graph.Arlist）的循环，导致出错了。
3.seq的问题。还是由于递归的关系，各个递归DFS函数无法共享变量，不过考虑到这个每次递归的返回值是添加到后面，就弄了个nseq，再拼接上就好了。


![图的搜索所使用的图](https://github.com/allanruin/algorithm/blob/master/illustrates/Graph_for_search.PNG)


###拓扑排序
测试用的例子是高分P173，类似这样：


![图的拓扑排序所使用的图](https://github.com/allanruin/algorithm/blob/master/illustrates/Graph_for_topologysort.PNG)


&emsp;&emsp;拓扑排序的思想是，循环选择入度为0的点，进行输出，同时把改点的邻接点的入度都减一（因为点去掉了，所有射出的边也就去掉了）。这里我写了个stack来作为容器装载度为0的点。当时遇到的一个问题就实现的时候选择的是出度为零的点了！
```
for i := range vs {
	t := vs[i]
	degree[t] = degree[t] - 1
	if degree[t] == 0 {
		//有新的点入度为0，扔到入度零的队列zeros中
		st.Push(t)
	}
}
```
TopologySort函数里的这个部分出过几个错误。第一个是混淆了用于遍历的i与邻接点数组储存的点序号，直接用的i带入到degree[]里，应该用vs[i].再者是入队的时候，一时没反应过来我都用了stack了，结果还是append到那个初始的zeros里去，我说怎么始终没入队呢。由于挑选入度0的点时顺序不唯一，所以拓扑排序结果不唯一。


###HEAP堆
其实简简单单的东西，还是花了好几个小时才自己做出来。大致原理其实是很快弄明白了，真正的问题是到了heapsort的实现的时候出现的，adjust函数的接口设计不对，`func adjust(seq *[]int, s int) {` ，测试结果不对，后来自己打印出过程才发现被挪移到末尾作为有序部分的居然也在变动，才发觉时因为adjust函数里求min的时候，是根据slice的长度作为上限的！而我们把min值调到末尾，还是在这个上限以内！内牛满面。。。果然还是要学天勤的，在参数里加个上限hi。
