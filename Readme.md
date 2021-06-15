#### 公共参数

##### 简要描述

- 公共参数说明

##### 请求URL
- ` http://39.108.110.77/DangoMark/api `

##### 请求方式
- POST

##### 返回示例

``` 
{
    "Bearer": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImV4cCI6MTYxODg1NTYyMSwianRpIjoiMSIsImlhdCI6MTYxODc2OTIyMSwiaXNzIjoiQ29udHJvbGxlciIsIm5iZiI6MTYxODc2OTIyMSwic3ViIjoiTG9nSW4ifQ.Mrbm5tEIeLWiK49dQf9l4LqVzKcYN8rsxOCpB9Jeuds",
    "Response": {
        "RequestId": "e3fcb873-5be1-488a-b190-c3f6b97aeb40",
        "Result": {},
        "Status": "Success"
    },
    "RetCode": 0,
    "RetMsg": "Success"
}
```

##### 返回参数说明

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|Bearer |String   |鉴权token，有效期24h  |
|RetCode |Int   |状态码，0表示正常  |
|RetMsg |String   |状态信息  |
|Response |ResponseMap   |响应信息  |

###### Response

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|RequestId |String   |本次请求唯一ID  |
|Status |String   |Success 或 Fail  |
|Result |Interface   |接口具体返回的详细数据，有错误则不返回  |
|Error |String   |错误信息，没有错误则不返回  |

##### 错误码

|错误码|说明|
|:-----  |:-----|
|0 |请求正常   |
|5001 |Token过期   |
|5002 |访问拒绝   |
|5003 |请求参数非法   |
|5004 |请求数据库错误   |
|5005 |密码错误   |
|5006 |无Token   |
|5007 |Token   |
|5008 |非法的动作请求   |
|5009|参数不全   |



#### 用户登录

##### 简要描述

- 用户登录

##### 请求URL
- ` http://39.108.110.77/DangoMark/api `

##### 请求方式
- POST

##### 请求示例

``` 
{
    "Module": "Controller",
    "Action": "Login",
    "User": "111",
    "Password": "222"
}
```

##### 请求参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|Module |是  |String |模块名   |
|Action |是  |String | 动作名    |
|User |是  |String | 用户名    |
|Password |是  |String | 密码   |

##### 返回示例

``` 
{
    "Bearer": "",
    "Response": {
        "RequestId": "396ab4b1-6db5-4c20-a073-7e3975dc7c66",
        "Result": {
            "Bearer": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMTEiLCJleHAiOjE2MjMwNDcwNTAsImp0aSI6IjEiLCJpYXQiOjE2MjI5NjA2NTAsImlzcyI6IkRhbmdvTWFyayIsIm5iZiI6MTYyMjk2MDY1MCwic3ViIjoiTG9naW4ifQ.DIi6JehZJSr1Wq4N-9-A7RpWu3pXZL1bG-6qs_I5Kxk",
            "ID": 1,
            "User": "111"
        },
        "Status": "Success"
    },
    "RetCode": 0,
    "RetMsg": "Success"
}
```

##### 返回参数说明

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|Bearer |String   |鉴权token  |
|ID |Int   |用户操作ID  |
|User |String   |用户名  |



#### 获取图片数据

##### 简要描述

- 获取图片数据

##### 请求URL
- ` http://39.108.110.77/DangoMark/api `

##### 请求方式
- POST

##### 请求示例

``` 
{
    "Module": "Controller",
    "Action": "GetImageData",
    "Language": "JAP",
    "Status": 0
}
```

##### 请求参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|Module |是  |String |模块名   |
|Action |是  |String | 动作名    |
|Language |是  |String | 语种: CH-中文, ENG-英语, JAP-日文, KOR-韩语    |
|Status |是  |Int | 数据状态: 0-未标注, 1-已标注未复检, 2-已标注已复检, 3-已用于训练, 4-无意义数据   |

##### 返回示例

``` 
{
    "Bearer": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMTEiLCJleHAiOjE2MjQzMzQ0MTEsImp0aSI6IjEiLCJpYXQiOjE2MjM3Mjk2MTEsImlzcyI6IkNvbnRyb2xsZXIiLCJuYmYiOjE2MjM3Mjk2MTEsInN1YiI6IkxvZ2luIn0.YB01m023aN5muj3_wpSpMWfU-kgMoQb1I0L9trb1UMc",
    "Response": {
        "RequestId": "bfa626fc-09d8-4f03-a073-7f758ad8a09a",
        "Result": {
            "Total": 98085,
            "Data": {
                "ID": 1,
                "Url": "http://39.108.110.77/group1/default/20210614/01/21/5/919350568b9210e1e87544c23829e0d7",
                "Language": "JAP",
                "Suggestion": {
                    "Words": [
                        "でも、痛みは?」",
                        "隐感"
                    ]
                },
                "MarkResult": "",
                "QualityResult": "",
                "CoordinateJson": "",
                "Status": 0
            }
        },
        "Status": "Success"
    },
    "RetCode": 0,
    "RetMsg": "Success"
}
```

##### 返回参数说明

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|Total |Int   |数据总量  |
|Data |Map   |具体的图片数据  |
|ID |Int   |图片唯一索引  |
|Url |String   |图片链接  |
|Language |是  |String | 语种: CH-中文, ENG-英语, JAP-日文, KOR-韩语    |
|Suggestion |Map   |预标注建议内容  |
|Words |StringArray   |文字内容  |
|MarkResult |String   |标注结果  |
|QualityResult |String   |复检结果  |
|CoordinateJson |String   |文字坐标json字符串  |
|Status |Int | 数据状态: 0-未标注, 1-已标注未复检, 2-已标注已复检, 3-已用于训练, 4-无意义数据   |