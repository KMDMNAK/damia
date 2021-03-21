# Damia

Damia is a code generator for development environment.
You can setupt dev configs like Dockerfile, .vscode(settings.json,extensions.json), webpack.config.js etc...
Damia constructs lint and formatter based on Docker.

# requirement 
- Docker
- docker-compose

# install

```
pip install damia
```

```
npm i -g damia
```

# commands

## help
### lan
show list of supported languages


___


## docker
### dev
#### options
- product


```
damia docker dev
```



___


## template
### generate
### add


```
damia generate -t python
damia generate -t vscode
damia generate -t python vscode
```

#### options
- root(r) : root directory
- tag(t) 

```
damia generate --t python vscode --root .
```


___


## vscode

### open    
arg : directory path, service name
**Options**
- --type (dir;default, service,container)
- --path **work only for type service**
- --work_dir **work only for type service**