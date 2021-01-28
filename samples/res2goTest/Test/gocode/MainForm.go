// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    Panel1        *vcl.TPanel
    BtnShowAbout  *vcl.TButton
    BtnShowFrame1 *vcl.TButton
    BtnShowFrame2 *vcl.TButton
    Splitter1     *vcl.TSplitter
    Panel2        *vcl.TPanel

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// 以字节形式加载
// vcl.Application.CreateForm(&MainForm)

func NewMainForm(owner vcl.IComponent) (root *TMainForm)  {
    vcl.CreateResForm(owner, &root)
    return
}

var mainFormBytes = []byte("\x54\x50\x46\x30\x09\x54\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\xA3\x01\x06\x48\x65\x69\x67\x68\x74\x03\x0C\x02\x03\x54\x6F\x70\x03\xFA\x00\x05\x57\x69\x64\x74\x68\x03\x4C\x03\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x0C\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x4C\x03\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x0C\x02\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xAA\x00\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x0C\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xAA\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x42\x74\x6E\x53\x68\x6F\x77\x41\x62\x6F\x75\x74\x04\x4C\x65\x66\x74\x02\x30\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x88\x00\x05\x57\x69\x64\x74\x68\x02\x58\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x53\x68\x6F\x77\x41\x62\x6F\x75\x74\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x11\x42\x74\x6E\x53\x68\x6F\x77\x41\x62\x6F\x75\x74\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0D\x42\x74\x6E\x53\x68\x6F\x77\x46\x72\x61\x6D\x65\x31\x04\x4C\x65\x66\x74\x02\x30\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xC0\x00\x05\x57\x69\x64\x74\x68\x02\x58\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0A\x53\x68\x6F\x77\x46\x72\x61\x6D\x65\x31\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x12\x42\x74\x6E\x53\x68\x6F\x77\x46\x72\x61\x6D\x65\x31\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0D\x42\x74\x6E\x53\x68\x6F\x77\x46\x72\x61\x6D\x65\x32\x04\x4C\x65\x66\x74\x02\x30\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xF8\x00\x05\x57\x69\x64\x74\x68\x02\x58\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0A\x53\x68\x6F\x77\x46\x72\x61\x6D\x65\x32\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x12\x42\x74\x6E\x53\x68\x6F\x77\x46\x72\x61\x6D\x65\x32\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00\x09\x54\x53\x70\x6C\x69\x74\x74\x65\x72\x09\x53\x70\x6C\x69\x74\x74\x65\x72\x31\x04\x4C\x65\x66\x74\x03\xAA\x00\x06\x48\x65\x69\x67\x68\x74\x03\x0C\x02\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x02\x05\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x03\xAF\x00\x06\x48\x65\x69\x67\x68\x74\x03\x0C\x02\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x9D\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(MainForm, &mainFormBytes)
