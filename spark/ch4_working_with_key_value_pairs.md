 # Key Value pairs Use Cases

* count reviews by product
* group together data with same key
* group together two different RDDs

* Spark provides special operations on RDDs containing key value pairs. These RDDs are called Pair RDDs.

## Creating Pair RDDs

```
inputRDD = sc.textFile("README.md")
pairs = inputRDD.map(lambda line: line.split(" ")[0], line)
pairs.collect()
```

## Transformation on Pair RDDs


