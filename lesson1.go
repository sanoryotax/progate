package main

func main() {

    n := 100
    // n����10��������n�ɑ�����܂��傤
    n -= 10
    
    println(n)
    
    
    m := 10
    // m��1�𑫂���m�ɑ�����܂��傤
    m ++
    
    println(m)


    n := 3
    switch n {
        // case��ǉ����An��0�̏ꍇ�A"���ł�"�Əo�͂��Ă�������
        case 0:
            println("���ł�")
        // case��ǉ����An��1�܂���2�̏ꍇ�A"�g�ł�"�Əo�͂��Ă�������
        case 1, 2:
            println("�g�ł�")
        // case��ǉ����An��3�܂���4�̏ꍇ�A"���g�ł�"�Əo�͂��Ă�������
        case 3, 4:
            println("���g�ł�")
        // case��ǉ����An��5�̏ꍇ�A"��g�ł�"�Əo�͂��Ă�������
        case 5:
            println("��g�ł�")
    }
}
