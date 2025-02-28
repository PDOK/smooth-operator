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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OwnerInfoSpec defines the desired state of OwnerInfo.
type OwnerInfoSpec struct {
	MetadataUrls MetadataUrls `json:"metadataUrls,omitempty"`
	Atom         Atom         `json:"atom,omitempty"`
	WFS          WFS          `json:"wfs,omitempty"`
}

// MetadataUrls contains various URL templates for metadata access
type MetadataUrls struct {
	CSV        MetadataURL `json:"csv,omitempty"`
	OpenSearch MetadataURL `json:"opensearch,omitempty"`
	HTML       MetadataURL `json:"html,omitempty"`
}

// MetadataURL holds information about URL templates for specific metadata formats
type MetadataURL struct {
	HrefTemplate string `json:"hrefTemplate,omitempty"`
	Type         string `json:"type,omitempty"`
}

// Atom contains information about the dataset's author/owner
type Atom struct {
	Author Author `json:"author,omitempty"`
}

// Author represents the author or owner of the dataset
type Author struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// WFS contains Web Feature Service related information
type WFS struct {
	ServiceProvider ServiceProvider `json:"serviceprovider,omitempty"`
}

// ServiceProvider describes the provider of the WFS service
type ServiceProvider struct {
	ProviderName   string         `json:"providername,omitempty"`
	ProviderSite   ProviderSite   `json:"providersite,omitempty"`
	ServiceContact ServiceContact `json:"servicecontact,omitempty"`
}

// ProviderSite holds information about the provider's site
type ProviderSite struct {
	Type string `json:"type,omitempty"`
	Href string `json:"href,omitempty"`
}

// ServiceContact provides contact information for the service
type ServiceContact struct {
	IndividualName string      `json:"individualname,omitempty"`
	PositionName   string      `json:"positionname,omitempty"`
	ContactInfo    ContactInfo `json:"contactinfo,omitempty"`
	Role           string      `json:"role,omitempty"`
}

// ContactInfo encapsulates various types of contact information
type ContactInfo struct {
	Text                Text           `json:"text,omitempty"`
	Phone               string         `json:"phone,omitempty"`
	Address             Address        `json:"address,omitempty"`
	OnlineResource      OnlineResource `json:"onlineresource,omitempty"`
	HoursOfService      string         `json:"hoursofservice,omitempty"`
	ContactInstructions string         `json:"contactinstructions,omitempty"`
}

// Text contains voice and facsimile numbers
type Text struct {
	Voice     string `json:"voice,omitempty"`
	Facsimile string `json:"facsmile,omitempty"`
}

// Address provides physical address details
type Address struct {
	DeliveryPoint         string `json:"deliverypoint,omitempty"`
	City                  string `json:"city,omitempty"`
	AdministrativeArea    string `json:"administrativearea,omitempty"`
	PostalCode            string `json:"postalcode,omitempty"`
	Country               string `json:"country,omitempty"`
	ElectronicMailAddress string `json:"electronicmailaddress,omitempty"`
}

// OnlineResource describes an online resource associated with the contact
type OnlineResource struct {
	Type string `json:"type,omitempty"`
	Href string `json:"href,omitempty"`
}

// OwnerInfoStatus defines the observed state of OwnerInfo.
type OwnerInfoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OwnerInfo is the Schema for the ownerinfoes API.
type OwnerInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OwnerInfoSpec   `json:"spec,omitempty"`
	Status OwnerInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OwnerInfoList contains a list of OwnerInfo.
type OwnerInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OwnerInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OwnerInfo{}, &OwnerInfoList{})
}
