# Multiverse OS Configuration & Runtime Files
This repository should be downloaded on the host as part of the installation process for now until an alpha installer is released. 

And all the files should be installed into the two key folders:


The core multiverse OS files will be located in:

Run-time files
```
/var/multiverse
```

and

System wide configuration files
```
/etc/multiverse
``` 

and per user configuration will be located: 

```
/home/user/.local/share/multiverse
```
#### Portal Gun 

Right now portalgun files are located inside the multiverse folders defined
above. 


#### Scramble Suite 
The scramble suite files been moved in the last few versions and are now
located: 

```
/home/user/.user/
```

This folder has fairly complex structure, but the key concept is that this one
folder will provide a location for all your personal settings; which will be
installed via symbolic linking from this folder making it a central location for
changing files, it will concentrate everything making it easier to backup, and
providing a ideally more logical and consistent way to access this data. It will
also be encrypted at rest for now, and eventually always remain partially
encrypted, only decrypting specific values that are actively being used as the
settings files become services that provide a file-like UI to change the values. 

This will allow us to experiment and extend the existing system while remaining
completely backwards compatible. 


