### 0818
- 表格设置列宽，直接设置width: '220px'即可
### 0812
- 使用vite时，要在.env环境中定义全局变量，必须使用VITE_开头定义，才能在vue中使用import.meta.env.VITE_*******引入环境变量 

### 0806 
- 进行合同添加，实现输入时实时搜索

### 0805
- 添加左上角notify
 
### 0804
- 添加后台url资源的api

### 0803
- 添加后台url资源的api

### 0802
- 添加瀑布流
- 添加cardlist页面
### 0801
- 封装一个message组件
- 修复上传功能，能传到服务器上去了，但是如何返回并刷新界面，还有待改进。

### 0731
- 使用gitemoji ,先安装 npm i -g gitmoji-cli

### 0730
- 添加product页面
- 完成表格编辑功能
- 使表格内可以显示缩略图
- 需要上传图片，待完成
 
### 0729
- 修復分頁功能，將分頁總數的數據在table頁麪獲取
- 使用dayjs來formatter表格中的日期格式
- 添加customer管理頁面
- 配置权限，目前还有问题，暂时全放开



### 0728
- 开始调contract接口，在vue中拼接参数
- 为了让github上的contribution生效，修改本地邮箱

从第三方式复制过来的vue，样式有问题，需要通过prettier进行下格式化，命令如下：
npx prettier --write "src/**/*.js" "src/**/*.vue"
还有一种方式是在vscode里安装prettier插件，直接使用alt+shift+F即可全部格式化

然后再打开就不会有格式问题了。

git config pull.rebase false
git pull origin main

