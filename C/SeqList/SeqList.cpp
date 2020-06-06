#include <stdio.h>
#include <stdlib.h>
#include "SeqList.h"

// ��ʼ��˳���
bool InitList(SeqList &L) {
    L.data = (ElemType*)malloc(sizeof(ElemType) * InitSize);
    if(L.data == NULL) return false;

    L.capacity = InitSize;
    L.length = 0;
    return true;
}


// ���
int Length(SeqList L) {
    return L.length;
}


// ��ֵ����, ��ΧԪ�ص�λ��
int LocateElem(SeqList L, ElemType e) {
    for (int i = 0; i < L.length; ++i) {
        if (L.data[i] == e) {
            return i + 1;
        }
    }
    return -1;
}

// ��λ�ò���,C���Բ�֧�ֶ෵��ֵ������Լ�����󷵻�-1��
ElemType GetElem(SeqList L, int i) {
    if(i < 1 || i > L.length) return -1;
    return L.data[i - 1];
}


// �������
bool ListInsert(SeqList &L, int i, ElemType e) {
    if(i < 1 || i > L.length + 1) return false; // Խ��
    if(L.length >= L.capacity) { // �������㣬��̬����
        printf("�������ݻ���\n");
        ElemType *tmp = (ElemType*)malloc(sizeof(ElemType) * (L.capacity * 2)); // ����Ϊԭ����2��
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

// ɾ������
bool ListDelete(SeqList &L, int i, ElemType &e) {
    if(i < 1 || i > L.length) return false;
    e = L.data[i - 1];
    for(int k = i; k < L.length; ++k) L.data[k - 1] = L.data[k];
    --L.length;
    return true;
}

// �������
void PrintList(SeqList L) {
    for(int i = 0; i < L.length; ++i)
        printf("%d ", L.data[i]);
    puts("");
}

// �ж�˳����Ƿ�Ϊ��
bool Empty(SeqList L) {
    return L.length == 0;
}

// ���ٲ������������Ա����ͷ���������пռ�
bool DestroyList(SeqList &L) {
    free(L.data);
    L.length = 0;
    L.capacity = 0;
    return true;
}
