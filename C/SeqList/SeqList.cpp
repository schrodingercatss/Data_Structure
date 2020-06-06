#include <stdio.h>
#include <stdlib.h>
#include "SeqList.h"

// 初始化顺序表
bool InitList(SeqList &L) {
    L.data = (ElemType*)malloc(sizeof(ElemType) * InitSize);
    if(L.data == NULL) return false;

    L.capacity = InitSize;
    L.length = 0;
    return true;
}


// 求表长
int Length(SeqList L) {
    return L.length;
}


// 按值查找, 范围元素的位置
int LocateElem(SeqList L, ElemType e) {
    for (int i = 0; i < L.length; ++i) {
        if (L.data[i] == e) {
            return i + 1;
        }
    }
    return -1;
}

// 按位置查找,C语言不支持多返回值，这里约定错误返回-1。
ElemType GetElem(SeqList L, int i) {
    if(i < 1 || i > L.length) return -1;
    return L.data[i - 1];
}


// 插入操作
bool ListInsert(SeqList &L, int i, ElemType e) {
    if(i < 1 || i > L.length + 1) return false; // 越界
    if(L.length >= L.capacity) { // 容量不足，动态扩容
        printf("触发扩容机制\n");
        ElemType *tmp = (ElemType*)malloc(sizeof(ElemType) * (L.capacity * 2)); // 扩容为原来的2倍
        for(int k = 0; k < L.length; ++k) tmp[k] = L.data[k];
        free(L.data);
        L.data = tmp;
        L.capacity *= 2;
    }
    for(int k = L.length; k >= i; --k) {
        L.data[k] = L.data[k - 1];
    }
    L.data[i - 1] = e;
    ++L.length;
    return true;
}

// 删除操作
bool ListDelete(SeqList &L, int i, ElemType &e) {
    if(i < 1 || i > L.length) return false;
    e = L.data[i - 1];
    for(int k = i; k < L.length; ++k) L.data[k - 1] = L.data[k];
    --L.length;
    return true;
}

// 输出操作
void PrintList(SeqList L) {
    for(int i = 0; i < L.length; ++i)
        printf("%d ", L.data[i]);
    puts("");
}

// 判断顺序表是否为空
bool Empty(SeqList L) {
    return L.length == 0;
}

// 销毁操作，销毁线性表，并释放申请的所有空间
bool DestroyList(SeqList &L) {
    free(L.data);
    L.length = 0;
    L.capacity = 0;
    return true;
}
