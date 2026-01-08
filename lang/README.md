# go
wget https://golang.google.cn/dl/go1.25.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

docker run --rm -v $GOPATH:/go -v $PWD:/work -w /work -e GOPROXY=$GOPROXY $GOIMAGE go build  -trimpath -o /work/build/$output 