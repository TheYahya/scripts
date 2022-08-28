# Remove ext files
This script scan a given directory and remove all file with specific file extension.


# Install
```bash
go install ./
```

# Run
The following command will remove all `.log` files in `/path/to/files` and sub directories recursively.
```bash
remove-ext-files /path/to/files .log 
```

