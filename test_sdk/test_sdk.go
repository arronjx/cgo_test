/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: testsdk.i

package test_sdk

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_test_sdk_68a13c115c613b51(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_test_sdk_68a13c115c613b51(swig_intgo arg1);
extern void _wrap_swig_m_test_set_test_sdk_68a13c115c613b51(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_swig_m_test_get_test_sdk_68a13c115c613b51(uintptr_t arg1);
extern uintptr_t _wrap_new_swig_test_sdk_68a13c115c613b51(void);
extern void _wrap_delete_swig_test_sdk_68a13c115c613b51(uintptr_t arg1);
extern void _wrap_swig_SetVector_test_sdk_68a13c115c613b51(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_swig_PrintHello_test_sdk_68a13c115c613b51(uintptr_t arg1);
#undef intgo
*/
import "C"

import "syscall"
import "unsafe"
import "sync"


type _ syscall.Sockaddr




type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_test_sdk_68a13c115c613b51(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_test_sdk_68a13c115c613b51(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrSwig uintptr

func (p SwigcptrSwig) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSwig) SwigIsSwig() {
}

func (arg1 SwigcptrSwig) SetM_test(arg2 Vector_Sl_int_Sg_) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_swig_m_test_set_test_sdk_68a13c115c613b51(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrSwig) GetM_test() (_swig_ret Vector_Sl_int_Sg_) {
	var swig_r Vector_Sl_int_Sg_
	_swig_i_0 := arg1
	swig_r = (Vector_Sl_int_Sg_)(SwigcptrVector_Sl_int_Sg_(C._wrap_swig_m_test_get_test_sdk_68a13c115c613b51(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewSwig() (_swig_ret Swig) {
	var swig_r Swig
	swig_r = (Swig)(SwigcptrSwig(C._wrap_new_swig_test_sdk_68a13c115c613b51()))
	return swig_r
}

func DeleteSwig(arg1 Swig) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_swig_test_sdk_68a13c115c613b51(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrSwig) SetVector(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_swig_SetVector_test_sdk_68a13c115c613b51(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrSwig) PrintHello() {
	_swig_i_0 := arg1
	C._wrap_swig_PrintHello_test_sdk_68a13c115c613b51(C.uintptr_t(_swig_i_0))
}

type Swig interface {
	Swigcptr() uintptr
	SwigIsSwig()
	SetM_test(arg2 Vector_Sl_int_Sg_)
	GetM_test() (_swig_ret Vector_Sl_int_Sg_)
	SetVector(arg2 int)
	PrintHello()
}


type SwigcptrVector_Sl_int_Sg_ uintptr
type Vector_Sl_int_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVector_Sl_int_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}



var swigDirectorTrack struct {
	sync.Mutex
	m map[int]interface{}
	c int
}

func swigDirectorAdd(v interface{}) int {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m == nil {
		swigDirectorTrack.m = make(map[int]interface{})
	}
	swigDirectorTrack.c++
	ret := swigDirectorTrack.c
	swigDirectorTrack.m[ret] = v
	return ret
}

func swigDirectorLookup(c int) interface{} {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	ret := swigDirectorTrack.m[c]
	if ret == nil {
		panic("C++ director pointer not found (possible	use-after-free)")
	}
	return ret
}

func swigDirectorDelete(c int) {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m[c] == nil {
		if c > swigDirectorTrack.c {
			panic("C++ director pointer invalid (possible memory corruption")
		} else {
			panic("C++ director pointer not found (possible use-after-free)")
		}
	}
	delete(swigDirectorTrack.m, c)
}


