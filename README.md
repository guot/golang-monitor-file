# golang-monitor-file「文件监控器」
ver 1.0 
> golang-monitor-file 是用go语言开的监控文件是否变化，如果变化自动执行 所设定的命令，可以是
系统命令或shell脚本等。

## 示例
可在Ubuntu14下直接运行已经编译好的可执行文件。也可自行下载源码进行编译运行。
> git clone https://github.com/guot/golang-monitor-file.git<br/>
$build<br/>
./golang-monitor-file --name RREADME.md --cmd   ./family_start.sh

### 参数说明
name: 需要监测的文件名
cmd :如果监测文件发行改动后，需要执行的程序脚本。
	
