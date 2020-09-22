### 简介
一个过滤器，可以在使用 Git 提交代码时通过注释忽略指定行或用任意内容替换指定行

### 使用

#### 1. build  
```
$ git clone https://github.com/Limgmk/git-filter
$ cd git-filter
$ go build
$ sudo cp git-filter /usr/local/bin/git-filter
```

#### 2. 测试  
示例文件 ```test.txt``` 内容:
```
第1行
#GITIGNORE<<<
第2行
第3行
#GITIGNORE>>>
第4行
第5行
第6行 #GITIGNORE
第7行
#GITREPLACE with 第八行
第8行


第11行
//GITIGNORE<<<
第12行
第13行
//GITIGNORE>>>
第14行
第15行
第16行 //GITIGNORE
第17行
//GITREPLACE with 第十八行
第18行
```
通过管道传入，支持多个注释符参数
```
$ cat test.txt | git-filter "#" "//"
```
过滤结果
```
第1行
第4行
第5行
第7行
第八行


第11行
第14行
第15行
第17行
第十八行
```

#### 3. 配合 Git 使用
以 yaml 文件为例  
添加过滤器规则:
```
$ git config --global filter.gitignore-yaml.clean "/usr/local/bin/git-filter '#' '//'"
```
在项目目录下新建并编辑 ```.gitattributes``` 文件，加入以下内容:
```
*.yaml filter=gitignore-yaml
*.yml filter=gitignore-yaml
```