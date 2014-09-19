
###PULL
&emsp;&emsp;由于还不是很清楚Github风格的markdown对应的效果（其实是完全不记得markdown啦），所以在github的在线编辑里直接查看修改的效果，修改了之后，有个问题就是怎么把这些内容同步回来我本地。找了下，应该用`git pull`,但具体怎么用呢？看了下`git help pull`弹出的网页，自己先是试了下`git pull master`，发觉报错说master不是一个远程库，于是猜想到了是要用clone时的那个地址，于是最终这样成功了：
```
git pull https://github.com/allanruin/algorithm.git
```