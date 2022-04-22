## 项目背景
参考
- [vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
- [vue-admin-template](https://github.com/PanJiaChen/vue-admin-template)

`vue-admin-template` 是一个极简的vue骨架，使用 `vue cli 3.0` 构建，它只包含 Element UI & axios & iconfont & permission control & lint，这些搭建后台必要的东西


## 获取源码
获取源码
>`git clone https://github.com/wangpengliang815/vue.admin.sample.git`

安装npm依赖包
> `npm install` 

运行
> `npm run dev`

##  项目结构
```
├──  build                       // 构建相关  
├──  mock                        // 使用mock.js构建Mock数据前后端分离开发
├──  public                      // 主页，打开网页后最先访问的页面
├──  src                         // 源代码 
├──  ├──  api                    // 所有请求
├──  ├──  assets                 // 主题 字体等静态资源
├──  ├──  components             // 全局公用组件
├──  ├──  icons                  // 项目所有 svg icons
├──  ├──  router                 // 路由
├──  ├──  store                  // 全局 store管理
├──  ├──  styles                 // 全局样式
├──  ├──  utils                  // 全局公用方法 
├──  ├──  views                  // view
├──  ├──  App.vue                // 项目入口文件
├──  ├──  main.js                // 入口 加载组件 初始化等
├──  tests                       // 前端测试
├──  .editorconfig               // 代码格式控制
├──  .eslintignore               // （配置）eslint 检查中需忽略的文件（夹）
├──  .eslintrc.js                // eslint,提供不同IDE统一代码规范
├──  .gitignore                  //（配置）在上传中需被 Git 忽略的文件（夹）
├──  .babelrc                    // babel 转码配置
├──  jest.config.js              // 单元测试相关配置
├──  package.json                // 本项目的配置信息，启动方式
├──  package-lock.json           // 记录当前状态下实际安装的各个npm package的具体来源和版本号
├──  postcss.config              //
├──  vue.config.js               // 

```
## 调用关系
用到最多的就是src文件夹，编写的代码要放在里面，在public目录下存在index.html文件，在src目录下存在`main.js`，`App.vue`以及在router文件夹下存在`index.js`，搞清楚这几个文件之间的关系将后续开发理清思路

**index.html**：主页，项目入口

`index.html` 为项目访问的首站点，一般只定义一个空的根节点，在`main.js`里面定义的实例将挂载到根节点下，内容都通过vue组件进行填充


**App.vue**：根组件

在Vue中要经常建立后缀名为.vue的文件，.vue的文件通常由三部分构成，分别用`<template></template>`，`<script></script>`与`<style></style`>标签包裹，可以理解为前端的 `html`，`javascript` 与 `css` 三个部分，其中，<template></template>通常建立我们要用的网页界面，<script></script>通常与数据打交道，定义数据的首发方式等，面向逻辑，而<style></style>主要负责<template></template>标签中的样式

补充：

`【template】`

其中模板只能包含一个父节点，也就是说顶层的div只能有一个（例如上图，父节点为#app的div，其没有兄弟节点）

<router-view></router-view>是子路由视图，后面的路由页面都显示在此处

<router-view>类似于一个插槽，跳转某个路由时，该路由下的页面就插在这个插槽中渲染显示

`【script】`

vue通常用es6来写，用export default导出，其下面可以包含数据data，生命周期(mounted等)，方法(methods)等

`【style】`

样式通过style标签<style></style>包裹，默认是影响全局的，如需定义作用域只在该组件下起作用，需在标签上加scoped，<style scoped></style>

如要引入外部css文件，首先需给项目安装css-loader依赖包，打开cmd，进入项目目录，输入npm install css-loader,回车。安装完成后，就可以在style标签下import所需的css文件。这样，我们就可以把style下的样式封装起来，写到css文件夹，再引入到页面使用，整个vue页面也看上去更简洁
```
<style> 
 import './assets/css/public.css'  
</style> 
 ```
**main.js**： 入口文件

main.js主要是引入vue框架，根组件及路由设置，并且定义vue实例

`components:{App}`就是引入的根组件 `App.vue`

后期还可以引入插件，当然首先得安装插件

**router.js**: 路由配置

在router文件夹下，有一个index.js文件，即为路由配置文件

## vue-cli
参考
- https://cli.vuejs.org/zh/guide/ 
> vuecli 3.0 和 2.0 区别参考
- https://zhuanlan.zhihu.com/p/154106807

## vuex
参考
- https://vuex.vuejs.org/zh/ 
- https://www.jianshu.com/p/2e5973fe1223 
- https://www.cnblogs.com/junjun-001/p/12546023.html


## mock.js 
参考、
- https://www.cnblogs.com/tzm-001/p/10382534.html 
- https://www.cnblogs.com/gaosong-shuhong/p/10517342.html

## 单元测试
`npm run test:unit` 

代码规范格式验证 : `npm run test:ci`


## Build

`npm run build:stage` 

`npm run build:prod` 

## 部署