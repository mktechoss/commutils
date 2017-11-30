说明: proto作为内部的公共协议

go 中proto文件的使用命令:

    1. 进入proto目录,输入命令: protoc --go_out=. *.proto
    2. 输入命令: go build; go install 