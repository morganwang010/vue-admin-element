从第三方式复制过来的vue，样式有问题，需要通过pretttier进行下格式化，命令如下：
npx prettier --write "src/**/*.js" "src/**/*.vue"
然后再打开就不会有格式问题了。

git config pull.rebase false
git pull origin main
