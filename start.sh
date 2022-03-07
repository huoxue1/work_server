/bin/bash
echo "正在下载作业管理系统"
wget https://github.com/huoxue1/work_server/releases/download/v1.0.1/work_server_linux_amd64.tar.gz
echo "文件下载完毕，开始解压文件"
tar -xzvf work_server_linux_amd64.tar.gz
chmod -R 777 ./work_server
read -n "请输入程序运行的端口号：" port
read -n "请输入你的token: " token
nohup ./work_server -p $port -t $token >work.log 2&1 & echo $!>pid.pid