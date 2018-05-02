# go-hosts
switch hosts in go

## 功能
追加单条hosts记录

    append [ip] [domain]

根据域名删除host记录

    del [host]

查看当前系统host列表
	
    list

查看当前系统host分组列表
	
    listGrp

使用host分组列表中的某一个分组
	
    enableGrp [grpName]

取消使用host分组列表中的某一个分组
	
    disableGrp [grpName]

## 使用

    make
    make install

    eg:
	append 127.0.1.1 ubuntu

	del ubuntu

	enableGrp ubuntu

	disableGrp ubuntu
