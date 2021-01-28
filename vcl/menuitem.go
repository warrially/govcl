
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TMenuItem struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewMenuItem(owner IComponent) *TMenuItem {
    m := new(TMenuItem)
    m.instance = MenuItem_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    // 不是TComponent应该是可以考虑加上的
    // runtime.SetFinalizer(m, (*TMenuItem).Free)
    return m
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsMenuItem(obj interface{}) *TMenuItem {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMenuItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsMenuItem.
func MenuItemFromInst(inst uintptr) *TMenuItem {
    return AsMenuItem(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsMenuItem.
func MenuItemFromObj(obj IObject) *TMenuItem {
    return AsMenuItem(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMenuItem.
func MenuItemFromUnsafePointer(ptr unsafe.Pointer) *TMenuItem {
    return AsMenuItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (m *TMenuItem) Free() {
    if m.instance != 0 {
        MenuItem_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (m *TMenuItem) Instance() uintptr {
    return m.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (m *TMenuItem) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (m *TMenuItem) IsValid() bool {
    return m.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (m *TMenuItem) Is() TIs {
    return TIs(m.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (m *TMenuItem) As() TAs {
//    return TAs(m.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TMenuItemClass() TClass {
    return MenuItem_StaticClassType()
}

func (m *TMenuItem) Insert(Index int32, Item IComponent) {
    MenuItem_Insert(m.instance, Index , CheckPtr(Item))
}

func (m *TMenuItem) Delete(Index int32) {
    MenuItem_Delete(m.instance, Index)
}

// 清除。
func (m *TMenuItem) Clear() {
    MenuItem_Clear(m.instance)
}

// 单击。
func (m *TMenuItem) Click() {
    MenuItem_Click(m.instance)
}

func (m *TMenuItem) IndexOf(Item IComponent) int32 {
    return MenuItem_IndexOf(m.instance, CheckPtr(Item))
}

// 是否有父容器。
//
// Is there a parent container.
func (m *TMenuItem) HasParent() bool {
    return MenuItem_HasParent(m.instance)
}

func (m *TMenuItem) Add(Item IComponent) {
    MenuItem_Add(m.instance, CheckPtr(Item))
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (m *TMenuItem) FindComponent(AName string) *TComponent {
    return AsComponent(MenuItem_FindComponent(m.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (m *TMenuItem) GetNamePath() string {
    return MenuItem_GetNamePath(m.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (m *TMenuItem) Assign(Source IObject) {
    MenuItem_Assign(m.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (m *TMenuItem) ClassType() TClass {
    return MenuItem_ClassType(m.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (m *TMenuItem) ClassName() string {
    return MenuItem_ClassName(m.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (m *TMenuItem) InstanceSize() int32 {
    return MenuItem_InstanceSize(m.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (m *TMenuItem) InheritsFrom(AClass TClass) bool {
    return MenuItem_InheritsFrom(m.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (m *TMenuItem) Equals(Obj IObject) bool {
    return MenuItem_Equals(m.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (m *TMenuItem) GetHashCode() int32 {
    return MenuItem_GetHashCode(m.instance)
}

// 文本类信息。
//
// Text information.
func (m *TMenuItem) ToString() string {
    return MenuItem_ToString(m.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (m *TMenuItem) Handle() HMENU {
    return MenuItem_GetHandle(m.instance)
}

func (m *TMenuItem) Count() int32 {
    return MenuItem_GetCount(m.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (m *TMenuItem) Parent() *TMenuItem {
    return AsMenuItem(MenuItem_GetParent(m.instance))
}

func (m *TMenuItem) Action() *TAction {
    return AsAction(MenuItem_GetAction(m.instance))
}

func (m *TMenuItem) SetAction(value IComponent) {
    MenuItem_SetAction(m.instance, CheckPtr(value))
}

func (m *TMenuItem) AutoCheck() bool {
    return MenuItem_GetAutoCheck(m.instance)
}

func (m *TMenuItem) SetAutoCheck(value bool) {
    MenuItem_SetAutoCheck(m.instance, value)
}

func (m *TMenuItem) Bitmap() *TBitmap {
    return AsBitmap(MenuItem_GetBitmap(m.instance))
}

func (m *TMenuItem) SetBitmap(value *TBitmap) {
    MenuItem_SetBitmap(m.instance, CheckPtr(value))
}

// 获取控件标题。
//
// Get the control title.
func (m *TMenuItem) Caption() string {
    return MenuItem_GetCaption(m.instance)
}

// 设置控件标题。
//
// Set the control title.
func (m *TMenuItem) SetCaption(value string) {
    MenuItem_SetCaption(m.instance, value)
}

// 获取是否选中。
func (m *TMenuItem) Checked() bool {
    return MenuItem_GetChecked(m.instance)
}

// 设置是否选中。
func (m *TMenuItem) SetChecked(value bool) {
    MenuItem_SetChecked(m.instance, value)
}

func (m *TMenuItem) Default() bool {
    return MenuItem_GetDefault(m.instance)
}

func (m *TMenuItem) SetDefault(value bool) {
    MenuItem_SetDefault(m.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (m *TMenuItem) Enabled() bool {
    return MenuItem_GetEnabled(m.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (m *TMenuItem) SetEnabled(value bool) {
    MenuItem_SetEnabled(m.instance, value)
}

// 获取团组索引。
func (m *TMenuItem) GroupIndex() uint8 {
    return MenuItem_GetGroupIndex(m.instance)
}

// 设置团组索引。
func (m *TMenuItem) SetGroupIndex(value uint8) {
    MenuItem_SetGroupIndex(m.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (m *TMenuItem) Hint() string {
    return MenuItem_GetHint(m.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (m *TMenuItem) SetHint(value string) {
    MenuItem_SetHint(m.instance, value)
}

// 获取图像在images中的索引。
func (m *TMenuItem) ImageIndex() int32 {
    return MenuItem_GetImageIndex(m.instance)
}

// 设置图像在images中的索引。
func (m *TMenuItem) SetImageIndex(value int32) {
    MenuItem_SetImageIndex(m.instance, value)
}

func (m *TMenuItem) RadioItem() bool {
    return MenuItem_GetRadioItem(m.instance)
}

func (m *TMenuItem) SetRadioItem(value bool) {
    MenuItem_SetRadioItem(m.instance, value)
}

// 获取快捷键。
func (m *TMenuItem) ShortCut() TShortCut {
    return MenuItem_GetShortCut(m.instance)
}

// 设置快捷键。
func (m *TMenuItem) SetShortCut(value TShortCut) {
    MenuItem_SetShortCut(m.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (m *TMenuItem) Visible() bool {
    return MenuItem_GetVisible(m.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (m *TMenuItem) SetVisible(value bool) {
    MenuItem_SetVisible(m.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
    MenuItem_SetOnClick(m.instance, fn)
}

func (m *TMenuItem) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
    MenuItem_SetOnMeasureItem(m.instance, fn)
}

// 获取组件总数。
//
// Get the total number of components.
func (m *TMenuItem) ComponentCount() int32 {
    return MenuItem_GetComponentCount(m.instance)
}

// 获取组件索引。
//
// Get component index.
func (m *TMenuItem) ComponentIndex() int32 {
    return MenuItem_GetComponentIndex(m.instance)
}

// 设置组件索引。
//
// Set component index.
func (m *TMenuItem) SetComponentIndex(value int32) {
    MenuItem_SetComponentIndex(m.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (m *TMenuItem) Owner() *TComponent {
    return AsComponent(MenuItem_GetOwner(m.instance))
}

// 获取组件名称。
//
// Get the component name.
func (m *TMenuItem) Name() string {
    return MenuItem_GetName(m.instance)
}

// 设置组件名称。
//
// Set the component name.
func (m *TMenuItem) SetName(value string) {
    MenuItem_SetName(m.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (m *TMenuItem) Tag() int {
    return MenuItem_GetTag(m.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (m *TMenuItem) SetTag(value int) {
    MenuItem_SetTag(m.instance, value)
}

func (m *TMenuItem) Items(Index int32) *TMenuItem {
    return AsMenuItem(MenuItem_GetItems(m.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (m *TMenuItem) Components(AIndex int32) *TComponent {
    return AsComponent(MenuItem_GetComponents(m.instance, AIndex))
}

