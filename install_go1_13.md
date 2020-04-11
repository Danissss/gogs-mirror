wget https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz
sudo tar -zxvf go1.13.6.linux-amd64.tar.gz -C /usr/local
rm go1.13.6.linux-amd64.tar.gz -f



echo 'export GOROOT=/usr/local/go' | sudo tee -a /etc/profile
echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile
source /etc/profile



setup gogs mysql db

mysql> create database gogs
    -> ;
Query OK, 1 row affected (0.01 sec)

mysql> create user 'gogs'@'localhost' identified by '858688258';
Query OK, 0 rows affected (0.01 sec)

mysql> grant all privileges on *gogs.* to 'gogs'@'localhost';
ERROR 1046 (3D000): No database selected
mysql> grant all privileges on gogs.* to 'gogs'@'localhost';
Query OK, 0 rows affected (0.02 sec)

mysql> flush privileges;
Query OK, 0 rows affected (0.01 sec)

mysql> exit
Bye

set GOPATH
put
```
export GOPATH="$HOME/go"
PATH="$PATH:$GOPATH/bin"
```
to `~/.bashrc`


install go-bindata
if you made change in tmpl, you need to run `make build` to regenerate the `templates_gen.go` and `public_gen.go` file 


Make change of templates:
`go generate internal/assets/templates/templates.go`

`go build`
