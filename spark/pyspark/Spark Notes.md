# Spark Notes


1. Difference between Spark Session vs Spark Context?

* Spark Session is a unified entry point to access all of spark functionality post Spark 2.0. 
* Prior to Spark 2.0, Spark Context was the main entry point and to access SQL or Hive metastores, developers were 
required to create SQLContext or HiveContext

2. What are the different ways through which we can create Spark Dataframes?


* existing RDD
* loading external data
* using Hive tables

3.


## Spark Tuning

* understand spark internals

* How does a user program get translated into physical execution in
spark?

jobs, stages and tasks

action triggers creating of DAG and execution of job
A job is divided into stages. 
Transformations are grouped together into a stage. The grouping is
decided by the narrow vs wide transformations.
Narrow transformations can be performed independently
on each executor of worker node.
Wide transformation require data from other executors. Hence separated
into a different stage.
Each stage is divided into tasks and tasks are executed by cores
on executors.



 



## Spark Common Issues

1. ```Exception: Python in worker has different version 3.6 than that in driver 3.7, PySpark cannot run with different minor versions.Please check environment variables PYSPARK_PYTHON and PYSPARK_DRIVER_PYTHON are correctly set.```

* Usually caused by different python versions referenced in driver and worker.


