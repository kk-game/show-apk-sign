# 写在最前面

1. 项目是在 https://github.com/Ed1s0nZ/APK-SignCheck 基础上优化的,了解更多可以查看原作者
2. 因为原作者的页面只能点击按钮上传, 比较不方便, 我优化了一版
3. 还有因为原来的页面跳转, 上传会出现意外, 还有上传过程中没有响应,体验较差
4. 基于以上原因优化, 目前只能在windows平台使用, 请不要构建linux
5. 如果有优化,或者需求,可以提交issues,我可能会修改

# APK签名检查工具

此工具旨在简化APK文件签名的查看过程，通过一个简单易用的Web界面，使非技术人员也能轻松查看APK文件的签名信息。

## 功能简介

本工具通过Web端接收上传的APK文件，后端运行`keytool -printcert -jarfile <apk>`命令后，将签名信息显示在Web页面上，简化了签名检查流程。

## 快速开始

`git clone https://github.com/kk-game/show-apk-sign.git`

运行build.bat (windows平台)

## 缘起

为了给非技术人查看apk的签名

## 页面效果

<img src="./image/view.jpg">




