#ifndef SEQLIST_SEQLIST_H
#define SEQLIST_SEQLIST_H

typedef int ElemType; // ��������

const int InitSize = 100; // ��ĳ�ʼ����

typedef struct {
    ElemType *data; //��̬���������ָ��
    int capacity, length; // �������������͵�ǰ����
}SeqList;

bool InitList(SeqList &L); // ��ʼ��������һ���յ�˳���
int Length(SeqList L); // ���
int LocateElem(SeqList L, ElemType e); // ��ֵ����, ��ΧԪ�ص�λ��
ElemType GetElem(SeqList L, int i); // ��λ�ò��ң���ȡ���е�i��Ԫ�ص�ֵ
bool ListInsert(SeqList &L, int i, ElemType e); // ����������ڵ�i��λ�ò���Ԫ��e
bool ListDelete(SeqList &L, int i, ElemType &e); // ɾ��������ɾ����i��λ�õ�Ԫ��
void PrintList(SeqList L); // �����������ӡ����˳���
bool Empty(SeqList L); // �ж�˳����Ƿ�Ϊ��
bool DestroyList(SeqList &L); // ���ٲ���������˳������ͷ���������пռ�

#endif //SEQLIST_SEQLIST_H
