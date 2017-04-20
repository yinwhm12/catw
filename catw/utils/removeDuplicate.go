package utils

import "reflect"

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

//func ArrayDuplicate(a []int) (array []int) {
//	va := reflect.ValueOf(a)
//	for i := 0; i < va.Len(); i++ {
//		if i > 0 && reflect.DeepEqual(va.Index(i-1).Int(), va.Index(i).Int()) {
//			continue
//		}
//		//ret = append(ret, va.Index(i).Interface())
//		array[i] = va.Index(i)
//	}
//	return array
//}