package main

/*
// Generated by rust2go. Please DO NOT edit this C part manually.

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct StringRef {
  const uint8_t *ptr;
  uintptr_t len;
} StringRef;

typedef struct ListRef {
  const void *ptr;
  uintptr_t len;
} ListRef;

typedef struct LoginResponseRef {
  bool succ;
  struct StringRef message;
  struct ListRef token;
} LoginResponseRef;

typedef struct UserRef {
  uint32_t id;
  struct StringRef name;
  uint8_t age;
} UserRef;

typedef struct LoginRequestRef {
  struct UserRef user;
  struct StringRef password;
} LoginRequestRef;

typedef struct FriendsListResponseRef {
  struct ListRef users;
} FriendsListResponseRef;

typedef struct FriendsListRequestRef {
  struct ListRef token;
  struct ListRef user_ids;
} FriendsListRequestRef;

typedef struct PMFriendRequestRef {
  uint32_t user_id;
  struct ListRef token;
  struct StringRef message;
} PMFriendRequestRef;

typedef struct PMFriendResponseRef {
  bool succ;
  struct StringRef message;
} PMFriendResponseRef;

typedef struct LogoutRequestRef {
  struct ListRef token;
  struct ListRef user_ids;
} LogoutRequestRef;

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void TestCall_ping_cb(const void *f_ptr, uintptr_t resp, const void *slot) {
((void (*)(uintptr_t, const void*))f_ptr)(resp, slot);
}

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void TestCall_login_cb(const void *f_ptr, struct LoginResponseRef resp, const void *slot) {
((void (*)(struct LoginResponseRef, const void*))f_ptr)(resp, slot);
}

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void TestCall_add_friends_cb(const void *f_ptr, struct FriendsListResponseRef resp, const void *slot) {
((void (*)(struct FriendsListResponseRef, const void*))f_ptr)(resp, slot);
}

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void TestCall_delete_friends_cb(const void *f_ptr, struct FriendsListResponseRef resp, const void *slot) {
((void (*)(struct FriendsListResponseRef, const void*))f_ptr)(resp, slot);
}

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void TestCall_pm_friend_cb(const void *f_ptr, struct PMFriendResponseRef resp, const void *slot) {
((void (*)(struct PMFriendResponseRef, const void*))f_ptr)(resp, slot);
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

var TestCallImpl TestCall

type TestCall interface {
	ping(n uint) uint
	login(req LoginRequest) LoginResponse
	logout(req User)
	add_friends(req FriendsListRequest) FriendsListResponse
	delete_friends(req FriendsListRequest) FriendsListResponse
	pm_friend(req PMFriendRequest) PMFriendResponse
}

//export CTestCall_ping
func CTestCall_ping(n C.uintptr_t, slot *C.void, cb *C.void) {
	resp := TestCallImpl.ping(newC_uintptr_t(n))
	resp_ref, buffer := cvt_ref(cntC_uintptr_t, refC_uintptr_t)(&resp)
	C.TestCall_ping_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
	runtime.KeepAlive(resp)
	runtime.KeepAlive(buffer)
}

//export CTestCall_login
func CTestCall_login(req C.LoginRequestRef, slot *C.void, cb *C.void) {
	resp := TestCallImpl.login(newLoginRequest(req))
	resp_ref, buffer := cvt_ref(cntLoginResponse, refLoginResponse)(&resp)
	C.TestCall_login_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
	runtime.KeepAlive(resp)
	runtime.KeepAlive(buffer)
}

//export CTestCall_logout
func CTestCall_logout(req C.UserRef) {
	TestCallImpl.logout(newUser(req))
}

//export CTestCall_add_friends
func CTestCall_add_friends(req C.FriendsListRequestRef, slot *C.void, cb *C.void) {
	_new_req := newFriendsListRequest(req)
	go func() {
		resp := TestCallImpl.add_friends(_new_req)
		resp_ref, buffer := cvt_ref(cntFriendsListResponse, refFriendsListResponse)(&resp)
		C.TestCall_add_friends_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
		runtime.KeepAlive(resp)
		runtime.KeepAlive(buffer)
	}()
}

//export CTestCall_delete_friends
func CTestCall_delete_friends(req C.FriendsListRequestRef, slot *C.void, cb *C.void) {
	_new_req := newFriendsListRequest(req)
	go func() {
		resp := TestCallImpl.delete_friends(_new_req)
		resp_ref, buffer := cvt_ref(cntFriendsListResponse, refFriendsListResponse)(&resp)
		C.TestCall_delete_friends_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
		runtime.KeepAlive(resp)
		runtime.KeepAlive(buffer)
	}()
}

//export CTestCall_pm_friend
func CTestCall_pm_friend(req C.PMFriendRequestRef, slot *C.void, cb *C.void) {
	_new_req := newPMFriendRequest(req)
	go func() {
		resp := TestCallImpl.pm_friend(_new_req)
		resp_ref, buffer := cvt_ref(cntPMFriendResponse, refPMFriendResponse)(&resp)
		C.TestCall_pm_friend_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
		runtime.KeepAlive(resp)
		runtime.KeepAlive(buffer)
	}()
}

func newString(s_ref C.StringRef) string {
	return unsafe.String((*byte)(unsafe.Pointer(s_ref.ptr)), s_ref.len)
}
func refString(s *string, _buffer *[]byte) C.StringRef {
	return C.StringRef{
		ptr: (*C.uint8_t)(unsafe.StringData(*s)),
		len: C.uintptr_t(len(*s)),
	}
}

func cntString(s *string, cnt *uint) []C.StringRef { return []C.StringRef{} }
func new_list_mapper[T1, T2 any](f func(T1) T2) func(C.ListRef) []T2 {
	return func(x C.ListRef) []T2 {
		input := unsafe.Slice((*T1)(unsafe.Pointer(x.ptr)), x.len)
		output := make([]T2, len(input))
		for i, v := range input {
			output[i] = f(v)
		}
		return output
	}
}
func new_list_mapper_primitive[T1, T2 any](f func(T1) T2) func(C.ListRef) []T2 {
	return func(x C.ListRef) []T2 {
		return unsafe.Slice((*T2)(unsafe.Pointer(x.ptr)), x.len)
	}
}

// only handle non-primitive type T
func cnt_list_mapper[T, R any](f func(s *T, cnt *uint) [0]R) func(s *[]T, cnt *uint) [0]C.ListRef {
	return func(s *[]T, cnt *uint) [0]C.ListRef {
		for _, v := range *s {
			f(&v, cnt)
		}
		*cnt += uint(len(*s)) * size_of[R]()
		return [0]C.ListRef{}
	}
}

// only handle primitive type T
func cnt_list_mapper_primitive[T, R any](f func(s *T, cnt *uint) [0]R) func(s *[]T, cnt *uint) [0]C.ListRef {
	return func(s *[]T, cnt *uint) [0]C.ListRef { return [0]C.ListRef{} }
}

// only handle non-primitive type T
func ref_list_mapper[T, R any](f func(s *T, buffer *[]byte) R) func(s *[]T, buffer *[]byte) C.ListRef {
	return func(s *[]T, buffer *[]byte) C.ListRef {
		if len(*buffer) == 0 {
			return C.ListRef{
				ptr: unsafe.Pointer(nil),
				len: C.uintptr_t(len(*s)),
			}
		}
		ret := C.ListRef{
			ptr: unsafe.Pointer(&(*buffer)[0]),
			len: C.uintptr_t(len(*s)),
		}
		children_bytes := int(size_of[R]()) * len(*s)
		children := (*buffer)[:children_bytes]
		*buffer = (*buffer)[children_bytes:]
		for _, v := range *s {
			child := f(&v, buffer)
			len := unsafe.Sizeof(child)
			copy(children, unsafe.Slice((*byte)(unsafe.Pointer(&child)), len))
			children = children[len:]
		}
		return ret
	}
}

// only handle primitive type T
func ref_list_mapper_primitive[T, R any](f func(s *T, buffer *[]byte) R) func(s *[]T, buffer *[]byte) C.ListRef {
	return func(s *[]T, buffer *[]byte) C.ListRef {
		if len(*s) == 0 {
			return C.ListRef{
				ptr: unsafe.Pointer(nil),
				len: C.uintptr_t(0),
			}
		}
		return C.ListRef{
			ptr: unsafe.Pointer(&(*s)[0]),
			len: C.uintptr_t(len(*s)),
		}
	}
}
func size_of[T any]() uint {
	var t T
	return uint(unsafe.Sizeof(t))
}
func cvt_ref[R, CR any](cnt_f func(s *R, cnt *uint) [0]CR, ref_f func(p *R, buffer *[]byte) CR) func(p *R) (CR, []byte) {
	return func(p *R) (CR, []byte) {
		var cnt uint
		cnt_f(p, &cnt)
		buffer := make([]byte, cnt)
		return ref_f(p, &buffer), buffer
	}
}

func newC_uint8_t(n C.uint8_t) uint8    { return uint8(n) }
func newC_uint16_t(n C.uint16_t) uint16 { return uint16(n) }
func newC_uint32_t(n C.uint32_t) uint32 { return uint32(n) }
func newC_uint64_t(n C.uint64_t) uint64 { return uint64(n) }
func newC_int8_t(n C.int8_t) int8       { return int8(n) }
func newC_int16_t(n C.int16_t) int16    { return int16(n) }
func newC_int32_t(n C.int32_t) int32    { return int32(n) }
func newC_int64_t(n C.int64_t) int64    { return int64(n) }
func newC_bool(n C.bool) bool           { return bool(n) }
func newC_uintptr_t(n C.uintptr_t) uint { return uint(n) }
func newC_intptr_t(n C.intptr_t) int    { return int(n) }
func newC_float(n C.float) float32      { return float32(n) }
func newC_double(n C.double) float64    { return float64(n) }

func cntC_uint8_t(s *uint8, cnt *uint) [0]C.uint8_t    { return [0]C.uint8_t{} }
func cntC_uint16_t(s *uint16, cnt *uint) [0]C.uint16_t { return [0]C.uint16_t{} }
func cntC_uint32_t(s *uint32, cnt *uint) [0]C.uint32_t { return [0]C.uint32_t{} }
func cntC_uint64_t(s *uint64, cnt *uint) [0]C.uint64_t { return [0]C.uint64_t{} }
func cntC_int8_t(s *int8, cnt *uint) [0]C.int8_t       { return [0]C.int8_t{} }
func cntC_int16_t(s *int16, cnt *uint) [0]C.int16_t    { return [0]C.int16_t{} }
func cntC_int32_t(s *int32, cnt *uint) [0]C.int32_t    { return [0]C.int32_t{} }
func cntC_int64_t(s *int64, cnt *uint) [0]C.int64_t    { return [0]C.int64_t{} }
func cntC_bool(s *bool, cnt *uint) [0]C.bool           { return [0]C.bool{} }
func cntC_uintptr_t(s *uint, cnt *uint) [0]C.uintptr_t { return [0]C.uintptr_t{} }
func cntC_intptr_t(s *int, cnt *uint) [0]C.intptr_t    { return [0]C.intptr_t{} }
func cntC_float(s *float32, cnt *uint) [0]C.float      { return [0]C.float{} }
func cntC_double(s *float64, cnt *uint) [0]C.double    { return [0]C.double{} }

func refC_uint8_t(p *uint8, buffer *[]byte) C.uint8_t    { return C.uint8_t(*p) }
func refC_uint16_t(p *uint16, buffer *[]byte) C.uint16_t { return C.uint16_t(*p) }
func refC_uint32_t(p *uint32, buffer *[]byte) C.uint32_t { return C.uint32_t(*p) }
func refC_uint64_t(p *uint64, buffer *[]byte) C.uint64_t { return C.uint64_t(*p) }
func refC_int8_t(p *int8, buffer *[]byte) C.int8_t       { return C.int8_t(*p) }
func refC_int16_t(p *int16, buffer *[]byte) C.int16_t    { return C.int16_t(*p) }
func refC_int32_t(p *int32, buffer *[]byte) C.int32_t    { return C.int32_t(*p) }
func refC_int64_t(p *int64, buffer *[]byte) C.int64_t    { return C.int64_t(*p) }
func refC_bool(p *bool, buffer *[]byte) C.bool           { return C.bool(*p) }
func refC_uintptr_t(p *uint, buffer *[]byte) C.uintptr_t { return C.uintptr_t(*p) }
func refC_intptr_t(p *int, buffer *[]byte) C.intptr_t    { return C.intptr_t(*p) }
func refC_float(p *float32, buffer *[]byte) C.float      { return C.float(*p) }
func refC_double(p *float64, buffer *[]byte) C.double    { return C.double(*p) }

type User struct {
	id   uint32
	name string
	age  uint8
}

func newUser(p C.UserRef) User {
	return User{
		id:   newC_uint32_t(p.id),
		name: newString(p.name),
		age:  newC_uint8_t(p.age),
	}
}
func cntUser(s *User, cnt *uint) [0]C.UserRef {
	return [0]C.UserRef{}
}
func refUser(p *User, buffer *[]byte) C.UserRef {
	return C.UserRef{
		id:   refC_uint32_t(&p.id, buffer),
		name: refString(&p.name, buffer),
		age:  refC_uint8_t(&p.age, buffer),
	}
}

type LoginRequest struct {
	user     User
	password string
}

func newLoginRequest(p C.LoginRequestRef) LoginRequest {
	return LoginRequest{
		user:     newUser(p.user),
		password: newString(p.password),
	}
}
func cntLoginRequest(s *LoginRequest, cnt *uint) [0]C.LoginRequestRef {
	return [0]C.LoginRequestRef{}
}
func refLoginRequest(p *LoginRequest, buffer *[]byte) C.LoginRequestRef {
	return C.LoginRequestRef{
		user:     refUser(&p.user, buffer),
		password: refString(&p.password, buffer),
	}
}

type LoginResponse struct {
	succ    bool
	message string
	token   []uint8
}

func newLoginResponse(p C.LoginResponseRef) LoginResponse {
	return LoginResponse{
		succ:    newC_bool(p.succ),
		message: newString(p.message),
		token:   new_list_mapper_primitive(newC_uint8_t)(p.token),
	}
}
func cntLoginResponse(s *LoginResponse, cnt *uint) [0]C.LoginResponseRef {
	return [0]C.LoginResponseRef{}
}
func refLoginResponse(p *LoginResponse, buffer *[]byte) C.LoginResponseRef {
	return C.LoginResponseRef{
		succ:    refC_bool(&p.succ, buffer),
		message: refString(&p.message, buffer),
		token:   ref_list_mapper_primitive(refC_uint8_t)(&p.token, buffer),
	}
}

type LogoutRequest struct {
	token    []uint8
	user_ids []uint32
}

func newLogoutRequest(p C.LogoutRequestRef) LogoutRequest {
	return LogoutRequest{
		token:    new_list_mapper_primitive(newC_uint8_t)(p.token),
		user_ids: new_list_mapper_primitive(newC_uint32_t)(p.user_ids),
	}
}
func cntLogoutRequest(s *LogoutRequest, cnt *uint) [0]C.LogoutRequestRef {
	return [0]C.LogoutRequestRef{}
}
func refLogoutRequest(p *LogoutRequest, buffer *[]byte) C.LogoutRequestRef {
	return C.LogoutRequestRef{
		token:    ref_list_mapper_primitive(refC_uint8_t)(&p.token, buffer),
		user_ids: ref_list_mapper_primitive(refC_uint32_t)(&p.user_ids, buffer),
	}
}

type FriendsListRequest struct {
	token    []uint8
	user_ids []uint32
}

func newFriendsListRequest(p C.FriendsListRequestRef) FriendsListRequest {
	return FriendsListRequest{
		token:    new_list_mapper_primitive(newC_uint8_t)(p.token),
		user_ids: new_list_mapper_primitive(newC_uint32_t)(p.user_ids),
	}
}
func cntFriendsListRequest(s *FriendsListRequest, cnt *uint) [0]C.FriendsListRequestRef {
	return [0]C.FriendsListRequestRef{}
}
func refFriendsListRequest(p *FriendsListRequest, buffer *[]byte) C.FriendsListRequestRef {
	return C.FriendsListRequestRef{
		token:    ref_list_mapper_primitive(refC_uint8_t)(&p.token, buffer),
		user_ids: ref_list_mapper_primitive(refC_uint32_t)(&p.user_ids, buffer),
	}
}

type FriendsListResponse struct {
	users []User
}

func newFriendsListResponse(p C.FriendsListResponseRef) FriendsListResponse {
	return FriendsListResponse{
		users: new_list_mapper(newUser)(p.users),
	}
}
func cntFriendsListResponse(s *FriendsListResponse, cnt *uint) [0]C.FriendsListResponseRef {
	cnt_list_mapper(cntUser)(&s.users, cnt)
	return [0]C.FriendsListResponseRef{}
}
func refFriendsListResponse(p *FriendsListResponse, buffer *[]byte) C.FriendsListResponseRef {
	return C.FriendsListResponseRef{
		users: ref_list_mapper(refUser)(&p.users, buffer),
	}
}

type PMFriendRequest struct {
	user_id uint32
	token   []uint8
	message string
}

func newPMFriendRequest(p C.PMFriendRequestRef) PMFriendRequest {
	return PMFriendRequest{
		user_id: newC_uint32_t(p.user_id),
		token:   new_list_mapper_primitive(newC_uint8_t)(p.token),
		message: newString(p.message),
	}
}
func cntPMFriendRequest(s *PMFriendRequest, cnt *uint) [0]C.PMFriendRequestRef {
	return [0]C.PMFriendRequestRef{}
}
func refPMFriendRequest(p *PMFriendRequest, buffer *[]byte) C.PMFriendRequestRef {
	return C.PMFriendRequestRef{
		user_id: refC_uint32_t(&p.user_id, buffer),
		token:   ref_list_mapper_primitive(refC_uint8_t)(&p.token, buffer),
		message: refString(&p.message, buffer),
	}
}

type PMFriendResponse struct {
	succ    bool
	message string
}

func newPMFriendResponse(p C.PMFriendResponseRef) PMFriendResponse {
	return PMFriendResponse{
		succ:    newC_bool(p.succ),
		message: newString(p.message),
	}
}
func cntPMFriendResponse(s *PMFriendResponse, cnt *uint) [0]C.PMFriendResponseRef {
	return [0]C.PMFriendResponseRef{}
}
func refPMFriendResponse(p *PMFriendResponse, buffer *[]byte) C.PMFriendResponseRef {
	return C.PMFriendResponseRef{
		succ:    refC_bool(&p.succ, buffer),
		message: refString(&p.message, buffer),
	}
}
func main() {}
