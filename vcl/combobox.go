
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

type TComboBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewComboBox(owner IComponent) *TComboBox {
    c := new(TComboBox)
    c.instance = ComboBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(c, (*TComboBox).Free)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsComboBox(obj interface{}) *TComboBox {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TComboBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsComboBox.
func ComboBoxFromInst(inst uintptr) *TComboBox {
    return AsComboBox(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsComboBox.
func ComboBoxFromObj(obj IObject) *TComboBox {
    return AsComboBox(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsComboBox.
func ComboBoxFromUnsafePointer(ptr unsafe.Pointer) *TComboBox {
    return AsComboBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TComboBox) Free() {
    if c.instance != 0 {
        ComboBox_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComboBox) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComboBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComboBox) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TComboBox) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TComboBox) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComboBoxClass() TClass {
    return ComboBox_StaticClassType()
}

func (c *TComboBox) AddItem(Item string, AObject IObject) {
    ComboBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// CN: 清除。
// EN: .
func (c *TComboBox) Clear() {
    ComboBox_Clear(c.instance)
}

// CN: 清除选择。
// EN: .
func (c *TComboBox) ClearSelection() {
    ComboBox_ClearSelection(c.instance)
}

func (c *TComboBox) DeleteSelected() {
    ComboBox_DeleteSelected(c.instance)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TComboBox) Focused() bool {
    return ComboBox_Focused(c.instance)
}

// CN: 全选。
// EN: .
func (c *TComboBox) SelectAll() {
    ComboBox_SelectAll(c.instance)
}

// CN: 是否可以获得焦点。
// EN: .
func (c *TComboBox) CanFocus() bool {
    return ComboBox_CanFocus(c.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TComboBox) ContainsControl(Control IControl) bool {
    return ComboBox_ContainsControl(c.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TComboBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ComboBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TComboBox) DisableAlign() {
    ComboBox_DisableAlign(c.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TComboBox) EnableAlign() {
    ComboBox_EnableAlign(c.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (c *TComboBox) FindChildControl(ControlName string) *TControl {
    return AsControl(ComboBox_FindChildControl(c.instance, ControlName))
}

func (c *TComboBox) FlipChildren(AllLevels bool) {
    ComboBox_FlipChildren(c.instance, AllLevels)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TComboBox) HandleAllocated() bool {
    return ComboBox_HandleAllocated(c.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (c *TComboBox) InsertControl(AControl IControl) {
    ComboBox_InsertControl(c.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (c *TComboBox) Invalidate() {
    ComboBox_Invalidate(c.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (c *TComboBox) RemoveControl(AControl IControl) {
    ComboBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (c *TComboBox) Realign() {
    ComboBox_Realign(c.instance)
}

// CN: 重绘。
// EN: Repaint.
func (c *TComboBox) Repaint() {
    ComboBox_Repaint(c.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (c *TComboBox) ScaleBy(M int32, D int32) {
    ComboBox_ScaleBy(c.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TComboBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ComboBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TComboBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ComboBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TComboBox) SetFocus() {
    ComboBox_SetFocus(c.instance)
}

// CN: 控件更新。
// EN: Update.
func (c *TComboBox) Update() {
    ComboBox_Update(c.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TComboBox) BringToFront() {
    ComboBox_BringToFront(c.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TComboBox) ClientToScreen(Point TPoint) TPoint {
    return ComboBox_ClientToScreen(c.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TComboBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ComboBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TComboBox) Dragging() bool {
    return ComboBox_Dragging(c.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TComboBox) HasParent() bool {
    return ComboBox_HasParent(c.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (c *TComboBox) Hide() {
    ComboBox_Hide(c.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (c *TComboBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ComboBox_Perform(c.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (c *TComboBox) Refresh() {
    ComboBox_Refresh(c.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TComboBox) ScreenToClient(Point TPoint) TPoint {
    return ComboBox_ScreenToClient(c.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TComboBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ComboBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TComboBox) SendToBack() {
    ComboBox_SendToBack(c.instance)
}

// CN: 显示控件。
// EN: Show control.
func (c *TComboBox) Show() {
    ComboBox_Show(c.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TComboBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ComboBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TComboBox) GetTextLen() int32 {
    return ComboBox_GetTextLen(c.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TComboBox) SetTextBuf(Buffer string) {
    ComboBox_SetTextBuf(c.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TComboBox) FindComponent(AName string) *TComponent {
    return AsComponent(ComboBox_FindComponent(c.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComboBox) GetNamePath() string {
    return ComboBox_GetNamePath(c.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComboBox) Assign(Source IObject) {
    ComboBox_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComboBox) ClassType() TClass {
    return ComboBox_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComboBox) ClassName() string {
    return ComboBox_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComboBox) InstanceSize() int32 {
    return ComboBox_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComboBox) InheritsFrom(AClass TClass) bool {
    return ComboBox_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComboBox) Equals(Obj IObject) bool {
    return ComboBox_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComboBox) GetHashCode() int32 {
    return ComboBox_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TComboBox) ToString() string {
    return ComboBox_ToString(c.instance)
}

func (c *TComboBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ComboBox_AnchorToNeighbour(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (c *TComboBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ComboBox_AnchorParallel(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (c *TComboBox) AnchorHorizontalCenterTo(ASibling IControl) {
    ComboBox_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (c *TComboBox) AnchorVerticalCenterTo(ASibling IControl) {
    ComboBox_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TComboBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ComboBox_AnchorAsAlign(c.instance, ATheAlign , ASpace)
}

func (c *TComboBox) AnchorClient(ASpace int32) {
    ComboBox_AnchorClient(c.instance, ASpace)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TComboBox) Align() TAlign {
    return ComboBox_GetAlign(c.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TComboBox) SetAlign(value TAlign) {
    ComboBox_SetAlign(c.instance, value)
}

func (c *TComboBox) AutoComplete() bool {
    return ComboBox_GetAutoComplete(c.instance)
}

func (c *TComboBox) SetAutoComplete(value bool) {
    ComboBox_SetAutoComplete(c.instance, value)
}

func (c *TComboBox) AutoDropDown() bool {
    return ComboBox_GetAutoDropDown(c.instance)
}

func (c *TComboBox) SetAutoDropDown(value bool) {
    ComboBox_SetAutoDropDown(c.instance, value)
}

func (c *TComboBox) Style() TComboBoxStyle {
    return ComboBox_GetStyle(c.instance)
}

func (c *TComboBox) SetStyle(value TComboBoxStyle) {
    ComboBox_SetStyle(c.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (c *TComboBox) Anchors() TAnchors {
    return ComboBox_GetAnchors(c.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (c *TComboBox) SetAnchors(value TAnchors) {
    ComboBox_SetAnchors(c.instance, value)
}

func (c *TComboBox) BiDiMode() TBiDiMode {
    return ComboBox_GetBiDiMode(c.instance)
}

func (c *TComboBox) SetBiDiMode(value TBiDiMode) {
    ComboBox_SetBiDiMode(c.instance, value)
}

func (c *TComboBox) CharCase() TEditCharCase {
    return ComboBox_GetCharCase(c.instance)
}

func (c *TComboBox) SetCharCase(value TEditCharCase) {
    ComboBox_SetCharCase(c.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (c *TComboBox) Color() TColor {
    return ComboBox_GetColor(c.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (c *TComboBox) SetColor(value TColor) {
    ComboBox_SetColor(c.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (c *TComboBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ComboBox_GetConstraints(c.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (c *TComboBox) SetConstraints(value *TSizeConstraints) {
    ComboBox_SetConstraints(c.instance, CheckPtr(value))
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TComboBox) DoubleBuffered() bool {
    return ComboBox_GetDoubleBuffered(c.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TComboBox) SetDoubleBuffered(value bool) {
    ComboBox_SetDoubleBuffered(c.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TComboBox) DragCursor() TCursor {
    return ComboBox_GetDragCursor(c.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TComboBox) SetDragCursor(value TCursor) {
    ComboBox_SetDragCursor(c.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TComboBox) DragKind() TDragKind {
    return ComboBox_GetDragKind(c.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TComboBox) SetDragKind(value TDragKind) {
    ComboBox_SetDragKind(c.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TComboBox) DragMode() TDragMode {
    return ComboBox_GetDragMode(c.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TComboBox) SetDragMode(value TDragMode) {
    ComboBox_SetDragMode(c.instance, value)
}

func (c *TComboBox) DropDownCount() int32 {
    return ComboBox_GetDropDownCount(c.instance)
}

func (c *TComboBox) SetDropDownCount(value int32) {
    ComboBox_SetDropDownCount(c.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TComboBox) Enabled() bool {
    return ComboBox_GetEnabled(c.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TComboBox) SetEnabled(value bool) {
    ComboBox_SetEnabled(c.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (c *TComboBox) Font() *TFont {
    return AsFont(ComboBox_GetFont(c.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (c *TComboBox) SetFont(value *TFont) {
    ComboBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TComboBox) ItemHeight() int32 {
    return ComboBox_GetItemHeight(c.instance)
}

func (c *TComboBox) SetItemHeight(value int32) {
    ComboBox_SetItemHeight(c.instance, value)
}

func (c *TComboBox) ItemIndex() int32 {
    return ComboBox_GetItemIndex(c.instance)
}

func (c *TComboBox) SetItemIndex(value int32) {
    ComboBox_SetItemIndex(c.instance, value)
}

// CN: 获取最大长度。
// EN: .
func (c *TComboBox) MaxLength() int32 {
    return ComboBox_GetMaxLength(c.instance)
}

// CN: 设置最大长度。
// EN: .
func (c *TComboBox) SetMaxLength(value int32) {
    ComboBox_SetMaxLength(c.instance, value)
}

// CN: 获取使用父容器颜色。
// EN: Get parent color.
func (c *TComboBox) ParentColor() bool {
    return ComboBox_GetParentColor(c.instance)
}

// CN: 设置使用父容器颜色。
// EN: Set parent color.
func (c *TComboBox) SetParentColor(value bool) {
    ComboBox_SetParentColor(c.instance, value)
}

// CN: 获取使用父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TComboBox) ParentDoubleBuffered() bool {
    return ComboBox_GetParentDoubleBuffered(c.instance)
}

// CN: 设置使用父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TComboBox) SetParentDoubleBuffered(value bool) {
    ComboBox_SetParentDoubleBuffered(c.instance, value)
}

// CN: 获取使用父容器字体。
// EN: Get Parent container font.
func (c *TComboBox) ParentFont() bool {
    return ComboBox_GetParentFont(c.instance)
}

// CN: 设置使用父容器字体。
// EN: Set Parent container font.
func (c *TComboBox) SetParentFont(value bool) {
    ComboBox_SetParentFont(c.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (c *TComboBox) ParentShowHint() bool {
    return ComboBox_GetParentShowHint(c.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (c *TComboBox) SetParentShowHint(value bool) {
    ComboBox_SetParentShowHint(c.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TComboBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ComboBox_GetPopupMenu(c.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TComboBox) SetPopupMenu(value IComponent) {
    ComboBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TComboBox) ShowHint() bool {
    return ComboBox_GetShowHint(c.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TComboBox) SetShowHint(value bool) {
    ComboBox_SetShowHint(c.instance, value)
}

func (c *TComboBox) Sorted() bool {
    return ComboBox_GetSorted(c.instance)
}

func (c *TComboBox) SetSorted(value bool) {
    ComboBox_SetSorted(c.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TComboBox) TabOrder() TTabOrder {
    return ComboBox_GetTabOrder(c.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TComboBox) SetTabOrder(value TTabOrder) {
    ComboBox_SetTabOrder(c.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TComboBox) TabStop() bool {
    return ComboBox_GetTabStop(c.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TComboBox) SetTabStop(value bool) {
    ComboBox_SetTabStop(c.instance, value)
}

// CN: 获取文本。
// EN: .
func (c *TComboBox) Text() string {
    strLen := c.GetTextLen()
    if strLen != 0 {
        var buffStr string
        c.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// CN: 设置文本。
// EN: .
func (c *TComboBox) SetText(value string) {
    ComboBox_SetText(c.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TComboBox) Visible() bool {
    return ComboBox_GetVisible(c.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TComboBox) SetVisible(value bool) {
    ComboBox_SetVisible(c.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (c *TComboBox) SetOnChange(fn TNotifyEvent) {
    ComboBox_SetOnChange(c.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TComboBox) SetOnClick(fn TNotifyEvent) {
    ComboBox_SetOnClick(c.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TComboBox) SetOnContextPopup(fn TContextPopupEvent) {
    ComboBox_SetOnContextPopup(c.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (c *TComboBox) SetOnDblClick(fn TNotifyEvent) {
    ComboBox_SetOnDblClick(c.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TComboBox) SetOnDragDrop(fn TDragDropEvent) {
    ComboBox_SetOnDragDrop(c.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TComboBox) SetOnDragOver(fn TDragOverEvent) {
    ComboBox_SetOnDragOver(c.instance, fn)
}

func (c *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
    ComboBox_SetOnDrawItem(c.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TComboBox) SetOnEndDrag(fn TEndDragEvent) {
    ComboBox_SetOnEndDrag(c.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TComboBox) SetOnEnter(fn TNotifyEvent) {
    ComboBox_SetOnEnter(c.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TComboBox) SetOnExit(fn TNotifyEvent) {
    ComboBox_SetOnExit(c.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TComboBox) SetOnKeyDown(fn TKeyEvent) {
    ComboBox_SetOnKeyDown(c.instance, fn)
}

// CN: 设置键键下事件。
// EN: .
func (c *TComboBox) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBox_SetOnKeyPress(c.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TComboBox) SetOnKeyUp(fn TKeyEvent) {
    ComboBox_SetOnKeyUp(c.instance, fn)
}

func (c *TComboBox) SetOnMeasureItem(fn TMeasureItemEvent) {
    ComboBox_SetOnMeasureItem(c.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
    ComboBox_SetOnMouseEnter(c.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
    ComboBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TComboBox) SetOnSelect(fn TNotifyEvent) {
    ComboBox_SetOnSelect(c.instance, fn)
}

func (c *TComboBox) Items() *TStrings {
    return AsStrings(ComboBox_GetItems(c.instance))
}

func (c *TComboBox) SetItems(value IObject) {
    ComboBox_SetItems(c.instance, CheckPtr(value))
}

// CN: 获取选择的文本。
// EN: .
func (c *TComboBox) SelText() string {
    return ComboBox_GetSelText(c.instance)
}

// CN: 设置选择的文本。
// EN: .
func (c *TComboBox) SetSelText(value string) {
    ComboBox_SetSelText(c.instance, value)
}

// CN: 获取画布。
// EN: .
func (c *TComboBox) Canvas() *TCanvas {
    return AsCanvas(ComboBox_GetCanvas(c.instance))
}

func (c *TComboBox) DroppedDown() bool {
    return ComboBox_GetDroppedDown(c.instance)
}

func (c *TComboBox) SetDroppedDown(value bool) {
    ComboBox_SetDroppedDown(c.instance, value)
}

// CN: 获取选择的长度。
// EN: .
func (c *TComboBox) SelLength() int32 {
    return ComboBox_GetSelLength(c.instance)
}

// CN: 设置选择的长度。
// EN: .
func (c *TComboBox) SetSelLength(value int32) {
    ComboBox_SetSelLength(c.instance, value)
}

// CN: 获取选择的启始位置。
// EN: .
func (c *TComboBox) SelStart() int32 {
    return ComboBox_GetSelStart(c.instance)
}

// CN: 设置选择的启始位置。
// EN: .
func (c *TComboBox) SetSelStart(value int32) {
    ComboBox_SetSelStart(c.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (c *TComboBox) DockClientCount() int32 {
    return ComboBox_GetDockClientCount(c.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TComboBox) DockSite() bool {
    return ComboBox_GetDockSite(c.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TComboBox) SetDockSite(value bool) {
    ComboBox_SetDockSite(c.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TComboBox) MouseInClient() bool {
    return ComboBox_GetMouseInClient(c.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TComboBox) VisibleDockClientCount() int32 {
    return ComboBox_GetVisibleDockClientCount(c.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TComboBox) Brush() *TBrush {
    return AsBrush(ComboBox_GetBrush(c.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TComboBox) ControlCount() int32 {
    return ComboBox_GetControlCount(c.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TComboBox) Handle() HWND {
    return ComboBox_GetHandle(c.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TComboBox) ParentWindow() HWND {
    return ComboBox_GetParentWindow(c.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TComboBox) SetParentWindow(value HWND) {
    ComboBox_SetParentWindow(c.instance, value)
}

func (c *TComboBox) Showing() bool {
    return ComboBox_GetShowing(c.instance)
}

// CN: 获取使用停靠管理。
// EN: .
func (c *TComboBox) UseDockManager() bool {
    return ComboBox_GetUseDockManager(c.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (c *TComboBox) SetUseDockManager(value bool) {
    ComboBox_SetUseDockManager(c.instance, value)
}

func (c *TComboBox) Action() *TAction {
    return AsAction(ComboBox_GetAction(c.instance))
}

func (c *TComboBox) SetAction(value IComponent) {
    ComboBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TComboBox) BoundsRect() TRect {
    return ComboBox_GetBoundsRect(c.instance)
}

func (c *TComboBox) SetBoundsRect(value TRect) {
    ComboBox_SetBoundsRect(c.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (c *TComboBox) ClientHeight() int32 {
    return ComboBox_GetClientHeight(c.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (c *TComboBox) SetClientHeight(value int32) {
    ComboBox_SetClientHeight(c.instance, value)
}

func (c *TComboBox) ClientOrigin() TPoint {
    return ComboBox_GetClientOrigin(c.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TComboBox) ClientRect() TRect {
    return ComboBox_GetClientRect(c.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TComboBox) ClientWidth() int32 {
    return ComboBox_GetClientWidth(c.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TComboBox) SetClientWidth(value int32) {
    ComboBox_SetClientWidth(c.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (c *TComboBox) ControlState() TControlState {
    return ComboBox_GetControlState(c.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (c *TComboBox) SetControlState(value TControlState) {
    ComboBox_SetControlState(c.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (c *TComboBox) ControlStyle() TControlStyle {
    return ComboBox_GetControlStyle(c.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (c *TComboBox) SetControlStyle(value TControlStyle) {
    ComboBox_SetControlStyle(c.instance, value)
}

func (c *TComboBox) Floating() bool {
    return ComboBox_GetFloating(c.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TComboBox) Parent() *TWinControl {
    return AsWinControl(ComboBox_GetParent(c.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TComboBox) SetParent(value IWinControl) {
    ComboBox_SetParent(c.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (c *TComboBox) Left() int32 {
    return ComboBox_GetLeft(c.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (c *TComboBox) SetLeft(value int32) {
    ComboBox_SetLeft(c.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TComboBox) Top() int32 {
    return ComboBox_GetTop(c.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TComboBox) SetTop(value int32) {
    ComboBox_SetTop(c.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (c *TComboBox) Width() int32 {
    return ComboBox_GetWidth(c.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (c *TComboBox) SetWidth(value int32) {
    ComboBox_SetWidth(c.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (c *TComboBox) Height() int32 {
    return ComboBox_GetHeight(c.instance)
}

// CN: 设置高度。
// EN: Set height.
func (c *TComboBox) SetHeight(value int32) {
    ComboBox_SetHeight(c.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TComboBox) Cursor() TCursor {
    return ComboBox_GetCursor(c.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TComboBox) SetCursor(value TCursor) {
    ComboBox_SetCursor(c.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TComboBox) Hint() string {
    return ComboBox_GetHint(c.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TComboBox) SetHint(value string) {
    ComboBox_SetHint(c.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TComboBox) ComponentCount() int32 {
    return ComboBox_GetComponentCount(c.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (c *TComboBox) ComponentIndex() int32 {
    return ComboBox_GetComponentIndex(c.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (c *TComboBox) SetComponentIndex(value int32) {
    ComboBox_SetComponentIndex(c.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TComboBox) Owner() *TComponent {
    return AsComponent(ComboBox_GetOwner(c.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (c *TComboBox) Name() string {
    return ComboBox_GetName(c.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (c *TComboBox) SetName(value string) {
    ComboBox_SetName(c.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TComboBox) Tag() int {
    return ComboBox_GetTag(c.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TComboBox) SetTag(value int) {
    ComboBox_SetTag(c.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (c *TComboBox) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ComboBox_GetAnchorSideLeft(c.instance))
}

// CN: 设置左边锚点。
// EN: .
func (c *TComboBox) SetAnchorSideLeft(value *TAnchorSide) {
    ComboBox_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (c *TComboBox) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ComboBox_GetAnchorSideTop(c.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (c *TComboBox) SetAnchorSideTop(value *TAnchorSide) {
    ComboBox_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (c *TComboBox) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ComboBox_GetAnchorSideRight(c.instance))
}

// CN: 设置右边锚点。
// EN: .
func (c *TComboBox) SetAnchorSideRight(value *TAnchorSide) {
    ComboBox_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (c *TComboBox) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ComboBox_GetAnchorSideBottom(c.instance))
}

// CN: 设置底边锚点。
// EN: .
func (c *TComboBox) SetAnchorSideBottom(value *TAnchorSide) {
    ComboBox_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TComboBox) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ComboBox_GetChildSizing(c.instance))
}

func (c *TComboBox) SetChildSizing(value *TControlChildSizing) {
    ComboBox_SetChildSizing(c.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (c *TComboBox) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ComboBox_GetBorderSpacing(c.instance))
}

// CN: 设置边框间距。
// EN: .
func (c *TComboBox) SetBorderSpacing(value *TControlBorderSpacing) {
    ComboBox_SetBorderSpacing(c.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TComboBox) DockClients(Index int32) *TControl {
    return AsControl(ComboBox_GetDockClients(c.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (c *TComboBox) Controls(Index int32) *TControl {
    return AsControl(ComboBox_GetControls(c.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TComboBox) Components(AIndex int32) *TComponent {
    return AsComponent(ComboBox_GetComponents(c.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (c *TComboBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ComboBox_GetAnchorSide(c.instance, AKind))
}

