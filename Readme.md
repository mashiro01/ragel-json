# Simple Json Parser with ragel

最近工作上开始使用ragel来进行语法分析，这里记录下学习过程以及使用ragel完成一个简单的json-parser的流程

## RoadMap

### Component
1. [x] j_number
2. [x] j_string
3. [x] j_whitespace
4. [x] j_null
5. [x] j_boolen

6. [ ] j_object
7. [x] j_array
8. [x] j_value

### Unit Test
1. [ ] j_object
2. [ ] j_array
3. [ ] j_number
4. [ ] j_string
5. [ ] j_whitespace
6. [ ] j_null
7. [ ] j_boolen

### Build
1. [ ] Bazel / CMake / Make

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
