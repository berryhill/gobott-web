# Gobott API

### Greetings to Canna-bot

This guide is here to help you start calling the Canna-bot API that is the deployed


## API Info
### ip
```
http://138.68.5.215
```

### endpoints
```
/reports
```
This endpoint currently returns all reports saved in the database
```
/resume_report
```
This endpoint currently tells the machine to resume outbound reports (granted there's a gateway online)
```
/halt_report
```
This endpoint currently tells the machine to halt outbound reports (granted there's a gateway online)


## API Development Environment Setup Guide
### Download and Install Go 1.6
Follow the instructions for the relevant operating system [here](https://golang.org/dl/)

### Configure Environmental Variables
```
something here
```

### Clone Project
```
cd $GOPATH/src/github.com/
git clone git@github.com:berryhill/gobott-web.git
```

### Install Dependencies
```
cd $GOPATH/src/github.com/canna-bot
go get
```

## Building and Running Canna-bot
### Build Project
```
cd $GOPATH/src/github.com/canna-bot
go build main.go
```

### Run Project
```
cd $GOPATH/src/github.com/canna-bot
go run main.go
```
