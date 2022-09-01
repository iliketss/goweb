# goweb
go web learning

#安装grpc
go get google.golang.org/grpc
go get google.golang.org/protobuf


#个人使用
1. protoc --go_out=../services/ prod.proto 
2. protoc --go-grpc_out=.  prod.proto


#认证
TLS（安全传输曾），TLS是建立在tcp协议上的协议服务于应用层，它的前身是SSL
####1.保密性，保密通过encryption实现
####2.完整性，MAC校验机制，一旦被篡改，立刻被发现
####3.认证，双向认证
######   
    1.openssl genrsa -des3 -out ca.key 2048
    2.openssl req -new -key ca.key -out ca.csr
    3.openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt
    4.openssl genpkey -algorithm RSA -out server.key
    5.openssl req -new -nodes -key server.key -out server.csr -days 3650 -config ./openssl.cnf -extensions v3_req
    6.openssl x509 -req -days 365 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAreateserial -extfile ./openssl.cnf -extensions v3_req
#####修改openssl.cnf文件
        copy_extensions = copy #去掉注释符号
        req_extensions = v3_req #去掉注释符号
        找到[ v3_req ]添加subjectAltName=@alt_names
        添加[alt_names]，DNS.1= *.mszlu.com
####生成服务端公钥
    openssl genrsa -out server.key 1024
    根据以上生成的server.key生成一个私钥
    openssl req -x509 -key server.key -out server.pem