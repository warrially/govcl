//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/types"
)

func Form_Create2(owner uintptr, initScale bool) uintptr {
	ret, _, _ := form_Create2.Call(owner, GoBoolToDBool(initScale))
	return ret
}

func Form_EnabledMaximize(obj uintptr, val bool) {
	form_EnabledMaximize.Call(obj, GoBoolToDBool(val))
}

func Form_EnabledMinimize(obj uintptr, val bool) {
	form_EnabledMinimize.Call(obj, GoBoolToDBool(val))
}

func Form_EnabledSystemMenu(obj uintptr, val bool) {
	form_EnabledSystemMenu.Call(obj, GoBoolToDBool(val))
}

func Form_SetOnDestroy(obj uintptr, fn interface{}) {
	form_SetOnDestroy.Call(obj, addEventToMap(fn))
}

func Form_SetOnConstrainedResize(obj uintptr, fn interface{}) {
	form_SetOnConstrainedResize.Call(obj, addEventToMap(fn))
}

func Form_SetOnWndProc(obj uintptr, fn interface{}) {
	form_SetOnWndProc.Call(obj, addMessageEventToMap(fn))
}

func Form_ShowInTaskBar(obj uintptr) types.TShowInTaskbar {
	r, _, _ := form_ShowInTaskBar.Call(obj)
	return types.TShowInTaskbar(r)
}

// 下面两个函数放在Application下面吧，直接调用实例类

func Form_ScaleForPPI(obj uintptr, newPPI int32) {
	form_ScaleForPPI.Call(obj, uintptr(newPPI))
}

func Form_ScaleControlsForDpi(obj uintptr, newPPI int32) {
	form_ScaleControlsForDpi.Call(obj, uintptr(newPPI))
}

func Form_ScaleForCurrentDpi(obj uintptr) {
	form_ScaleForCurrentDpi.Call(obj)
}

func Form_InheritedWndProc(obj uintptr, msg *types.TMessage) {
	form_InheritedWndProc.Call(obj, uintptr(unsafe.Pointer(msg)))
}
