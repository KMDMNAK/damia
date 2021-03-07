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

# usage
## help
### lan
show list of supported languages

## docker
### dev
#### options
- product


```
damia docker dev
```

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
