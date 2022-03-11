# Refactoring_IDEC_golang

Refactoring Improving the Design of Existing Code Second Edition

### 重构 改善既有代码的设计（第2版） Golang版本的代码示例

#### 第6章 第一组重构
- 6.1 提炼函数：
  - 重构前 [extract_function.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.1/extract_function.go)
  - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.1/refactoring/refactoring.go)
- 6.2 内联函数：
  - 重构前 [inline_function.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.2/inline_function.go)
  - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.2/refactoring/refactoring.go)
- 6.3 提炼变量：
  - 重构前 [extract_variable.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.3/extract_variable.go)
  - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.3/refactoring/refactoring.go)
- 6.4 内联变量：
  - 重构前 [inline_variable.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.4/inline_variable.go)
  - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.4/refactoring/refactoring.go)
- 6.5 改变函数声明：
  - 简单做法
    - 重构前 [change_function_declaration.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.5/simple/change_function_declaration.go)
    - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.5/simple/refactoring/refactoring.go)
  - 迁移式做法
    - 重构前 [change_function_declaration.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.5/migration/change_function_declaration.go)
    - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.5/migration/refactoring/refactoring.go)
- 6.9 函数组合成类：
  - 重构前 [combine_functions_into_class.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.9/combine_functions_into_class.go)
  - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter6/6.9/refactoring/refactoring.go)

#### 第7章 封装
- 7.3 以对象取代基本类型：
  - 重构前 [replace_primitive_with_obj.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter7/7.3/replace_primitive_with_obj.go)
  - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter7/7.3/refactoring/refactoring.go)

#### 第10章 简化条件逻辑
- 10.4 用多态处理变体逻辑：
  - 2
    - 重构前 [using_polymorphism_for_variation.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter10/10.4/2/using_polymorphism_for_variation.go)
    - 重构后 [refactoring.go](https://github.com/dgqypl/Refactoring_IDEC_golang/blob/main/chapter10/10.4/2/refactoring/refactoring.go)