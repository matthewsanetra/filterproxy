# filterproxy
A fast in-browser proxy that can replace any expression/key words from HTML.  
Pretty functional; you can browse old.reddit.com perfectly well, not to mention less comlicated sites.

## Deploying
Everything is set up with docker.  
Deploying is as easy as cloning this repo and running `docker-compose build && docker-compose up`  
The proxy is then deployed at 127.0.0.1:8080

### Known Issues
- XHR Requests don't go through proxy
- Only GET is supported
- HTML forms don't work
- You have to modify source code in a few places to get the proxy to run on your own domain/ip rather than on localhost

### Possible future features
- Solutions for the issues mentioned above
- Allow users to enter their own word that censored expressions get replaced with
- Make tests and setup continuous integration
