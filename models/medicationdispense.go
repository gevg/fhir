// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MedicationDispense struct {
	DomainResource            `bson:",inline"`
	Identifier                *Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                    string                                   `bson:"status,omitempty" json:"status,omitempty"`
	MedicationCodeableConcept *CodeableConcept                         `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                               `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Patient                   *Reference                               `bson:"patient,omitempty" json:"patient,omitempty"`
	SupportingInformation     []Reference                              `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Dispenser                 *Reference                               `bson:"dispenser,omitempty" json:"dispenser,omitempty"`
	DispensingOrganization    *Reference                               `bson:"dispensingOrganization,omitempty" json:"dispensingOrganization,omitempty"`
	AuthorizingPrescription   []Reference                              `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	Type                      *CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Quantity                  *Quantity                                `bson:"quantity,omitempty" json:"quantity,omitempty"`
	DaysSupply                *Quantity                                `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	WhenPrepared              *FHIRDateTime                            `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	WhenHandedOver            *FHIRDateTime                            `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination               *Reference                               `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver                  []Reference                              `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Note                      []Annotation                             `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction         []DosageInstruction                      `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	Substitution              *MedicationDispenseSubstitutionComponent `bson:"substitution,omitempty" json:"substitution,omitempty"`
	EventHistory              []Reference                              `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationDispense) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicationDispense"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MedicationDispense), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MedicationDispense) GetBSON() (interface{}, error) {
	x.ResourceType = "MedicationDispense"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "medicationDispense" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationDispense MedicationDispense

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationDispense) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationDispense{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationDispense(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationDispense) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationDispense"
	} else if x.ResourceType != "MedicationDispense" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicationDispense, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationDispenseSubstitutionComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Reason           []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ResponsibleParty []Reference       `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
}

type MedicationDispensePlus struct {
	MedicationDispense                     `bson:",inline"`
	MedicationDispensePlusRelatedResources `bson:",inline"`
}

type MedicationDispensePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByReceiver           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByReceiver,omitempty"`
	IncludedPatientResourcesReferencedByReceiver                *[]Patient               `bson:"_includedPatientResourcesReferencedByReceiver,omitempty"`
	IncludedLocationResourcesReferencedByDestination            *[]Location              `bson:"_includedLocationResourcesReferencedByDestination,omitempty"`
	IncludedMedicationResourcesReferencedByMedication           *[]Medication            `bson:"_includedMedicationResourcesReferencedByMedication,omitempty"`
	IncludedPractitionerResourcesReferencedByResponsibleparty   *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByResponsibleparty,omitempty"`
	IncludedPractitionerResourcesReferencedByDispenser          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByDispenser,omitempty"`
	IncludedMedicationRequestResourcesReferencedByPrescription  *[]MedicationRequest     `bson:"_includedMedicationRequestResourcesReferencedByPrescription,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPractitionerResourcesReferencedByReceiver() (practitioners []Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByReceiver == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *m.IncludedPractitionerResourcesReferencedByReceiver
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPatientResourcesReferencedByReceiver() (patients []Patient, err error) {
	if m.IncludedPatientResourcesReferencedByReceiver == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *m.IncludedPatientResourcesReferencedByReceiver
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedLocationResourceReferencedByDestination() (location *Location, err error) {
	if m.IncludedLocationResourcesReferencedByDestination == nil {
		err = errors.New("Included locations not requested")
	} else if len(*m.IncludedLocationResourcesReferencedByDestination) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*m.IncludedLocationResourcesReferencedByDestination))
	} else if len(*m.IncludedLocationResourcesReferencedByDestination) == 1 {
		location = &(*m.IncludedLocationResourcesReferencedByDestination)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedMedicationResourceReferencedByMedication() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByMedication == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByMedication))
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByMedication)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPractitionerResourcesReferencedByResponsibleparty() (practitioners []Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByResponsibleparty == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *m.IncludedPractitionerResourcesReferencedByResponsibleparty
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPractitionerResourceReferencedByDispenser() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByDispenser == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByDispenser) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByDispenser))
	} else if len(*m.IncludedPractitionerResourcesReferencedByDispenser) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByDispenser)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedMedicationRequestResourcesReferencedByPrescription() (medicationRequests []MedicationRequest, err error) {
	if m.IncludedMedicationRequestResourcesReferencedByPrescription == nil {
		err = errors.New("Included medicationRequests not requested")
	} else {
		medicationRequests = *m.IncludedMedicationRequestResourcesReferencedByPrescription
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByPatient))
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if m.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *m.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if m.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *m.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if m.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *m.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if m.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *m.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if m.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *m.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if m.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *m.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByReceiver {
			rsc := (*m.IncludedPractitionerResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByReceiver {
			rsc := (*m.IncludedPatientResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedLocationResourcesReferencedByDestination != nil {
		for idx := range *m.IncludedLocationResourcesReferencedByDestination {
			rsc := (*m.IncludedLocationResourcesReferencedByDestination)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByResponsibleparty != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByResponsibleparty {
			rsc := (*m.IncludedPractitionerResourcesReferencedByResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByDispenser != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByDispenser {
			rsc := (*m.IncludedPractitionerResourcesReferencedByDispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationRequestResourcesReferencedByPrescription != nil {
		for idx := range *m.IncludedMedicationRequestResourcesReferencedByPrescription {
			rsc := (*m.IncludedMedicationRequestResourcesReferencedByPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByReceiver {
			rsc := (*m.IncludedPractitionerResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByReceiver {
			rsc := (*m.IncludedPatientResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedLocationResourcesReferencedByDestination != nil {
		for idx := range *m.IncludedLocationResourcesReferencedByDestination {
			rsc := (*m.IncludedLocationResourcesReferencedByDestination)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByResponsibleparty != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByResponsibleparty {
			rsc := (*m.IncludedPractitionerResourcesReferencedByResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByDispenser != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByDispenser {
			rsc := (*m.IncludedPractitionerResourcesReferencedByDispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationRequestResourcesReferencedByPrescription != nil {
		for idx := range *m.IncludedMedicationRequestResourcesReferencedByPrescription {
			rsc := (*m.IncludedMedicationRequestResourcesReferencedByPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
