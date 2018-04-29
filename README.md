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

查看当前系统host分组列表
	
	./main listGrp

	eg:
	./main listGrp

添加host分组列表中某一个分组到系统hosts中
	
	./main enableGrp [grpName]

	eg:
	./main enableGrp ubuntu
