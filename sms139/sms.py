# coding=utf-8
import requests
import sys
import json
import time

reload(sys)
sys.setdefaultencoding('utf8')

if len(sys.argv) < 3:
    print '参数错误'
    exit()
# 判断参数
if len(sys.argv[1]) != 11:
    print '手机号输入有误'
    exit()
if len(sys.argv[2]) == 0:
    print '短信内容不能为空'
    exit()
phoneNum = sys.argv[1]
content = sys.argv[2]

# 登陆139

s = requests.Session()
response = s.post(
    "https://mail.10086.cn/Login/Login.ashx?_fv=4&_=980d90ee3c9fd7c2bccce499e70c2db17393edf9&resource=indexLogin",
    data={
        'UserName': phoneNum,
        'verifyCode': '',
        'auto': 'on',
        'webVersion': '25',
        'Password': '你的加密后的密码',
        'authType': '2'
    })
sid = s.cookies.get('Os_SSo_Sid')
if len(sid) < len('00UwMDk3OTc5MTAwMjAzMTg105BFD9B5000007'):
    print '发送失败，time:' + time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    exit()
else:
    # 发短信
    response = s.post(
        "http://smsrebuild1.mail.10086.cn/sms/sms?func=sms:sendSms&sid=" + sid + "&rnd=0.45051754841189506"
        , data='<object>'
               '<int name="doubleMsg">0</int>'
               '<int name="submitType">1</int>'
               '<string name="smsContent">'+content+'</string>'
               '<string name="receiverNumber">86'+phoneNum+'</string>'
               '<string name="comeFrom">104</string>'
               '<int name="sendType">0</int>'
               '<int name="smsType">2</int>'
               '<int name="serialId">-1</int>'
               '<int name="isShareSms">0</int>'
               '<string name="sendTime"/>'
               '<string name="validImg"/>'
               '<int name="groupLength">50</int>'
               '<int name="isSaveRecord">1</int>'
               '<int name="smsIds"/>'
               '<array name="receiverList"/>'
               '</object>')
    jsonResponse = json.loads(response.content)
    print response.content
    if jsonResponse['code'] == 'S_OK':
        print '短信发送成功,time:' + time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
