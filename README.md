# Get Fake Faces
This tool will query and download images from: https://this-person-does-not-exist.com/en

# Usage
| Param | Description                           |
| ----- | ------------------------------------- |
| -n    | Number of requests to make            |
| -t    | Time between requests in milliseconds | 

Example:<br/>
```
go run ./makesomefaces.go -n 50 -t 1000
```

# Duplicate Images
 Some duplicate images will be generated. This can be solved by increasing the time between requests, but the number of duplicate images are low enough that I didn't want to bother doing so. 
 
 Instead, Use the following tool to remove duplicate images:
 https://github.com/cw-somil/Duplicate-Remover
