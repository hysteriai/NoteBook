# 一、编译和安装
## 1 编译
- 进入OpenSIPS源码文件夹，执行
  ```
  make all
  ```
  core和其他配置过要编译的模块，会被编译。

### 1.1 配置编译选项
- OpenSIPS有很多编译选项。不同的选项对应不同的功能。比如，你可以选择启用内存分析工具，或者默认不启用的TLS模块。
- 你需要用‘menuconfig’工具来修改编译选项。此工具是基于curses的，所以你在使用之前需要安装ncurses的库文件。在Debian系操作系统中，你可以通过“apt-get install libncurses5-dev”来安装。
- 安装完成后可以进入OpenSIPS根目录，执行“make menuconfig”进入配置菜单，在这里你可以用方向键浏览所有的选项（在控制台底部会显示当前选项的简单介绍）。
- 你可以通过空格键来更改配置项的使能状态，配置完成后可以用q键退出，并点击“Save Changes”。
- 更改完编译选项后，你需要重新编译并重新安装OpenSIPS
### 1.2 编译需要外部依赖的模块
- 有些模块需要一些你操作系统中没有的外部依赖项，所以OpenSIPS中一些模块是默认不编译的。因此，这些模块在安装的时候需要多加注意。比如依赖mysql devel library的DB_MYSQL模块，依赖外部JSON解析器的JSON模块等。
- 若要编译这些模块，你同样需要在menuconfig中进行配置。执行“'make menuconfig”，选择“'Configure Excluded Modules”，在这里你可以看到所有默认不启用的模块，同样，在控制台底部会显示模块的简单介绍。通过空格键来选择你需要的模块，选择完成后，按q返回上一层菜单，点击“Save Changes”。工具会显示出你选择的模块，并且会告诉你有哪些依赖项。同样，更改完编译选项后，你需要重新编译并重新安装OpenSIPS。

## 2 安装
- 进入OpenSIPS文件夹，执行
```
make install
```
- OpenSIPS会默认安装在 / 路径下

### 2.1 减少编译时间
- 可以用FASTER变量来减少编译时间。这个功能通过使用make -jNR_OF_CORES来使多核处理器并行的编译所有模块。注意，这个方法会使用大量的资源，并且开启的线程数需要等于或小于机器的核数。并且，这个变量会屏蔽大部分编译输出。
- 例如，在一个4核机器安装OpenSIPS，可以使用如下指令
```
FASTER=1 make -j4 install
```

### 2.2 配置安装路径
- 由于诸多原因（比如，需要在一台机器上部署两个OpenSIPS实例），你想改变OpenSIPS的安装路径。你需要使用上文提到的menuconfig tool来完成这件事。
- 执行'make menuconfig' 进入'Configure Install Prefix'选项，然后输入你想安装的文件夹，回车。然后选择'Save Changes' ，回车。然后再执行'make install'。

# 二、数据库部署
- 安装完OpenSIPS之后，你很可能需要部署一个数据库，以便于你能使用其他一些特性（用户鉴权、重复注册、会话等）。你可以通过 opensips-cli来部署数据库（opensips-cli需要额外安装）。
## 1 配置OpenSIPS CLI
- 打开OpenSIPS CLI 配置文件，确认如下几个参数
    - database_path：（默认 /usr/share/opensips/）需要更改为[OpenSIPS安装路径]/share/opensips/
    - database_url：链接数据库的URL，如果没配置，你需要在数据库部署时提供URL
    - database_name:（默认 opensips）使用的数据库名称
    - database_modules：（默认标准模块）你想部署的模块
- 关于OpenSIPS CLI的更多信息，你可以访问https://github.com/OpenSIPS/opensips-cli/blob/master/docs/modules/database.md#configuration

- **注意**:OpenSIPS CLI默认在 ~/.opensips-cli.cfg, /etc/opensips-cli.cfg, /etc/opensips/opensips-cli.cfg 路径下搜索配置文件，你可以使用-f参数来设置自己的配置文件。

## 2 创建数据库
- 使用如下指令创建数据库
```
opensips-cli -x database create
```
- 如果你想增加新的模块（比如presence），执行
```
opensips-cli -x database add presence
```
- 你也可以在撞见数据库的时候指定一个新名字（比如opensips_test），执行
```
opensips-cli -x database create opensips_test
```
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

## Load Balancer模块
### 1、设置目的地
- 在Load Balancer模块看来，各个目的地是不一样的。LB模块只关心各个目的地能提供哪些资源/服务，配置时只需告知LB模块目的地的地址，以及其能力和对应的最大并发量。
- 例：

| id | group_id | dst_uri | resources |
| ---- | ---- | ---- | ---- |
|  1   |  1   |  1.1.1.1 (A) | reg=1000;push=5000; |
|  2   |  1   |  2.2.2.2 (B)| reg=1000;vedio=50; |
|  3   |  1   |  3.3.3.3 (C)| reg=1000;recv=5000; |
|  4   |  1   |  4.4.4.4 (D)| reg=1000;recv=5000;vedio=200; |

- 可以通过MI（ Management Interface）来修改上述配置

### 2、使用Load Balancer模块
- 在OpenSIPS脚本中可以用如下方式使用LB模块的功能
```
  if (!load_balance("1","reg;push")) {
      sl_send_reply("500","Service full");
      exit;
  }
```
- 第一个参数为目的地的组（group_id）。第二个参数为此次会话需要的能力列表。也可以传入第三个可选参数，来指定优先选择剩余某个值或某个百分比能力的目的地。
- 会话结束LB模块会自动释放资源

### 3、目的选择逻辑
- 1、根据第一个参数，获取目的地集合
- 2、筛选出能提供需要的能力（第二个参数）的集合
- 3、根据每个目的地剩余的能力，结合需要的能力，选择合适的目的地
- 4、最终胜出的是拥有最大的剩余能力最小值的节点
- 例：
```
load_balance("1","reg;vedio")
```
- 1、获取目的地集合，group_id=1的，A B C D
- 2、需要reg和vedio，剩余B D
- 3、假设B剩余reg=900;vedio=45;D剩余reg=100;vedio=150;记录剩余的最小能力的值，B是vedio为45，D是reg为100。
- 4、100>45，所以选择D。

### 4、Disabling and Pinging
- 可以在脚本中根据自己的逻辑，通过lb_disable()将某个目的设置为disable状态。在之后的LB逻辑中，不会再选择设置为disable状态的目的地
- 可以通过如下两种方法将目的地设置为enable 状态
- 1、使用MI指令：lb_status手动启用
- 2、使用SIP Ping，OpenSIPS给目的地发送OPTIONS消息（可以设置在某条件下发送或者一直发送，一直发送情况下，如果没有回复会将目的地标记为disable状态），目的地回复200 OK，则将目的地修改为enable 状态。

### 5、实时控制LB模块
- lb_reload：强制重新从DB中重新加载所有配置
- lb_resize：修改某个目的地的能力值
- lb_status：查看/修改目的地的状态（enable/disable）
- lb_list：查看每个目的地最大能力值和已使用的能力
```
$ ./opensipsctl fifo lb_list
Destination:: sip:127.0.0.1:5100 id=1 enabled=yes auto-re=on
        Resource:: pstn max=3 load=0
        Resource:: transc max=5 load=1
        Resource:: vm max=5 load=2
Destination:: sip:127.0.0.1:5200 id=2 enabled=no auto-re=on
        Resource:: pstn max=6 load=0
        Resource:: trans max=57 load=0
        Resource:: vm max=5 load=0
```
