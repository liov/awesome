前言：

网站早已经不流行，赶快回来到我身边。--周杰伦《好久不见》

那就做app吧，选来选去，无疑只能锁定flutter，跨端，高性能，大厂背书，有坑好搞。可惜了学了好久的kotlin，竟用武之地不大。

part1-跑demo：

先来个最简单的demo，下载flutter，AndroidSDK，我用的是IDEA直接新建了个fultter工程，打开模拟器，直接——run

nice，run了半天没动静，换源

添加环境变量

export PUB_HOSTED_URL=https://pub.flutter-io.cn
export FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn

此外你可能需要添加一系列的path

${AndroidSdk}\tools;${AndroidSdk}\tools\bin;${AndroidSdk}\platform-tools;${flutter}\bin;${flutter}\bin\cache\dart-sdk\bin;

如果用到protobuf

$Home\AppData\Roaming\Pub\Cache\bin;$protoc\bin;
继续run

这时候看运气，有时候run的起来，有时候不行，大多时候卡在assembleDebug or assembleRelease，

然后报错，xxx-debug.xxxx下载不下来




flutter upgrade到1.20.x+；1.20版本的flutter.gradle有改动，直接取环境变量的源下载，至于1.20以下怎么解决，自己搜索（干嘛不升级！！！）

接着run

成功在屏幕中央显示一个0，右下角一个+号按钮，可喜可贺，成功1%也就是成功了99%（反正都没成功）

part2-改界面：

先折腾折腾flutter吧，看看文档，学学组件，自己写个小页面，在国内还是用国内的网站吧，感谢这些大佬的贡献

https://flutter.cn/，https://dart.cn/，https://flutterchina.club/

刚开始你会被这层层嵌套恶心的想疯狂骂shit，但又敬畏于大厂的作品怀着谦卑的心继续尝试了下来，然后你会觉得也就那样，div不也是一层套一层吗，只要你拆的够细，封装度够高，其实代码也是很简洁的




于是，你得到了这样的界面




是的，实现了减的功能，并且两个按钮移到了中间

part3-赋予动态能力：

现在多数app都采用内嵌h5的方式来快速多端开发，实现热更新，我的app怎么能少了这样的能力，虽然flutter有热重载的能力，但毕竟不是热更新，想到热更新，第一时间想到的就是lua和webview，不过lua只能用来实现逻辑的热更，当然后面有意外收获，后话了

webview

webview可以实现页面热更，我也正准备把pc版的hoper改h5（最初就是h5的版本，亿种循环），因为现在用电脑看网站的太少了，除了从业相关的网站，显然我的网站并不是，当然我是先尝试的lua，但是这里讲webview，因为简单




官方有集成好的插件，直接在依赖引入webview_flutter，然后照着官方文档把webview当做组件集成就好了

刚好把我的https://hoper.xyz放进去。

lua

最开始的思路是搜索安卓lua热更新，搜索结果并不理想，然后在GitHub上找找有没有啥成熟案例，关键词flutter lua，别说还真被我找到了，最开始找的是一个简单集成lua的插件，使用很简单，像webview一样加依赖，然后用插件，不过貌似flutter的插件调用都会被封装成async函数，简单封装了个text，接着尝试跑一下，gg，再跑一下，gg，这里不得不说一下dart的async函数，异步函数必须在异步函数里调用，带有await的函数必须有async关键字，和js一样，然鹅，你不用await就不会调用，js不是，所以在dart的同步函数中调用异步函数只能用then+回调的形式(更正：并不是这样的，异步函数就是异步函数)




回到上面的crash，看报错，打印了一堆内存地址，猜也知道是因为依赖的so库跟模拟器cpu水土不服，当然我不是猜出来的，我是上真机能跑才知道的，按理说到这里应该就结束了，然而我又发现了一个星星比较多的插件




flutter_luakit_plugin，国人开发，看了介绍，功能强大，集成n多第三方库，虽然都有点旧，上手试试，诶，跟上面那个一样的问题，上真机，没问题，就是集成太多so库大了点，4M，还行，接着拆包看了下，咦，咋还有libgojni.so，啥玩意，我也没用go啊，看了上面那个lua库得源代码，有少量的go，点进去，原来用的是go lua5.2的实现，一个开源库，库中说比clua慢6倍，还是5.2的就放弃了，只留个luakit




按理说这里总该结束了吧，还没，不过是后话了




part4-混合开发

一个富有探索精神的开发者是不会止步于此的，下个目标是，混合开发，按例先看有没有成熟方案，先找到了闲鱼的flutter_boost混合栈管理，看了文档我也是云里雾里，毕竟不是做移动开发的，我适用了ios的闲鱼app，说实话，很流畅的，于是开搞，但是flutter_boost是在已有原生app的基础上添加flutter页面，看代码加的还是fragment，我是flutter项目加原生，不管那么多，依瓢画葫芦，把MainActivity都改成继承Activity了，然后添加其他Activity，最后居然奇迹般的跑了起来，页面也都能跳转，然而使用体验并不好，主要是白屏时间接受不了，此过程中发现了flutter官方的add-to-app方案，也是像已经存在的app加flutter页面，这次我决定综合二者，依瓢画葫芦，在flutter中添加原生页面，自己写跳转，没想到，成了，现在我已经掌握了flutter和原生页面相互跳转的核心技术🙊，原生中直接跳转到flutterAcitvity，flutter中利用methodChannel调用原生跳转方法，原生跳fultter出现的白屏时间过久情况，在Application中缓存flutterEngine

part5-FFI

如果以上都做了，那怎么会想不到FFI呢，做FFI当然得用最钟爱的语言Rust，dart做UI，底层Rust，想想就激动呢




Rust提供了完善的工具链，这使得FFI开发简直不要太简单

$Home/.cargo/config文件中添加

[target.x86_64-linux-android]
linker = "ndkbinpath\\x86_64-linux-android30-clang.cmd"
ar =  "ndkbinpath\\x86_64-linux-android-ar"

[target.aarch64-linux-android]
linker = "ndkbinpath\\aarch64-linux-android30-clang.cmd"
ar =  "ndkbinpath\\aarch64-linux-android-ar"

[target.armv7-linux-androideabi]
linker = "ndkbinpath\\armv7a-linux-androideabi30-clang.cmd"
ar =  "ndkbinpath\\armv7a-linux-androideabi-ar"
rustup target add x86_64-linux-android armv7-linux-androideabi aarch64-linux-android
cargo new rust --lib

Cargo.toml添加

[lib]
crate-type = ["staticlib", "cdylib"]

写个简单的导出Cabi的函数

use std::os::raw::{c_char};
use std::ffi::{CString, CStr};

#[no_mangle]
pub extern "C" fn rust_greeting(to: *const c_char) -> *mut c_char {
let c_str = unsafe { CStr::from_ptr(to) };
let recipient = match c_str.to_str() {
Err(_) => "there",
Ok(string) => string,
};

    CString::new("Hello ".to_owned() + recipient).unwrap().into_raw()
}
cargo build --release --target=armv7-linux-androideabi

为什么是armv7呢，我也不想啊，鸡精的开发者当然想用64位，但是好多第三方库都是32位啊，这坑我先替你们踩了

调用的方式不是JNI+MethodChannel，而是Dart直接FFI，dart有内建的dart:ffi库，但是只能操作数值类型，必须引入额外的ffi库，才能传字符串过去

封装个组件看结果，自豪感满满

part5-lua后续

无意间发了一个库MLN，陌陌团队开发的跨平台基于lua的动态UI，虽然最近好像不怎么更新了，但是爱折腾的你怎么能放过