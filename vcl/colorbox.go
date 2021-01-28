
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

type TColorBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewColorBox(owner IComponent) *TColorBox {
    c := new(TColorBox)
    c.instance = ColorBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    // 不是TComponent应该是可以考虑加上的
    // runtime.SetFinalizer(c, (*TColorBox).Free)
    return c
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsColorBox(obj interface{}) *TColorBox {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TColorBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsColorBox.
func ColorBoxFromInst(inst uintptr) *TColorBox {
    return AsColorBox(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsColorBox.
func ColorBoxFromObj(obj IObject) *TColorBox {
    return AsColorBox(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsColorBox.
func ColorBoxFromUnsafePointer(ptr unsafe.Pointer) *TColorBox {
    return AsColorBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (c *TColorBox) Free() {
    if c.instance != 0 {
        ColorBox_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TColorBox) Instance() uintptr {
    return c.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TColorBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TColorBox) IsValid() bool {
    return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TColorBox) Is() TIs {
    return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TColorBox) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TColorBoxClass() TClass {
    return ColorBox_StaticClassType()
}

func (c *TColorBox) AddItem(Item string, AObject IObject) {
    ColorBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// 清除。
func (c *TColorBox) Clear() {
    ColorBox_Clear(c.instance)
}

// 清除选择。
func (c *TColorBox) ClearSelection() {
    ColorBox_ClearSelection(c.instance)
}

// 删除选择的。
func (c *TColorBox) DeleteSelected() {
    ColorBox_DeleteSelected(c.instance)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (c *TColorBox) Focused() bool {
    return ColorBox_Focused(c.instance)
}

// 全选。
func (c *TColorBox) SelectAll() {
    ColorBox_SelectAll(c.instance)
}

// 是否可以获得焦点。
func (c *TColorBox) CanFocus() bool {
    return ColorBox_CanFocus(c.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (c *TColorBox) ContainsControl(Control IControl) bool {
    return ColorBox_ContainsControl(c.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (c *TColorBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ColorBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (c *TColorBox) DisableAlign() {
    ColorBox_DisableAlign(c.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (c *TColorBox) EnableAlign() {
    ColorBox_EnableAlign(c.instance)
}

// 查找子控件。
//
// Find sub controls.
func (c *TColorBox) FindChildControl(ControlName string) *TControl {
    return AsControl(ColorBox_FindChildControl(c.instance, ControlName))
}

func (c *TColorBox) FlipChildren(AllLevels bool) {
    ColorBox_FlipChildren(c.instance, AllLevels)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TColorBox) HandleAllocated() bool {
    return ColorBox_HandleAllocated(c.instance)
}

// 插入一个控件。
//
// Insert a control.
func (c *TColorBox) InsertControl(AControl IControl) {
    ColorBox_InsertControl(c.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (c *TColorBox) Invalidate() {
    ColorBox_Invalidate(c.instance)
}

// 移除一个控件。
//
// Remove a control.
func (c *TColorBox) RemoveControl(AControl IControl) {
    ColorBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (c *TColorBox) Realign() {
    ColorBox_Realign(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TColorBox) Repaint() {
    ColorBox_Repaint(c.instance)
}

// 按比例缩放。
//
// Scale by.
func (c *TColorBox) ScaleBy(M int32, D int32) {
    ColorBox_ScaleBy(c.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (c *TColorBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ColorBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TColorBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (c *TColorBox) SetFocus() {
    ColorBox_SetFocus(c.instance)
}

// 控件更新。
//
// Update.
func (c *TColorBox) Update() {
    ColorBox_Update(c.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TColorBox) BringToFront() {
    ColorBox_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TColorBox) ClientToScreen(Point TPoint) TPoint {
    return ColorBox_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TColorBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ColorBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TColorBox) Dragging() bool {
    return ColorBox_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TColorBox) HasParent() bool {
    return ColorBox_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TColorBox) Hide() {
    ColorBox_Hide(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TColorBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorBox_Perform(c.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TColorBox) Refresh() {
    ColorBox_Refresh(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TColorBox) ScreenToClient(Point TPoint) TPoint {
    return ColorBox_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TColorBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ColorBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TColorBox) SendToBack() {
    ColorBox_SendToBack(c.instance)
}

// 显示控件。
//
// Show control.
func (c *TColorBox) Show() {
    ColorBox_Show(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TColorBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ColorBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TColorBox) GetTextLen() int32 {
    return ColorBox_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TColorBox) SetTextBuf(Buffer string) {
    ColorBox_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TColorBox) FindComponent(AName string) *TComponent {
    return AsComponent(ColorBox_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TColorBox) GetNamePath() string {
    return ColorBox_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TColorBox) Assign(Source IObject) {
    ColorBox_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TColorBox) ClassType() TClass {
    return ColorBox_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TColorBox) ClassName() string {
    return ColorBox_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TColorBox) InstanceSize() int32 {
    return ColorBox_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TColorBox) InheritsFrom(AClass TClass) bool {
    return ColorBox_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TColorBox) Equals(Obj IObject) bool {
    return ColorBox_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TColorBox) GetHashCode() int32 {
    return ColorBox_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TColorBox) ToString() string {
    return ColorBox_ToString(c.instance)
}

func (c *TColorBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ColorBox_AnchorToNeighbour(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (c *TColorBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ColorBox_AnchorParallel(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TColorBox) AnchorHorizontalCenterTo(ASibling IControl) {
    ColorBox_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TColorBox) AnchorVerticalCenterTo(ASibling IControl) {
    ColorBox_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TColorBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ColorBox_AnchorAsAlign(c.instance, ATheAlign , ASpace)
}

func (c *TColorBox) AnchorClient(ASpace int32) {
    ColorBox_AnchorClient(c.instance, ASpace)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TColorBox) Align() TAlign {
    return ColorBox_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TColorBox) SetAlign(value TAlign) {
    ColorBox_SetAlign(c.instance, value)
}

func (c *TColorBox) AutoComplete() bool {
    return ColorBox_GetAutoComplete(c.instance)
}

func (c *TColorBox) SetAutoComplete(value bool) {
    ColorBox_SetAutoComplete(c.instance, value)
}

func (c *TColorBox) AutoDropDown() bool {
    return ColorBox_GetAutoDropDown(c.instance)
}

func (c *TColorBox) SetAutoDropDown(value bool) {
    ColorBox_SetAutoDropDown(c.instance, value)
}

func (c *TColorBox) DefaultColorColor() TColor {
    return ColorBox_GetDefaultColorColor(c.instance)
}

func (c *TColorBox) SetDefaultColorColor(value TColor) {
    ColorBox_SetDefaultColorColor(c.instance, value)
}

func (c *TColorBox) NoneColorColor() TColor {
    return ColorBox_GetNoneColorColor(c.instance)
}

func (c *TColorBox) SetNoneColorColor(value TColor) {
    ColorBox_SetNoneColorColor(c.instance, value)
}

func (c *TColorBox) Selected() TColor {
    return ColorBox_GetSelected(c.instance)
}

func (c *TColorBox) SetSelected(value TColor) {
    ColorBox_SetSelected(c.instance, value)
}

func (c *TColorBox) Style() TColorBoxStyle {
    return ColorBox_GetStyle(c.instance)
}

func (c *TColorBox) SetStyle(value TColorBoxStyle) {
    ColorBox_SetStyle(c.instance, value)
}

// 获取四个角位置的锚点。
func (c *TColorBox) Anchors() TAnchors {
    return ColorBox_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TColorBox) SetAnchors(value TAnchors) {
    ColorBox_SetAnchors(c.instance, value)
}

func (c *TColorBox) BiDiMode() TBiDiMode {
    return ColorBox_GetBiDiMode(c.instance)
}

func (c *TColorBox) SetBiDiMode(value TBiDiMode) {
    ColorBox_SetBiDiMode(c.instance, value)
}

// 获取颜色。
//
// Get color.
func (c *TColorBox) Color() TColor {
    return ColorBox_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TColorBox) SetColor(value TColor) {
    ColorBox_SetColor(c.instance, value)
}

// 获取约束控件大小。
func (c *TColorBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ColorBox_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TColorBox) SetConstraints(value *TSizeConstraints) {
    ColorBox_SetConstraints(c.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (c *TColorBox) DoubleBuffered() bool {
    return ColorBox_GetDoubleBuffered(c.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (c *TColorBox) SetDoubleBuffered(value bool) {
    ColorBox_SetDoubleBuffered(c.instance, value)
}

func (c *TColorBox) DropDownCount() int32 {
    return ColorBox_GetDropDownCount(c.instance)
}

func (c *TColorBox) SetDropDownCount(value int32) {
    ColorBox_SetDropDownCount(c.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TColorBox) Enabled() bool {
    return ColorBox_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TColorBox) SetEnabled(value bool) {
    ColorBox_SetEnabled(c.instance, value)
}

// 获取字体。
//
// Get Font.
func (c *TColorBox) Font() *TFont {
    return AsFont(ColorBox_GetFont(c.instance))
}

// 设置字体。
//
// Set Font.
func (c *TColorBox) SetFont(value *TFont) {
    ColorBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TColorBox) ItemHeight() int32 {
    return ColorBox_GetItemHeight(c.instance)
}

func (c *TColorBox) SetItemHeight(value int32) {
    ColorBox_SetItemHeight(c.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (c *TColorBox) ParentColor() bool {
    return ColorBox_GetParentColor(c.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (c *TColorBox) SetParentColor(value bool) {
    ColorBox_SetParentColor(c.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (c *TColorBox) ParentDoubleBuffered() bool {
    return ColorBox_GetParentDoubleBuffered(c.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (c *TColorBox) SetParentDoubleBuffered(value bool) {
    ColorBox_SetParentDoubleBuffered(c.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TColorBox) ParentFont() bool {
    return ColorBox_GetParentFont(c.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TColorBox) SetParentFont(value bool) {
    ColorBox_SetParentFont(c.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (c *TColorBox) ParentShowHint() bool {
    return ColorBox_GetParentShowHint(c.instance)
}

// 设置以父容器的ShowHint属性为准。
func (c *TColorBox) SetParentShowHint(value bool) {
    ColorBox_SetParentShowHint(c.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (c *TColorBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ColorBox_GetPopupMenu(c.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (c *TColorBox) SetPopupMenu(value IComponent) {
    ColorBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TColorBox) ShowHint() bool {
    return ColorBox_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TColorBox) SetShowHint(value bool) {
    ColorBox_SetShowHint(c.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (c *TColorBox) TabOrder() TTabOrder {
    return ColorBox_GetTabOrder(c.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (c *TColorBox) SetTabOrder(value TTabOrder) {
    ColorBox_SetTabOrder(c.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (c *TColorBox) TabStop() bool {
    return ColorBox_GetTabStop(c.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (c *TColorBox) SetTabStop(value bool) {
    ColorBox_SetTabStop(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TColorBox) Visible() bool {
    return ColorBox_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TColorBox) SetVisible(value bool) {
    ColorBox_SetVisible(c.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (c *TColorBox) SetOnChange(fn TNotifyEvent) {
    ColorBox_SetOnChange(c.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (c *TColorBox) SetOnClick(fn TNotifyEvent) {
    ColorBox_SetOnClick(c.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (c *TColorBox) SetOnContextPopup(fn TContextPopupEvent) {
    ColorBox_SetOnContextPopup(c.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (c *TColorBox) SetOnDragDrop(fn TDragDropEvent) {
    ColorBox_SetOnDragDrop(c.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (c *TColorBox) SetOnDragOver(fn TDragOverEvent) {
    ColorBox_SetOnDragOver(c.instance, fn)
}

func (c *TColorBox) SetOnDropDown(fn TNotifyEvent) {
    ColorBox_SetOnDropDown(c.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (c *TColorBox) SetOnEndDrag(fn TEndDragEvent) {
    ColorBox_SetOnEndDrag(c.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (c *TColorBox) SetOnEnter(fn TNotifyEvent) {
    ColorBox_SetOnEnter(c.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (c *TColorBox) SetOnExit(fn TNotifyEvent) {
    ColorBox_SetOnExit(c.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (c *TColorBox) SetOnKeyDown(fn TKeyEvent) {
    ColorBox_SetOnKeyDown(c.instance, fn)
}

// 设置键键下事件。
func (c *TColorBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorBox_SetOnKeyPress(c.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (c *TColorBox) SetOnKeyUp(fn TKeyEvent) {
    ColorBox_SetOnKeyUp(c.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorBox_SetOnMouseEnter(c.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TColorBox) SetOnSelect(fn TNotifyEvent) {
    ColorBox_SetOnSelect(c.instance, fn)
}

func (c *TColorBox) CharCase() TEditCharCase {
    return ColorBox_GetCharCase(c.instance)
}

func (c *TColorBox) SetCharCase(value TEditCharCase) {
    ColorBox_SetCharCase(c.instance, value)
}

// 获取选择的文本。
func (c *TColorBox) SelText() string {
    return ColorBox_GetSelText(c.instance)
}

// 设置选择的文本。
func (c *TColorBox) SetSelText(value string) {
    ColorBox_SetSelText(c.instance, value)
}

// 获取画布。
func (c *TColorBox) Canvas() *TCanvas {
    return AsCanvas(ColorBox_GetCanvas(c.instance))
}

func (c *TColorBox) DroppedDown() bool {
    return ColorBox_GetDroppedDown(c.instance)
}

func (c *TColorBox) SetDroppedDown(value bool) {
    ColorBox_SetDroppedDown(c.instance, value)
}

func (c *TColorBox) Items() *TStrings {
    return AsStrings(ColorBox_GetItems(c.instance))
}

func (c *TColorBox) SetItems(value IStrings) {
    ColorBox_SetItems(c.instance, CheckPtr(value))
}

// 获取选择的长度。
func (c *TColorBox) SelLength() int32 {
    return ColorBox_GetSelLength(c.instance)
}

// 设置选择的长度。
func (c *TColorBox) SetSelLength(value int32) {
    ColorBox_SetSelLength(c.instance, value)
}

// 获取选择的启始位置。
func (c *TColorBox) SelStart() int32 {
    return ColorBox_GetSelStart(c.instance)
}

// 设置选择的启始位置。
func (c *TColorBox) SetSelStart(value int32) {
    ColorBox_SetSelStart(c.instance, value)
}

func (c *TColorBox) ItemIndex() int32 {
    return ColorBox_GetItemIndex(c.instance)
}

func (c *TColorBox) SetItemIndex(value int32) {
    ColorBox_SetItemIndex(c.instance, value)
}

// 获取依靠客户端总数。
func (c *TColorBox) DockClientCount() int32 {
    return ColorBox_GetDockClientCount(c.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (c *TColorBox) DockSite() bool {
    return ColorBox_GetDockSite(c.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (c *TColorBox) SetDockSite(value bool) {
    ColorBox_SetDockSite(c.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (c *TColorBox) MouseInClient() bool {
    return ColorBox_GetMouseInClient(c.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (c *TColorBox) VisibleDockClientCount() int32 {
    return ColorBox_GetVisibleDockClientCount(c.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (c *TColorBox) Brush() *TBrush {
    return AsBrush(ColorBox_GetBrush(c.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (c *TColorBox) ControlCount() int32 {
    return ColorBox_GetControlCount(c.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (c *TColorBox) Handle() HWND {
    return ColorBox_GetHandle(c.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (c *TColorBox) ParentWindow() HWND {
    return ColorBox_GetParentWindow(c.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (c *TColorBox) SetParentWindow(value HWND) {
    ColorBox_SetParentWindow(c.instance, value)
}

func (c *TColorBox) Showing() bool {
    return ColorBox_GetShowing(c.instance)
}

// 获取使用停靠管理。
func (c *TColorBox) UseDockManager() bool {
    return ColorBox_GetUseDockManager(c.instance)
}

// 设置使用停靠管理。
func (c *TColorBox) SetUseDockManager(value bool) {
    ColorBox_SetUseDockManager(c.instance, value)
}

func (c *TColorBox) Action() *TAction {
    return AsAction(ColorBox_GetAction(c.instance))
}

func (c *TColorBox) SetAction(value IComponent) {
    ColorBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TColorBox) BoundsRect() TRect {
    return ColorBox_GetBoundsRect(c.instance)
}

func (c *TColorBox) SetBoundsRect(value TRect) {
    ColorBox_SetBoundsRect(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TColorBox) ClientHeight() int32 {
    return ColorBox_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TColorBox) SetClientHeight(value int32) {
    ColorBox_SetClientHeight(c.instance, value)
}

func (c *TColorBox) ClientOrigin() TPoint {
    return ColorBox_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TColorBox) ClientRect() TRect {
    return ColorBox_GetClientRect(c.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TColorBox) ClientWidth() int32 {
    return ColorBox_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TColorBox) SetClientWidth(value int32) {
    ColorBox_SetClientWidth(c.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (c *TColorBox) ControlState() TControlState {
    return ColorBox_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TColorBox) SetControlState(value TControlState) {
    ColorBox_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TColorBox) ControlStyle() TControlStyle {
    return ColorBox_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TColorBox) SetControlStyle(value TControlStyle) {
    ColorBox_SetControlStyle(c.instance, value)
}

func (c *TColorBox) Floating() bool {
    return ColorBox_GetFloating(c.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TColorBox) Parent() *TWinControl {
    return AsWinControl(ColorBox_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TColorBox) SetParent(value IWinControl) {
    ColorBox_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TColorBox) Left() int32 {
    return ColorBox_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TColorBox) SetLeft(value int32) {
    ColorBox_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TColorBox) Top() int32 {
    return ColorBox_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TColorBox) SetTop(value int32) {
    ColorBox_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TColorBox) Width() int32 {
    return ColorBox_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TColorBox) SetWidth(value int32) {
    ColorBox_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TColorBox) Height() int32 {
    return ColorBox_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TColorBox) SetHeight(value int32) {
    ColorBox_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TColorBox) Cursor() TCursor {
    return ColorBox_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TColorBox) SetCursor(value TCursor) {
    ColorBox_SetCursor(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TColorBox) Hint() string {
    return ColorBox_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TColorBox) SetHint(value string) {
    ColorBox_SetHint(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TColorBox) ComponentCount() int32 {
    return ColorBox_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TColorBox) ComponentIndex() int32 {
    return ColorBox_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TColorBox) SetComponentIndex(value int32) {
    ColorBox_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TColorBox) Owner() *TComponent {
    return AsComponent(ColorBox_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TColorBox) Name() string {
    return ColorBox_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TColorBox) SetName(value string) {
    ColorBox_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TColorBox) Tag() int {
    return ColorBox_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TColorBox) SetTag(value int) {
    ColorBox_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TColorBox) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ColorBox_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TColorBox) SetAnchorSideLeft(value *TAnchorSide) {
    ColorBox_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TColorBox) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ColorBox_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TColorBox) SetAnchorSideTop(value *TAnchorSide) {
    ColorBox_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TColorBox) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ColorBox_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TColorBox) SetAnchorSideRight(value *TAnchorSide) {
    ColorBox_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TColorBox) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ColorBox_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TColorBox) SetAnchorSideBottom(value *TAnchorSide) {
    ColorBox_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TColorBox) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ColorBox_GetChildSizing(c.instance))
}

func (c *TColorBox) SetChildSizing(value *TControlChildSizing) {
    ColorBox_SetChildSizing(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TColorBox) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ColorBox_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TColorBox) SetBorderSpacing(value *TControlBorderSpacing) {
    ColorBox_SetBorderSpacing(c.instance, CheckPtr(value))
}

func (c *TColorBox) Colors(Index int32) TColor {
    return ColorBox_GetColors(c.instance, Index)
}

func (c *TColorBox) ColorNames(Index int32) string {
    return ColorBox_GetColorNames(c.instance, Index)
}

// 获取指定索引停靠客户端。
func (c *TColorBox) DockClients(Index int32) *TControl {
    return AsControl(ColorBox_GetDockClients(c.instance, Index))
}

// 获取指定索引子控件。
func (c *TColorBox) Controls(Index int32) *TControl {
    return AsControl(ColorBox_GetControls(c.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TColorBox) Components(AIndex int32) *TComponent {
    return AsComponent(ColorBox_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TColorBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ColorBox_GetAnchorSide(c.instance, AKind))
}

