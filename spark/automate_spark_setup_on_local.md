# How to automate spark setup on local mac?

## Download latest spark

```
mkdir -p $HOME/software && \
cd $HOME/software && \
curl -O https://www.apache.org/dyn/closer.lua/spark/spark-2.4.6/spark-2.4.6-bin-hadoop2.6.tgz && \
tar -xzvf spark-2.4.6-bin-hadoop2.6.tgz
```

## Setup Spark home as an environment

```
echo "export SPARK_HOME=$HOME/software/spark-2.4.6-bin-hadoop2.6" >> ~/.bash_profile
source ~/.bash_profile
```

## Add below functions to .bash_profile

```
spark_up() {
  export HOSTNAME=$(hostname)
  sh $SPARK_HOME/sbin/start-master.sh
  sh $SPARK_HOME/sbin/start-slave.sh spark://${HOSTNAME}:7077
  sh $SPARK_HOME/sbin/start-history-server.sh
}

spark_down() {
  sh $SPARK_HOME/sbin/stop-master.sh
  sh $SPARK_HOME/sbin/stop-slave.sh
  sh $SPARK_HOME/sbin/stop-history-server.sh
}
```

## Adding .bash_profile to .zshrc (optional)

In case you have .zshrc 

```
if [ -f ~/.bash_profile ]; then
    source ~/.bash_profile;
fi
```
## source .zshrc

```
source ~/.zshrc
```

## Usage

```
spark_up # will bring up spark master, slave and history server
spark_down # will bring spark_master, slave and history server
```

