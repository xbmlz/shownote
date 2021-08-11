# shownote
一款简洁、高效、免费的云笔记软件

<div align=left>
<img src="https://github.com/viodo/shownote/workflows/build/badge.svg"/>
<img src="https://img.shields.io/badge/golang-1.15-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.2-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.0.5-brightgreen"/>
</div>

### TODO

- [ ] 新建文件时有默认默认名称【新建文件.md】,如果当前目录已经有同名文件，则遵循重名判断逻辑自动修改名称
- [x] 增加`ctrl+s`保存功能
- [x] 默认笔记内容为预览模式，点击编辑按钮时再切换为编辑模式
- [x] 右键目录空白处新建文件夹
- [x] 删除提示
- [ ] 目录自定义排序
- [x] 目录中文件显示对应类型的图标
- [ ] 编辑器字体默认设置为16px
- [x] 增加注销账号功能
- [ ] 增加分享功能

### BUG

- [x] 目录和笔记内容分割把手不起作用
- [x] 新增文件或目录时，重名检测不生效
- [x] 登录成功后，再次进入还需要登录
- [x] 切换到新的笔记内容时按下`ctrl+z`，会显示上次点击的笔记内容
- [x] 编辑区域预览和编辑状态时，状态栏和文本区域样式错乱
- [x] token失效后无法重新登录，应自动刷新token
- [x] 新增笔记后，处于禁用状态，无法编辑

