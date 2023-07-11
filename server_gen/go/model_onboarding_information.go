/*
 * CAPIF_API_Invoker_Management_API
 *
 * API for API invoker management. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OnboardingInformation struct {
	// The API Invoker’s public key
	ApiInvokerPublicKey string `json:"apiInvokerPublicKey"`
	// The API Invoker’s generic client certificate, provided by the CAPIF core function.
	ApiInvokerCertificate string `json:"apiInvokerCertificate,omitempty"`
	// The API Invoker’s onboarding secret, provided by the CAPIF core function.
	OnboardingSecret string `json:"onboardingSecret,omitempty"`
}
