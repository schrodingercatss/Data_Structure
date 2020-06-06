#include <cstdio>
#include <cstdlib>
#include "SeqList.h"
using namespace std;

int main() {
    SeqList array;
    InitList(array);

    // ����10��Ԫ��
    for(int i = 0; i < 10; ++i) {
        ListInsert(array, i + 1, i + 1);
    }

    // ��
    printf("��Ϊ��%d\n", Length(array));

    // ����ֵΪ8��Ԫ�ص�λ��
    printf("ֵΪ8Ԫ�ص�λ���ǣ�%d\n", LocateElem(array, 8));

    // ��ȡ��6��λ���ϵ�Ԫ��
    printf("��6��λ���ϵ�Ԫ���ǣ�%d\n", GetElem(array, 6));

    // �����
    PrintList(array);

    // ɾ����5��λ���ϵ�Ԫ��
    ElemType e;
    ListDelete(array, 5, e);
    printf("��5��λ���ϵ�Ԫ��ֵΪ%d\n", e);

    PrintList(array);

    // �ж��Ƿ�Ϊ�ձ�
    printf("�Ƿ�Ϊ�ձ�%d\n", Empty(array));

    // ��������˳���
    DestroyList(array);

    printf("�Ƿ�Ϊ�ձ�%d\n", Empty(array));
}

/*
��������
��Ϊ��10
ֵΪ8Ԫ�ص�λ���ǣ�8
��6��λ���ϵ�Ԫ���ǣ�6
1 2 3 4 5 6 7 8 9 10
��5��λ���ϵ�Ԫ��ֵΪ5
1 2 3 4 6 7 8 9 10
�Ƿ�Ϊ�ձ�0
�Ƿ�Ϊ�ձ�1
*/