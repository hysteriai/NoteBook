# 三、数据库的schema
- 数据库表的格式，直接参考官网即可
# 四、配置文件
- OpenSIPS配置文件包含控制核心和外部模块的所有参数以及路由SIP消息的路由逻辑。
- 安装完成后，默认配置文件路径是：
```
[INSTALL_PATH]/etc/opensips/opensips.cfg
```
- 配置文件是基于文本的，使用OpenSIPS特有的语法编写（语法类似C语言）。脚本中可以有很多变量（在文档后面会进一步介绍），你可以写if / while / switch等这样的结构。你也可以通过参数调用一些子例程，所以对有SIP基础的人，脚本应该是十分易读的。
- **注意**：如果你更改了配置文件，你必须重启OpenSIPS来试更改生效
- 因为每次修改之后需要重启，修改时保证没有语法错误时十分重要的。你可以通过如下指令来检验配置文件合法性：
```
[INSTALL_PATH]/sbin/opensips -C [PATH_TO_CFG]
```
- 如果配置文件合法，会返回0.如果非法，会打印错误，返回-1

## 1、生成配置文件
- 可以使用menuconfig tool来生成配置文件。图形界面是基于ncurses 的，所以需要安装ncurses库（一般是 libncurses5-dev）
### 1.1使用menuconfig tool
- menuconfig tool在源文件夹和安装文件夹下都可以执行
  - 在源文件夹内，执行：
   ```
    make menuconfig
    ```
  - 在安装后，可以在安装路径下执行
  ```
  [install_path]/sbin/osipsconfig
  ```
- 在menuconfig tool中选择 'Generate OpenSIPS Script'，然后选择你想生成的脚本类型。选择脚本类型后能够配置很多有用的选项（具体会在文档下面描述）。使用空格键来激活脚本中的选项。配置完成后，用q键退出，然后选择 'Save Changes'。然后你可以通过你的配置生成脚本。最后图形界面会返回给你新生成配置文件的路径（比如：Config generated : /usr/local/etc/opensips/opensips_residential_2013-5-21_12:39:48.cfg）

### 1.2配置的类型
- 目前可以生成3种类型的脚本（menuconfig1.12），每种脚本的可选参数如下：
  - Residential Script：
    - ENABLE_TCP：OpenSIPS会监听TCP的SIP请求
    - ENABLE_TLS ：OpenSIPS会监听TCP的SIP请求
    - USE_ALIASES：OpenSIPS允许SIP用户使用别名
    - USE_AUTH：OpenSIPS会对用户注册和invite进行鉴权
    - USE_DBACC：在DB中记录所有会话的ACC登入
    - USE_DBUSRLOC：在DB中记录用户的登陆地点
    - USE_DIALOG：持续跟踪会话
    - USE_MULTIDOMAIN：为订阅者处理多域
    - USE_NAT：尝试修改SIP消息和使用RTPProxy来处理NAT
    - USE_PRESENCE：OpenSIPS可以当作一个Presence server 
    - USE_DIALPLAN：在本地号码的转换中使用呼叫计划
    - VM_DIVERSION：重定向发不到订阅者的VM呼叫（OpenSIPS will redirect to VM calls not reaching the subscribers）
    - HAVE_INBOUND_PSTN：接受来自PSTN网关的呼叫
    - HAVE_OUTBOUND_PSTN：发送数字会话到PSTN网关
    - USE_DR_PSTN：PSTN互联时的动态路由支持
  - Trunking Script
    - ENABLE_TCP：OpenSIPS会监听TCP的SIP请求  
    - ENABLE_TLS ：OpenSIPS会监听TCP的SIP请求
    - USE_DBACC：在DB中记录所有会话的ACC登入
    - USE_DIALPLAN：在本地号码的转换中使用呼叫计划
    - USE_DIALOG：持续跟踪会话
    - DO_CALL_LIMITATION：限制每条干线的并发呼叫量
  - Load-Balancer Scrip
    - ENABLE_TCP：OpenSIPS会监听TCP的SIP请求  
    - ENABLE_TLS ：OpenSIPS会监听TCP的SIP请求
    - USE_DBACC：在DB中记录所有会话的ACC登入
    - USE_DISPATCHER：使用DISPATCHER 而不是Load-Balancer来分发流量
    - DISABLE_PINGING：不ping目的（否则发现失败后会ping）
### 1.3 已生成脚本的修改
- 生成脚本后，你需要用编辑器打开脚本，浏览所有的 '# CUSTOMIZE ME'注释。这些注释标记出了用户需要注意的地方，以及监听端口和数据库的URL
- 全部确认完成后，你可以保存你的脚本，然后进行测试 
## 2、模板生成opensips.cfg文件
### 2.1 支持泛型预处理
- OpenSIPS 3.1+支持在opensips.cfg文件（包括其他被导入的文件）中使用泛型预处理指令。当opensips.cfg必须被参数化（比如监听接口、端口，数据库链接器等等），或者需要自动化部署到多台服务器上时，泛型支持会十分有用。系统管理员可以通过"-p \<cmdline\>" 来实现，例如：
```
opensips -f opensips.cfg -p /bin/cat
```
- 这是"-p"的基本应用：通过提交他到接收标准输入并映射到标准输出的"echo"预处理器上。
- 用其他支持模板的语言也可以替代，只要满足开发需要即可。比如sed:
```
opensips -f opensips.cfg -p "/bin/sed s/PRIVATE_IP/10.0.0.10/g"
```

### 2.2 常用的模板语言+示例
