# Word Count PySpark Example

import os
from pyspark import SparkContext

sc = SparkContext()

lines = sc.textFile(os.environ["SPARK_HOME"]+"/README.md")

readme_line_count=lines.count()

print("Total number of lines in README: ",readme_line_count)

python_lines = lines.filter(lambda line: "Python" in line)

python_lines_in_readme = python_lines.count()

print("Total number of python lines: ",python_lines_in_readme)

sc.stop()
