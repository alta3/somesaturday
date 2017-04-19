# The somesaturday server golang project

Somesaturday is a frameshop where a FRAMER on RED TOP road does museum quality framing

1. Be in the home directory  
   `cd`

2. Set up directories  
  `mkdir git`  
  `mkdir git/go`
  `mkdir go`
  `mkdir go/src`

3. Install go  
   `wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz`  
   `sudo tar -C /usr/local -xzf  go1.8.1.linux-amd64.tar.gz`  
   `export PATH=$PATH:/usr/local/go/bin`  
   
4. clone somesaturday  
   `cd ~/git/go`  
   `git clone git@github.com:alta3/somesaturday.git`  
   `cd ~/git/go/somesaturday`  

5.  get viper  
   `go get github.com/spf13/viper`  

6. Familiarize yourself with the config file, it will issue a git pull example  
   `cat deploy/config.yaml`  
   
7. start up screen to run somesaturday.go  
    `screen`  
    
8. compile somesaturday.go  
   `go build somesaturday.go`  
   
You should see the web page now http://somesaturday.com
