# filterproxy
A fast in-browser proxy that can replace any expression/key words from HTML.

~~I initially won't be working on adding CSS & JS support, but once I am able to modify links from HTML source it will be trivial~~

CSS and Images work, now we can think about injecting javascript into page source to hook AJAX calls and replace the host with proxy


### TODO
- [ ] Nginx/caddy (Only passthrough /proxy)
- [ ] Allow users to enter their own word that censored items get replaced with
- [ ] Containerize