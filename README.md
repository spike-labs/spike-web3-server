## spike-web3-server quick start

### Prequisition
#### SDK
1. Go
https://go.dev/
2. Morails
https://moralis.io/
3. bscscan
https://bscscan.com/
4. Redis
https://redis.io/
5. MySQL
https://www.mysql.com/
#### Services
1. Signature
https://github.com/spike-engine/spike-signature-server


### Build and install spike-web3-server

1. Clone the repository
```shell
git clone https://github.com/spike-engine/spike-web3-server.git
cd spike-web3-server/
```
2. Install all the dependencies
```shell
go mod tidy
```
3. Install swagger
```shell
go install github.com/swaggo/swag/cmd/swag
sudo mv $GOPATH/bin/swag /usr/local/bin
swag -v
```
if you want to update swagger doc, please execute :
```shell
swag init
```
4. Make build
```shell
go build -o spike-web3-server ./main.go
```
5. Update Config
```shell
cp config-example.toml config.toml
```
6. Run
```
./spike-web3-server
```

### Register spike as a system service
Startup script
```shell
vim /etc/systemd/system/spike-web3-server.service
```
Specify the path to the binary
```markdown
[Service] 
ExecStart=PATH-TO-SPIKE-WEB3-SERVER/spike-web3-server
Environment=SPIKE_CONFIG_PATH=/etc/spike/config.toml
Restart=always
RestartSec=5 
```
```shell
systemctl daemon-reload
systemctl start spike-web3-server.service
journalctl -u spike-web3-server.service -f
```

### config
By default, spike-web3-server reads configuration from config.toml under the current folder. 
If it is not created, you may copy from config-example.config.
If you need to specifiy the config file, you may set the enviornment as follows:
```
export SPIKE_CONFIG_PATH=~/spike_home/config.toml
```

### swagger
If you run the project successfully, you can visit http://localhost:3000/swagger/index.html.
And you can see some interface information.
