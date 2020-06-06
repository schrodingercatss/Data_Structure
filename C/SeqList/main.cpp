#include <cstdio>
#include <cstdlib>
#include "SeqList.h"
using namespace std;

int main() {
    SeqList array;
    InitList(array);

    // 插入10个元素
    for(int i = 0; i < 10; ++i) {
        ListInsert(array, i + 1, i + 1);
    }

    // 表长
    printf("表长为：%d\n", Length(array));

    // 查找值为8的元素的位置
    printf("值为8元素的位置是：%d\n", LocateElem(array, 8));

    // 获取第6个位置上的元素
    printf("第6个位置上的元素是：%d\n", GetElem(array, 6));

    // 输出表
    PrintList(array);

    // 删除第5个位置上的元素
    ElemType e;
    ListDelete(array, 5, e);
    printf("第5个位置上的元素值为%d\n", e);

    PrintList(array);

    // 判断是否为空表
    printf("是否为空表：%d\n", Empty(array));

    // 销毁整个顺序表
    DestroyList(array);

    printf("是否为空表：%d\n", Empty(array));
}

/*
输出结果：
表长为：10
值为8元素的位置是：8
第6个位置上的元素是：6
1 2 3 4 5 6 7 8 9 10
第5个位置上的元素值为5
1 2 3 4 6 7 8 9 10
是否为空表：0
是否为空表：1
*/