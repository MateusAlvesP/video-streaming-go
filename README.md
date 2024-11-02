In order to run this app, follow the instructions:

**1) Install GO**

Download go from https://go.dev/dl/
After, run this:
```sudo cd ~/Downloads && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz```

For more details on installing Go, access: https://go.dev/doc/install


**2) Install docker**

For details on installing docker, access https://docs.docker.com/


**3)Run RabbitMQ container**

To stream the data from the webcam we will use a RabbitMQ container
```docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management```


**4) Install gocv**

To open the webcam and show images we will use OpenCV implementation for GO
Details on installation here: https://github.com/hybridgroup/gocv


**5) Run the program**

First run ```go run main.go``` to execute the publisher to stream webcam data using RabbitMQ
After run ```go run consumer/consumer.go``` to execute the consumer that will receive the data in RabbitMQ ant show it in a window.
