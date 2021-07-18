1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 10 -t get,set    
SET: 44843.05 requests per second, p50=0.567 msec         
GET: 45662.10 requests per second, p50=0.559 msec                  

[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 20 -t get,set  
SET: 42016.80 requests per second, p50=0.607 msec         
GET: 41666.67 requests per second, p50=0.607 msec                  

[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 50 -t get,set  
SET: 45871.56 requests per second, p50=0.551 msec         
GET: 46948.36 requests per second, p50=0.543 msec                  

[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 100 -t get,set  
SET: 44444.45 requests per second, p50=0.575 msec         
GET: 45454.55 requests per second, p50=0.559 msec                  

[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 200 -t get,set   
SET: 45454.55 requests per second, p50=0.559 msec         
GET: 46082.95 requests per second, p50=0.551 msec                  

[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 1000 -t get,set   
SET: 38167.94 requests per second, p50=0.663 msec                   
GET: 37878.79 requests per second, p50=0.663 msec                   

[root@VM-0-5-centos src]# ./redis-benchmark -n 10000 -q -d 5000 -t get,set    
SET: 38910.51 requests per second, p50=0.647 msec                   
GET: 38759.69 requests per second, p50=0.647 msec 

