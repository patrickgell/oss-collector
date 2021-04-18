# oss-collector

This tool collects all the sources from the configured open source projects and produces a big zip file containing all sources.

### Configuration file
The structure of the JSON file:
```json
{
  "project": "oss-test",
  "components": [
    {
      "type": "JAVA",
      "groupID": "org.apache.commons",
      "artifactID": "commons-lang3",
      "version": "3.12.0",
      "svmID": 123,
      "sources": "https://github.com/apache/commons-lang/archive/refs/tags/rel/commons-lang-3.12.0.zip"
    } 
  ]
}
```
| Key | Description
| --- | --- 
| project | Name of the project
| components |List of open source components
| type | language of the component e.g. JAVA
| groupID | Maven groupId 
| artifactID | Maven artifactId
| version | version of the component
| sources | URL to original component sources 
| svmID | optional ID

### Build

```shell
go build cmd/collector/collector.go
```