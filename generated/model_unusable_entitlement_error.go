/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type UnusableEntitlementError struct {
	// The tag for violating an entitlement action.
	ReasonTag string `json:"reason_tag"`
	// Short tag describing the unusable entitlement action taken by the user.
	EntitlementTag string `json:"entitlement_tag"`
	// Message displayed to the user on why their action failed.
	Message string `json:"message"`
}
