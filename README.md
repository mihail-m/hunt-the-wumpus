# hunt-the-wumpus

Run with a json config file:
```
go run *.go -jsonConfig=<path to json configuration file>
```
Or run by passing command line arguments (all are optional):
```
go run *.go -wHP=<wombat healt> -arrows=<starting arrows> -bpChance=<chanse for a pit>  -arrowChance=<chance for a dropped arrow> -batsChance=<chance for bats> -rows=<rows of the map> -cols=<columns of the map>
```