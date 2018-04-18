# go-hosts
switch hosts in go

## 功能
追加单条hosts记录

        sudo ./main append [ip] [domain]

        eg:
        sudo ./main append 127.0.0.1 www.baidu.com

根据域名删除host记录

	sudo ./main del [host]

	eg:
	sudo ./main del localhost

查看当前系统host列表
	
	./main list

	eg:
	./main list
