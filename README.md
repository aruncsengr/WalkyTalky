# WalkyTalky

WalkyTalky on the Go!
A chat application built using the Go(lang)

## Go setup on Ubuntu
Follow the guideline steps to install GO on your local machine

```
sudo curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
sudo tar -xvf go1.6.linux-amd64.tar.gz
sudo mv go $HOME
```

Create a GO workspace
```
mkdir -p $HOME/Development/GoWorkspace/src
```

Set GOROOT and GOPATH
```
sudo nano ~/.profile
```

Copy and paste the following commands after opening ~/.profile file
```
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/Development/GoWorkspace
```

Read and execute commands from ~/.profile using source command in terminal
```
source ~/.profile
```


## App setup

Follow the guideline steps to run this app on your local machine
```
cd $HOME/Development/GoWorkspace/src
git clone https://github.com/aruncsengr/WalkyTalky.git
cd WalkyTalky
```

Setup HOST connection

```
go run main.go -listen <ip>
```

Setup GUEST connection

```
go run main.go <ip>
```

Cool, Now you are up and running the WalkyTalky app :sparkles: :wink:



To create binary file and setup connections
```
go build
```

Setup HOST connection
```
./WalkyTalky -listen <ip>
```

Setup GUEST connection
```
./WalkyTalky <ip>
```

To create binary file compatible to WINDOWS
```
GOOS=windows GOARCH=amd64 go build -o WalkyTalky.exe
```

To check file type
```
file WalkyTalky
```
