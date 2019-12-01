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
