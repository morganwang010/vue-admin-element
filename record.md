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

