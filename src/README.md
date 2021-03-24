# cosmosdb
vi ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin
export GOBIN="$HOME/go_projects/bin"
export GOPATH="$HOME/go_projects"
export http_proxy=
export https_proxy=
export ftp_proxy=
no_proxy=hc-eu-west-aws-artifactory.cloud.health.ge.com,.ge.com,*.ge.com,10.

source /etc/profile && source ~/.bash_profile

cd /tmp
curl -LO https://storage.googleapis.com/golang/go1.13.linux-amd64.tar.gz

cd /etc/pki/ca-trust/source/anchors/
wget https://static.gecirtnotification.com/browser_remediation/packages/GE_External_Root_CA_2_1.crt
update-ca-trust

cd /tmp
wget https://github.com/Masterminds/glide/releases/download/v0.13.3/glide-v0.13.3-linux-amd64.tar.gz
tar -zxvf glide-v0.13.3-linux-amd64.tar.gz

tar -C /usr/local -xvzf go1.13.linux-amd64.tar.gz
mkdir -p ~/go_projects/{bin,pkg,src}

cp /tmp/linux-amd64/glide /usr/local/bin/

cd $GOPATH/src
git clone https://github.com/nuthankumar/cosmos_glide.git -b develop

cd $GOPATH/src/cosmos_glide/src
glide up --skip-test
glide install 
go build *.go
go run *.go



glide


cd /etc/pki/ca-trust/source/anchors/
wget https://static.gecirtnotification.com/browser_remediation/packages/GE_External_Root_CA_2_1.crt

cd /tmp
wget https://github.com/Masterminds/glide/releases/download/v0.13.3/glide-v0.13.3-linux-amd64.tar.gz
tar -zxvf glide-v0.13.3-linux-amd64.tar.gz

mkdir -p /root/glide/{bin,pkg,src}
cd src/
git clone https://github.com/fabianlee/go-vendortest1.git
cd go-vendortest1/vendortest
glide init --non-interactive
glide up
glide install
go build