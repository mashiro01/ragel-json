# Simple Json Parser with ragel

最近工作上开始使用ragel来进行语法分析，这里记录下学习过程以及使用ragel完成一个简单的json-parser的流程

## RoadMap

### Component
[x] j_number
[x] j_string
[x] j_whitespace
[x] j_null
[x] j_boolen

[ ] j_object
[x] j_array
[x] j_value

### Unit Test
[ ] j_object
[ ] j_array
[ ] j_number
[ ] j_string
[ ] j_whitespace
[ ] j_null
[ ] j_boolen

### Build
[ ] Bazel / CMake / Make

### TODO

1. [x] 完成单个组件的编写
2. [x] 完成单组件的单元测试
3. [x] 组合单组件为Parser整体
4. [ ] 使用CMake等构建工具完成具体的解析构建

## Summary

Ragel在语法分析的过程中确实是一个趁手的工具，能以原生语言的速度完成字符串的分块解析，可以看作是regexpp，值得深入了解

## Reference

* [JSON Parsing from Scratch in Haskell](https://abhinavsarkar.net/posts/json-parsing-from-scratch-in-haskell/)
* [The JavaScript Object Notation (JSON) Data Interchange Format](https://datatracker.ietf.org/doc/html/rfc7159#section-2)
* [Introducing JSON](https://www.json.org/json-en.html)

## Extend
<!-- 具有深度的知识还是要外国原文 -->

<!-- CMake -->
* [An Introduction to Modern CMake](https://cliutils.gitlab.io/modern-cmake/)
* [Effective Modern CMake](https://gist.github.com/mbinna/c61dbb39bca0e4fb7d1f73b0d66a4fd1)
