Algorithm 
[![Circle CI](https://circleci.com/gh/ThreeBearsDan/Algorithm.svg?style=shield)](https://circleci.com/gh/ThreeBearsDan/Algorithm)
-----------------------
Practise one algorithm one day.

## 基本概念
>一个排序算法是一种将一串数据依照特定排序方式进行排列的一种算法，最常用到的排序方式是数值顺序以及字典顺序。    
排序算法的输出必须遵守下列两个原则：
* 输出结果为递增序列（递增是针对所需的排序顺序而言）
* 输出结果是原输入的一种排列或是重组


### [Quick Sort]
#### 快速排序的基本思想：
1. 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分数据比另一部分数据都要小。
2. 按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

### [Bubble Sort]
#### 冒泡排序基本思想：
1. 重复地走访排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
2. 走访数列的工作重复地进行直到再没有需要交换的元素。

### [Select Sort]
#### 简单选择排序基本思想：
1. 从未排序的序列中找到最大（最小）值，放到序列的末尾（开头）。
2. 重复步骤1，从未排序序列中找到最大（最小）值，放到已排序列的前面（后面）。

### [Insertion Sort]
#### 插入排序基本思想：
1. 通过构建有序序列，将未排序数据，在已排序序列中从后向前扫描，找到相应的位置并插入。

### [Shell Sort]
#### 希尔排序基本思想：
希尔排序是对直接插入排序的一种改进，增大了元素移动的步长。
1. 选定一个步长，一般为序列长度的一半。
2. 按指定的步长将序列分组，使得每组以直接插入排序的方式排序。
3. 逐渐缩短步长，重复1，2两步，使得最终步长为1。

### [Heap Sort]
#### 堆排序基本思想：
1. 将堆的末端子节点进行调整，使得子节点永远小于父节点。
2. 将堆的所有节点重新排序。
3. 移除跟节点，并作最大堆调整递归运算。

### [Merge Sort]
#### 归并排序基本思想：
1. 将序列相邻的两个元素进行归并操作，形成ceil(n/2)个序列，每个序列包含2或1个元素。
2. 如果此时序列数不为1则将上述序列再次归并，形成ceil(n/4)个序列，每个序列包含4或3个元素。
3. 重复步骤2，直到所有的元素排序完毕，即序列数为1。