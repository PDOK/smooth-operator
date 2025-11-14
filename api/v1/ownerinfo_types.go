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
	model "github.com/pdok/smooth-operator/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OwnerInfoSpec defines the desired state of OwnerInfo.
// +kubebuilder:validation:XValidation:message="providerSite is required when using wms or wfs",rule="(!has(self.wms) && !has(self.wfs)) || has(self.providerSite)"
type OwnerInfoSpec struct {
	MetadataUrls      *MetadataUrls `json:"metadataUrls,omitempty"`
	Atom              *Atom         `json:"atom,omitempty"`
	WFS               *WFS          `json:"wfs,omitempty"`
	WMS               *WMS          `json:"wms,omitempty"`
	NamespaceTemplate *string       `json:"namespaceTemplate,omitempty"`
	ProviderSite      *ProviderSite `json:"providerSite,omitempty"`
}

// MetadataUrls contains various URL templates for metadata access
type MetadataUrls struct {
	CSW        *MetadataURL `json:"csw,omitempty"`
	OpenSearch *MetadataURL `json:"openSearch,omitempty"`
	HTML       *MetadataURL `json:"html,omitempty"`
}

// MetadataURL holds information about URL templates for specific metadata formats
type MetadataURL struct {
	HrefTemplate string `json:"hrefTemplate"`
	Type         string `json:"type"`
}

// Atom contains information about the dataset's author/owner
type Atom struct {
	Author model.Author `json:"author"`

	DefaultStylesheet *model.URL `json:"defaultStylesheet,omitempty"`
}

// WFS contains Web Feature Service related information
type WFS struct {
	ServiceProvider ServiceProvider `json:"serviceProvider"`
}

// ServiceProvider describes the provider of the WFS service
type ServiceProvider struct {
	ProviderName   *string         `json:"providerName,omitempty"`
	ServiceContact *ServiceContact `json:"serviceContact,omitempty"`
}

// ProviderSite holds information about the provider's site
type ProviderSite struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

// ServiceContact provides contact information for the service
type ServiceContact struct {
	IndividualName *string      `json:"individualName,omitempty"`
	PositionName   *string      `json:"positionName,omitempty"`
	ContactInfo    *ContactInfo `json:"contactInfo,omitempty"`
	Role           *string      `json:"role,omitempty"`
}

// ContactInfo encapsulates various types of contact information
type ContactInfo struct {
	Text                *string         `json:"text,omitempty"`
	Phone               *Phone          `json:"phone,omitempty"`
	Address             *Address        `json:"address,omitempty"`
	OnlineResource      *OnlineResource `json:"onlineResource,omitempty"`
	HoursOfService      *string         `json:"hoursOfService,omitempty"`
	ContactInstructions *string         `json:"contactInstructions,omitempty"`
}

// Phone contains voice and facsimile numbers
type Phone struct {
	Voice     *string `json:"voice,omitempty"`
	Facsimile *string `json:"facsimile,omitempty"`
}

// Address provides physical address details
type Address struct {
	DeliveryPoint         *string `json:"deliveryPoint,omitempty"`
	City                  *string `json:"city,omitempty"`
	AdministrativeArea    *string `json:"administrativeArea,omitempty"`
	PostalCode            *string `json:"postalCode,omitempty"`
	Country               *string `json:"country,omitempty"`
	ElectronicMailAddress *string `json:"electronicMailAddress,omitempty"`
}

// OnlineResource describes an online resource associated with the contact
type OnlineResource struct {
	Type *string `json:"type,omitempty"`
	Href *string `json:"href,omitempty"`
}

// WMS contains Web Map Service related information
type WMS struct {
	ContactInformation ContactInformation `json:"contactInformation"`
}

// Information about a contact person for the service
type ContactInformation struct {
	ContactPersonPrimary         *ContactPersonPrimary `json:"contactPersonPrimary,omitempty"`
	ContactPosition              *string               `json:"contactPosition,omitempty"`
	ContactAddress               *ContactAddress       `json:"contactAddress,omitempty"`
	ContactVoiceTelephone        *string               `json:"contactVoiceTelephone,omitempty"`
	ContactFacsimileTelephone    *string               `json:"contactFacsimileTelephone,omitempty"`
	ContactElectronicMailAddress *string               `json:"contactElectronicMailAddress,omitempty"`
}

// The primary contact person
type ContactPersonPrimary struct {
	ContactPerson       *string `json:"contactPerson,omitempty"`
	ContactOrganization *string `json:"contactOrganization,omitempty"`
}

// The address for the contact supplying the service
type ContactAddress struct {
	AddressType     *string `json:"addressType,omitempty"`
	Address         *string `json:"address,omitempty"`
	City            *string `json:"city,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	PostCode        *string `json:"postCode,omitempty"`
	Country         *string `json:"country,omitempty"`
}

// OwnerInfoStatus defines the observed state of OwnerInfo.
type OwnerInfoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=pdok
// +kubebuilder:resource:path=ownerinfo

// OwnerInfo is the Schema for the ownerinfo API.
type OwnerInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OwnerInfoSpec    `json:"spec"`
	Status *OwnerInfoStatus `json:"status,omitempty"`
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
