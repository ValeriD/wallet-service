wget https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz
tar -zxvf go1.13.6.linux-amd64.tar.gz -C /usr/local
rm go1.13.6.linux-amd64.tar.gz -f
echo 'export GOROOT=/usr/local/go' | tee -a /etc/profile
echo 'export PATH=$PATH:/usr/local/go/bin' | tee -a /etc/profile
source /etc/profile
go version
