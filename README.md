# acm-runner
listen aliyun acm change

**step1** 

copy yaml file

```
$ cp ./acm-runner.yaml.example ./acm-runner.yaml
```

**step2** 

defind your acm info

```
namespace:
  id: your namespace id
  end_point: acm.aliyun.com:8080
  access_key: your accessKey
  secret_key: your secretKey

list:
  - data_id: your data id
    group: your groupon
    filename: ./out/{data_id}.json


```

**step3**

run

```
$ ./acm-runner
```

