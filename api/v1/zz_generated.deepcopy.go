//go:build !ignore_autogenerated

/*
MIT License

Copyright (c) 2024 Publieke Dienstverlening op de Kaart

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Address) DeepCopyInto(out *Address) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Address.
func (in *Address) DeepCopy() *Address {
	if in == nil {
		return nil
	}
	out := new(Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Atom) DeepCopyInto(out *Atom) {
	*out = *in
	out.Author = in.Author
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Atom.
func (in *Atom) DeepCopy() *Atom {
	if in == nil {
		return nil
	}
	out := new(Atom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactAddress) DeepCopyInto(out *ContactAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactAddress.
func (in *ContactAddress) DeepCopy() *ContactAddress {
	if in == nil {
		return nil
	}
	out := new(ContactAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactInfo) DeepCopyInto(out *ContactInfo) {
	*out = *in
	out.Text = in.Text
	out.Address = in.Address
	out.OnlineResource = in.OnlineResource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactInfo.
func (in *ContactInfo) DeepCopy() *ContactInfo {
	if in == nil {
		return nil
	}
	out := new(ContactInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactInformation) DeepCopyInto(out *ContactInformation) {
	*out = *in
	out.ContactPersonPrimary = in.ContactPersonPrimary
	out.ContactAddress = in.ContactAddress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactInformation.
func (in *ContactInformation) DeepCopy() *ContactInformation {
	if in == nil {
		return nil
	}
	out := new(ContactInformation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPersonPrimary) DeepCopyInto(out *ContactPersonPrimary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPersonPrimary.
func (in *ContactPersonPrimary) DeepCopy() *ContactPersonPrimary {
	if in == nil {
		return nil
	}
	out := new(ContactPersonPrimary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataURL) DeepCopyInto(out *MetadataURL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataURL.
func (in *MetadataURL) DeepCopy() *MetadataURL {
	if in == nil {
		return nil
	}
	out := new(MetadataURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataUrls) DeepCopyInto(out *MetadataUrls) {
	*out = *in
	out.CSW = in.CSW
	out.OpenSearch = in.OpenSearch
	out.HTML = in.HTML
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataUrls.
func (in *MetadataUrls) DeepCopy() *MetadataUrls {
	if in == nil {
		return nil
	}
	out := new(MetadataUrls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnlineResource) DeepCopyInto(out *OnlineResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnlineResource.
func (in *OnlineResource) DeepCopy() *OnlineResource {
	if in == nil {
		return nil
	}
	out := new(OnlineResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerInfo) DeepCopyInto(out *OwnerInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerInfo.
func (in *OwnerInfo) DeepCopy() *OwnerInfo {
	if in == nil {
		return nil
	}
	out := new(OwnerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OwnerInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerInfoList) DeepCopyInto(out *OwnerInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OwnerInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerInfoList.
func (in *OwnerInfoList) DeepCopy() *OwnerInfoList {
	if in == nil {
		return nil
	}
	out := new(OwnerInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OwnerInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerInfoSpec) DeepCopyInto(out *OwnerInfoSpec) {
	*out = *in
	out.MetadataUrls = in.MetadataUrls
	out.Atom = in.Atom
	out.WFS = in.WFS
	out.WMS = in.WMS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerInfoSpec.
func (in *OwnerInfoSpec) DeepCopy() *OwnerInfoSpec {
	if in == nil {
		return nil
	}
	out := new(OwnerInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerInfoStatus) DeepCopyInto(out *OwnerInfoStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerInfoStatus.
func (in *OwnerInfoStatus) DeepCopy() *OwnerInfoStatus {
	if in == nil {
		return nil
	}
	out := new(OwnerInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSite) DeepCopyInto(out *ProviderSite) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSite.
func (in *ProviderSite) DeepCopy() *ProviderSite {
	if in == nil {
		return nil
	}
	out := new(ProviderSite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceContact) DeepCopyInto(out *ServiceContact) {
	*out = *in
	out.ContactInfo = in.ContactInfo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceContact.
func (in *ServiceContact) DeepCopy() *ServiceContact {
	if in == nil {
		return nil
	}
	out := new(ServiceContact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceProvider) DeepCopyInto(out *ServiceProvider) {
	*out = *in
	out.ProviderSite = in.ProviderSite
	out.ServiceContact = in.ServiceContact
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceProvider.
func (in *ServiceProvider) DeepCopy() *ServiceProvider {
	if in == nil {
		return nil
	}
	out := new(ServiceProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Text) DeepCopyInto(out *Text) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Text.
func (in *Text) DeepCopy() *Text {
	if in == nil {
		return nil
	}
	out := new(Text)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WFS) DeepCopyInto(out *WFS) {
	*out = *in
	out.ServiceProvider = in.ServiceProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WFS.
func (in *WFS) DeepCopy() *WFS {
	if in == nil {
		return nil
	}
	out := new(WFS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WMS) DeepCopyInto(out *WMS) {
	*out = *in
	out.ContactInformation = in.ContactInformation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WMS.
func (in *WMS) DeepCopy() *WMS {
	if in == nil {
		return nil
	}
	out := new(WMS)
	in.DeepCopyInto(out)
	return out
}
