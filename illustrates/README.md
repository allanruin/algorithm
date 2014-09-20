
大部分黑白无水印的图都是我用DOT写出，并通过在线工具http://stamm-wilbrandt.de/GraphvizFiddle/ 生成。比如我生成拓扑排序时的图，


![图的拓扑排序所使用的图](https://github.com/allanruin/algorithm/blob/master/illustrates/Graph_for_topologysort.PNG)


用的就是如下的DOT代码
```
digraph {
	0->1;
	0->2;
	0->3;
	1->2;
	1->4;
	2->4;
	2->5;
	3->5;
	5->4;
	4->6;
	5->6;
}

```

将`digraph`改成`graph`既生成无向图。

加上`[label=xx]`就可以得到边上的标示了，比如
```
digraph {
	0->1 [label=11];
	0->2 [label=21];
	0->3 [label=31];
	1->2 [label=41];
	1->4 [label=51];
	2->4 [label=61];
	2->5 [label=71];
	3->5 [label=81];
	5->4 [label=91];
	4->6 [label=15];
	5->6 [label=15];
}
```


![带权值的图](https://github.com/allanruin/algorithm/blob/master/illustrates/Graph_with_edge_label.PNG)