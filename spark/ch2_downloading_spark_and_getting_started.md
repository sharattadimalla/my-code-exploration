
* Spark is written in Scala and run on Java Virtual Machine (JVM)

* Java8 is required to run Spark on local

* download the spark-2.4.6-bin-hadoop2.7.tgz from spark website

### Setup Spark locally

```
tar -xf path_to_download/spark-2.4.6-bin-hadoop2.7.tgz
cd spark-2.4.6-bin-hadoop2.7
echo "export SPARK_HOME=path_to_download/spark-2.4.6-bin-hadoop2.7" >> ~/.bash_profile
source ~/.bash_profile
``` 

### Scala Spark Shell
```
spark-shell
```

### PySpark Shell
```
pyspark
```

* spark-shell is a driver program. 

* In spark driver program access spark via SparkContext object.

* SparkContext is a connection to the computing cluster.

* To run operations, driver program manages a number of nodes called executors.
