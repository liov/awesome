GERBER是一种光绘文件格式，用于描述光绘机进行各种绘制或运动行为。
GERBER格式是EIA 标准RS-274D的子集；扩展GERBER格式是EIA标准RS-274D格式的超集，又叫RS-274X。RS-274X增强了处理多边形填充，正负图组合和自定义D码及其它功能。它还定义了GERBER数据文件中嵌入光圈表的规则。 所以，RS-274D类型的Gerber文件不包含Aperture(光圈)数据，即需要同时附带D码文件，才能完整描述一张图形；而RS-274X类型的Gerber文件则不用附带。

GERBER格式解析
GERBER格式文件由一系列数据块组成。所有的数据块以结束（EOB）符结尾，EOB字符通常是星号（*），而每个数据块包括了一个或多个参数组成，例如X10000Y0DO1*。

数据块的类型主要包括以下几种类型：

坐标数据（Coordinate Data）
功能码（Function Codes)
RS-274X参数（Parameters)
标准的RS-274D中，数据类型包括了坐标数据以及功能码，如D码，G码，M码等。

坐标数据（Coordinate Data）
坐标数据主要是定义在平面的中点数据，在RS274D的术语中称为地址。坐标数据可能是：
1）X和Y坐标定义的点，
2）相对于X，Y方向的便移量数据，称为I，J数据

坐标数据（Coordinate Data）：

D：指定当前使用的光圈编号。
X 和 Y：指定坐标位置。
I 和 J：用于圆弧插补的相对坐标。
R：用于圆弧插补的半径。

FS(Format Specification) 格式定义指示了数字如何被解释的。
坐标系采用右手坐标系。坐标是模态（modal) 的，如果一个X被忽略，则X将保留上一次的X坐标值，如果在当前层的第一个X被忽略，因为没有上一次的X的坐标值，那么X坐标将被视为零。类似地，Y坐标也是这样处理的。
偏移量不是模态上的，如果I或J被忽略，则缺省值为零。

注意：GERBER的读者有时候会错误地处理缺省值零。为了清晰和鲁棒性，推荐总是显式地指定第一个坐标（即便就是零，也显式指定），这样就不用考虑缺省的零。

示例：
X100Y200*             坐标点 (+100, +200)
Y-300*                     坐标点 (+100, -300)
I200J100*               平移 (+200, +100)
X300Y200I150J50* 坐标点(+300, +200) 且平移(+150, +50)
X+100I-50*              坐标点 (+100, +200) 且 平移 (-50, 0)

功能码（Function Codes)
功能码描述的是如何解析相关联的坐标数据，如画一条线或画一个圆。（通常，但不是所有，这些代码是延续已经过时的RS-274D的格式，它们被称为字(words)或码(codes)），如
G04 PC Circuitry*
G54D10*
G54D11*
G01X466000Y240000D02*
X474000D01*
X470000Y236000D02*
....
G74*
X0Y0D02*
M02*

每个指令都会影响到其后的数据块，直到遇到另外一个相同类型的代码或生成新层时结束。我们称这种持续性的活动为模态(modal)。例如G02指示的是顺时针圆弧插补。在遇到另外一个插补指令或生成新层之前，该指令后的所有坐标数据都被解释为顺时针圆弧插补。

N码：顺序码，命名数据块顺序。（0-99999）  
D码：绘图码，选择，控制光圈，指定线型。
D01 划线，开光圈。 不能用自定义光圈划线
D02 关光圈
D03 闪绘光圈，反复曝光
D04 绘画笔提取，快速移到
D05 结束D04的动作
D10-D999 选择由AD命令定义的光圈

（后补:
G01：直线插补。
G02：顺时针圆弧插补。
G03：逆时针圆弧插补。
G04：暂停（延时）。
G90：绝对坐标模式。
G91：相对坐标模式。
）
G码：通用码，用于坐标定位。  
G00        快速移动             格式：G00[Xsn][Ysn]D02
G01/G1  1:1线性运动        格式：G01[Xsn][Ysn][Dn]*   D码控制光圈并运动至指定坐标点
G02/G2  顺时针圆周运动  格式：G02[Xsn][Ysn][In][Jn][Dn]*   I/J表示圆心到坐标的距离
G03/G3  逆时针圆周运动
G04  忽略当前数据块
G05  更换镜头
G10  10倍线性比例
G11   0.1倍线性比例
G12   0.01倍线性比例
G20   指定英寸单位
G30   指定毫米单位
G36  打开多边形填充
G37 关闭多边形填充
G54  选择光圈
G70  指定英寸单位
G71  指定毫米单位
G74  四分之一圆周运动模式   格式同G02
G75  360度圆周运动模式
G84  用1/3孔径大小的钻头在XnYn处钻直径为M的大孔，格式：XnYnG84XM
G85  在两个坐标点之间钻出槽孔，格式：XnYnG85XmYm
G90  指定绝对坐标格式
G91  指定相对坐标格式
M码：指定文件结束等。
M00  程序停止
M01  条件停止
M02  文件结束
M03  结束磁带程序或回带
M08  结束重复指令
M25  重复指令中，定义数据块的起始
M30  程序结束指令
M48  程序起始指令
M64  设定当前位置为图档的原点，并继续绘图
M71:表示公制
M72:表示英制
M97  在指定坐标点沿X轴纂刻文本
M98  在指定坐标点沿Y轴纂刻文本
RS-274X参数（Parameters)
参数定义了整个图像或单层的各种特征。它们被用于解释其他的数据类型，（通常，这些参数被称为Mass 参数）。控制整个图像的参数通常会放在文件的开始处。产生新层的参数被放置在文件恰当的位置。参数由两个字符加一个或多个紧随其后的可选修改符组成。参数的限定符号为“%”.每个包含在数据块内的参数必须以“*”结束。并且参数限定符必须立即跟在块结束符后面，不允许插入空格，例如：
%FSLAX23Y23*%
参数必须是在成对的参数限定符内，限定符内可以放一个或多个参数，两个限定符之间最大的字符数为4096个，例如：
%SFA1.0B1.0*ASAXBY*%
为了提高可读性，两个参数间允许换行，如：
%SFA1.0B1.0*
ASAXBY*
%
当然，为了简化和可读性，推荐每行是只设置一个参数。与参数联合的所有数值都使用显式的小数点，如果不使用小数点，数值应当认为是整数。
参数的语法为：
%参数指令<必选修饰符>[可选修饰符]%

语法	说明
参数指令 （parameter code)	两个字符的指令，如AD,AM,FS等
必选修饰符(required modifiers)	必须是完整的定义
可选修饰符(optional modifiers)	依赖必选修饰符的定义
具体参数分类如下：

提示性参数
AS （Axis Select）坐标选择
格式：%ASA[X|Y]B[X|Y]*%  
其中，A B 输出设备坐标轴
X Y 数据文件坐标轴
FS （Format Statement）格式描述
格式：%FS[L|T][A|I][Nn][Gn]XnnYnn[Dn][Mn]*%  
其中，L/ T --  L 省略前导零 T省略尾零
A / I --  A 绝对坐标 I 相对坐标
Nn Gn Dn Mn  -- 设定N G D M码的长度/范围， n=2 表示00-99
Xnn Ynn  -- 坐标数据格式，例如X23表示X轴坐标含两位整数位和三位小数位
MI （Mirror Image）镜像图像
格式：%MI[A[0|1]B[0|1>*%   
其中，0 -- 不镜像，1 --  镜像

MO （Mode）单位
格式：%MO[IN|MM]*%   
其中，IN -- 英寸 ，MM -- 毫米
OF （Offset ）偏移
格式：%OFA<n>B<n>*%   
其中，A<n> n定义输出设备A轴向的偏移，5.5格式
B<n> n定义输出设备B轴向的偏移，5.5格式
SF （Scale Factor）比例因子
格式：%SF[A<n>][B<n>]*%   
其中，A<n> n定义输出设备A轴向的比例
B<n> n定义输出设备B轴向的比例
图像参数
IJ（Image Justify）图像对齐
格式：%IJ[A[L|C]B[L|C>[<offset>]*%  
其中，A A轴对齐
L 左或下对齐
C 中心对齐
B B轴对齐
<offset> 偏移
IN （Image Name）图像名称
格式：%IN<name>*%
命名当前图像为name

IO （Image Offset）图像偏移
格式：%IOA<n>B<n>*%   
其中，A<n> n定义输出设备A轴向的偏移
B<n> n定义输出设备B轴向的偏移
IP （Image Polarity）图像正负性
格式：%IP[NEG|POS]*%   
其中，IPNEG 设置为负图
IPPOS 设置为正图
IR （Image Rotate）图像旋转
格式：%IR[90|180|270]*%
表示逆时针旋转图像  
PF （Plot Film）绘图胶片名
格式：%PF<name>*%
表示提示操作员胶片名为name
光圈参数
AD（Aperture Definition）光圈描述
格式：%ADD<n1><type>,<n2>[X<n3>]*%  
其中，<n1> 为D码编号(10-9999)  
<type>     <n2>     <n3>         <n4>         <n5>         <n6>  
C(圆)         外径      X向孔径     Y向孔径      
R(长方)     X向大小 Y向大小     X向孔径     Y向孔径    
O(椭圆)     X向大小 Y向大小     X向孔径     Y向孔径    
P(正多边)  外径     边数         旋转角度     X向孔径     Y向孔径  
AM （Aperture Macro） 自定义光圈  
格式：%AM<name>*<type>,<$1>,<$2>,[<…>]*
[<type>,<$1>,<$2>,[<…>>*…*%  
其中，<name> 为当前自定义光圈定义一个名称
<type>        $1         $2         $3         $4         $5         $6         $7         $8         $9  
1(圆)           Exp         直径     圆心X     圆心Y            
2/20(线)      Exp         线宽     起点X     起点Y     终点X     终点Y     角度      
21(长方形)  Exp     宽         高         中心Ｘ     中心Ｙ     角度        
22(长方形)  Exp     宽         高         左下X     左下Y     角度        
4(多边形)    Exp     点个数     起点X     起点Y     X1         Y1 。。。    角度  
5(正多边形) Exp     顶点数     中心X     中心Y     直径          
6(Moire)      X0         Y0         外径     环宽     环间距     环个数     十宽     十长     角度  
7(散热形)    X0         Y0         外径     内径     口尺寸     角度

层参数
KO （KnockOut）挖除
格式：%KO[C|D][XnYnInJn]*%
其中，C Clear 挖除矩形块
D Dark 添补矩形块  
XnYn 矩形块左下角坐标
In 矩形块宽度
Jn 矩形块高度  
LN （Layer Name）层名
格式：%LN<name>*% 命名当前层为name
LP （Layer Polarity）层正负性
格式：%IP[C|D]*%
其中，IPC 设置为负图
IPD 设置为正图
SR （Step & Repeat）移动与复制
格式：%SR[Xn][Yn][In][Jn]*%  
其中，Xn In X方向移动复制的数量和步长
Yn Jn Y方向移动复制的数量和步长
其他参数
IF（Include File） 嵌入文件
格式：%IF<filename>*%
表示把filename中的内容放到当前位置  
实例分析

```gerber

*
*
G04 PADS VX.2.7 Build Number: 15549477 generated Gerber (RS-274-X) file*
G04 PC Version=2.1*  //G04表示本行是注释描述
*
%IN "mcuplane.pcb"*%  //Image Name图形名称
*
%MOIN*%    //模式单位，IN:inch  MM:milimeter
*
%FSLAX35Y35*% //格式描述：忽略前导零，XY轴数据格式都为3个整数+5个小数（数值单位为inch）
*
*
G04 PC Standard Apertures*
*
*
G04 Thermal Relief Aperture macro.*
%AMTER*  //光圈自定义,命名为TER
1,1,$1,0,0*  //第一个参数为形状类型，1表示圆形
1,0,$1-$2,0,0*
21,0,$3,$4,0,0,45*
21,0,$3,$4,0,0,135*
%
*
*
G04 Annular Aperture macro.*
%AMANN*  //光圈自定义,命名为ANN
1,1,$1,0,0*
1,0,$2,0,0*
%
*
*
G04 Odd Aperture macro.*
%AMODD* //光圈自定义,命名为ODD
1,1,$1,0,0*
1,0,$1-0.005,0,0*
%
*
*
G04 PC Custom Aperture Macros*
*
*
*
*
*
*
G04 PC Aperture Table*
*
%ADD010C,0.001*%  //设置D码为10的光圈，圆图形，直径为1 mil
%ADD011C,0.01*%    //设置D码为11的光圈，圆图形，直径为10 mil
*
*
*
*
G04 PC Circuitry*
G04 Layer Name mcuplane.pcb - circuitry*
%LPD*%   //设置Layout层为正
*
*
G04 PC Custom Flashes*
G04 Layer Name mcuplane.pcb - flashes*
%LPD*%  //设置Layout层为正
*
*
G04 PC Circuitry*
G04 Layer Name mcuplane.pcb - circuitry*
%LPD*%  //设置Layout层为正
*
G54D10*   //选择D码为10的光圈
G54D11*   //选择D码为11的光圈
G01X466000Y240000D02*  //1倍线性运动，关闭光圈，移到坐标点（4660.00 mil,2400.00 mil)
X474000D01*            //打开光圈，移到坐标点（4740.00 mil,2400.00 mil)
X470000Y236000D02*     //关闭光圈，移到坐标点（4700.00 mil,2360.00 mil)
Y244000D01*            //打开光圈，移到坐标点（4700.00 mil,2440.00 mil)
....
X348909Y113875*
X348364*
G74*                   //关闭圆周运动
X0Y0D02*               //关闭光圈并回到原点
M02*                   //文件结束
```
转自 https://blog.csdn.net/Golden_Chen/article/details/127734439

const (
GTL Loayer = "GTL" //顶层走线
GBL Loayer = "GBL" //底层走线
GTO Loayer = "GTO" //顶层丝印
GBO Loayer = "GBO" //底层丝印
GTS Loayer = "GTS" // 顶层阻焊
GBS Loayer = "GBS" //底层阻焊
GPT Loayer = "GPT" //顶层主焊盘
GPB Loayer = "GPB" //底层主焊盘
G1  Loayer = "G1"  //内部走线层1
G2  Loayer = "G2"  //内部走线层2
G3  Loayer = "G3"  //内部走线层3
G4  Loayer = "G4"  //内部走线层4
GP1 Loayer = "GP1" //内平面1(负片)
GP2 Loayer = "GP2" //内平面2(负片)
GM1 Loayer = "GM1" //机械层1
GM2 Loayer = "GM2" //机械层2
GM3 Loayer = "GM3" //机械层3
GM4 Loayer = "GM4" //机械层4
GKO Loayer = "GKO" //禁止布线层(可做板子外形)
)
