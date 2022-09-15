## spike-frame-server 

### Build and install spike-frame-server
1. Clone the repository
```shell
git clone https://github.com/spike-engine/spike-frame-server.git
cd spike-frame-server/
```
2. Install all the dependencies
```shell
go mod tidy
```
3. Install swagger
```shell
go get -u github.com/swaggo/swag/cmd/swag
swag -v
sudo mv $GOPATH/bin/swag /usr/local/bin
```
if you want to update swagger doc, please execute :
```shell
swag init
```
4. Make build
```shell
go build -o spike-frame-server ./main.go
```
5. Startup script
```shell
vim /etc/systemd/system/spike-frame-server.service
```
Specify the path to the binary
```markdown
[Service] 
ExecStart=/root/go/src/github.com/spike-engine/spike-frame-server/spike-frame-server
Environment=CONFIG_PATH=/etc/config.toml
Restart=always
RestartSec=5 
```
```shell
systemctl daemon-reload
systemctl start spike-frame-server.service
journalctl -u chain-server.service -f
```
Of course, you can click the build icon in your IDE to run the project instead of startup script.
But, we recommend using system script in mainnet.

### config
If you don't specify the path to the configuration file in the environment variable in the startup script, 
config.toml is the default.And config-example.toml is a demo.

You should configure some information about system port, mysql , redis, contract address etc.

### swagger
If you run the project successfully, you can visit http://localhost:3000/swagger/index.html.
And you can see some interface information.
