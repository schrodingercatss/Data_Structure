#ifndef SEQLIST_SEQLIST_H
#define SEQLIST_SEQLIST_H

typedef int ElemType; // 数据类型

const int InitSize = 100; // 表的初始长度

typedef struct {
    ElemType *data; //动态分配的数组指针
    int capacity, length; // 数组的最大容量和当前长度
}SeqList;

bool InitList(SeqList &L); // 初始化表，构造一个空的顺序表
int Length(SeqList L); // 求表长
int LocateElem(SeqList L, ElemType e); // 按值查找, 范围元素的位置
ElemType GetElem(SeqList L, int i); // 按位置查找，获取表中第i个元素的值
bool ListInsert(SeqList &L, int i, ElemType e); // 插入操作，在第i个位置插入元素e
bool ListDelete(SeqList &L, int i, ElemType &e); // 删除操作，删除第i个位置的元素
void PrintList(SeqList L); // 输出操作，打印整个顺序表
bool Empty(SeqList L); // 判断顺序表是否为空
bool DestroyList(SeqList &L); // 销毁操作，销毁顺序表，并释放申请的所有空间

#endif //SEQLIST_SEQLIST_H
