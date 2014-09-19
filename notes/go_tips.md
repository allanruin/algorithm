
1. 数组指针
指向数组的parray指针要想再次进行下标取值操作，既`[3]`这样，得先解引用，转回数组才行。比如dfs里的这段
```go
if (*visited)[s] != true {
	(*visited)[s] = true
	
	...

	}
}
```
visited的类型是`*[]int`,要先解引用为数组，才能用index取值，而且得加括号。
