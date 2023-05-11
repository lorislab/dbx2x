# dbx2x

Convert excel to db-unit xml

[![License](https://img.shields.io/github/license/lorislab/dbx2x?style=for-the-badge&logo=apache)](https://www.apache.org/licenses/LICENSE-2.0)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/lorislab/dbx2x?logo=github&style=for-the-badge)](https://github.com/lorislab/dbx2x/releases/latest)

## Getting Started

Convert excel to db-unit xml
```shell script
dbx2x convert --file data.xlsx --output data.xml
```

Run it from docker
```shell script
docker run --rm --name dbx2x ghcr.io/lorislab/dbx2x:latest dbx2x convert --help
```

## Dev

### Test release packages
```
goreleaser release --snapshot --clean
```