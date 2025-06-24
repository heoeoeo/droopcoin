package elliptic_curve

import (
	"fmt"
)

/*
	유한체는 유한한 수의 원소를 가지는 집합으로, 일반적으로 정수의 집합을 모듈로 연산을 통해 정의
	유한체의 원소는 필드 엘리먼트(Field Element)라고 하며, 이들은 특정 오더(order)를 가지며, 해당 오더에 대한 덧셈과 곱셈 연산이 정의
	필드 엘리먼트는 유한체의 원소로, 특정 오더를 가지며, 해당 오더에 대한 덧셈과 곱셈 연산이 정의
*/

type FieldElement struct {
	order uint64 // 필드 오더 / modulus / 유한체의 크기 / 일반적으로 소수의 거듭제곱
	num   uint64 // value of the given element in the field / 실제값 / 0 <= num < order
}

// order는 소수(prime number)
func NewFieldElement(order uint64, num uint64) *FieldElement {
	if num >= order {
		err := fmt.Sprintf("Num not in the range of 0 to %d\nNum의 값은 0부터 %d까지입니다.", order-1, order-1)
		panic(err)
	}

	return &FieldElement{
		order: order,
		num:   num,
	}
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement{order: %d, num: %d}\n", f.order, f.num)
}

/*
두 FieldElment의 order와 num을 각각 비교하여
같은 유한체에 속하며 같은 원소인지 판단.
*/
func (f *FieldElement) EqualTo(other *FieldElement) bool {
	return f.order == other.order && f.num == other.num
}

/*
 */
func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	if f.order != other.order {
		panic("add need to do on the field element with the same order\n더하기는 같은 오더의 필드 엘리먼트에서만 수행해야 합니다")
	}
	return NewFieldElement(f.order, (f.num+other.num)%f.order)
}

/*
덧셈 역원을 구하는 함수
a + (-a) = 0 mod p
예) 3의 역원 in mod 7 : 7 - 3 = 4 -> 3 + 4 = 0 = 0 mod 7
*/
func (f *FieldElement) Negative() *FieldElement {
	// b, (a+b) % order = 0, b = order - a, (a+b) % order => (a + order -a) => order => order % order = 0
	return NewFieldElement(f.order, f.order-f.num)
}

/*
other의 역원을 구한 후 f와 더함
*/
func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	// a, b element of the finite set, c = a - b, given b how can we find c, (b + c) % order = a, a - b => (a + (-b)) % order

	return f.Add(other.Negative())
}
