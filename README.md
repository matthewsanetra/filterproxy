# filterproxy
A fast in-browser proxy that can replace any expression/key words from HTML.

Currently pretty functional, but work on patching javascript AJAX/XHR requests needs to be worked on.
Only supports GET requests, and form submissions don't work.


### TODO
- [x] Nginx ~~/caddy~~ (Only passthrough /proxy)
- [ ] Allow users to enter their own word that censored items get replaced with
- [x] Containerize
- [ ] More than one method (currently only GET is supported)
- [ ] Patch javascript XHR requests
- [ ] Modify forms to use the proxy and work