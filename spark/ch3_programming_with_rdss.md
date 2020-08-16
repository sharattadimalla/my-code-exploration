 * RDDs - Resilient Distributed Dataset

* RDDs are core fundamental abstraction in Spark

* RDDs are a collection of items in memory whch are available in computation

* RDDs offer two types of operations
   - Transformations
   - Actions

* Transformations construct a new RDD from a previous one.

* Actions compute a result based on an RDD and either return the result to the driver program or save it to an external storage system.

* Transformations and Actions are computed differently in Spark

* Spark computes only when an action is called.

* Spark RDD by default are recomputed each time an action is executed.

* RDDs are created either by parallelizing a collection or loading from an external dataset.

* Transformation RDDs are computed lazily i.e. only when an action is called upon them.

* Actions force the evaluation of the transformations required for the RDD they were called on, since they actually need to produce
output

* Lazy Evaluation - Spark does not evaluate transformations until an action is called upon it. This behavior is called lazy evaluation.
RDD 

* Passing Functions to Spark - One issue with passing functions in spark is serializing the object containing the function. When you pass a 
function that is a member of an object, Spark sends the entire object to worker nodes which may be large the bit of information you need.
Sol:- Extract the fields you need from the object into local variable and pass it in.

## Common Transformations

* map 

```
sc.parallelize(range(1,11)).map(lambda num: num * num).collect()
```

* flatMap

```
sc.parallelize(["hello world","john doe"]).flatMap(lambda line: line.split(" ")).collect()
```

* pseduo set operations
```
rdd1 = sc.parallelize(["a","b","c","d","c"])
rdd2 = sc.parallelize(["c","d","e","f"])
```

* distinct
```
rdd1.distinct().collect()

```

* union
```
rdd2.union(rdd1).collect()
```

* intersection
```
rdd1.intersection(rdd2).collect()
```

* subtract
```
rdd1.subtract(rdd2).collect()
```

* cartesian
```
rdd1.cartesian(rdd2).collect()
```

* sample
```
rdd1.sample(false, 0.5).collect()
```

## Common Actions

* reduce
```
sc.parallelize(range(1,11)).reduce(lambda a,b: a+b)
```

* collect
```
sc.parallelize(range(1,6)).collect()
```

* count
```
sc.parallelize(range(1,6)).count()
```

* countByValue
```
sc.parallelize([1,1,1,1,2,4,3,2,3,4,4]).countByValue()
```

* take
```
sc.parallelize(range(1,6)).take(2)
```

* top
```
sc.parallelize(range(1,11)).top(2)
```

* takeOrdered
```
sc.parallelize(range(1,11)).takeOrdered(2)
```

* takeSample
```
sc.parallelize(range(1,11)).takeSample(False, 2, 999)
```

## Persistence

For iterative algorithms, it may be expensive to recompute an RDD everytime an action is called on it. To avoid re-computing an RDD multiple times, we can ask Spark to persist the data.

```
sampleRDD = sc.parallelize(range(1,11)).persist()
sampleRDD.is_cached
sampleRDD.unpersist()
```


