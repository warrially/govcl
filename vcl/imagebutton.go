
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

type TImageButton struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewImageButton(owner IComponent) *TImageButton {
    i := new(TImageButton)
    i.instance = ImageButton_Create(CheckPtr(owner))
    i.ptr = unsafe.Pointer(i.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(i, (*TImageButton).Free)
    return i
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsImageButton(obj interface{}) *TImageButton {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TImageButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsImageButton.
func ImageButtonFromInst(inst uintptr) *TImageButton {
    return AsImageButton(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsImageButton.
func ImageButtonFromObj(obj IObject) *TImageButton {
    return AsImageButton(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsImageButton.
func ImageButtonFromUnsafePointer(ptr unsafe.Pointer) *TImageButton {
    return AsImageButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (i *TImageButton) Free() {
    if i.instance != 0 {
        ImageButton_Free(i.instance)
        i.instance, i.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TImageButton) Instance() uintptr {
    return i.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TImageButton) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TImageButton) IsValid() bool {
    return i.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (i *TImageButton) Is() TIs {
    return TIs(i.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (i *TImageButton) As() TAs {
//    return TAs(i.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TImageButtonClass() TClass {
    return ImageButton_StaticClassType()
}

// CN: 单击。
// EN: .
func (i *TImageButton) Click() {
    ImageButton_Click(i.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (i *TImageButton) BringToFront() {
    ImageButton_BringToFront(i.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (i *TImageButton) ClientToScreen(Point TPoint) TPoint {
    return ImageButton_ClientToScreen(i.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (i *TImageButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ImageButton_ClientToParent(i.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (i *TImageButton) Dragging() bool {
    return ImageButton_Dragging(i.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (i *TImageButton) HasParent() bool {
    return ImageButton_HasParent(i.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (i *TImageButton) Hide() {
    ImageButton_Hide(i.instance)
}

// CN: 要求重绘。
// EN: Redraw.
func (i *TImageButton) Invalidate() {
    ImageButton_Invalidate(i.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (i *TImageButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ImageButton_Perform(i.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (i *TImageButton) Refresh() {
    ImageButton_Refresh(i.instance)
}

// CN: 重绘。
// EN: Repaint.
func (i *TImageButton) Repaint() {
    ImageButton_Repaint(i.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (i *TImageButton) ScreenToClient(Point TPoint) TPoint {
    return ImageButton_ScreenToClient(i.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (i *TImageButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ImageButton_ParentToClient(i.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (i *TImageButton) SendToBack() {
    ImageButton_SendToBack(i.instance)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (i *TImageButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ImageButton_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 显示控件。
// EN: Show control.
func (i *TImageButton) Show() {
    ImageButton_Show(i.instance)
}

// CN: 控件更新。
// EN: Update.
func (i *TImageButton) Update() {
    ImageButton_Update(i.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (i *TImageButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ImageButton_GetTextBuf(i.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (i *TImageButton) GetTextLen() int32 {
    return ImageButton_GetTextLen(i.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (i *TImageButton) SetTextBuf(Buffer string) {
    ImageButton_SetTextBuf(i.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (i *TImageButton) FindComponent(AName string) *TComponent {
    return AsComponent(ImageButton_FindComponent(i.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (i *TImageButton) GetNamePath() string {
    return ImageButton_GetNamePath(i.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (i *TImageButton) Assign(Source IObject) {
    ImageButton_Assign(i.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TImageButton) ClassType() TClass {
    return ImageButton_ClassType(i.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TImageButton) ClassName() string {
    return ImageButton_ClassName(i.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TImageButton) InstanceSize() int32 {
    return ImageButton_InstanceSize(i.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TImageButton) InheritsFrom(AClass TClass) bool {
    return ImageButton_InheritsFrom(i.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TImageButton) Equals(Obj IObject) bool {
    return ImageButton_Equals(i.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TImageButton) GetHashCode() int32 {
    return ImageButton_GetHashCode(i.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (i *TImageButton) ToString() string {
    return ImageButton_ToString(i.instance)
}

func (i *TImageButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ImageButton_AnchorToNeighbour(i.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (i *TImageButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ImageButton_AnchorParallel(i.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (i *TImageButton) AnchorHorizontalCenterTo(ASibling IControl) {
    ImageButton_AnchorHorizontalCenterTo(i.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (i *TImageButton) AnchorVerticalCenterTo(ASibling IControl) {
    ImageButton_AnchorVerticalCenterTo(i.instance, CheckPtr(ASibling))
}

func (i *TImageButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ImageButton_AnchorAsAlign(i.instance, ATheAlign , ASpace)
}

func (i *TImageButton) AnchorClient(ASpace int32) {
    ImageButton_AnchorClient(i.instance, ASpace)
}

func (i *TImageButton) Action() *TAction {
    return AsAction(ImageButton_GetAction(i.instance))
}

func (i *TImageButton) SetAction(value IComponent) {
    ImageButton_SetAction(i.instance, CheckPtr(value))
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (i *TImageButton) Align() TAlign {
    return ImageButton_GetAlign(i.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (i *TImageButton) SetAlign(value TAlign) {
    ImageButton_SetAlign(i.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (i *TImageButton) Anchors() TAnchors {
    return ImageButton_GetAnchors(i.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (i *TImageButton) SetAnchors(value TAnchors) {
    ImageButton_SetAnchors(i.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (i *TImageButton) AutoSize() bool {
    return ImageButton_GetAutoSize(i.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (i *TImageButton) SetAutoSize(value bool) {
    ImageButton_SetAutoSize(i.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (i *TImageButton) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ImageButton_GetConstraints(i.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (i *TImageButton) SetConstraints(value *TSizeConstraints) {
    ImageButton_SetConstraints(i.instance, CheckPtr(value))
}

// CN: 获取控件标题。
// EN: Get the control title.
func (i *TImageButton) Caption() string {
    return ImageButton_GetCaption(i.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (i *TImageButton) SetCaption(value string) {
    ImageButton_SetCaption(i.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (i *TImageButton) DragCursor() TCursor {
    return ImageButton_GetDragCursor(i.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (i *TImageButton) SetDragCursor(value TCursor) {
    ImageButton_SetDragCursor(i.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (i *TImageButton) DragKind() TDragKind {
    return ImageButton_GetDragKind(i.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (i *TImageButton) SetDragKind(value TDragKind) {
    ImageButton_SetDragKind(i.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (i *TImageButton) DragMode() TDragMode {
    return ImageButton_GetDragMode(i.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (i *TImageButton) SetDragMode(value TDragMode) {
    ImageButton_SetDragMode(i.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (i *TImageButton) Enabled() bool {
    return ImageButton_GetEnabled(i.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (i *TImageButton) SetEnabled(value bool) {
    ImageButton_SetEnabled(i.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (i *TImageButton) Font() *TFont {
    return AsFont(ImageButton_GetFont(i.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (i *TImageButton) SetFont(value *TFont) {
    ImageButton_SetFont(i.instance, CheckPtr(value))
}

func (i *TImageButton) ImageCount() int32 {
    return ImageButton_GetImageCount(i.instance)
}

func (i *TImageButton) SetImageCount(value int32) {
    ImageButton_SetImageCount(i.instance, value)
}

// CN: 获取模态对话框显示结果。
// EN: .
func (i *TImageButton) ModalResult() TModalResult {
    return ImageButton_GetModalResult(i.instance)
}

// CN: 设置模态对话框显示结果。
// EN: .
func (i *TImageButton) SetModalResult(value TModalResult) {
    ImageButton_SetModalResult(i.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (i *TImageButton) ParentShowHint() bool {
    return ImageButton_GetParentShowHint(i.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (i *TImageButton) SetParentShowHint(value bool) {
    ImageButton_SetParentShowHint(i.instance, value)
}

// CN: 获取使用父容器字体。
// EN: Get Parent container font.
func (i *TImageButton) ParentFont() bool {
    return ImageButton_GetParentFont(i.instance)
}

// CN: 设置使用父容器字体。
// EN: Set Parent container font.
func (i *TImageButton) SetParentFont(value bool) {
    ImageButton_SetParentFont(i.instance, value)
}

// CN: 获取图片。
// EN: .
func (i *TImageButton) Picture() *TPicture {
    return AsPicture(ImageButton_GetPicture(i.instance))
}

// CN: 设置图片。
// EN: .
func (i *TImageButton) SetPicture(value *TPicture) {
    ImageButton_SetPicture(i.instance, CheckPtr(value))
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (i *TImageButton) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ImageButton_GetPopupMenu(i.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (i *TImageButton) SetPopupMenu(value IComponent) {
    ImageButton_SetPopupMenu(i.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (i *TImageButton) ShowHint() bool {
    return ImageButton_GetShowHint(i.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (i *TImageButton) SetShowHint(value bool) {
    ImageButton_SetShowHint(i.instance, value)
}

func (i *TImageButton) ShowCaption() bool {
    return ImageButton_GetShowCaption(i.instance)
}

func (i *TImageButton) SetShowCaption(value bool) {
    ImageButton_SetShowCaption(i.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (i *TImageButton) Visible() bool {
    return ImageButton_GetVisible(i.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (i *TImageButton) SetVisible(value bool) {
    ImageButton_SetVisible(i.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (i *TImageButton) SetOnClick(fn TNotifyEvent) {
    ImageButton_SetOnClick(i.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (i *TImageButton) SetOnContextPopup(fn TContextPopupEvent) {
    ImageButton_SetOnContextPopup(i.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (i *TImageButton) SetOnDblClick(fn TNotifyEvent) {
    ImageButton_SetOnDblClick(i.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (i *TImageButton) SetOnDragDrop(fn TDragDropEvent) {
    ImageButton_SetOnDragDrop(i.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (i *TImageButton) SetOnDragOver(fn TDragOverEvent) {
    ImageButton_SetOnDragOver(i.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (i *TImageButton) SetOnEndDock(fn TEndDragEvent) {
    ImageButton_SetOnEndDock(i.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (i *TImageButton) SetOnEndDrag(fn TEndDragEvent) {
    ImageButton_SetOnEndDrag(i.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (i *TImageButton) SetOnMouseDown(fn TMouseEvent) {
    ImageButton_SetOnMouseDown(i.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (i *TImageButton) SetOnMouseEnter(fn TNotifyEvent) {
    ImageButton_SetOnMouseEnter(i.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (i *TImageButton) SetOnMouseLeave(fn TNotifyEvent) {
    ImageButton_SetOnMouseLeave(i.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (i *TImageButton) SetOnMouseMove(fn TMouseMoveEvent) {
    ImageButton_SetOnMouseMove(i.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (i *TImageButton) SetOnMouseUp(fn TMouseEvent) {
    ImageButton_SetOnMouseUp(i.instance, fn)
}

func (i *TImageButton) BiDiMode() TBiDiMode {
    return ImageButton_GetBiDiMode(i.instance)
}

func (i *TImageButton) SetBiDiMode(value TBiDiMode) {
    ImageButton_SetBiDiMode(i.instance, value)
}

func (i *TImageButton) BoundsRect() TRect {
    return ImageButton_GetBoundsRect(i.instance)
}

func (i *TImageButton) SetBoundsRect(value TRect) {
    ImageButton_SetBoundsRect(i.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (i *TImageButton) ClientHeight() int32 {
    return ImageButton_GetClientHeight(i.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (i *TImageButton) SetClientHeight(value int32) {
    ImageButton_SetClientHeight(i.instance, value)
}

func (i *TImageButton) ClientOrigin() TPoint {
    return ImageButton_GetClientOrigin(i.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (i *TImageButton) ClientRect() TRect {
    return ImageButton_GetClientRect(i.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (i *TImageButton) ClientWidth() int32 {
    return ImageButton_GetClientWidth(i.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (i *TImageButton) SetClientWidth(value int32) {
    ImageButton_SetClientWidth(i.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (i *TImageButton) ControlState() TControlState {
    return ImageButton_GetControlState(i.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (i *TImageButton) SetControlState(value TControlState) {
    ImageButton_SetControlState(i.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (i *TImageButton) ControlStyle() TControlStyle {
    return ImageButton_GetControlStyle(i.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (i *TImageButton) SetControlStyle(value TControlStyle) {
    ImageButton_SetControlStyle(i.instance, value)
}

func (i *TImageButton) Floating() bool {
    return ImageButton_GetFloating(i.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (i *TImageButton) Parent() *TWinControl {
    return AsWinControl(ImageButton_GetParent(i.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (i *TImageButton) SetParent(value IWinControl) {
    ImageButton_SetParent(i.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (i *TImageButton) Left() int32 {
    return ImageButton_GetLeft(i.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (i *TImageButton) SetLeft(value int32) {
    ImageButton_SetLeft(i.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (i *TImageButton) Top() int32 {
    return ImageButton_GetTop(i.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (i *TImageButton) SetTop(value int32) {
    ImageButton_SetTop(i.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (i *TImageButton) Width() int32 {
    return ImageButton_GetWidth(i.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (i *TImageButton) SetWidth(value int32) {
    ImageButton_SetWidth(i.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (i *TImageButton) Height() int32 {
    return ImageButton_GetHeight(i.instance)
}

// CN: 设置高度。
// EN: Set height.
func (i *TImageButton) SetHeight(value int32) {
    ImageButton_SetHeight(i.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (i *TImageButton) Cursor() TCursor {
    return ImageButton_GetCursor(i.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (i *TImageButton) SetCursor(value TCursor) {
    ImageButton_SetCursor(i.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (i *TImageButton) Hint() string {
    return ImageButton_GetHint(i.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (i *TImageButton) SetHint(value string) {
    ImageButton_SetHint(i.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (i *TImageButton) ComponentCount() int32 {
    return ImageButton_GetComponentCount(i.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (i *TImageButton) ComponentIndex() int32 {
    return ImageButton_GetComponentIndex(i.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (i *TImageButton) SetComponentIndex(value int32) {
    ImageButton_SetComponentIndex(i.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (i *TImageButton) Owner() *TComponent {
    return AsComponent(ImageButton_GetOwner(i.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (i *TImageButton) Name() string {
    return ImageButton_GetName(i.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (i *TImageButton) SetName(value string) {
    ImageButton_SetName(i.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (i *TImageButton) Tag() int {
    return ImageButton_GetTag(i.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (i *TImageButton) SetTag(value int) {
    ImageButton_SetTag(i.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (i *TImageButton) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ImageButton_GetAnchorSideLeft(i.instance))
}

// CN: 设置左边锚点。
// EN: .
func (i *TImageButton) SetAnchorSideLeft(value *TAnchorSide) {
    ImageButton_SetAnchorSideLeft(i.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (i *TImageButton) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ImageButton_GetAnchorSideTop(i.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (i *TImageButton) SetAnchorSideTop(value *TAnchorSide) {
    ImageButton_SetAnchorSideTop(i.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (i *TImageButton) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ImageButton_GetAnchorSideRight(i.instance))
}

// CN: 设置右边锚点。
// EN: .
func (i *TImageButton) SetAnchorSideRight(value *TAnchorSide) {
    ImageButton_SetAnchorSideRight(i.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (i *TImageButton) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ImageButton_GetAnchorSideBottom(i.instance))
}

// CN: 设置底边锚点。
// EN: .
func (i *TImageButton) SetAnchorSideBottom(value *TAnchorSide) {
    ImageButton_SetAnchorSideBottom(i.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (i *TImageButton) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ImageButton_GetBorderSpacing(i.instance))
}

// CN: 设置边框间距。
// EN: .
func (i *TImageButton) SetBorderSpacing(value *TControlBorderSpacing) {
    ImageButton_SetBorderSpacing(i.instance, CheckPtr(value))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (i *TImageButton) Components(AIndex int32) *TComponent {
    return AsComponent(ImageButton_GetComponents(i.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (i *TImageButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ImageButton_GetAnchorSide(i.instance, AKind))
}

